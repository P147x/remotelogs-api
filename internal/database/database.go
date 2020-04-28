package database

import (
	"log"

	"github.com/P147x/remotelogs-api/internal/model"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

var db *gorm.DB

//GetDatabase is a singleton used to get the database instance
func GetDatabase() *gorm.DB {
	if db == nil {
		InitDatabase()
	}
	return db
}

// InitDatabase is used for the first instanciation of the database, connect to the server and creating the missing tables.
func InitDatabase() {
	cstr := "user:password@tcp(database:3306)/mydatabase?charset=utf8&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open("mysql", cstr)
	log.Println("Connection to database started")
	if err != nil {
		log.Fatalln("Error occured: " + err.Error())
		return
	}
	log.Println("Success")
	db.LogMode(true)
	db.AutoMigrate(&model.User{})
	return
}

//Login is used to verify if the user currently exist in the database using password and username
func Login(username string, password string) (model.User) {
	var user model.User
	GetDatabase().Where("username = ? AND password = ?", username, password).First(&user)
	return user
}

//GetUser is used to get user information from his username
func GetUser(username string) (model.User) {
	var user model.User
	GetDatabase().Where("username = ?", username).First(&user)
	return user
}

//InsertUser is used to insert a new user in database. Returns false if an error occurs.
func InsertUser(u model.User) bool {
		if GetDatabase().Create(&u).Error == nil {
			return true
		}
		return false
}

//Close is used to close the database instance
func Close() {
	if db != nil {
		db.Close()
	}
}