package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	title := "Jenkins X golang http example"

	from := ""
	if r.URL != nil {
		from = r.URL.String()
	}
	if from != "/favicon.ico" {
		log.Printf("title: %s\n", title)
	}

	fmt.Fprintf(w, "<h1 style='color:navy'>Cute Hello from:  "+title+" </h1><br><br><img src='https://ssl.gstatic.com/onebox/media/sports/photos/reuters/tag:reuters.com,2019:newsml_RC14EF376110:2026799446_768x432.jpg'>\n")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
