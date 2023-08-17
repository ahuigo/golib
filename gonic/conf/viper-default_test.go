package conf

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/spf13/viper"
)

func TestDefault(t *testing.T) {
	var confs = map[string][]byte{
		"conf": []byte(`
env: local
app: ahui
`),
		"conf_dev": []byte(`
env: dev
`),
	}
	overrideLoad := func(in string) {
		fmt.Println("=================")
		err := viper.ReadConfig(bytes.NewBuffer(confs[in]))
		viper.SetDefault("env", "default_env")
		viper.SetDefault("app", "default_app")
		//viper.Set("env", "dy_new_value")
		if err != nil {
			fmt.Println(err)
		}
		keys := []string{"env", "app"}
		for _, key := range keys {
			fmt.Println(key+":", viper.GetString(key))
		}

	}
	viper.SetConfigType("yaml") // or viper.SetConfigType("YAML")
	overrideLoad("conf")
	overrideLoad("conf_dev")
}
