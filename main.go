package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)




func InitDb(){
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}

func main() {
	router:=InitRouter()
	http.ListenAndServe(":8080", router)
}
