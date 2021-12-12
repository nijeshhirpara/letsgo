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

func RegisterRoutes(ctx context.Context) *mux.Router {
	r := mux.NewRouter()

	apiRouter = r.PathPrefix("/api/v1").Subrouter()

	apiRouter.HandleFunc("/company/list", api.ListCompanies).Methods("GET")
	apiRouter.HandleFunc("/company", api.CreateCompany).Methods("POST")

	return r
}
