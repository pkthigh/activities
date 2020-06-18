package storage

import (
	"activities/common"
	"activities/library/config"
	"activities/library/logger"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/redis.v5"
)

var storage *Storage

// Storage 存储模块
type Storage struct {
	dbs map[common.MYSQL]*gorm.DB
	rds map[common.REDIS]*redis.Client
}

func init() {
	storage = &Storage{
		dbs: make(map[common.MYSQL]*gorm.DB),
		rds: make(map[common.REDIS]*redis.Client),
	}

	// 载入Mysql
	for name, addr := range config.GetStorageConf().SQL.DBs {
		db, err := gorm.Open("mysql", addr)
		if err != nil {
			logger.ErrorF("connect %v error: %v", name, err)
		}

		db.LogMode(true)
		db.Callback().Create().Replace("gorm:update_time_stamp", updateTimestampForCreateCallback)
		db.Callback().Update().Replace("gorm:update_time_stamp", updateTimestampForUpdateCallback)
		db.Callback().Delete().Replace("gorm:delete", updateTimestampForDeleteCallback)

		storage.dbs[common.MYSQL(name)] = db
	}

	// 载入Redis
	if config.GetStorageConf().Rds.Addr != "" {
		// 默认
		storage.rds[0] = redis.NewClient(&redis.Options{
			Addr:     config.GetStorageConf().Rds.Addr,
			Password: config.GetStorageConf().Rds.Password,
			DB:       common.Default.Int(),
		})
		// 道具
		storage.rds[1] = redis.NewClient(&redis.Options{
			Addr:     config.GetStorageConf().Rds.Addr,
			Password: config.GetStorageConf().Rds.Password,
			DB:       common.ItemRecordStore.Int(),
		})
		// 每手
		storage.rds[2] = redis.NewClient(&redis.Options{
			Addr:     config.GetStorageConf().Rds.Addr,
			Password: config.GetStorageConf().Rds.Password,
			DB:       common.HandOverRecordStore.Int(),
		})
		// 保险
		storage.rds[3] = redis.NewClient(&redis.Options{
			Addr:     config.GetStorageConf().Rds.Addr,
			Password: config.GetStorageConf().Rds.Password,
			DB:       common.InsuranceRecordStore.Int(),
		})
	}

}

// GetSQLDB 获取Mysql
func GetSQLDB(mysql common.MYSQL) (*gorm.DB, bool) {
	db, ok := storage.dbs[mysql]
	return db, ok
}

// GetRdsDB 获取Redis
func GetRdsDB(redis common.REDIS) (*redis.Client, bool) {
	db, ok := storage.rds[redis]
	return db, ok
}

// updateTimestampForCreateCallback will set `CreatedAt`, `UpdatedAt` when creating
func updateTimestampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreatedAt"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("UpdatedAt"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}
}

// updateTimestampForCreateCallback will set `UpdatedAt` when updating
func updateTimestampForUpdateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if modifyTimeField, ok := scope.FieldByName("UpdatedAt"); ok {
			modifyTimeField.Set(nowTime)
		}
	}
}

// updateTimestampForCreateCallback will set `DeletedAt` when deleting
func updateTimestampForDeleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedAt")

		if !scope.Search.Unscoped && hasDeletedOnField {
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedOnField.DBName),
				scope.AddToVars(time.Now().Unix()),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
