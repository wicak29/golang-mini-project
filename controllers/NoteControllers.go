package controllers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	// "golang-mini-project/views"
	"golang-mini-project/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
	"text/template"
)

type NoteControllers struct{}

func (controller *NoteControllers) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// fmt.Fprint(w, "Welcome!\n")

	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	files := []string{
		"./views/base.html",
		"./views/index.html",
	}

	htmlTemplate, err := template.ParseFiles(files...)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var notes []models.Note
	db.Find(&notes)

	// datas := map[string]interface{}{}
	datas := map[string]interface{}{
		"Notes": notes,
	}

	// println(datas)
	err = htmlTemplate.ExecuteTemplate(w, "base", datas)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err.Error())
	}
}

func (controller *NoteControllers) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// fmt.Println(r)
	// return
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	if r.Method == "POST" {
		// fmt.Println(r.FormValue("content"))
		note := models.Note{
			Assignee: r.FormValue("assignee"),
			Content:  r.FormValue("content"),
			Date:     r.FormValue("deadline"),
		}
		fmt.Println(note)

		result := db.Create(&note)
		if result.Error != nil {
			log.Println(result.Error)
			fmt.Println(result.Error)
			return
		} else {
			fmt.Println("Gagal")
			log.Println(result.Error)
			fmt.Println(result.Error)
			return
		}

		http.Redirect(w, r, "/", http.StatusFound)

	} else {
		files := []string{
			"./views/base.html",
			"./views/create.html",
		}

		htmlTemplate, err := template.ParseFiles(files...)

		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		datas := map[string]interface{}{}

		err = htmlTemplate.ExecuteTemplate(w, "base", datas)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println(err.Error())
		}
	}

}
