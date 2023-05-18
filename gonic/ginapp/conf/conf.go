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
	InitViper()
}

var _inited = false

func InitViper() {
	if !_inited {
		_inited = true
		_loadConfig("conf")
	}
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
