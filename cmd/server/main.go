package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/torwald-sergesson/app-a/pkg/dto"
)

type ObjectFactory func() interface{}

type baseHandler struct {
	newObject ObjectFactory
}

func (h *baseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	obj := h.newObject()
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

var meUserHandler = &baseHandler{
	newObject: func() interface{} {
		return dto.User{
			ID:   1,
			Name: "Odin",
			Age:  45,
			Tags: []string{"hallo", "world", "new"},
		}
	},
}

var myGroupHandler = &baseHandler{
	newObject: func() interface{} {
		return dto.Group{
			ID:   1,
			Name: "Asgard",
		}
	},
}

func main() {
	http.Handle("/api/me", meUserHandler)
	http.Handle("/api/group/my", myGroupHandler)

	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatalf("error: %s", err)
	}
}
