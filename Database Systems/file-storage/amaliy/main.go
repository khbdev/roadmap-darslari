package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)


var (
    rdb = redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
    })
    ctx = context.Background()
)



func uploadHandler(w http.ResponseWriter, r *http.Request){
	file, header, err := r.FormFile("file")
	if err != nil {
		    http.Error(w, "File not found in request", http.StatusBadRequest)
        return
	}
	defer file.Close()

	fileId := uuid.New().String()
	filename := fileId + filepath.Ext(header.Filename)
	filepath := filepath.Join("uploads", filename)

	out , err := os.Create(filepath)
	if  err != nil {
		     http.Error(w, "Unable to create file", http.StatusInternalServerError)
        return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		       http.Error(w, "Unable to save file", http.StatusInternalServerError)
        return
	}
	err = rdb.Set(ctx, fileId, filepath, 24*time.Hour).Err()
	if err != nil {
		   http.Error(w, "Unable to save mapping in Redis", http.StatusInternalServerError)
        return
	}
	 fmt.Fprintf(w, "File uploaded successfully! File ID: %s\n", fileId)
}


func dowlandHandler(w http.ResponseWriter, r *http.Request){
	    fileID := r.URL.Query().Get("id")
    if fileID == "" {
        http.Error(w, "File ID missing", http.StatusBadRequest)
        return
    }

	    path, err := rdb.Get(ctx, fileID).Result()
    if err == redis.Nil {
        http.Error(w, "File not found", http.StatusNotFound)
        return
    } else if err != nil {
        http.Error(w, "Redis error", http.StatusInternalServerError)
        return
    }

	http.ServeFile(w, r, path)
}

func main(){
	  http.HandleFunc("/upload", uploadHandler)
    http.HandleFunc("/download", dowlandHandler)

    fmt.Println("Server running at http://localhost:8081")
    log.Fatal(http.ListenAndServe(":8081", nil))
}