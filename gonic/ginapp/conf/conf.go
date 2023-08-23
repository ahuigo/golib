package conf

import (
	"flag"
	"ginapp/fslib"
	"log"
	"os"
	"path"
	"runtime"

	"github.com/spf13/viper"
)

func init() {
	initWkDir()
	// InitViper()
}

type Config struct {
	Http Http `mapstructure:"http"`
}
type Http struct {
	ReadTimeout  int `mapstructure:"read_timeout"`
	WriteTimeout int `mapstructure:"write_timeout"`
}

var _inited = false
var _config = Config{}

func GetConf() *Config {
	if !_inited {
		_inited = true
		_loadConfig("conf")
		if err := viper.Unmarshal(&_config); err != nil {
			panic(err)
		}
	}
	return &_config
}

// LoadConfig 载入配置
func _loadConfig(in string) {
	if !fslib.IsValidPath("./config") {
		return
	}
	viper.SetConfigName(in)
	viper.AddConfigPath("config/")
	err := viper.ReadInConfig() // Find and read the config file

	if err != nil { // Handle errors reading the config file
		log.Fatal("fail to load config file:", err)
	}
}

func initWkDir() {
	if !isInTest() {
		println("in normal mode")
		return
	}
	println("in test mode")
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}

func isInTest() bool {
	// strings.HasSuffix(os.Args[0], ".test")
	//Or strings.Contains(os.Args[0], "/_test/")
	return flag.Lookup("test.v") != nil
}
