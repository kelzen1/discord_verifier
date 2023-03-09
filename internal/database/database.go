package database

import (
	"database/sql"
	"github.com/glebarez/sqlite"
	"github.com/yoonaowo/discord_verifier/internal/models/database"
	"github.com/yoonaowo/discord_verifier/internal/models/rest"
	"github.com/yoonaowo/discord_verifier/internal/utils"
	"gorm.io/gorm"
	"sync"
)

type T struct {
	raw *gorm.DB
}

var dbStruct T

var (
	mutex sync.Mutex
	once  sync.Once
)

func initOnce() {
	utils.Logger().Println("connecting to database")

	db, err := gorm.Open(sqlite.Open("verifier.db"), &gorm.Config{
		PrepareStmt: true,
	})
	dbStruct.raw = db

	if err != nil {
		utils.Logger().Panicln("cannot connect to database:", err)
		return
	}

	err = db.AutoMigrate(&databaseTables.Codes{}, &databaseTables.Roles{}, &databaseTables.Users{})
	if err != nil {
		utils.Logger().Panicln("cannot migrate database:", err)
		return
	}

	utils.Logger().Println("database connected")
}

// Get singleton
func Get() T {

	mutex.Lock()
	defer mutex.Unlock()

	once.Do(initOnce)

	return dbStruct
}

func (db *T) GetRoleID(roleName string) (string, error) {

	roleData := databaseTables.Roles{}
	res := db.raw.Where(&databaseTables.Roles{Name: roleName}).Limit(1).Find(&roleData) // lifehack to prevent console spam caused by gorm

	err := res.Error

	if err == nil && res.RowsAffected == 0 {
		err = sql.ErrNoRows
	}

	return roleData.Role, err
}

func (db *T) GetCodeInfo(code string) (databaseTables.Codes, error) {
	var dest databaseTables.Codes
	res := db.raw.Where(&databaseTables.Codes{Code: code}).Limit(1).Find(&dest)

	err := res.Error
	if err == nil && res.RowsAffected == 0 {
		err = sql.ErrNoRows
	}

	return dest, err
}

func (db *T) CreateOrGetCode(receiver restModels.VerifyReceiver) (string, error) {

	code := utils.HashMD5(receiver.Username + receiver.Role)

	newCodeData := &databaseTables.Codes{
		Code:       code,
		Username:   receiver.Username,
		AssignRole: receiver.Role,
	}

	db.raw.FirstOrCreate(&databaseTables.Codes{}, &newCodeData)

	return code, nil
}

func (db *T) SetUsed(UserID string, codeData databaseTables.Codes) {

	db.raw.Where(databaseTables.Codes{Code: codeData.Code}).Updates(databaseTables.Codes{
		Used:   true,
		UsedBy: UserID,
	})

	if codeData.Username == "unknown" {
		return
	}

	db.raw.Create(&databaseTables.Users{
		Username:  codeData.Username,
		DiscordID: UserID,
	})
}
