package mysql

import (
	"log"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func TestInit(t *testing.T) {

	mysqlConf := MysqlConfig{
		User:        "admin",
		Password:    "eotPTBF9JJLVBWSn",
		Host:        "192.168.12.73",
		Port:        3306,
		Dbname:      "dr_pcms_db",
		MaxIdleConn: 1024,
		MaxOpenConn: 1024,
		Debug:       true,
	}

	Init(mysqlConf)

	db := GetConn()
	defer Close()

	var data []int
	err := db.Table("t_user_rate_info").Pluck("id", &data).Error
	if err != nil {
		t.Error(err)
	}
	log.Println(data)
}
