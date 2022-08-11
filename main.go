package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"golang-mini-project/controllers"
	"golang-mini-project/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})

	if err != nil {
		panic("Program error bos")
	}

	err = db.AutoMigrate(&models.Note{})
	if err != nil {
		panic(err.Error())
	}

	noteControllers := &controllers.NoteControllers{}

	router := httprouter.New()

	router.GET("/", noteControllers.Index)
	router.GET("/create", noteControllers.Create)
	router.POST("/create", noteControllers.Create)

	port := ":1234"
	fmt.Println("Aplikasi jalan di http://localhost:1234")

	// fmt.Println("aman boss")
	log.Fatal(http.ListenAndServe(port, router))

}
