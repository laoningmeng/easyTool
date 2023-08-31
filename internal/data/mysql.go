package data

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
	Db       string
}

type MysqlFaker struct {
	Data  dataSource
	table []table
	conf  string
	size  int
}

func (m *MysqlFaker) getConn() *gorm.DB {
	c := conf.NewConf(m.conf)
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
		mysqlConn.Port,
		mysqlConn.Db)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
		Logger:         newLogger,
	})
	if err != nil {
		fmt.Println("ERR:", err)
		os.Exit(1)
	}
	return db
}

func (m *MysqlFaker) Parse() {
	db := m.getConn()
	for _, t := range m.table {
		execNum := t.number / m.size
		if execNum != 0 {
			for i := 0; i < execNum; i++ {
				m.chunk(m.size, t, db)
			}
		}
		execRemain := t.number % m.size
		if execRemain != 0 {
			m.chunk(execRemain, t, db)
		}
	}
}

func (m *MysqlFaker) chunk(batch int, t table, db *gorm.DB) {
	var field []string
	for _, c := range t.columns {
		field = append(field, c.name)
	}
	var sqlRecord []string
	for i := 0; i < batch; i++ {
		var value []string
		for _, c := range t.columns {
			v := m.ParseColumn(c)
			value = append(value, fmt.Sprintf("'%v'", v))
		}
		valueStr := strings.Join(value, ",")
		record := "(" + valueStr + ")"
		sqlRecord = append(sqlRecord, record)
	}
	fieldStr := strings.Join(field, ",")
	sqlValueRecord := strings.Join(sqlRecord, ",")
	sql := fmt.Sprintf("insert into %s (%s) values %s;", t.tableName, fieldStr, sqlValueRecord)
	tx := db.Exec(sql)
	if tx.Error != nil {
		os.Exit(1)
	}
}
func (m *MysqlFaker) ParseColumn(c column) string {
	s := strings.Split(c.val, "-")
	if len(s) < 1 {
		os.Exit(1)
	}
	var args []interface{}
	for _, e := range s[1:] {
		i, err := strconv.Atoi(e)
		if err == nil {
			args = append(args, i)
		} else {
			args = append(args, e)
		}
	}
	res, err := m.Data.invoke(s[0], args...)
	if err != nil {
		os.Exit(1)
	}
	return fmt.Sprintf("%v", res)
}

func NewMysqlFaker() MysqlFaker {
	data := NewFakeIt()
	return MysqlFaker{Data: data}
}

func (m *MysqlFaker) Run(path string) {
	m.conf = path
	m.size = 10
	m.setTables()
	m.Parse()
}

type table struct {
	tableName string
	number    int
	confKey   string
	columns   []column
}
type column struct {
	name    string
	confKey string
	val     string
}

func (m *MysqlFaker) setTables() {
	c := conf.NewConf(m.conf)
	tables := c.Vp().Sub("data.tables").AllSettings()
	var tableList []table
	for tableName, _ := range tables {
		coluns := c.Vp().Sub("data.tables." + tableName).AllSettings()
		var columnsList []column
		var number int
		for columnName, columnItem := range coluns {
			n := c.Vp().Get("data.tables." + tableName + "._number")
			if _, ok := n.(int); !ok {
				os.Exit(1)
			}
			numberConf := n.(int)
			if numberConf == 0 {
				number = 10
			} else {
				number = numberConf
			}
			if columnName == "_number" {
				continue
			}
			columnsList = append(columnsList, column{
				name:    columnName,
				confKey: "data.tables." + tableName + "." + columnName,
				val:     fmt.Sprintf("%v", columnItem),
			})
		}
		tableList = append(tableList, table{
			tableName: tableName,
			number:    number,
			confKey:   "data.tables." + tableName,
			columns:   columnsList,
		})
	}
	m.table = tableList
}
