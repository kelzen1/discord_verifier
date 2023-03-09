package database

import databaseTables "github.com/yoonaowo/discord_verifier/internal/models/database"

func (db *T) EditRole(roleName, roleID string) error {

	var actualData databaseTables.Roles

	scanRes := db.raw.Where(&databaseTables.Roles{
		Name: roleName,
	}).Limit(1).Find(&actualData)

	if scanRes.Error != nil {
		return scanRes.Error
	}

	if scanRes.RowsAffected == 0 {

		res := db.raw.Create(&databaseTables.Roles{
			Name: roleName,
			Role: roleID,
		})

		return res.Error
	}

	res := db.raw.Where(databaseTables.Roles{ID: actualData.ID}).Updates(databaseTables.Roles{
		Role: roleID,
	})

	return res.Error
}

func (db *T) ListRoles() []databaseTables.Roles {
	rows, _ := db.raw.Model(&databaseTables.Roles{}).Rows()
	res := make([]databaseTables.Roles, 0)

	for rows.Next() {
		var tmp databaseTables.Roles
		_ = db.raw.ScanRows(rows, &tmp)

		res = append(res, tmp)
	}

	return res
}

func (db *T) DeleteRole(roleName string) error {

	res := db.raw.Where(&databaseTables.Roles{Name: roleName}).Unscoped().Delete(&databaseTables.Roles{})

	return res.Error
}
