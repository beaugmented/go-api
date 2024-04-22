package api

import (
	"github.com/eyecuelab/kit/web/server"
	"github.com/beaugmented/go-api/cmd/api/routes"
	"github.com/beaugmented/go-api/cmd/middleware"
	"github.com/spf13/cobra"
)

// APICmd ...
var APICmd *cobra.Command

func Init() {
	routes.Init()
	cobra.OnInitialize(func() {
		server.AddMiddleWare(middleware.Authed())
		server.AddMiddleWare(middleware.Cors())
	})
}
