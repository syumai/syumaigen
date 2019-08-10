package main

import (
	"bytes"
	"fmt"
	"image/png"
	"io"
	"log"
	"net/http"

	"github.com/syumai/syumaigen"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png")
		img, err := syumaigen.GenerateImage(
			syumaigen.Pattern,
			syumaigen.GenerateRandomColorMap(),
			10,
		)
		if err != nil {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(500)
			fmt.Fprintf(w, "Internal Server Error")
			return
		}
		var buf bytes.Buffer
		err = png.Encode(&buf, img)
		if err != nil {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(500)
			fmt.Fprintf(w, "Internal Server Error")
			return
		}
		w.WriteHeader(200)
		if _, err := io.Copy(w, &buf); err != nil {
			log.Fatal(err)
		}
	})
	fmt.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
