package route

import (
	"letsgo/controller/api"

	"github.com/gorilla/mux"
	"golang.org/x/net/context"
)

var (
	router    = mux.NewRouter()
	apiRouter = router.PathPrefix("/v1").Subrouter()
)

type Router struct {
	apiHandler *api.BaseHandler
}

func NewRouter(h *api.BaseHandler) *Router {
	return &Router{
		apiHandler: h,
	}
}

func (h *Router) RegisterRoutes(ctx context.Context) *mux.Router {
	r := mux.NewRouter()

	apiRouter = r.PathPrefix("/api/v1").Subrouter()

	// Company API endpoints
	apiRouter.HandleFunc("/company/list", h.apiHandler.ListCompanies).Methods("GET")
	apiRouter.HandleFunc("/company", h.apiHandler.CreateCompany).Methods("POST")

	// Team API endpoints
	apiRouter.HandleFunc("/company/{id}/team/list", h.apiHandler.ListTeams).Methods("GET")
	apiRouter.HandleFunc("/company/{id}/team", h.apiHandler.CreateTeam).Methods("POST")
	apiRouter.HandleFunc("/company/{id}/team/{teamid}/users", h.apiHandler.AddUsersToTeam).Methods("POST")

	// User API endpoints
	apiRouter.HandleFunc("/company/{id}/user/list", h.apiHandler.ListUsers).Methods("GET")
	apiRouter.HandleFunc("/company/{id}/user", h.apiHandler.CreateUser).Methods("POST")
	apiRouter.HandleFunc("/company/{id}/user/{userid}/teams", h.apiHandler.AddUserToTeams).Methods("POST")

	// Task API endpoints
	apiRouter.HandleFunc("/company/{id}/team/{teamid}/tasks/list", h.apiHandler.ListTasksByTeam).Methods("GET")
	apiRouter.HandleFunc("/company/{id}/team/{teamid}/tasks", h.apiHandler.CreateTask).Methods("POST")

	return r
}
