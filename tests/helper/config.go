package helper

import "github.com/MOZGIII/evans/config"

func TestConfig() *config.Config {
	return config.Get()
}
