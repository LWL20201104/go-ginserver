package global

import (
	"ginserver/config"
	"go.uber.org/zap"
)

type Application struct {
	Config config.Configuration `json:"config" yaml:"config"`
	Log    *zap.Logger
}

var AppConfig = &Application{}
