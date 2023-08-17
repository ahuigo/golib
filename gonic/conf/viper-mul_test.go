package conf

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"
)

// 不支持multi load, 会被override
func TestMultiLoad(t *testing.T) {
	overrideLoad := func(in string) {
		viper.SetConfigName(in)
		viper.AddConfigPath("./")
		err := viper.ReadInConfig()
		if err != nil {
			fmt.Println(err)
		}
		keys := []string{"env", "app"}
		for _, key := range keys {
			fmt.Println(key+":", viper.GetString(key))
		}

	}

	overrideLoad("conf")
	overrideLoad("conf_dev")
}
