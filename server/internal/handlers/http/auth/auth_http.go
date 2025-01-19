package authHandler

import (
	"encoding/json"
	"github.com/nnniyaz/blog/internal/handlers/http/response"
	authService "github.com/nnniyaz/blog/internal/services/auth"
	"github.com/nnniyaz/blog/pkg/logger"
	"github.com/nnniyaz/blog/pkg/web"
	"net/http"
)

type HttpDelivery struct {
	logger  logger.Logger
	service authService.AuthService
}

func NewHttpDelivery(logger logger.Logger, service authService.AuthService) *HttpDelivery {
	return &HttpDelivery{logger: logger, service: service}
}

type LoginIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Login godoc
//
//	@Summary		Login
//	@Description	This can be done by a user who has an account
//	@Tags			Authorization
//	@Accept			json
//	@Produce		json
//	@Param			data			body	LoginIn		true	"Login information"
//	@Success		200						{object}	response.Success
//	@Failure		default					{object}	response.Error
//	@Router			/auth/login [post]
func (hd *HttpDelivery) Login(w http.ResponseWriter, r *http.Request) {
	requestInfo := r.Context().Value("requestInfo").(web.RequestInfo)

	var in LoginIn
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}

	token, err := hd.service.Login(r.Context(), in.Email, in.Password, requestInfo.UserAgent.String)
	if err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "blog-app-session",
		Value:    token.String(),
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	})
	response.NewSuccess(hd.logger, w, r, token)
}

// Logout godoc
//
//	@Summary		Logout
//	@Description	This can only be done by the logged-in user.
//	@Tags			Authorization
//	@Accept			json
//	@Produce		json
//	@Success		200						{object}	response.Success
//	@Failure		default					{object}	response.Error
//	@Router			/auth/logout [post]
func (hd *HttpDelivery) Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("blog-app-session")
	if err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}

	err = hd.service.Logout(r.Context(), cookie.Value)
	if err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "blog-app-session",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	})
	response.NewSuccess(hd.logger, w, r, nil)
}
