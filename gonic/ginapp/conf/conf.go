package conf

import (
	"flag"
	"ginapp/fslib"
	"log"
	"os"
	"path"
	"runtime"
	"runtime/debug"

	"github.com/spf13/viper"
)

func initGcConf() {
	debug.SetGCPercent(1000)
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.Println("max procs:", runtime.GOMAXPROCS(0))
}
func init() {
	initGcConf()
	initWkDir()
}

type Config struct {
	Http Http `mapstructure:"http"`
}
type Http struct {
	ReadTimeout  int `mapstructure:"read_timeout"`
	WriteTimeout int `mapstructure:"write_timeout"`
}

var _config = &Config{}

func GetConf() *Config {
	if _config == nil {
		_loadConfig("conf")
		if err := viper.Unmarshal(&_config); err != nil {
			panic(err)
		}
	}
	return _config
}

// LoadConfig 载入配置
func _loadConfig(in string) {
	if !fslib.IsValidPath("./config") {
		return
	}
	viper.SetConfigName(in)
	viper.AddConfigPath("config/")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal("fail to load config file:", err)
	}
}

func initWkDir() {
	if !isInTest() {
		println("in normal mode(not test)")
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
