package mysql

import (
	"fmt"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db   *gorm.DB  // mysql链接
	once sync.Once // 保证至连接一次
)

// mysql配置
type MysqlConfig struct {
	User        string `json:"user" yaml:"user"`                                                 // 用户名
	Password    string `json:"password" yaml:"password"`                                         // 密码
	Host        string `json:"host" yaml:"host"`                                                 // 主机地址
	Port        int    `json:"port" yaml:"port"`                                                 // 端口号
	Dbname      string `json:"dbname" yaml:"dbname"`                                             // 数据库名
	MaxIdleConn int    `json:"max_idle_conn" yaml:"max_idle_conn" mapstructure:"max_idle_conn"`  // 最大空闲连接
	MaxOpenConn int    `json:"max_open_conn" yaml:"max_open_conn" mapstructure:"max_open_conn" ` // 最大活跃连接
	Debug       bool   `json:"debug" yaml:"debug"`                                               // 是否开启Debug（开启Debug会打印数据库操作日志）
}

func Init(mysqlConf MysqlConfig) (err error) {
	once.Do(func() {
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local",
			mysqlConf.User,
			mysqlConf.Password,
			mysqlConf.Host,
			mysqlConf.Port,
			mysqlConf.Dbname,
		)
		conn, err := gorm.Open("mysql", dsn)
		if err != nil {
			fmt.Println("mysql connect failed: ", err)
			return
		}
		conn.DB().SetMaxIdleConns(mysqlConf.MaxIdleConn)
		conn.DB().SetMaxOpenConns(mysqlConf.MaxOpenConn)

		conn.LogMode(mysqlConf.Debug)
		if err = conn.DB().Ping(); err != nil {
			fmt.Printf("database heartbeat failed: %v", err)
			return
		}
		db = conn
		fmt.Println("mysql connect successfully")
	})
	return
}

// 获取mysql连接
func GetConn() *gorm.DB {
	return db
}

// 关闭mysql连接
func Close() error {
	if db != nil {
		if err := db.Close(); err != nil {
			return err
		}
	}
	//logrus.Info("mysql connect closed")
	return nil
}
