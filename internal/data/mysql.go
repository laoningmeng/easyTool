package data

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"github.com/laoningmeng/fusion/internal/pkg/conf"
)

type Mysql struct {
	Host     string
	Username string
	Password string
	Port     int32
	DbName   string
}

type MysqlFaker struct {
	dataSource
}

func (m *MysqlFaker) getConn(path string) *gorm.DB {
	c := conf.NewConf(path)
	var mysqlConn Mysql
	err := c.Vp().Sub("data.mysql").Unmarshal(&mysqlConn)
	if err != nil {
		os.Exit(1)
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,  // Slow SQL threshold
			LogLevel:                  logger.Error, // Log level
			IgnoreRecordNotFoundError: true,         // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,        // Don't include params in the SQL log
			Colorful:                  true,         // Disable color
		},
	)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlConn.Username,
		mysqlConn.Password,
		mysqlConn.Host,
		mysqlConn.Password,
		mysqlConn.DbName)
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 设置单表表名字
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
		Logger:         newLogger,
	})
	if err != nil {
		fmt.Println("ERR:", err)
		os.Exit(1)
	}
	return db
}

func (m *MysqlFaker) generateSql() {
	// 拼装sql

	//生成数据

	//生成一套sql

}
func (m *MysqlFaker) Parse(path string) {
	//conn := m.getConn(path)

}

func NewMysqlFaker() MysqlFaker {
	data := NewFakeIt()
	return MysqlFaker{dataSource: data}
}

func (m *MysqlFaker) Run() {

}
