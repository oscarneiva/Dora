package main

import (
	"Dora/internal/pdfletter"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/lettergen", func(w http.ResponseWriter, r *http.Request) {
		arr, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return
		}
		letter := pdfletter.Letter{}
		err = json.Unmarshal(arr, &letter)
		if err != nil{
			return
		}
		pdfletter.LetterGen(letter)
	})
	http.ListenAndServe(":3000", r)
}