package controllers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	// "golang-mini-project/views"
	"golang-mini-project/config"
	"golang-mini-project/models"
	"log"
	"net/http"
	"text/template"
)

type NoteControllers struct{}

func (controller *NoteControllers) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db, err := config.ConnectionDatabase()
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
	db.Order("id desc").Find(&notes)

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
	_, err := config.ConnectionDatabase()
	if err != nil {
		panic(err.Error())
	}

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

func (controller *NoteControllers) Edit(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	db, err := config.ConnectionDatabase()
	if err != nil {
		panic(err.Error())
	}

	files := []string{
		"./views/base.html",
		"./views/edit.html",
	}

	htmlTemplate, err := template.ParseFiles(files...)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	var note models.Note
	db.Where("ID= ?", params.ByName("id")).Find(&note)

	datas := map[string]interface{}{
		"Note": note,
		"ID":   params.ByName("id"),
	}

	err = htmlTemplate.ExecuteTemplate(w, "base", datas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err.Error())
	}

}

func (controller *NoteControllers) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	db, err := config.ConnectionDatabase()
	if err != nil {
		panic(err.Error())
	}

	// noteID := params.ByName("id")
	var note models.Note
	// db.Where("ID = ?", noteID).First(&note)
	db.First(&note, params.ByName("id"))

	note.Assignee = r.FormValue("assignee")
	note.Date = r.FormValue("deadline")
	note.Content = r.FormValue("content")

	db.Save(&note)

	http.Redirect(w, r, "/", http.StatusFound)

}

func (controller *NoteControllers) Store(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db, err := config.ConnectionDatabase()
	if err != nil {
		panic(err.Error())
	}

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
		// return
	} else {
		fmt.Println("Gagal")
		log.Println(result.Error)
		fmt.Println(result.Error)
		// return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func (controller *NoteControllers) Done(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	db, err := config.ConnectionDatabase()
	if err != nil {
		panic(err.Error())
	}

	var note models.Note
	db.Find(&note, params.ByName("id"))
	note.IsDone = !note.IsDone
	db.Save(&note)

	http.Redirect(w, r, "/", http.StatusFound)

}

func (controller *NoteControllers) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	db, err := config.ConnectionDatabase()
	if err != nil {
		panic(err.Error())
	}

	var note models.Note
	db.Delete(&note, params.ByName("id"))

	http.Redirect(w, r, "/", http.StatusFound)

}
