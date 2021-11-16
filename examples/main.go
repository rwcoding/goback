package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rwcoding/goback"
	"github.com/rwcoding/goback/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	goback.SetLang("zh")
	goback.SetDev(true)
	goback.SetDb(Db())

	g := gin.Default()
	g.POST("/api", func(context *gin.Context) {
		goback.Run(context)
	})
	_ = g.Run("0.0.0.0:9090")
}

type DbConfig struct {
	Host        string
	Port        int
	Username    string
	Password    string
	Dbname      string
	Charset     string
	MaxOpenConn int
	MaxIdleConn int
	MaxLifetime int
}

func Db() *gorm.DB {

	lvl := logger.Warn
	if config.IsDev() {
		lvl = logger.Info
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), //需要文件的替换为文件
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  lvl,         // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	conf := &DbConfig{
		Host:     "127.0.0.1",
		Port:     3306,
		Username: "root",
		Password: "root",
		Dbname:   "gobui",
		Charset:  "utf8mb4",
	}
	if conf.Username == "" && conf.Password == "" {
		return nil
	}
	dsn := conf.Username + ":" + conf.Password + "@tcp(" +
		conf.Host + ":" + strconv.Itoa(conf.Port) + ")/" +
		conf.Dbname + "?charset=" + conf.Charset + "&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Fatal(err)
	}

	//连接池
	sd, _ := db.DB()
	if conf.MaxOpenConn == 0 {
		conf.MaxOpenConn = 100
	}
	if conf.MaxIdleConn == 0 {
		conf.MaxIdleConn = 50
	}
	if conf.MaxLifetime == 0 {
		conf.MaxLifetime = 600
	}
	sd.SetMaxOpenConns(conf.MaxOpenConn)
	sd.SetMaxIdleConns(conf.MaxIdleConn)
	sd.SetConnMaxLifetime(time.Second * time.Duration(conf.MaxLifetime))

	return db
}
