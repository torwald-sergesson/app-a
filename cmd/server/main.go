package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/torwald-sergesson/app-a/pkg/dto"
)

type baseHandler struct{}

func (h *baseHandler) Handle(obj interface{}, w http.ResponseWriter, r *http.Request) {
	buf, err := json.Marshal(obj)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("%s", err)))
		log.Printf("URL: %s; error: %s", r.URL.String(), err.Error())
		return
	}

	log.Printf("URL: %s; ok", r.URL.String())
	w.Write(buf)
}

type meHandler struct {
	baseHandler
}

func (h *meHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := dto.User{
		ID:   1,
		Name: "Odin",
		Age:  45,
		Tags: []string{"hallo", "world", "new"},
	}
	h.Handle(&data, w, r)
}

type myGroupHandler struct {
	baseHandler
}

func (h *myGroupHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := dto.Group{
		ID:   1,
		Name: "Asgard",
	}
	h.Handle(&data, w, r)
}

func main() {
	http.Handle("/api/me", &meHandler{})
	http.Handle("/api/group/my", &myGroupHandler{})

	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatalf("error: %s", err)
	}
}
