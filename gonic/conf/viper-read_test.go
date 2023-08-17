package conf

import (
	"fmt"
	"testing"

	"bytes"

	"github.com/spf13/viper"
)

func TestRead(t *testing.T) {
	//GetDuration
	viper.SetConfigType("yaml") // or viper.SetConfigType("YAML")
	var yamlExample = []byte(`
# Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
# 1 year
d: 8640hs
`)

	viper.ReadConfig(bytes.NewBuffer(yamlExample))

	d := viper.GetDuration("d") // this would be "steve"
	fmt.Println(d)
	fmt.Println(d == 0)
}
