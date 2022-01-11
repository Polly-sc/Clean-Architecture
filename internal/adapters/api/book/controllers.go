package adapters

import (
	"Clean-Architecture/internal/adapters"
	"Clean-Architecture/internal/book"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const (
	bookURL  = "book/:book_id"
	booksURL = "/books"
)

type handler struct {
	bookService book.Service
}

func NewHandler(service book.Service) adapters.Handler {
	return &handler{bookService: service}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(booksURL, h.GetAllBooks)
}

func (h *handler) GetAllBooks(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("books"))
	w.WriteHeader(http.StatusOK)
}
