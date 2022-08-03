package main
import (
	"github.com/go-redis/redis"
    "fmt"

)

func main(){

	var client *redis.Client

    // 还可用universal 模式
    client = redis.NewFailoverClient(&redis.FailoverOptions{
        MasterName:    "master",
        SentinelAddrs: []string{"host1:port1","host2:port2"},
        Password:"password",
    })
    fmt.Println("echo:",client.Echo("abc"))
    client.Set("k","ahui",0)
    v:=client.Get("k")
    fmt.Println("get:",v)

}


/**
可用以下方式代替:

    redisClient = redis.NewUniversalClient(&redis.UniversalOptions{
        MasterName: viper.GetString("redisConfig.masterName"),
        Addrs:      viper.GetStringSlice("redisConfig.addrs"),
        Password:   viper.GetString("redisConfig.password"),
        DB:         0,
    })

原型：

    func NewUniversalClient(opts *UniversalOptions) UniversalClient {
        if opts.MasterName != "" {
            return NewFailoverClient(opts.Failover())
        } else if len(opts.Addrs) > 1 {
            return NewClusterClient(opts.Cluster())
        }
        return NewClient(opts.Simple())
    }


**/
