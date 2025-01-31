package bookHandler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/nnniyaz/blog/internal/domain/book"
	"github.com/nnniyaz/blog/internal/handlers/http/response"
	"github.com/nnniyaz/blog/internal/services/book"
	"github.com/nnniyaz/blog/pkg/core"
	"github.com/nnniyaz/blog/pkg/logger"
	"net/http"
	"time"
)

type HttpDelivery struct {
	logger  logger.Logger
	service bookService.BookService
}

func NewHttpDelivery(l logger.Logger, s bookService.BookService) *HttpDelivery {
	return &HttpDelivery{logger: l, service: s}
}

type CreateBookIn struct {
	Title       core.MlString `json:"title"`
	Description core.MlString `json:"description"`
	Author      core.MlString `json:"author"`
	CoverUri    string        `json:"coverUri"`
	EBookUri    string        `json:"eBookUri"`
}

// CreateBook godoc
//
//	@Summary		Create book
//	@Description	This can only be done by the logged-in user.
//	@Tags			Book
//	@Accept			json
//	@Produce		json
//	@Param			data	body		CreateBookIn	true	"Create Book Structure"
//	@Success		200		{object}	response.Success
//	@Failure		default	{object}	response.Error
//	@Router			/book [post]
func (hd *HttpDelivery) CreateBook(w http.ResponseWriter, r *http.Request) {
	var in CreateBookIn
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}
	err := hd.service.Create(r.Context(), in.Title, in.Description, in.Author, in.CoverUri, in.EBookUri)
	if err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}
	response.NewSuccess(hd.logger, w, r, nil)
}

type UpdateBookIn struct {
	Title       core.MlString `json:"title"`
	Description core.MlString `json:"description"`
	Author      core.MlString `json:"author"`
	CoverUri    string        `json:"coverUri"`
	EBookUri    string        `json:"eBookUri"`
}

// UpdateBook godoc
//
//	@Summary		Update book
//	@Description	This can only be done by the logged-in user.
//	@Tags			Book
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string			true	"Book ID"
//	@Param			data	body		UpdateBookIn	true	"Update Book Structure"
//	@Success		200		{object}	response.Success
//	@Failure		default	{object}	response.Error
//	@Router			/book/{id} [put]
func (hd *HttpDelivery) UpdateBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var in UpdateBookIn
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}
	err := hd.service.Update(r.Context(), id, in.Title, in.Description, in.Author, in.CoverUri, in.EBookUri)
	if err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}
	response.NewSuccess(hd.logger, w, r, nil)
}

// DeleteBook godoc
//
//	@Summary		Delete book
//	@Description	This can only be done by the logged-in user.
//	@Tags			Book
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string	true	"Book ID"
//	@Success		200		{object}	response.Success
//	@Failure		default	{object}	response.Error
//	@Router			/book/{id} [delete]
func (hd *HttpDelivery) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := hd.service.Delete(r.Context(), id)
	if err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}
	response.NewSuccess(hd.logger, w, r, nil)
}

// RestoreBook godoc
//
//	@Summary		Restore book
//	@Description	This can only be done by the logged-in user.
//	@Tags			Book
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string	true	"Book ID"
//	@Success		200		{object}	response.Success
//	@Failure		default	{object}	response.Error
//	@Router			/book/restore/{id} [put]
func (hd *HttpDelivery) RestoreBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := hd.service.Restore(r.Context(), id)
	if err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}
	response.NewSuccess(hd.logger, w, r, nil)
}

type GetBookOut struct {
	Id          string        `json:"id"`
	Title       core.MlString `json:"title"`
	Description core.MlString `json:"description"`
	Author      core.MlString `json:"author"`
	CoverUri    string        `json:"coverUri"`
	EBookUri    string        `json:"eBookUri"`
	IsDeleted   bool          `json:"isDeleted"`
	CreatedAt   string        `json:"createdAt"`
	UpdatedAt   string        `json:"updatedAt"`
}

func newGetBookOut(book *book.Book) *GetBookOut {
	return &GetBookOut{
		Id:          book.GetId().String(),
		Title:       book.GetTitle(),
		Description: book.GetDescription(),
		Author:      book.GetAuthor(),
		CoverUri:    book.GetCoverUri(),
		EBookUri:    book.GetEBookUri(),
		IsDeleted:   book.GetIsDeleted(),
		CreatedAt:   book.GetCreatedAt().Format(time.RFC3339),
		UpdatedAt:   book.GetUpdatedAt().Format(time.RFC3339),
	}
}

// GetBook godoc
//
//	@Summary		Get book
//	@Description	This can only be done by the logged-in user.
//	@Tags			Book
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string	true	"Book ID"
//	@Success		200		{object}	response.Success{GetBookOut}
//	@Failure		default	{object}	response.Error
//	@Router			/book/{id} [get]
func (hd *HttpDelivery) GetBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	book, err := hd.service.FindById(r.Context(), id)
	if err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}
	response.NewSuccess(hd.logger, w, r, newGetBookOut(book))
}

type GetAllBooksOut struct {
	Books []*GetBookOut `json:"books"`
	Count int64         `json:"count"`
}

func newGetAllBooksOut(books []*book.Book, count int64) *GetAllBooksOut {
	var responseBooks []*GetBookOut
	for _, book := range books {
		responseBooks = append(responseBooks, newGetBookOut(book))
	}
	return &GetAllBooksOut{Books: responseBooks, Count: count}
}

// GetAllBooks godoc
//
//	@Summary		Get books by filters
//	@Description	This can only be done by the logged-in user.
//	@Tags			Book
//	@Accept			json
//	@Produce		json
//	@Param			offset		query		int		false	"Offset"
//	@Param			limit		query		int		false	"Limit"
//	@Param			is_deleted	query		bool	false	"Is deleted"
//	@Param			search		query		string	false	"Search Title, Description, Author"
//	@Success		200			{object}	response.Success{GetAllBooksOut}
//	@Failure		default		{object}	response.Error
//	@Router			/book [get]
func (hd *HttpDelivery) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	offset := r.Context().Value("offset").(int64)
	limit := r.Context().Value("limit").(int64)
	isDeleted := r.Context().Value("is_deleted").(bool)
	search := r.Context().Value("search").(string)
	books, count, err := hd.service.FindAll(r.Context(), offset, limit, isDeleted, search)
	if err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}
	response.NewSuccess(hd.logger, w, r, newGetAllBooksOut(books, count))
}
