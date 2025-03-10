package main

import (
	"api-shortener/db"
	"api-shortener/restapi"
	"api-shortener/shortreq"
	"net/http"

	"go.uber.org/fx"
)

func main() {
	LoadEnv()
	fx.New(
		fx.Provide(
			NewHTTPServer,

			shortreq.NewIncomingRequestProcessor,
			shortreq.NewResponseShorteningService,
			shortreq.NewLoopLimiter,
			shortreq.NewLoopLimiterSettings,

			db.NewDBConnectionSettings,
			db.NewDB,
			db.NewMigrator,

			restapi.NewShortenedAPIDAO,
			restapi.NewOutgoingRequestConfigDAO,
			restapi.NewOutgoingRequestHeaderDAO,
			restapi.NewOutgoingRequestParamDAO,
			restapi.NewShorteningRuleDAO,

			shortreq.NewOutgoingRequestProcessor,
			shortreq.NewJsonResponseShortener,

			shortreq.NewOutgoingRequestClientSettings,
			shortreq.NewOutgoingRequestClient,

			restapi.NewRESTService,

			shortreq.NewValidate,
		),
		fx.Invoke(
			func(migrator db.IDBMigrator) { migrator.Migrate() },
			func(*http.Server) {},
		),
	).Run()
}
