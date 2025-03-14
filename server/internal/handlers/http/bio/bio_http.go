package bioHandler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/nnniyaz/blog/internal/domain/bio"
	"github.com/nnniyaz/blog/internal/handlers/http/response"
	"github.com/nnniyaz/blog/internal/services/bio"
	"github.com/nnniyaz/blog/pkg/core"
	"github.com/nnniyaz/blog/pkg/logger"
	"net/http"
	"time"
)

type HttpDelivery struct {
	logger  logger.Logger
	service bioService.BioService
}

func NewHttpDelivery(l logger.Logger, s bioService.BioService) *HttpDelivery {
	return &HttpDelivery{logger: l, service: s}
}

type CreateBioIn struct {
	Bio core.MlString `json:"bio"`
}

// CreateBio godoc
//
//	@Summary		Create bio
//	@Description	This can only be done by the logged-in user.
//	@Tags			Bio
//	@Accept			json
//	@Produce		json
//	@Param			data	body		CreateBioIn	true	"Create Bio Structure"
//	@Success		200		{object}	response.Success
//	@Failure		default	{object}	response.Error
//	@Router			/bio [post]
func (hd *HttpDelivery) CreateBio(w http.ResponseWriter, r *http.Request) {
	var in CreateBioIn
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}
	err := hd.service.Create(r.Context(), in.Bio)
	if err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}
	response.NewSuccess(hd.logger, w, r, nil)
}

type UpdateBioIn struct {
	Bio core.MlString `json:"bio"`
}

// UpdateBio godoc
//
//	@Summary		Update bio
//	@Description	This can only be done by the logged-in user.
//	@Tags			Bio
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string		true	"Bio ID"
//	@Param			data	body		UpdateBioIn	true	"Update Bio Structure"
//	@Success		200		{object}	response.Success
//	@Failure		default	{object}	response.Error
//	@Router			/bio/{id} [put]
func (hd *HttpDelivery) UpdateBio(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var in UpdateBioIn
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}
	err := hd.service.Update(r.Context(), id, in.Bio)
	if err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}
	response.NewSuccess(hd.logger, w, r, nil)
}

// DeleteBio godoc
//
//	@Summary		Delete bio
//	@Description	This can only be done by the logged-in user.
//	@Tags			Bio
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string	true	"Bio ID"
//	@Success		200		{object}	response.Success
//	@Failure		default	{object}	response.Error
//	@Router			/bio/{id} [delete]
func (hd *HttpDelivery) DeleteBio(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := hd.service.Delete(r.Context(), id)
	if err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}
	response.NewSuccess(hd.logger, w, r, nil)
}

// RestoreBio godoc
//
//	@Summary		Restore bio
//	@Description	This can only be done by the logged-in user.
//	@Tags			Bio
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string	true	"Bio ID"
//	@Success		200		{object}	response.Success
//	@Failure		default	{object}	response.Error
//	@Router			/bio/{id} [put]
func (hd *HttpDelivery) RestoreBio(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := hd.service.Restore(r.Context(), id)
	if err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}
	response.NewSuccess(hd.logger, w, r, nil)
}

type GetBioOut struct {
	Id        string        `json:"id"`
	Bio       core.MlString `json:"bio"`
	Active    bool          `json:"active"`
	CreatedAt string        `json:"createdAt"`
	UpdatedAt string        `json:"updatedAt"`
	Version   int           `json:"version"`
}

func newGetBioOut(bio *bio.Bio) *GetBioOut {
	return &GetBioOut{
		Id:        bio.GetId().String(),
		Bio:       bio.GetBio(),
		Active:    bio.GetActive(),
		CreatedAt: bio.GetCreatedAt().Format(time.RFC3339),
		UpdatedAt: bio.GetUpdatedAt().Format(time.RFC3339),
		Version:   bio.GetVersion(),
	}
}

// GetActiveBio godoc
//
//	@Summary		Get active bio
//	@Description	This can only be done by the logged-in user.
//	@Tags			Bio
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	response.Success{GetBioOut}
//	@Failure		default	{object}	response.Error
//	@Router			/bio/active [get]
func (hd *HttpDelivery) GetActiveBio(w http.ResponseWriter, r *http.Request) {
	bio, err := hd.service.FindByActive(r.Context())
	if err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}
	response.NewSuccess(hd.logger, w, r, newGetBioOut(bio))
}

// GetBio godoc
//
//	@Summary		Get bio
//	@Description	This can only be done by the logged-in user.
//	@Tags			Bio
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string	true	"Bio ID"
//	@Success		200		{object}	response.Success{GetBioOut}
//	@Failure		default	{object}	response.Error
//	@Router			/bio/{id} [get]
func (hd *HttpDelivery) GetBio(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	bio, err := hd.service.FindById(r.Context(), id)
	if err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}
	response.NewSuccess(hd.logger, w, r, newGetBioOut(bio))
}

type GetAllBiosOut struct {
	Bios  []*GetBioOut `json:"bios"`
	Count int64        `json:"count"`
}

func newGetAllBiosOut(bios []*bio.Bio, count int64) *GetAllBiosOut {
	var responseBios []*GetBioOut
	for _, bio := range bios {
		responseBios = append(responseBios, newGetBioOut(bio))
	}
	return &GetAllBiosOut{Bios: responseBios, Count: count}
}

// GetAllBios godoc
//
//	@Summary		Get bios
//	@Description	This can only be done by the logged-in user.
//	@Tags			Bio
//	@Accept			json
//	@Produce		json
//	@Param			offset	query		int	false	"Offset"
//	@Param			limit	query		int	false	"Limit"
//	@Success		200		{object}	response.Success{GetAllBiosOut}
//	@Failure		default	{object}	response.Error
//	@Router			/bio [get]
func (hd *HttpDelivery) GetAllBios(w http.ResponseWriter, r *http.Request) {
	offset := r.Context().Value("offset").(int64)
	limit := r.Context().Value("limit").(int64)
	bios, count, err := hd.service.FindAll(r.Context(), offset, limit)
	if err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}
	response.NewSuccess(hd.logger, w, r, newGetAllBiosOut(bios, count))
}
