package inmemmory

import "skill-review/internal/config"

func ConfigLoaderMock() (c config.Config, err error) {
	return config.Config{
		Environment: "test",
		Version:     "1",
	}, nil
}
