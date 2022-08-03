package main

import (
	"fmt"
    "bytes"

	"github.com/spf13/viper"
)


var confs = map[string][]byte{
    "conf": []byte(`
env: local
app: ahui
`),
    "conf_dev": []byte(`
env: dev
`),
}




func overrideLoad(in string) {
    fmt.Println("=================")
    //viper.Set("env", "dy_new_value")
    err :=viper.ReadConfig(bytes.NewBuffer(confs[in]))
    viper.SetDefault("env", "default_env")
    viper.SetDefault("app", "default_app")

	if err != nil {
        fmt.Println(err)
	}
    keys := []string{"env", "app"}
    for _, key := range keys{
        fmt.Println(key+":",viper.GetString(key) )
    }

}


func main(){
    viper.SetConfigType("yaml") // or viper.SetConfigType("YAML")
    overrideLoad("conf")
    overrideLoad("conf_dev")
}
