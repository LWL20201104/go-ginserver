package bootstrap

import (
	"ginserver/global"
	"github.com/bytedance/sonic"
	"testing"
)

func TestInitializeConfig(t *testing.T) {
	InitializeConfig("../config.yaml")

	appConfig, _ := sonic.MarshalString(global.AppConfig)
	t.Logf("Config: %s", appConfig)
}
