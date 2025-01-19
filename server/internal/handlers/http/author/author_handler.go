package authorHandler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/nnniyaz/blog/internal/domain/author"
	response2 "github.com/nnniyaz/blog/internal/handlers/http/response"
	"github.com/nnniyaz/blog/internal/services/author"
	"github.com/nnniyaz/blog/pkg/logger"
	"net/http"
	"time"
)

type HttpDelivery struct {
	logger  logger.Logger
	service authorService.AuthorService
}

func NewHttpDelivery(l logger.Logger, s authorService.AuthorService) *HttpDelivery {
	return &HttpDelivery{logger: l, service: s}
}

type CreateAuthorIn struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	AvatarUri string `json:"avatarUri"`
}

// CreateAuthor godoc
//
//	@Summary		Create author
//	@Description	This can only be done by the logged-in user.
//	@Tags			Author
//	@Accept			json
//	@Produce		json
//	@Param			data		body		CreateAuthorIn		true	"Create Author Structure"
//	@Success		200						{object}	response.Success
//	@Failure		default					{object}	response.Error
//	@Router			/author [post]
func (hd *HttpDelivery) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var in CreateAuthorIn
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		response2.NewError(hd.logger, w, r, err)
		return
	}
	err := hd.service.Create(r.Context(), in.FirstName, in.LastName, in.AvatarUri)
	if err != nil {
		response2.NewError(hd.logger, w, r, err)
		return
	}
	response2.NewSuccess(hd.logger, w, r, nil)
}

type UpdateAuthorIn struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	AvatarUri string `json:"avatarUri"`
}

// UpdateAuthor godoc
//
//	@Summary		Update author
//	@Description	This can only be done by the logged-in user.
//	@Tags			Author
//	@Accept			json
//	@Produce		json
//	@Param			id			path		string				true	"Author ID"
//	@Param			data		body		UpdateAuthorIn		true	"Update Author Structure"
//	@Success		200						{object}	response.Success
//	@Failure		default					{object}	response.Error
//	@Router			/author/{id} [put]
func (hd *HttpDelivery) UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var in UpdateAuthorIn
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		response2.NewError(hd.logger, w, r, err)
		return
	}
	err := hd.service.Update(r.Context(), id, in.FirstName, in.LastName, in.AvatarUri)
	if err != nil {
		response2.NewError(hd.logger, w, r, err)
		return
	}
	response2.NewSuccess(hd.logger, w, r, nil)
}

// DeleteAuthor godoc
//
//	@Summary		Delete author
//	@Description	This can only be done by the logged-in user.
//	@Tags			Author
//	@Accept			json
//	@Produce		json
//	@Param			id			path		string				true	"Author ID"
//	@Success		200						{object}	response.Success
//	@Failure		default					{object}	response.Error
//	@Router			/author/{id} [delete]
func (hd *HttpDelivery) DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := hd.service.Delete(r.Context(), id)
	if err != nil {
		response2.NewError(hd.logger, w, r, err)
		return
	}
	response2.NewSuccess(hd.logger, w, r, nil)
}

// RestoreAuthor godoc
//
//	@Summary		Restore author
//	@Description	This can only be done by the logged-in user.
//	@Tags			Author
//	@Accept			json
//	@Produce		json
//	@Param			id			path		string				true	"Author ID"
//	@Success		200						{object}	response.Success
//	@Failure		default					{object}	response.Error
//	@Router			/author/{id} [put]
func (hd *HttpDelivery) RestoreAuthor(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := hd.service.Restore(r.Context(), id)
	if err != nil {
		response2.NewError(hd.logger, w, r, err)
		return
	}
	response2.NewSuccess(hd.logger, w, r, nil)
}

type GetAuthorOut struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	AvatarUri string `json:"avatarUri"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	Version   int    `json:"version"`
}

func newGetAuthorOut(author *author.Author) *GetAuthorOut {
	return &GetAuthorOut{
		Id:        author.GetId().String(),
		FirstName: author.GetFirstName(),
		LastName:  author.GetLastName(),
		AvatarUri: author.GetAvatarUri(),
		CreatedAt: author.GetCreatedAt().Format(time.RFC3339),
		UpdatedAt: author.GetUpdatedAt().Format(time.RFC3339),
		Version:   author.GetVersion(),
	}
}

// GetAuthor godoc
//
//	@Summary		Get author
//	@Description	This can only be done by the logged-in user.
//	@Tags			Author
//	@Accept			json
//	@Produce		json
//	@Param			id			path		string				true	"Author ID"
//	@Success		200						{object}	response.Success{GetAuthorOut}
//	@Failure		default					{object}	response.Error
//	@Router			/author/{id} [get]
func (hd *HttpDelivery) GetAuthor(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	author, err := hd.service.FindById(r.Context(), id)
	if err != nil {
		response2.NewError(hd.logger, w, r, err)
		return
	}
	response2.NewSuccess(hd.logger, w, r, newGetAuthorOut(author))
}

type GetAuthorsOut struct {
	Authors []*GetAuthorOut `json:"authors"`
	Count   int64           `json:"count"`
}

func newGetAuthorsOut(authors []*author.Author, count int64) *GetAuthorsOut {
	var out []*GetAuthorOut
	for _, a := range authors {
		out = append(out, newGetAuthorOut(a))
	}
	return &GetAuthorsOut{
		Authors: out,
		Count:   count,
	}
}

// GetAllAuthors godoc
//
//	@Summary		Get authors
//	@Description	This can only be done by the logged-in user.
//	@Tags			Author
//	@Accept			json
//	@Produce		json
//	@Success		200						{object}	response.Success{GetAuthorsOut}
//	@Failure		default					{object}	response.Error
//	@Router			/author [get]
func (hd *HttpDelivery) GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	authors, count, err := hd.service.FindAll(r.Context())
	if err != nil {
		response2.NewError(hd.logger, w, r, err)
		return
	}
	response2.NewSuccess(hd.logger, w, r, newGetAuthorsOut(authors, count))
}
