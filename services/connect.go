// services/database.go

package services

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var credentials map[string]string

func init() {
	godotenv.Load()
	credentials = make(map[string]string)
	credentials["driver"] = os.Getenv("DB_DRIVER")
	credentials["username"] = os.Getenv("DB_USERNAME")
	credentials["password"] = os.Getenv("DB_PASSWORD")
	credentials["database"] = os.Getenv("DB_NAME")
}

func Connect() *gorm.DB {
	db, _ := gorm.Open(credentials["driver"], fmt.Sprint(credentials["username"]+":"+credentials["password"]+"@/"+credentials["database"]+"?charset=utf8&parseTime=True&loc=Local"))
	return db
}
