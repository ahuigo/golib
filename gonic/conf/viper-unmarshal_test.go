package conf

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"
)

type Toleration struct {
	Key      string
	Effect   string
	Value    string
	Operator string
}

func TestUnmarshal(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}

	in := "conf"
	viper.SetConfigName(in)
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 1. UnmarshalKey
	tolerations := []Toleration{}
	err = viper.UnmarshalKey("tolerations", &tolerations)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("tolerations: %#v\n", tolerations)

	// 2. Unmarshal
	data := struct {
		Env  string `yaml:"env1"` // viper不使用 yaml
		App  string
		User User `mapstructure:"user1"` // mapstructure 不区分大小写
	}{}
	if err = viper.Unmarshal(&data); err != nil {
		t.Fatal(err)
	}
	fmt.Printf("conf: %#v\n", data)
}
