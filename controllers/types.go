package controllers

import (
	"html/template"

	cfclient "github.com/cloudfoundry-community/go-cfclient"
)

type Account struct {
	Username string
}
type Orgs = []cfclient.Org

type PageContent struct {
	Account
	Orgs
	Body template.HTML
}
