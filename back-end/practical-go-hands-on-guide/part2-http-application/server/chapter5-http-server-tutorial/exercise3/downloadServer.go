package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func downloadLocalFileHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	fileName := params.Get("filename")
	if len(fileName) == 0 {
		http.Error(w, "filename query parameter required", http.StatusBadRequest)
		return
	}

	fileName = filepath.Base(fileName)
	file, err := os.Open(fileName)
	if err != nil {
		http.Error(w, "no such file", http.StatusNotFound)
		return
	}
	defer file.Close()

	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		log.Printf("Error for read file(%s): %#v\n", fileName, err)
		http.Error(w, "error when reading file", http.StatusInternalServerError)
		return
	}

	file.Seek(0, 0)
	contentType := http.DetectContentType(buffer)

	log.Printf("File %s has contentType %s\n", fileName, contentType)
	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, fileName))

	io.Copy(w, file)
}

func main() {
	listenPort := ":7166"
	mux := http.NewServeMux()
	mux.HandleFunc("/download", downloadLocalFileHandler)

	err := http.ListenAndServe(listenPort, mux)
	if err != nil {
		log.Fatalf("Server could not start on %s.\nError:%#v\n", listenPort, err)
	}
}
