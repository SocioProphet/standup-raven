package controller

import (
	"net/http"

	"github.com/standup-raven/standup-raven/server/controller/middleware"
	"github.com/standup-raven/standup-raven/server/util"
)

type Endpoint struct {
	Path        string
	Method      string
	Execute     func(w http.ResponseWriter, r *http.Request) error
	Middlewares []middleware.Middleware
}

var Endpoints = map[string]*Endpoint{
	getEndpointKey(hook):                     hook,
	getEndpointKey(getStandup):               getStandup,
	getEndpointKey(saveStandup):              saveStandup,
	getEndpointKey(getConfig):                getConfig,
	getEndpointKey(setConfig):                setConfig,
	getEndpointKey(getDefaultTimezone):       getDefaultTimezone,
	getEndpointKey(getActiveStandupChannels): getActiveStandupChannels,
	getEndpointKey(getPluginConfig):          getPluginConfig,
}

func getEndpointKey(endpoint *Endpoint) string {
	return util.GetKeyHash(endpoint.Path + "-" + endpoint.Method)
}

func GetEndpoint(r *http.Request) *Endpoint {
	return Endpoints[util.GetKeyHash(r.URL.Path+"-"+r.Method)]
}
