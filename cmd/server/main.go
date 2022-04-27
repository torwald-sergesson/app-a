package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/torwald-sergesson/app-a/pkg/client/v2"
	"github.com/torwald-sergesson/app-a/pkg/dto/v2"
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

var grpAsgard = dto.Group{
	ID:   1,
	Name: "Asgard",
}

var usrOdin = dto.User{
	ID:    1,
	Name:  "Odin",
	Age:   45,
	Group: grpAsgard,
	Tags:  []string{"hallo", "world", "new"},
}

var meUserHandler = &baseHandler{
	newObject: func() interface{} { return usrOdin },
}

var myGroupHandler = &baseHandler{
	newObject: func() interface{} { return grpAsgard },
}

func syncWorker(period time.Duration) {
	cli := client.NewClient("localhost:8080", time.Second)
	ticker := time.NewTicker(period)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			usr, err := cli.Me()
			if err != nil {
				log.Printf("sync error: %s\n", err)
				continue
			}
			log.Printf("SYNC user: %+v\n", usr)
			log.Printf("SYNC tags: [%s]\n", strings.Join(usr.Tags, " "))
		}
	}
}

func main() {
	http.Handle("/api/me", meUserHandler)
	http.Handle("/api/group/my", myGroupHandler)

	go syncWorker(time.Second * 5)

	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatalf("error: %s", err)
	}
}
