package main

import (
	"fmt"
	"log"

	"github.com/JIeeiroSst/store/config"
	"github.com/JIeeiroSst/store/middleware"
	"github.com/JIeeiroSst/store/pkg/bigcache"
	"github.com/JIeeiroSst/store/pkg/mysql"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	config, err := config.ReadConfig("config.yml")
	if err != nil {
		log.Fatal(err)
	}

	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Mysql.MysqlUser,
		config.Mysql.MysqlPassword,
		config.Mysql.MysqlHost,
		config.Mysql.MysqlPort,
		config.Mysql.MysqlDbname,
	)

	mysqlOrm, _ := mysql.InitMysql(dns)

	mysqlOrm.AutoMigrate(&CasbinRule{})

	cache := bigcache.NewBigCache()

	casbin := middleware.NewAuthorization(*cache)

	adapter, _ := gormadapter.NewAdapterByDB(mysqlOrm)

	resource := router.Group("/api")
	resource.Use(casbin.Authenticate())
	{
		router.GET("/", func(ctx *gin.Context) {
			router.GET("/", casbin.Authorize("/api/user/*", "POST", adapter), hello)
		})

	}

	router.Run(":1234")
}

func hello(ctx *gin.Context) {
	ctx.String(200, "hello world")
}

type CasbinRule struct {
	ID    int `gorm:"primaryKey;autoIncrement"`
	Ptype string
	V0    string
	V1    string
	V2    string
	V3    string
	V4    string
	V5    string
}
