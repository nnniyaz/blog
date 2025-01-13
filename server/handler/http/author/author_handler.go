package authorHandler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/nnniyaz/blog/handler/http/response"
	"github.com/nnniyaz/blog/pkg/logger"
	authorService "github.com/nnniyaz/blog/service/author"
	"net/http"
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
		response.NewError(hd.logger, w, r, err)
		return
	}
	err := hd.service.Create(r.Context(), in.FirstName, in.LastName, in.AvatarUri)
	if err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}
	response.NewSuccess(hd.logger, w, r, nil)
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
		response.NewError(hd.logger, w, r, err)
		return
	}
	err := hd.service.Update(r.Context(), id, in.FirstName, in.LastName, in.AvatarUri)
	if err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}
	response.NewSuccess(hd.logger, w, r, nil)
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
		response.NewError(hd.logger, w, r, err)
		return
	}
	response.NewSuccess(hd.logger, w, r, nil)
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
		response.NewError(hd.logger, w, r, err)
		return
	}
	response.NewSuccess(hd.logger, w, r, nil)
}

// GetAuthor godoc
//
//	@Summary		Get author
//	@Description	This can only be done by the logged-in user.
//	@Tags			Author
//	@Accept			json
//	@Produce		json
//	@Param			id			path		string				true	"Author ID"
//	@Success		200						{object}	response.Success
//	@Failure		default					{object}	response.Error
//	@Router			/author/{id} [get]
func (hd *HttpDelivery) GetAuthor(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	author, err := hd.service.FindById(r.Context(), id)
	if err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}
	response.NewSuccess(hd.logger, w, r, author)
}

// GetAllAuthors godoc
//
//	@Summary		Get authors
//	@Description	This can only be done by the logged-in user.
//	@Tags			Author
//	@Accept			json
//	@Produce		json
//	@Success		200						{object}	response.Success
//	@Failure		default					{object}	response.Error
//	@Router			/author [get]
func (hd *HttpDelivery) GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	authors, err := hd.service.FindAll(r.Context())
	if err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}
	response.NewSuccess(hd.logger, w, r, authors)
}
