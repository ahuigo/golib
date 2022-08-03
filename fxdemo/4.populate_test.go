package fxdemo

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/ahuigo/glogger"
	"go.uber.org/fx"
)

func TestPopulate(t *testing.T) {
	var log = glogger.GetLogger("proj", glogger.InfoLevel)
	var reader io.Reader

	app := fx.New(
		// io.reader的应用
		// 提供构造函数
		fx.Provide(func() io.Reader {
			log.Infof("1")
			return strings.NewReader("hello world")
		}),
		fx.Populate(&reader), // 通过依赖注入完成变量与具体类的映射
	)
	log.Infof("2")
	app.Start(context.Background())
	log.Infof("3")
	defer app.Stop(context.Background())

	// 使用
	// reader变量已与fx.Provide注入的实现类关联了
	bs, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Panic("read occur error, ", err)
	}
	fmt.Printf("the result is '%s' \n", string(bs))
}
