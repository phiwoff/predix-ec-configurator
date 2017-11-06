package services

import (
	cfclient "github.com/cloudfoundry-community/go-cfclient"
)

// PredixService ...
type PredixService struct {
	client *cfclient.Client
}

// NewPredixService ...
func NewPredixService(cfClient *cfclient.Client) (*PredixService, error) {
	return &PredixService{
		client: cfClient,
	}, nil
}

// GetOrgs ...
func (pc *PredixService) GetOrgs() ([]cfclient.Org, error) {
	orgs, err := pc.client.ListOrgs()
	if err != nil {
		return nil, err
	}
	return orgs, nil
}

// GetOrgSpaces ...
func (pc *PredixService) GetOrgSpaces(orgGUID string) ([]cfclient.Space, error) {
	spaces, err := pc.client.OrgSpaces(orgGUID)
	if err != nil {
		return nil, err
	}
	return spaces, nil
}

// GetSpaceApps ...
func (pc *PredixService) GetSpaceApps(spaceGUID string) ([]cfclient.AppSummary, error) {
	space, err := pc.client.GetSpaceByGuid(spaceGUID)
	if err != nil {
		return nil, err
	}
	spaceSummary, err := space.Summary()
	if err != nil {
		return nil, err
	}
	apps := spaceSummary.Apps
	return apps, nil
}

// GetAppServiceEnv ...
func (pc *PredixService) GetAppServiceEnv(appGUID string) (map[string]interface{}, error) {
	appEnv, err := pc.client.GetAppEnv(appGUID)
	if err != nil {
		return nil, err
	}
	return appEnv.SystemEnv, nil
}

// GetSpaceSummary ...
func (pc *PredixService) GetSpaceSummary(spaceGUID string) (cfclient.SpaceSummary, error) {
	space, err := pc.client.GetSpaceByGuid(spaceGUID)
	if err != nil {
		return cfclient.SpaceSummary{}, err
	}
	spaceSummary, err := space.Summary()
	if err != nil {
		return cfclient.SpaceSummary{}, err
	}
	return spaceSummary, nil
}
