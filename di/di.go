package di

import (
	"skill-review/internal/api"
	apihandlers "skill-review/internal/api/handlers"
	"skill-review/internal/config"
	"skill-review/internal/mainfeature"
)

func BaseParametersLoader() config.Loader {
	return config.LoadParameters
}

func ApiPostRoutes(c config.Loader) (routes []api.Route) {
	routes = append(routes, api.Route{
		Method:  "POST",
		Address: api.NamedParamAddressName + "/:message/",
		Handler: apihandlers.NewNamedParamHandler(MainFeatureProcessor(c)).Handler,
	})

	routes = append(routes, api.Route{
		Method:  "POST",
		Address: api.PostMessageAddressName,
		Handler: apihandlers.NewDefaultHandler(MainFeatureProcessor(c)).Handler,
	})

	return routes
}

func MainFeatureProcessor(c config.Loader) *mainfeature.Processor {
	return mainfeature.NewProcessor(c)
}
