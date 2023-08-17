package conf

import (
	"fmt"

	"github.com/spf13/viper"
)


type Toleration struct{
    Key string
    Effect string
    Value string
    Operator string
}

func main(){
    in := "conf"
	viper.SetConfigName(in)
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
        fmt.Println(err)
        return
	}

    // UnmarshalKey
    tolerations :=[]Toleration{}
    err=viper.UnmarshalKey("tolerations",&tolerations)
	if err != nil {
        fmt.Println(err)
        return
	}
    fmt.Printf("conf: %#v\n", tolerations)

    // Unmarshal
    data := struct{
        Env string
        App string
    }{}
    err=viper.Unmarshal(&data)
    fmt.Printf("conf: %#v\n", data)

}
