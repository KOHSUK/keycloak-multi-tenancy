package server

import (
	"api/identityaccess/application/command/tenant"
	"api/identityaccess/infrastructure/chi/tenantroute"
	"api/identityaccess/infrastructure/gorm/connector"
	"api/identityaccess/infrastructure/gorm/repository/tenantrepo"
	"api/identityaccess/interface/controller/tenantctl"
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

type Server struct {
	mux *chi.Mux
}

func NewServer(mux *chi.Mux) *Server {
	return &Server{mux: mux}
}

func NewRouter(lc fx.Lifecycle, routes *Routes) *chi.Mux {
	r := chi.NewRouter()

	r.Post("/provision-tenant", routes.provisionTenantRoute.ServeHTTP)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				fmt.Println("Starting HTTP server at :8080")
				http.ListenAndServe(":8080", r)
			}()
			return nil
		},
	})

	return r
}

func Start() {
	app := fx.New(
		fx.Provide(
			connector.NewPostgresConnector,
			tenantrepo.NewGormTenantFactory,
			tenantrepo.NewGormTenantRepository,
			tenant.NewProvisionTenantCommandHandler,
			tenantctl.NewProvisionTenantController,
			tenantroute.NewProvisionTenantRoute,
			NewRoutes,
			NewRouter,
			NewServer,
		),
		fx.Invoke(func(*Server) {}),
	)

	app.Run()
}
