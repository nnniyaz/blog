package userHandler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/nnniyaz/blog/internal/domain/user"
	"github.com/nnniyaz/blog/internal/handlers/http/response"
	userService "github.com/nnniyaz/blog/internal/services/user"
	"github.com/nnniyaz/blog/pkg/logger"
	"net/http"
	"time"
)

type HttpDelivery struct {
	logger  logger.Logger
	service userService.UserService
}

func NewHttpDelivery(logger logger.Logger, service userService.UserService) *HttpDelivery {
	return &HttpDelivery{logger: logger, service: service}
}

type CreateUserIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// CreateUser godoc
//
// @Summary		Create user
// @Description	This can only be done by the logged-in user.
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			data		body		CreateUserIn		true	"Create User Structure"
// @Success		200		{object}	response.Success
// @Failure		default	{object}	response.Error
// @Router			/user [post]
func (hd *HttpDelivery) CreateUser(w http.ResponseWriter, r *http.Request) {
	var in CreateUserIn
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}
	err := hd.service.Create(r.Context(), in.Email, in.Password)
	if err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}
	response.NewSuccess(hd.logger, w, r, nil)
}

type UpdateUserIn struct {
	Email string `json:"email"`
}

// UpdateUserEmail godoc
//
// @Summary		Update user
// @Description	This can only be done by the logged-in user.
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			id			path		string				true	"User ID"
// @Param			data		body		UpdateUserIn		true	"Update User Structure"
// @Success		200		{object}	response.Success
// @Failure		default	{object}	response.Error
// @Router			/user/email/{id} [put]
func (hd *HttpDelivery) UpdateUserEmail(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var in UpdateUserIn
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}
	err := hd.service.UpdateEmail(r.Context(), id, in.Email)
	if err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}
	response.NewSuccess(hd.logger, w, r, nil)
}

type UpdateUserPasswordIn struct {
	Password string `json:"password"`
}

// UpdateUserPassword godoc
//
// @Summary			Update user
// @Description		This can only be done by the logged-in user.
// @Tags			User
// @Accept			json
// @Produce			json
// @Param			id			path		string				true	"User ID"
// @Param			data		body		UpdateUserIn		true	"Update User Structure"
// @Success			200			{object}	response.Success
// @Failure			default		{object}	response.Error
// @Router			/user/password/{id} [put]
func (hd *HttpDelivery) UpdateUserPassword(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var in UpdateUserIn
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}
	err := hd.service.UpdatePassword(r.Context(), id, in.Email)
	if err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}
	response.NewSuccess(hd.logger, w, r, nil)
}

// DeleteUser godoc
//
// @Summary			Delete user
// @Description		This can only be done by the logged-in user.
// @Tags			User
// @Accept			json
// @Produce			json
// @Param			id			path		string	true	"User ID"
// @Success			200			{object}	response.Success
// @Failure			default		{object}	response.Error
// @Router			/user/{id} [delete]
func (hd *HttpDelivery) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := hd.service.Delete(r.Context(), id)
	if err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}
	response.NewSuccess(hd.logger, w, r, nil)
}

// RestoreUser godoc
//
// @Summary			Restore user
// @Description		This can only be done by the logged-in user.
// @Tags			User
// @Accept			json
// @Produce			json
// @Param			id			path		string	true	"User ID"
// @Success			200			{object}	response.Success
// @Failure			default		{object}	response.Error
// @Router			/user/{id}/restore [put]
func (hd *HttpDelivery) RestoreUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := hd.service.Restore(r.Context(), id)
	if err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}
	response.NewSuccess(hd.logger, w, r, nil)
}

type GetUserOut struct {
	Id        string `json:"id"`
	Email     string `json:"email"`
	IsDeleted bool   `json:"is_deleted"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func newGetUserOut(u *user.User) *GetUserOut {
	return &GetUserOut{
		Id:        u.GetId().String(),
		Email:     u.GetEmail(),
		IsDeleted: u.GetIsDeleted(),
		CreatedAt: u.GetCreatedAt().Format(time.RFC3339),
		UpdatedAt: u.GetUpdatedAt().Format(time.RFC3339),
	}
}

// GetUser godoc
//
// @Summary			Get user
// @Description		This can only be done by the logged-in user.
// @Tags			User
// @Accept			json
// @Produce			json
// @Param			id			path		string	true	"User ID"
// @Success			200			{object}	GetUserOut{GetUserOut}
// @Failure			default		{object}	response.Error
// @Router			/user/{id} [get]
func (hd *HttpDelivery) GetUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	u, err := hd.service.FindById(r.Context(), id)
	if err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}
	response.NewSuccess(hd.logger, w, r, newGetUserOut(u))
}

type GetUsersOut struct {
	Users []*GetUserOut `json:"users"`
	Count int64         `json:"count"`
}

func newGetUsersOut(users []*user.User, count int64) *GetUsersOut {
	var usersOut []*GetUserOut
	for _, u := range users {
		usersOut = append(usersOut, newGetUserOut(u))
	}
	return &GetUsersOut{
		Users: usersOut,
		Count: count,
	}
}

// GetUsers godoc
//
// @Summary			Get users
// @Description		This can only be done by the logged-in user.
// @Tags			User
// @Accept			json
// @Produce			json
// @Param			limit		query		int		false	"Limit"
// @Param			offset		query		int		false	"Offset"
// @Param			isDeleted	query		bool	false	"Is Deleted"
// @Param			search		query		string	false	"Search"
// @Success			200			{object}	GetUsersOut{GetUsersOut}
// @Failure			default		{object}	response.Error
// @Router			/user [get]
func (hd *HttpDelivery) GetUsers(w http.ResponseWriter, r *http.Request) {
	offset := r.Context().Value("offset").(int64)
	limit := r.Context().Value("limit").(int64)
	isDeleted := r.Context().Value("is_deleted").(bool)
	search := r.Context().Value("search").(string)
	users, count, err := hd.service.FindAll(r.Context(), offset, limit, isDeleted, search)
	if err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}
	response.NewSuccess(hd.logger, w, r, newGetUsersOut(users, count))
}
