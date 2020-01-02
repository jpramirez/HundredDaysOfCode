package web

import (
	models "github.com/jpramirez/HundredDaysOfCode/DayOne/pkg/models"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	//utils "github.com/jpramirez/HundredDaysOfCode/DayOne/pkg/utils"

	webapp "github.com/jpramirez/HundredDaysOfCode/DayOne/pkg/web/app"
)

type WebOne struct {
	webConfig models.Config
}

//StartServer Starts the server using the variable sip and port, creates anew instance.
func (W *WebOne) StartServer() {
	http.ListenAndServe(W.webConfig.WebAddress+":"+W.webConfig.WebPort, W.New())
}

//NewWebAgent creates new instance.
func NewWebAgent(config models.Config) (WebOne, error) {
	var webone WebOne
	log.Println("Starting News Agent")
	webone.webConfig = config
	// Stop the grpc verbose logging
	return webone, nil
}

//New creates a new handler
func (W *WebOne) New() http.Handler {

	app, err := webapp.NewApp(W.webConfig)

	if err != nil {
		log.Fatalln("Error creating WebApp", err)
		return nil
	}

	api := app.Mux.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/import", app.ImportFullDataSet)
	api.HandleFunc("/getdata", app.GetFullDataSet)
	api.HandleFunc("/status", app.Liveness)
	api.HandleFunc("/check/{country}", app.CheckCountry)

	ch := make(chan os.Signal, 2)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-ch
		log.Println("Closing Database")
		os.Exit(0)
	}()

	return &app
}
