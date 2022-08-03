package main

import (
	"fmt"

	"github.com/spf13/viper"
)


func overrideLoad(in string) {
	viper.SetConfigName(in)
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
        fmt.Println(err)
	}
    keys := []string{"env", "app"}
    for _, key := range keys{
        fmt.Println(key+":",viper.GetString(key) )
    }

}


func main(){
    overrideLoad("conf")
    overrideLoad("conf_dev")
}
