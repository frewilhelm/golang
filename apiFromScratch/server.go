package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

type ProgLang struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Fun        string `json:"fun"`
	Complexity string `json:"complexity"`
}

type progLangHandlers struct {
	sync.Mutex
	store map[string]ProgLang
}

func (h *progLangHandlers) progLangs(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.get(w, r)
		return
	case "POST":
		h.post(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}
}

func (h *progLangHandlers) get(w http.ResponseWriter, r *http.Request) {
	progLangs := make([]ProgLang, len(h.store))

	h.Lock()
	i := 0
	for _, progLang := range h.store {
		progLangs[i] = progLang
		i++
	}
	h.Unlock()

	jsonBytes, err := json.Marshal(progLangs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h *progLangHandlers) getRandomProgLang(w http.ResponseWriter, r *http.Request) {
	ids := make([]string, len(h.store))
	h.Lock()
	i := 0
	for id := range h.store {
		ids[i] = id
		i++
	}
	defer h.Unlock()

	var target string
	if len(ids) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if len(ids) == 1 {
		target = ids[0]
	} else {
		rand.Seed(time.Now().UnixNano())
		target = ids[rand.Intn(len(ids))]
	}

	w.Header().Add("location", fmt.Sprintf("/progLangs/%s", target))
	w.WriteHeader(http.StatusFound)
}

func (h *progLangHandlers) getProgLang(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.String(), "/")
	if len(parts) != 3 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if parts[2] == "random" {
		h.getRandomProgLang(w, r)
		return
	}

	h.Lock()
	progLang, ok := h.store[parts[2]]
	h.Unlock()
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	jsonBytes, err := json.Marshal(progLang)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h *progLangHandlers) post(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	ct := r.Header.Get("content-type")
	if ct != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte(fmt.Sprintf("need content-type 'application/json', but got '%s'", ct)))
		return
	}

	var progLang ProgLang
	err = json.Unmarshal(bodyBytes, &progLang)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	progLang.ID = fmt.Sprintf("%d", time.Now().UnixNano())

	h.Lock()
	h.store[progLang.ID] = progLang
	defer h.Unlock()
}

func newProgLangHandlers() *progLangHandlers {
	return &progLangHandlers{
		store: map[string]ProgLang{},
	}
}

type adminPortal struct {
	password string
}

func newAdminPortal() *adminPortal {
	password := os.Getenv("ADMIN_PASSWORD")
	if password == "" {
		panic("required env var ADMIN_PASSWORD not set")
	}

	return &adminPortal{password: password}
}

func (a adminPortal) handler(w http.ResponseWriter, r *http.Request) {
	user, pass, ok := r.BasicAuth()
	if !ok || user != "admin" || pass != a.password {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("401 - unauthorized"))
		return
	}

	w.Write([]byte("<html><h1>Super secret admin portal</1></html>"))
}

func main() {
	admin := newAdminPortal()
	progLangHandlers := newProgLangHandlers()
	http.HandleFunc("/progLangs", progLangHandlers.progLangs)
	http.HandleFunc("/progLangs/", progLangHandlers.getProgLang)
	http.HandleFunc("/admin", admin.handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
