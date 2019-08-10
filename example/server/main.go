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
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Internal Server Error")
			return
		}
		var buf bytes.Buffer
		err = png.Encode(&buf, img)
		if err != nil {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Internal Server Error")
			return
		}
		w.WriteHeader(http.StatusOK)
		if _, err := io.Copy(w, &buf); err != nil {
			log.Fatal(err)
		}
	})
	port := "8080"
	fmt.Printf("listening on http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
