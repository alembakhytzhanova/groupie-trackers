package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	infoLog := log.New(os.Stdout, "\033[92mINFO\033[0m\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "\033[91mERROR\033[0m\t", log.Ldate|log.Ltime|log.Lshortfile)

	port := "8989"
	if len(os.Args) == 2 {
		_, err := strconv.Atoi(os.Args[1])
		if err != nil {
			errorLog.Println("Not correct port. Rewrite it")
			return
		}
		port = os.Args[1]
	}

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./ui/static"))

	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/", HomePageHandler)
	mux.HandleFunc("/artist/", ArtistHandler)

	infoLog.Printf("Starting server on http://localhost:%s", port)
	err := http.ListenAndServe(":"+port, mux)
	errorLog.Fatal(err)
}

// func main() {
// 	addr := flag.String("addr", "1111", "HTTP network address")
// 	flag.Parse()

// 	mux := http.NewServeMux()

// 	fileServer := http.FileServer(http.Dir("./ui/static"))
// 	mux.Handle("/static/", http.StripPrefix("/static/", fileServer)) //удаляет префикс из url , чтобы запросы к статическим файлам не включали этот прифекс

// 	mux.HandleFunc("/", HomePageHandler)
// 	mux.HandleFunc("/artist/", ArtistHandler)

// 	log.Printf("Server is listening... http://localhost:%s", *addr)
// 	err := http.ListenAndServe(":"+*addr, mux)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
