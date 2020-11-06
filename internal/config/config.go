package config

import (
	"path/filepath"
	"runtime"

	"github.com/errntry/errntry/internal/errors"
	"github.com/errntry/errntry/pkg/dotenv"
)

var CachedConfig *dotenv.DotEnv
var ConfigErr error

func Init() (*dotenv.DotEnv, error) {
	_, b, _, _ := runtime.Caller(0)
	cfgFile := filepath.Join(filepath.Dir(b), "../../.env")

	CachedConfig, ConfigErr = dotenv.Init(cfgFile)
	return CachedConfig, ConfigErr
}

func Load() *dotenv.DotEnv {
	if CachedConfig != nil || ConfigErr != nil {
		if ConfigErr != nil {
			errors.Fatal(2, "failed to load config:", ConfigErr)
		}

		return CachedConfig
	}

	cfg, err := Init()
	if err != nil {
		errors.Fatal(2, "failed to load config:", ConfigErr)
	}

	return cfg
}
