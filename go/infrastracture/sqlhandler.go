package infrastracture

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jphacks/F_2002_1/go/entity/model"
	"github.com/jphacks/F_2002_1/go/interface/database"
	"github.com/labstack/echo/v4"
)

type SqlHandler struct {
	Conn *gorm.DB
}

func NewSqlHandler() database.SqlHandler {
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = dbConnect()
	return sqlHandler
}

func (handler *SqlHandler) Find(out interface{}, where ...interface{}) *gorm.DB {
	return handler.Conn.Find(out, where...)
}

func (handler *SqlHandler) Exec(sql string, values ...interface{}) *gorm.DB {
	return handler.Conn.Exec(sql, values...)
}

func (handler *SqlHandler) First(out interface{}, where ...interface{}) *gorm.DB {
	return handler.Conn.First(out, where...)
}

func (handler *SqlHandler) Raw(sql string, values ...interface{}) *gorm.DB {
	return handler.Conn.Raw(sql, values...)
}

func (handler *SqlHandler) Create(value interface{}) *gorm.DB {
	return handler.Conn.Create(value)
}

func (handler *SqlHandler) Save(value interface{}) *gorm.DB {
	return handler.Conn.Save(value)
}

func (handler *SqlHandler) Update(attrs ...interface{}) *gorm.DB {
	return handler.Conn.Update(attrs)
}

func (handler *SqlHandler) Delete(value interface{}) *gorm.DB {
	return handler.Conn.Delete(value)
}

func (handler *SqlHandler) Where(query interface{}, args ...interface{}) *gorm.DB {
	return handler.Conn.Where(query, args...)
}

func (handler *SqlHandler) Set(name string, value interface{}) *gorm.DB {
	return handler.Conn.Set(name, value)
}

func (handler *SqlHandler) Query() *gorm.DB {
	return handler.Conn
}

func (handler *SqlHandler) Model(value interface{}) *gorm.DB {
	return handler.Conn.Model(value)
}

func dbRetryConnect(count int) *gorm.DB {
	DBMS := "mysql"
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASSWORD")
	PROTOCOL := "tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")"
	DBNAME := os.Getenv("DB_NAME")
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	conn, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		log.Println("DB not ready. Retconnecting...")
		time.Sleep(1 * time.Second)
		count++
		log.Println("Try count: ", count)
		if count > 30 {
			log.Println("Cannot connect to DB.")
			panic(err.Error())
		}
		return dbRetryConnect(count)
	}
	return conn
}

func dbConnect() *gorm.DB {
	conn := dbRetryConnect(0)

	conn.LogMode(true)

	return conn
}

func dbMigrate(conn *gorm.DB) {
	log.Println("Start auto migration")
	conn.AutoMigrate(
		&model.Example{},
	)
	log.Println("Finish auto migration")

	log.Println("Start inserting data")
	for _, Example := range Examples {
		conn.Create(&Example)
	}
	log.Println("Finish inserting data")
}

func InitDBServer() {
	conn := dbConnect()
	if !conn.HasTable(&model.Example{}) {
		dbMigrate(conn)
	}
}

func dbDrop(conn *gorm.DB) {
	log.Println("Start drop")
	// conn.DropTable(&model.Hoge{})
	log.Println("Finish drop")
}

func ResetDB(c echo.Context) error {
	conn := dbConnect()
	dbDrop(conn)
	dbMigrate(conn)
	return c.String(http.StatusOK, "Reset Complete!")
}
