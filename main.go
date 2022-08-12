package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"golang-mini-project/config"
	"golang-mini-project/controllers"
	"golang-mini-project/models"
	"log"
	"net/http"
	"os"
)

func main() {
	db, err := config.ConnectionDatabase()

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
	router.POST("/create", noteControllers.Store)
	router.GET("/edit/:id", noteControllers.Edit)
	router.POST("/edit/:id", noteControllers.Update)
	router.POST("/done/:id", noteControllers.Done)
	router.POST("/delete/:id", noteControllers.Delete)

	// port := ":1234"
	port := os.Getenv("PORT")
	if port == "" {
		port = ":9000" // Default port if not specified
	}
	fmt.Println("Aplikasi jalan port " + port)

	// fmt.Println("aman boss")
	log.Fatal(http.ListenAndServe(port, router))

}
