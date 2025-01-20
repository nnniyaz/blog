package currentUserHandler

import (
	"encoding/json"
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

type GetCurrentUserOut struct {
	Id        string `json:"id"`
	Email     string `json:"email"`
	IsDeleted bool   `json:"isDeleted"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func newGetCurrentUserOut(user *user.User) GetCurrentUserOut {
	return GetCurrentUserOut{
		Id:        user.GetId().String(),
		Email:     user.GetEmail().String(),
		IsDeleted: user.GetIsDeleted(),
		CreatedAt: user.GetCreatedAt().Format(time.RFC3339),
		UpdatedAt: user.GetUpdatedAt().Format(time.RFC3339),
	}
}

// GetCurrentUser godoc
//
//	@Summary		Get current user
//	@Description	This can only be done by the logged-in user.
//	@Tags			Me
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	GetCurrentUserOut
//	@Failure		default	{object}	response.Error
//	@Router			/me [get]
func (hd *HttpDelivery) GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*user.User)
	response.NewSuccess(hd.logger, w, r, newGetCurrentUserOut(user))
}

type UpdateCurrentUserEmailIn struct {
	Email string `json:"email"`
}

// UpdateCurrentUserEmail godoc
//
//	@Summary		Update current user email
//	@Description	This can only be done by the logged-in user.
//	@Tags			Me
//	@Accept			json
//	@Produce		json
//	@Param			email	body		UpdateCurrentUserEmailIn	true	"Email"
//	@Success		200		{object}	response.Success
//	@Failure		default	{object}	response.Error
//	@Router			/me/email [put]
func (hd *HttpDelivery) UpdateCurrentUserEmail(w http.ResponseWriter, r *http.Request) {
	var in UpdateCurrentUserEmailIn
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}

	user := r.Context().Value("user").(*user.User)
	err := hd.service.UpdateEmail(r.Context(), user.GetId().String(), in.Email)
	if err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}
	response.NewSuccess(hd.logger, w, r, nil)
}

type UpdateCurrentUserPasswordIn struct {
	Password string `json:"password"`
}

// UpdateCurrentUserPassword godoc
//
//	@Summary		Update current user password
//	@Description	This can only be done by the logged-in user.
//	@Tags			Me
//	@Accept			json
//	@Produce		json
//	@Param			password	body		UpdateCurrentUserPasswordIn	true	"Password"
//	@Success		200			{object}	response.Success
//	@Failure		default		{object}	response.Error
//	@Router			/me/password [put]
func (hd *HttpDelivery) UpdateCurrentUserPassword(w http.ResponseWriter, r *http.Request) {
	var in UpdateCurrentUserPasswordIn
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}

	user := r.Context().Value("user").(*user.User)
	err := hd.service.UpdatePassword(r.Context(), user.GetId().String(), in.Password)
	if err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}
	response.NewSuccess(hd.logger, w, r, nil)
}
