package main

import (
	"log"
	"net/http"
	"os"
	"slices"
	"strings"

	"github.com/TetraDash/tetradash/helpers"
	_ "github.com/TetraDash/tetradash/migrations"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
	"github.com/pocketbase/pocketbase/ui"
)

func main() {
	app := pocketbase.New()

	// loosely check if it was executed using "go run"
	isGoRun := strings.HasPrefix(os.Args[0], os.TempDir())

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		// enable auto creation of migration files when making collection changes in the Admin UI
		// (the isGoRun check is to enable it only during development)
		Automigrate: isGoRun,
	})

	// serves static files from the provided public dir (if exists)
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		helpers.ApplySettings(app)

		if helpers.ConfigEnv.PbEnableWebUI {
			e.Router.GET("/_/", func(c echo.Context) error {
				if slices.Contains(helpers.ConfigEnv.PbAdminAcceptDomains, c.Request().Host) {
					return echo.StaticDirectoryHandler(ui.DistDirFS, false)(c)
				}

				return c.JSON(http.StatusOK, map[string]any{"error": "pocketbase web ui disabled"})
			})
		}

		e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
