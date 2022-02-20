package server

import (
	"fmt"
	"log"
	"net/http"
	"qlt/internal/database"
)

//App ...
type App struct {
	config   *Config
	Router   *Router
	database *database.Database
}

//CreateApp ...
func CreateApp(paymentsRepository *PaymentsRepository) App {
	return App{Router: NewRouter(paymentsRepository),
	config: NewConfig()}
}

//Run ...
func (a *App) Run() error {
	port := 8000
	log.Println(fmt.Sprintf("Listening on port %d", port))
	if err := a.ConfigureDB(); err != nil {
		return err
	}
	return http.ListenAndServe(fmt.Sprintf(":%d", port), a.Router.Router)
}

// ConfigureDB ...
func (a *App) ConfigureDB() error {
	db := database.NewDB(a.config.Database)
	if err := db.InitDB(); err != nil {
		return err
	}
	a.database = db //fill Server with DB instance
	return nil
}
