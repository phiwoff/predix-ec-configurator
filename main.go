package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/indaco/predix-ec-configurator/controllers"
	"github.com/indaco/predix-ec-configurator/helpers"
	"github.com/indaco/predix-ec-configurator/services"
)

func init() {
	outputFolderName := "output"
	if helpers.IsExist(outputFolderName) {
		log.Println("-> Removing old generated files...")
		err := os.RemoveAll(outputFolderName)
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	log.Println("-> Starting EC-CONFIGURATOR...")
	log.Println("-> Initializing the app...")
	// Load configurations
	appConfig := helpers.DefaultAppSettings()
	userConfig := helpers.UserConfig{}
	_ = helpers.LoadConfig("config.json", &userConfig)

	// Create a CF Predix Client and sign-in to Predix.io
	predixClientConfig := helpers.GetPredixClientConfig(&userConfig)
	predixClient := helpers.PredixLogin(predixClientConfig)
	predixService, _ := services.NewPredixService(predixClient)

	// Create services for both scenarios
	scenarioOneService, _ := services.NewScenarioOneService(appConfig, &userConfig)
	scenarioTwoService, _ := services.NewScenarioTwoService(appConfig, &userConfig)

	// Initialize the output folders structure
	helpers.InitAppStructure(appConfig)
	// Download latest version for EC-SDK form Github
	helpers.DownloadLatestECSDKVersion(appConfig)

	// Declare app controllers
	staticC := controllers.NewStatic()
	scenarioOneC := controllers.NewScenario(appConfig, &userConfig, predixService, scenarioOneService)
	scenarioTwoC := controllers.NewScenario(appConfig, &userConfig, predixService, scenarioTwoService)

	// Create a new Router
	r := mux.NewRouter()
	// Declare requests handlers
	r.HandleFunc("/", staticC.Home).Methods("GET")
	r.Handle("/start", staticC.Start).Methods("GET")
	// Scenario 1
	r.HandleFunc("/scenario-1", scenarioOneC.New).Methods("GET")
	r.HandleFunc("/scenario-1", scenarioOneC.Create).Methods("POST")
	// Scenario 2
	r.HandleFunc("/scenario-2", scenarioTwoC.New).Methods("GET")
	r.HandleFunc("/scenario-2", scenarioTwoC.Create).Methods("POST")
	// Ajax requests - works for both scenarios
	r.HandleFunc("/retrieveOrgSpaces", scenarioOneC.OrgSpaces).Methods("POST")
	r.HandleFunc("/retrieveSpaceApps", scenarioOneC.SpaceApps).Methods("POST")
	r.HandleFunc("/retrieveAppEnv", scenarioOneC.AppServicesEnv).Methods("POST")

	// Asset contents
	assetHandler := http.FileServer(http.Dir("./public/"))
	assetHandler = http.StripPrefix("/public/", assetHandler)
	r.PathPrefix("/public/").Handler(assetHandler)

	// Declare the address for the HTTP server
	addr := fmt.Sprintf("%s:%d", userConfig.WebServer.Host, userConfig.WebServer.Port)

	// Create the HTTP Server
	srv := &http.Server{
		Handler:      r,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second}
	log.Printf("==> EC-CONFIGURATOR is up and running on port %d <==\n", userConfig.WebServer.Port)

	// Starts the HTTP server with a given address and handler
	log.Fatal(srv.ListenAndServe())
}
