package di

import (
	"skill-review/internal/config"
)

func BaseParametersLoader() config.Loader {
	return config.LoadParameters
}
