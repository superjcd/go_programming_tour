package model

import (
	"project2/pkg/setting"
	"project2/global"
	"github.com/jinzhu/gorm"	
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Model struct {...}

func NewDBEngine(databaseSetting *setting.DabaseSetting) (*gorm.DB, error) {
	db, err := gorm.Open(database.DBType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,    
    ))
	if err != nil {
		return nil, err
	}

	if global.ServerSetting.Runmode == "debug" {
		db.LogMode(True)
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)   // 最大连接数
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)  // 最大同时连接数据

	return db, nil

}

