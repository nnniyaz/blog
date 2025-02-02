package uploadHandler

import (
	"github.com/nnniyaz/blog/internal/handlers/http/response"
	uploadService "github.com/nnniyaz/blog/internal/services/upload"
	"github.com/nnniyaz/blog/pkg/logger"
	"net/http"
)

type HttpDelivery struct {
	logger  logger.Logger
	service uploadService.UploadService
}

func NewHttpDelivery(l logger.Logger, s uploadService.UploadService) *HttpDelivery {
	return &HttpDelivery{logger: l, service: s}
}

type UploadOut struct {
	FileName string `json:"filename"`
}

// UploadAuthor godoc
//
//	@Summary		Upload author-avatar
//	@Description	This can only be done by the logged-in user.
//	@Tags			Upload
//	@Accept			json
//	@Produce		json
//	@Param			data	formData	file	true	"file to upload"
//	@Success		200		{object}	response.Success
//	@Failure		default	{object}	response.Error
//	@Router			/upload/author-avatar [post]
func (hd *HttpDelivery) UploadAuthor(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(5 << 20)
	if err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}
	defer file.Close()

	filename, err := hd.service.UploadImage("author-avatar", file, header)
	if err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}
	response.NewSuccess(hd.logger, w, r, UploadOut{FileName: filename})
}

// UploadProject godoc
//
//	@Summary		Upload project
//	@Description	This can only be done by the logged-in user.
//	@Tags			Upload
//	@Accept			json
//	@Produce		json
//	@Param			data	formData	file	true	"file to upload"
//	@Success		200		{object}	response.Success
//	@Failure		default	{object}	response.Error
//	@Router			/upload/project [post]
func (hd *HttpDelivery) UploadProject(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(5 << 20)
	if err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}
	defer file.Close()

	filename, err := hd.service.UploadImage("project", file, header)
	if err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}
	response.NewSuccess(hd.logger, w, r, UploadOut{FileName: filename})
}

// UploadArticle godoc
//
//	@Summary		Upload article
//	@Description	This can only be done by the logged-in user.
//	@Tags			Upload
//	@Accept			json
//	@Produce		json
//	@Param			data	formData	file	true	"file to upload"
//	@Success		200		{object}	response.Success
//	@Failure		default	{object}	response.Error
//	@Router			/upload/article [post]
func (hd *HttpDelivery) UploadArticle(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(5 << 20)
	if err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}
	defer file.Close()

	filename, err := hd.service.UploadImage("article", file, header)
	if err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}
	response.NewSuccess(hd.logger, w, r, UploadOut{FileName: filename})
}

// UploadBook godoc
//
//	@Summary		Upload book
//	@Description	This can only be done by the logged-in user.
//	@Tags			Upload
//	@Accept			json
//	@Produce		json
//	@Param			data	formData	file	true	"file to upload"
//	@Success		200		{object}	response.Success
//	@Failure		default	{object}	response.Error
//	@Router			/upload/book [post]
func (hd *HttpDelivery) UploadBook(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(5 << 20)
	if err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}
	defer file.Close()

	filename, err := hd.service.UploadImage("book", file, header)
	if err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}
	response.NewSuccess(hd.logger, w, r, UploadOut{FileName: filename})
}
