package main

import (
	"log"
	"net/http"
	"shortener/internal/dbrepo"

	"github.com/gin-gonic/gin"
)

type application struct {
}

func main() {

	if err := dbrepo.OpenDb(); err != nil {
		log.Println("连接数据库失败: ", err.Error())
	}

	if err := dbrepo.DbMigrate(); err != nil {
		log.Println("迁移Table失败: ", err.Error())
	}

	router := gin.Default()

	router.Static("/static", "./static/")

	router.LoadHTMLGlob("templates/**/*")

	router.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "views/index.html", gin.H{
			"content": "partials/homecard.html",
		})
	})

	router.GET("/login", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "views/index.html", nil)
	})

	if err := router.Run(":9000"); err != nil {
		log.Fatal(err)
	}
}
