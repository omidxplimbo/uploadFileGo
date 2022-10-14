package main

import (
	"fmt"
	"log"
	"net/http"
	"toolkit"
)

func main() {
	mux := routes()
	log.Println("Starting server on port 8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("."))))
	mux.HandleFunc("/upload", uploadFiles)

	return mux
}

func uploadFiles(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	t := toolkit.Tools{
		MaxFileSize:      1024 * 1024 * 1024,
		AllowedFileTypes: []string{"image/png", "image/jpeg", "image/jpg"},
	}

	uploadFile, err := t.UploadFiles(r, "./uploads", true)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	out := ""
	for _, file := range uploadFile {
		out += fmt.Sprintf("Uploaded %s to the uploads folder , renamed to %s", file.OriginalFileName, file.NewFileName)
	}

	_, _ = w.Write([]byte(out))

}
