package tests

import (
	"log"
	"os"
	"shortener/internal/dbrepo"
	"testing"

	"gorm.io/gorm"
)

var tDb *gorm.DB
var tRepo *dbrepo.Repository

func TestMain(m *testing.M) {
	err := dbrepo.OpenDb()
	if err != nil {
		log.Fatalln("连接数据库失败: ", err)
	}
	tDb = dbrepo.Db

	tRepo = &dbrepo.Repository{
		UserRepo:  dbrepo.NewUserRepo(),
		EntryRepo: dbrepo.NewEntryRepo(),
		ClickRepo: dbrepo.NewClickRepo(),
	}
	os.Exit(m.Run())
}
