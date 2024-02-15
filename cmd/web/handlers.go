package main

import (
	"fmt"
	"group/internal"
	"group/internal/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		Error(w, http.StatusNotFound)
		return
	}
	if r.Method == "GET" {

		artistResp, err := internal.GetArtists()
		if err != nil {
			Error(w, http.StatusInternalServerError)
			return
		}
		relationResp, err := internal.GetRelation()
		if err != nil {
			Error(w, http.StatusInternalServerError)
			return
		}

		for i := range artistResp {
			artistResp[i].Related = relationResp.GetNormalMap(i)
		}

		tmpl, err := template.ParseFiles("./ui/html/home.html")
		if err != nil {
			log.Println(err)
			Error(w, http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, artistResp)
		if err != nil {
			log.Println(err)
			Error(w, http.StatusInternalServerError)
			return
		}
	} else {
		Error(w, http.StatusBadRequest)
		return
	}
}
func ArtistHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		artistResp, err := internal.GetArtists()
		if err != nil {
			Error(w, http.StatusInternalServerError)
			fmt.Println("open")
			return
		}
		relationResp, err := internal.GetRelation()
		if err != nil {
			fmt.Println(fmt.Errorf("get relation error: %w", err))
			Error(w, http.StatusInternalServerError)
			return
		}
		if !strings.HasPrefix(r.URL.Path, "/artist/") { // Функция strings.HasPrefix используется для проверки, начинается ли URL-путь (r.URL.Path) с "/artist/".
			log.Print("err art")

			Error(w, http.StatusNotFound)
			return
		}
		idstr := strings.TrimPrefix(r.URL.Path, "/artist/") // извлекаем индентификатора артиста из url пути, TrimPrefix позволяет нам удалить "/artist/" и получить только идентификатор
		if len(idstr) == 0 {
			Error(w, http.StatusNotFound)
			return
		}
		// idstr := r.FormValue("id")
		// if len(idstr) == 0 {
		// 	Error(w, http.StatusNotFound)
		// 	return
		// }
		// id, err := strconv.Atoi(idstr)
		// if err != nil || len(artistResp) < id {
		// 	Error(w, http.StatusNotFound)
		// 	return
		// }
		id, err := strconv.Atoi(idstr) // Преобразует идентификатор в целое число id.
		if err != nil || id < 1 || id > len(artistResp) {
			Error(w, http.StatusNotFound)
			return
		}
		tmpl, err := template.ParseFiles("./ui/html/artist.html") // загружаем html шаблон
		if err != nil {
			log.Print(err)
			Error(w, http.StatusInternalServerError)
			return
		}
		artistInfo := &models.ArtistInfo{
			Artist:  artistResp[id-1],                  //
			Related: relationResp.GetNormalMap(id - 1), // информация о конкретном артисте с учетом номера
		}
		err = tmpl.Execute(w, artistInfo)
		if err != nil {
			fmt.Println(err)
			Error(w, http.StatusInternalServerError)
			return
		}
	} else {
		Error(w, http.StatusBadRequest)
		return
	}
}

// func ArtistHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/artist" {
// 		Error(w, http.StatusNotFound)
// 		return
// 	}
// 	if r.Method == "GET" {
// 		artistResp, err := internal.GetArtists()
// 		if err != nil {
// 			Error(w, http.StatusInternalServerError)
// 			return
// 		}
// 		relationResp, err := internal.GetRelation()
// 		if err != nil {
// 			Error(w, http.StatusInternalServerError)
// 			return
// 		}

// 		idstr := r.FormValue("id")
// 		if len(idstr) == 0 {
// 			Error(w, http.StatusNotFound)
// 			return
// 		}
// 		id, err := strconv.Atoi(idstr)
// 		if err != nil || len(artistResp) < id {
// 			Error(w, http.StatusNotFound)
// 			return
// 		}

// 		tmpl, err := template.ParseFiles("./ui/html/artist.html")
// 		if err != nil {
// 			Error(w, http.StatusInternalServerError)
// 			return
// 		}

// 		artistInfo := &models.ArtistInfo{
// 			Artist:  artistResp[id-1],
// 			Related: relationResp.GetNormalMap(id - 1),
// 		}

// 		err = tmpl.Execute(w, artistInfo)
// 		if err != nil {
// 			Error(w, http.StatusInternalServerError)
// 			return
// 		}
// 	} else {
// 		Error(w, http.StatusBadRequest)
// 		return
// 	}
// }
