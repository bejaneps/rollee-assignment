package transport

import (
	"errors"
	"log"
	"net/http"

	"github.com/bejaneps/rollee-assignment/test2/internal/storage"
	"github.com/gorilla/mux"
)

func handleServicePost(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		word := vars["word"]

		err := s.UpsertWord(word)
		if err != nil {
			log.Printf("handleServicePost: %v\n", err)

			http.Error(
				w,
				http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError,
			)

			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func handleServiceGet(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		word := vars["word"]

		mfw, err := s.GetMostFrequentWord(word)
		if errors.Is(err, storage.ErrWordNotFound) {
			http.Error(
				w,
				"null",
				http.StatusBadRequest,
			)

			return
		} else if err != nil {
			log.Printf("handleServiceGet: %v\n", err)

			http.Error(
				w,
				http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError,
			)

			return
		}

		w.Write([]byte(mfw))
	}
}

func handleNotFound() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Error(
			w,
			"Specified path doesn't exist or word is invalid.",
			http.StatusBadRequest,
		)
	}
}
