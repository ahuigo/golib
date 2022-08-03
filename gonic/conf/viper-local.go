package main

import (
	"fmt"

	Viper "github.com/spf13/viper"
)


type Toleration struct{
    Key string
    Effect string
    Value string
    Operator string
}

func main(){
    viper:=Viper.New() //非全局, local-viper
    in := "conf"
	viper.SetConfigName(in)
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
        fmt.Println(err)
        return
	}

    tolerations :=[]Toleration{}
    //tolerations :=[]corev1.Toleration{}
    err=viper.UnmarshalKey("tolerations",&tolerations)
	if err != nil {
        fmt.Println(err)
        return
	}
    fmt.Printf("conf: %#v\n", tolerations)

}
