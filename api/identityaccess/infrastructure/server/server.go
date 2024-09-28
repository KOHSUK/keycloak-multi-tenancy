package server

import (
	"api/identityaccess/application/command/tenant"
	"api/identityaccess/domain/identity/factory"
	"api/identityaccess/domain/identity/repository"
	"api/identityaccess/infrastructure/chi/tenantroute"
	"api/identityaccess/infrastructure/gorm/connector"
	"api/identityaccess/infrastructure/gorm/repository/tenantrepo"
	"api/identityaccess/interface/controller/tenantctl"
	"api/identityaccess/usecase/command/uctenant"
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

type Server struct {
	mux *chi.Mux
}

func NewServer(lc fx.Lifecycle, routes *Routes) *Server {
	mux := chi.NewRouter()

	mux.Post("/provision-tenant", routes.provisionTenantRoute.ServeHTTP)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				fmt.Println("Starting HTTP server at :8080")
				http.ListenAndServe(":8080", mux)
			}()
			return nil
		},
	})

	return &Server{mux: mux}
}

func Start() {
	app := fx.New(
		fx.Supply(connector.PostgresConnectorConfig{
			DSN: "host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable",
		}),
		fx.Provide(
			connector.NewPostgresConnector,
			tenantrepo.NewGormTenantFactory,
			fx.Annotate(
				connector.NewPostgresConnector,
				fx.As(new(connector.Connector)),
			),
			fx.Annotate(
				tenant.NewProvisionTenantCommandHandler,
				fx.As(new(uctenant.ProvisionTenantUseCase)),
			),
			fx.Annotate(
				tenantrepo.NewGormTenantFactory,
				fx.As(new(factory.TenantFactory)),
			),
			fx.Annotate(
				tenantrepo.NewGormTenantRepository,
				fx.As(new(repository.TenantRepository)),
			),
			tenant.NewProvisionTenantCommandHandler,
			tenantctl.NewProvisionTenantController,
			tenantroute.NewProvisionTenantRoute,
			NewRoutes,
			NewServer,
		),
		fx.Invoke(func(*Server) {}),
	)

	app.Run()
}
