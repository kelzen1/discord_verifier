package database

import (
	databaseTables2 "github.com/yoonaowo/discord_verifier/internal/models/database"
	restModels "github.com/yoonaowo/discord_verifier/internal/models/rest"
	utils2 "github.com/yoonaowo/discord_verifier/internal/utils"

	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"os"
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
	utils2.Logger().Println("connecting to database")

	databaseUrl := os.Getenv("DB_URL")

	db, err := gorm.Open(mysql.Open(databaseUrl), &gorm.Config{})
	dbStruct.raw = db

	if err != nil {
		utils2.Logger().Panicln("cannot connect to database:", err)
		return
	}

	utils2.Logger().Println("database connected")
}

// Get singleton
func Get() T {

	mutex.Lock()
	defer mutex.Unlock()

	once.Do(initOnce)

	return dbStruct
}

func (dbT *T) GetRoleID(roleName string) (string, error) {
	db := dbT.raw.Table("roles")

	roleData := databaseTables2.Roles{}
	res := db.Where("name = ?", roleName).Limit(1).Find(&roleData)

	err := res.Error

	if err == nil && res.RowsAffected == 0 {
		err = sql.ErrNoRows
	}

	return roleData.Role, err
}

func (dbT *T) GetCodeInfo(code string) (databaseTables2.Codes, error) {
	var dest databaseTables2.Codes
	query := dbT.raw.Table("codes").Where("code = ?", code)
	res := query.Limit(1).Find(&dest)

	err := res.Error
	if err == nil && res.RowsAffected == 0 {
		err = sql.ErrNoRows
	}

	return dest, err
}

func (dbT *T) CreateOrGetCode(receiver *restModels.VerifyReceiver) (string, error) {
	db := dbT.raw.Table("codes")

	code := utils2.HashMD5(receiver.Username + receiver.Role)

	newCodeData := &databaseTables2.Codes{
		Code:       code,
		Username:   receiver.Username,
		AssignRole: receiver.Role,
	}

	db.Where("code = ?", code).FirstOrCreate(&newCodeData)

	return code, nil
}

func (dbT *T) SetUsed(UserID string, codeData databaseTables2.Codes) {
	db := dbT.raw

	updateData := map[string]interface{}{
		"used":    true,
		"used_by": UserID,
	}

	db.Table("codes").Where("code = ?", codeData.Code).Updates(updateData)

	if codeData.Username == "unknown" {
		return
	}

	db = dbT.raw.Table("users")

	db.Create(&databaseTables2.Users{
		Username:  codeData.Username,
		DiscordID: UserID,
	})
}
