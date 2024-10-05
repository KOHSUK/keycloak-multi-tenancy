package server

import (
	"api/identityaccess/application/command/tenant"
	"api/identityaccess/domain/identity/factory"
	"api/identityaccess/domain/identity/repository"
	"api/identityaccess/infrastructure/chi/tenantroute"
	"api/identityaccess/infrastructure/gorm/connector"
	"api/identityaccess/infrastructure/gorm/repository/tenantrepo"
	"api/identityaccess/infrastructure/inmemory/repository/inmemtenantrepo"
	"api/identityaccess/interface/controller/tenantctl"
	"api/identityaccess/usecase/command/uctenant"
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

type Server struct {
	mux *chi.Mux
}

func NewServer(lc fx.Lifecycle, mux *chi.Mux, logger *zap.Logger) *Server {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				err := http.ListenAndServe(":7070", mux)
				if err != nil {
					logger.Error("Error starting HTTP server", zap.Error(err))
				}
				logger.Info("HTTP server started at :7070")
			}()
			return nil
		},
	})

	server := &Server{mux: mux}

	return server
}

func NewServeMux(routes *Routes) *chi.Mux {
	mux := chi.NewMux()

	mux.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Health check")
		w.Write([]byte("OK"))
	})
	mux.Post("/provision-tenant", routes.provisionTenantRoute.ServeHTTP)

	return mux
}

func Start() {
	app := fx.New(
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
		fx.Supply(connector.PostgresConnectorConfig{
			DSN: "host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable",
		}),
		fx.Provide(
			NewServer,
			NewServeMux,
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
				inmemtenantrepo.NewInMemTenantFactory,
				fx.As(new(factory.TenantFactory)),
			),
			fx.Annotate(
				inmemtenantrepo.NewInMemTenantRepository,
				fx.As(new(repository.TenantRepository)),
			),
			tenant.NewProvisionTenantCommandHandler,
			tenantctl.NewProvisionTenantController,
			tenantroute.NewProvisionTenantRoute,
			NewRoutes,
			zap.NewExample,
		),
		fx.Invoke(func(*Server) {}),
	)

	app.Run()
}
