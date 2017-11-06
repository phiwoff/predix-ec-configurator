package controllers

import (
	"encoding/json"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/indaco/predix-ec-configurator/helpers"
	"github.com/indaco/predix-ec-configurator/services"
	"github.com/indaco/predix-ec-configurator/views"
)

// NewScenario ...
func NewScenario(
	appConfig *helpers.AppSettings,
	userConfig *helpers.UserConfig,
	predixService *services.PredixService,
	scenarioService interface{}) *Scenario {

	var _view *views.View
	switch scenarioService.(type) {
	case *services.ScenarioOneService:
		_view = views.NewView("page", "scenarioOne/new")
	case *services.ScenarioTwoService:
		_view = views.NewView("page", "scenarioTwo/new")
	}
	return &Scenario{
		NewView:    _view,
		CreateView: views.NewMarkdownView("markdownpage"),
		ac:         appConfig,
		uc:         userConfig,
		ps:         predixService,
		sc:         scenarioService,
	}
}

// Scenario ...
type Scenario struct {
	NewView    *views.View
	CreateView *views.MarkdownView
	ac         *helpers.AppSettings
	uc         *helpers.UserConfig
	ps         *services.PredixService
	sc         interface{}
}

// New is used to render the form where a user
// can create a new configuration for Scenarios
// GET /scenario-1 and /scenario-2
func (s *Scenario) New(w http.ResponseWriter, r *http.Request) {
	orgs, _ := s.ps.GetOrgs()
	data := views.Data{
		Yield: PageContent{
			Account: Account{Username: s.uc.Predix.Username},
			Orgs:    orgs,
		},
	}
	if err := s.NewView.Render(w, data); err != nil {
		panic(err)
	}
}

// OrgSpaces ...
func (s *Scenario) OrgSpaces(w http.ResponseWriter, r *http.Request) {
	orgGUID := r.FormValue("ajax_post_data")
	//fmt.Println("Receive ajax post data string ", orgGUID)
	spaces, _ := s.ps.GetOrgSpaces(orgGUID)
	sp, _ := json.Marshal(spaces)
	w.Write([]byte(sp))
}

// SpaceApps ...
func (s *Scenario) SpaceApps(w http.ResponseWriter, r *http.Request) {
	spaceGUID := r.FormValue("ajax_post_data")
	//fmt.Println("Receive ajax post data string ", spaceGUID)
	apps, _ := s.ps.GetSpaceApps(spaceGUID)
	a, _ := json.Marshal(apps)
	w.Write([]byte(a))
}

// AppServicesEnv ...
func (s *Scenario) AppServicesEnv(w http.ResponseWriter, r *http.Request) {
	appGUID := r.FormValue("ajax_post_data")
	appSystemEnv, _ := s.ps.GetAppServiceEnv(appGUID)
	a, _ := json.Marshal(appSystemEnv)
	w.Write([]byte(string(a)))
}

// SpaceSummary ...
func (s *Scenario) SpaceSummary(w http.ResponseWriter, r *http.Request) {
	spaceGUID := r.FormValue("ajax_post_data")
	spaceSummary, _ := s.ps.GetSpaceSummary(spaceGUID)
	sp, _ := json.Marshal(spaceSummary)
	w.Write([]byte(sp))
}

// Create tries to create a new
// configuation set for Scenarios
// POST /scenario-1 or /scenario-2
func (s *Scenario) Create(w http.ResponseWriter, r *http.Request) {
	filename := filepath.Join(s.ac.Output.Root, s.ac.Output.FullDocFile)

	switch scenarioService := s.sc.(type) {
	case *services.ScenarioOneService:
		configurationForm := services.ScenarioOneConfigurationForm{
			TypeOfScenario:  r.URL.Path,
			UAAServiceURI:   r.FormValue("uaaServiceURI"),
			UAAClient:       r.FormValue("uaaClient"),
			UAAClientSecret: r.FormValue("uaaClientSecret"),

			ECZoneID:      r.FormValue("ecZoneID"),
			ECServiceURI:  r.FormValue("ecServiceURI"),
			ECAdminToken:  r.FormValue("ecAdmToken"),
			ECIDS:         strings.Split(r.FormValue("ecIDS"), ","),
			ECGatewayName: r.FormValue("ecGatewayName"),

			OnPremiseOS:     r.FormValue("onPremiseOS"),
			ResourceHost:    r.FormValue("resourceHost"),
			ResourcePort:    r.FormValue("resourcePort"),
			IsDebug:         helpers.GetDebugString(r.FormValue("isDebug")),
			HealthcheckPort: r.FormValue("healthcheckPort"),

			Proxy: helpers.GetProxyString(r.FormValue("proxyURL")),
		}
		scenarioService.GenerateAndSaveAllScenarioOne(filename, configurationForm)
		scenarioService.ExtractAndMoveECBinaryScenarioOne(configurationForm)
	case *services.ScenarioTwoService:
		configurationForm := services.ScenarioTwoConfigurationForm{
			TypeOfScenario:  r.URL.Path,
			UAAServiceURI:   r.FormValue("uaaServiceURI"),
			UAAClient:       r.FormValue("uaaClient"),
			UAAClientSecret: r.FormValue("uaaClientSecret"),

			ECZoneID:      r.FormValue("ecZoneID"),
			ECServiceURI:  r.FormValue("ecServiceURI"),
			ECAdminToken:  r.FormValue("ecAdmToken"),
			ECIDS:         strings.Split(r.FormValue("ecIDS"), ","),
			ECGatewayName: r.FormValue("ecGatewayName"),
			ECServerName:  r.FormValue("ecServerName"),

			ResourceHost: r.FormValue("resourceHost"),
			ResourcePort: r.FormValue("resourcePort"),

			OnPremiseOS:     r.FormValue("onPremiseOS"),
			LocalPort:       r.FormValue("localPort"),
			IsDebug:         helpers.GetDebugString(r.FormValue("isDebug")),
			HealthcheckPort: r.FormValue("healthcheckPort"),

			Proxy: helpers.GetProxyString(r.FormValue("proxyURL")),
		}
		scenarioService.GenerateAndSaveAllScenarioTwo(filename, configurationForm)
		scenarioService.ExtractAndMoveECBinaryScenarioTwo(configurationForm)
	}

	var vd views.Data
	vd.Yield = PageContent{
		Body: helpers.ParseMarkdown(filename),
	}
	if err := s.CreateView.Render(w, vd); err != nil {
		panic(err)
	}
}
