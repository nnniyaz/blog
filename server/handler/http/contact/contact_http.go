package contactHandler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/nnniyaz/blog/domain/contact"
	"github.com/nnniyaz/blog/handler/http/response"
	"github.com/nnniyaz/blog/pkg/logger"
	contactService "github.com/nnniyaz/blog/service/contact"
	"net/http"
	"time"
)

type HttpDelivery struct {
	logger  logger.Logger
	service contactService.ContactService
}

func NewHttpDelivery(l logger.Logger, s contactService.ContactService) *HttpDelivery {
	return &HttpDelivery{logger: l, service: s}
}

type CreateContactIn struct {
	Label string `json:"label"`
	Link  string `json:"link"`
}

// CreateContact godoc
//
//	@Summary		Create contact
//	@Description	This can only be done by the logged-in user.
//	@Tags			Contacts
//	@Accept			json
//	@Produce		json
//	@Param			data		body		CreateContactIn		true	"Create Contact Structure"
//	@Success		200						{object}	response.Success
//	@Failure		default					{object}	response.Error
//	@Router			/contact [post]
func (hd *HttpDelivery) CreateContact(w http.ResponseWriter, r *http.Request) {
	var in CreateContactIn
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}
	err := hd.service.Create(r.Context(), in.Label, in.Link)
	if err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}
	response.NewSuccess(hd.logger, w, r, nil)
}

type UpdateContactIn struct {
	Label string `json:"label"`
	Link  string `json:"link"`
}

// UpdateContact godoc
//
//	@Summary		Update contact
//	@Description	This can only be done by the logged-in user.
//	@Tags			Contacts
//	@Accept			json
//	@Produce		json
//	@Param			id			path		string				true		"Contact ID"
//	@Param			data		body		UpdateContactIn		true		"Update Contact Structure"
//	@Success		200						{object}	response.Success
//	@Failure		default					{object}	response.Error
//	@Router			/contact/{id} [put]
func (hd *HttpDelivery) UpdateContact(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var in UpdateContactIn
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}

	err := hd.service.Update(r.Context(), id, in.Label, in.Link)
	if err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}
	response.NewSuccess(hd.logger, w, r, nil)
}

// DeleteContact godoc
//
//	@Summary		Delete contact
//	@Description	This can only be done by the logged-in user.
//	@Tags			Contacts
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string		true	"Contact ID"
//	@Success		200		{object}	response.Success
//	@Failure		default	{object}	response.Error
//	@Router			/contact/{id} [delete]
func (hd *HttpDelivery) DeleteContact(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := hd.service.Delete(r.Context(), id)
	if err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}
	response.NewSuccess(hd.logger, w, r, nil)
}

// RestoreContact godoc
//
//	@Summary		Restore contact
//	@Description	This can only be done by the logged-in user.
//	@Tags			Contacts
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string		true	"Contact ID"
//	@Success		200		{object}	response.Success
//	@Failure		default	{object}	response.Error
//	@Router			/contact/restore/{id} [put]
func (hd *HttpDelivery) RestoreContact(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := hd.service.Restore(r.Context(), id)
	if err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}
	response.NewSuccess(hd.logger, w, r, nil)
}

type FindContactByIdOut struct {
	Id        string `json:"id"`
	Label     string `json:"label"`
	Link      string `json:"link"`
	IsDeleted bool   `json:"is_deleted"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func newFindContactByIdOut(contact *contact.Contact) *FindContactByIdOut {
	return &FindContactByIdOut{
		Id:        contact.GetId().String(),
		Label:     contact.GetLabel(),
		Link:      contact.GetLink(),
		IsDeleted: contact.GetIsDeleted(),
		CreatedAt: contact.GetCreatedAt().Format(time.RFC3339),
		UpdatedAt: contact.GetUpdatedAt().Format(time.RFC3339),
	}
}

// GetContact godoc
//
//	@Summary		Get contact by ID
//	@Description	This can only be done by the logged-in user.
//	@Tags			Contacts
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string		true	"Contact ID"
//	@Success		200		{object}	response.Success
//	@Failure		default	{object}	response.Error
//	@Router			/contact/{id} [get]
func (hd *HttpDelivery) GetContact(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	contact, err := hd.service.FindById(r.Context(), id)
	if err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}
	response.NewSuccess(hd.logger, w, r, newFindContactByIdOut(contact))
}

type FindAllContactsOut struct {
	Contacts []*FindContactByIdOut `json:"contacts"`
	Count    int64                 `json:"count"`
}

func newFindAllContactsOut(contacts []*contact.Contact, count int64) *FindAllContactsOut {
	var responseContacts []*FindContactByIdOut
	for _, contact := range contacts {
		responseContacts = append(responseContacts, newFindContactByIdOut(contact))
	}
	return &FindAllContactsOut{Contacts: responseContacts, Count: count}
}

// GetAllContacts godoc
//
//	@Summary		Get all contacts
//	@Description	This can only be done by the logged-in user.
//	@Tags			Contacts
//	@Accept			json
//	@Produce		json
//	@Param			offset		query		int		false	"Offset"
//	@Param			limit		query		int		false	"Limit"
//	@Param			isDeleted	query		bool	false	"Is Deleted"
//	@Param			search		query		string	false	"Search"
//	@Success		200		{object}	response.Success
//	@Failure		default	{object}	response.Error
//	@Router			/contact [get]
func (hd *HttpDelivery) GetAllContacts(w http.ResponseWriter, r *http.Request) {
	offset := r.Context().Value("offset").(int64)
	limit := r.Context().Value("limit").(int64)
	isDeleted := r.Context().Value("is_deleted").(bool)
	search := r.Context().Value("search").(string)
	contacts, total, err := hd.service.FindAll(r.Context(), offset, limit, isDeleted, search)
	if err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}
	response.NewSuccess(hd.logger, w, r, newFindAllContactsOut(contacts, total))
}
