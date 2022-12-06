package main

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type ImgChunk struct {
	ID   int    `json:"id"`
	Data string `json:"data"`
}

func main() {
	imgBytes, err := os.ReadFile("./static/img/arcane_jinx3.jpg")
	imgSrc := base64.StdEncoding.EncodeToString(imgBytes)
	if err != nil {
		log.Fatalln(err)
	}
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/thumbnails", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Transfer-Encoding", "chunked")
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		flusher, _ := w.(http.Flusher)

		for i := 0; i < 10; i++ {
			imgChunk := ImgChunk{ID: i + 1, Data: imgSrc}
			// log.Printf("Chunk#%d %s\n\n", i, imgSrc)
			log.Printf("Chunk#%d \n", i)
			resp, err := json.Marshal(imgChunk)
			if err != nil {
				log.Println(err)
				break
			}
			w.Write(resp)
			flusher.Flush()
			time.Sleep(100 * time.Millisecond)
		}
		log.Println("Finished")
	})
	log.Println("Server running on port 3000")
	http.ListenAndServe(":3000", nil)
}
