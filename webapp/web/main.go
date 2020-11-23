package main

import (
	"context"

	"github.com/samthor/nicehttp"

	"github.com/web-platform-tests/wpt.fyi/api"
	"github.com/web-platform-tests/wpt.fyi/api/azure"
	"github.com/web-platform-tests/wpt.fyi/api/checks"
	"github.com/web-platform-tests/wpt.fyi/api/query"
	"github.com/web-platform-tests/wpt.fyi/api/receiver"
	"github.com/web-platform-tests/wpt.fyi/api/screenshot"
	"github.com/web-platform-tests/wpt.fyi/api/taskcluster"
	"github.com/web-platform-tests/wpt.fyi/shared"
	"github.com/web-platform-tests/wpt.fyi/webapp"
)

func init() {
	// webapp.RegisterRoutes has a catch-all, so needs to go last.
	api.RegisterRoutes()
	azure.RegisterRoutes()
	checks.RegisterRoutes()
	query.RegisterRoutes()
	receiver.RegisterRoutes()
	screenshot.RegisterRoutes()
	taskcluster.RegisterRoutes()
	webapp.RegisterRoutes()
}

func main() {
	if err := shared.Clients.Init(context.Background()); err != nil {
		shared.Clients.Close()
		panic(err)
	}
	defer shared.Clients.Close()

	// This behaves differently in prod and locally:
	// * Prod: the provided app.yaml is not used; it simply starts the
	//   DefaultServerMux on $PORT (defaults to 8080).
	// * Local: in addition to the prod behaviour, it also starts some
	//   static handlers according to app.yaml, which effectively replaces
	//   dev_appserver.py.
	nicehttp.Serve("webapp/web/app.staging.yaml", nil)
}
