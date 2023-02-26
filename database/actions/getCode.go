package actions

import (
	"Verifier/database"
	databaseTables "Verifier/models/database"
	restModels "Verifier/models/rest"
	"Verifier/utils"
)

func GetRoleID(roleName string) (string, error) {
	db := database.Get().Table("roles")

	roleData := databaseTables.Roles{}
	err := db.Where("name = ?", roleName).First(&roleData)

	if err.Error != nil {
		return "", err.Error
	}

	return roleData.Role, nil
}

func GetCodeInfo(code string) (databaseTables.Codes, error) {
	db := database.Get()
	query := db.Table("codes").Where("code = ?", code)

	result, err := database.ScanMany[databaseTables.Codes](query)

	if err != nil {
		return databaseTables.Codes{}, err
	}

	return result[0], err
}

func CreateOrGetCode(receiver *restModels.VerifyReceiver) (string, error) {
	db := database.Get().Table("codes")

	code := utils.HashMD5(receiver.Username + receiver.Role)

	newCodeData := &databaseTables.Codes{
		Code:       code,
		Username:   receiver.Username,
		AssignRole: receiver.Role,
	}

	db.Where("code = ?", code).FirstOrCreate(&newCodeData)

	return code, nil
}
