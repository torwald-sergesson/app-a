package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/torwald-sergesson/app-a/pkg/dto"
)

type meHandler struct{}

func (h *meHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := dto.User{
		ID:   1,
		Name: "Tor",
	}

	buf, err := json.Marshal(&data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("%s", err)))
		log.Printf("URL: %s; error: %s", r.URL.String(), err.Error())
		return
	}

	log.Printf("URL: %s; ok", r.URL.String())
	w.Write(buf)
}

func main() {
	http.Handle("/api/me", &meHandler{})

	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatalf("error: %s", err)
	}
}
