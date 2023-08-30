package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

type Conf struct {
	vp *viper.Viper
}

func NewConf(path string) *Conf {
	conf := Conf{}
	load(path, &conf)
	return &conf
}

func load(path string, c *Conf) {
	v := viper.New()
	c.vp = v
	c.vp.SetConfigFile(path)
	err := c.vp.ReadInConfig() // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func (c *Conf) Vp() *viper.Viper {
	return c.vp
}

func (c *Conf) GetConf(path string) map[string]interface{} {
	return c.vp.AllSettings()
}

func (c *Conf) GetConfByKey(path, key string) interface{} {
	return c.vp.Get(key)
}
