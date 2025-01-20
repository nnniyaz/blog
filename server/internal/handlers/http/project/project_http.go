package projectHandler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/nnniyaz/blog/internal/domain/project"
	"github.com/nnniyaz/blog/internal/handlers/http/response"
	"github.com/nnniyaz/blog/internal/services/project"
	"github.com/nnniyaz/blog/pkg/logger"
	"net/http"
	"time"
)

type HttpDelivery struct {
	logger  logger.Logger
	service projectService.ProjectService
}

func NewHttpDelivery(logger logger.Logger, service projectService.ProjectService) *HttpDelivery {
	return &HttpDelivery{
		logger:  logger,
		service: service,
	}
}

type CreateProjectIn struct {
	Name           string `json:"name"`
	Description    string `json:"description"`
	CoverUri       string `json:"coverUri"`
	AppLink        string `json:"appLink"`
	SourceCodeLink string `json:"sourceCodeLink"`
}

// CreateProject godoc
//
//	@Summary		Create project
//	@Description	This can only be done by the logged-in user.
//	@Tags			Project
//	@Accept			json
//	@Produce		json
//	@Param			data	body		CreateProjectIn	true	"Create Project Structure"
//	@Success		200		{object}	response.Success
//	@Failure		default	{object}	response.Error
//	@Router			/project [post]
func (hd *HttpDelivery) CreateProject(w http.ResponseWriter, r *http.Request) {
	var in CreateProjectIn
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}
	err := hd.service.Create(r.Context(), in.Name, in.Description, in.CoverUri, in.AppLink, in.SourceCodeLink)
	if err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}
	response.NewSuccess(hd.logger, w, r, nil)
}

type UpdateProjectIn struct {
	Name           string `json:"name"`
	Description    string `json:"description"`
	CoverUri       string `json:"coverUri"`
	AppLink        string `json:"appLink"`
	SourceCodeLink string `json:"sourceCodeLink"`
}

// UpdateProject godoc
//
//	@Summary		Update project
//	@Description	This can only be done by the logged-in user.
//	@Tags			Project
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string			true	"Project ID"
//	@Param			data	body		UpdateProjectIn	true	"Update Project Structure"
//	@Success		200		{object}	response.Success
//	@Failure		default	{object}	response.Error
//	@Router			/project/{id} [put]
func (hd *HttpDelivery) UpdateProject(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var in UpdateProjectIn
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}
	err := hd.service.Update(r.Context(), id, in.Name, in.Description, in.CoverUri, in.AppLink, in.SourceCodeLink)
	if err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}
	response.NewSuccess(hd.logger, w, r, nil)
}

// DeleteProject godoc
//
//	@Summary		Delete project
//	@Description	This can only be done by the logged-in user.
//	@Tags			Project
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string	true	"Project ID"
//	@Success		200		{object}	response.Success
//	@Failure		default	{object}	response.Error
//	@Router			/project/{id} [delete]
func (hd *HttpDelivery) DeleteProject(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := hd.service.Delete(r.Context(), id)
	if err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}
	response.NewSuccess(hd.logger, w, r, nil)
}

// RestoreProject godoc
//
//	@Summary		Restore project
//	@Description	This can only be done by the logged-in user.
//	@Tags			Project
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string	true	"Project ID"
//	@Success		200		{object}	response.Success
//	@Failure		default	{object}	response.Error
//	@Router			/project/restore/{id} [put]
func (hd *HttpDelivery) RestoreProject(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := hd.service.Restore(r.Context(), id)
	if err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}
	response.NewSuccess(hd.logger, w, r, nil)
}

type GetProjectOut struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	CoverUri       string `json:"coverUri"`
	AppLink        string `json:"appLink"`
	SourceCodeLink string `json:"sourceCodeLink"`
	IsDeleted      bool   `json:"isDeleted"`
	CreatedAt      string `json:"createdAt"`
	UpdatedAt      string `json:"updatedAt"`
}

func newGetProjectOut(p *project.Project) *GetProjectOut {
	return &GetProjectOut{
		Id:             p.GetId().String(),
		Name:           p.GetName(),
		Description:    p.GetDescription(),
		CoverUri:       p.GetCoverUri(),
		AppLink:        p.GetAppLink(),
		SourceCodeLink: p.GetSourceCodeLink(),
		IsDeleted:      p.GetIsDeleted(),
		CreatedAt:      p.GetCreatedAt().Format(time.RFC3339),
		UpdatedAt:      p.GetUpdatedAt().Format(time.RFC3339),
	}
}

// GetProject godoc
//
//	@Summary		Find project by ID
//	@Description	This can only be done by the logged-in user.
//	@Tags			Project
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string	true	"Project ID"
//	@Success		200		{object}	response.Success{GetProjectOut}
//	@Failure		default	{object}	response.Error
//	@Router			/project/{id} [get]
func (hd *HttpDelivery) GetProject(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	project, err := hd.service.FindById(r.Context(), id)
	if err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}
	response.NewSuccess(hd.logger, w, r, newGetProjectOut(project))
}

type GetAllProjectsOut struct {
	Projects []*GetProjectOut `json:"projects"`
	Count    int64            `json:"count"`
}

func newGetAllProjectsOut(projects []*project.Project, count int64) *GetAllProjectsOut {
	var out []*GetProjectOut
	for _, p := range projects {
		out = append(out, newGetProjectOut(p))
	}
	return &GetAllProjectsOut{
		Projects: out,
		Count:    count,
	}
}

// GetAllProjects godoc
//
//	@Summary		Find all projects
//	@Description	This can only be done by the logged-in user.
//	@Tags			Project
//	@Accept			json
//	@Produce		json
//	@Param			offset		query		int64	false	"Offset"
//	@Param			limit		query		int64	false	"Limit"
//	@Param			isDeleted	query		bool	false	"Is Deleted"
//	@Param			search		query		string	false	"Search"
//	@Success		200			{object}	response.Success{GetAllProjectsOut}
//	@Failure		default		{object}	response.Error
//	@Router			/project [get]
func (hd *HttpDelivery) GetAllProjects(w http.ResponseWriter, r *http.Request) {
	offset := r.Context().Value("offset").(int64)
	limit := r.Context().Value("limit").(int64)
	isDeleted := r.Context().Value("is_deleted").(bool)
	search := r.Context().Value("search").(string)
	projects, count, err := hd.service.FindAll(r.Context(), offset, limit, isDeleted, search)
	if err != nil {
		response.NewError(hd.logger, w, r, err)
		return
	}
	response.NewSuccess(hd.logger, w, r, newGetAllProjectsOut(projects, count))
}
