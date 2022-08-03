package main

import (
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jinzhu/gorm"
    "github.com/spf13/viper"
    "sync"
    "time"
)



func multi_run(count int, fn func()){
    var wg sync.WaitGroup
    for i:=0; i<count; i++{
        wg.Add(1)
        go func(){
            defer wg.Done()
            fn()
        }()
    }
    wg.Wait()
    println(count, "tasks done")
}

var db *gorm.DB


type MonitorK8sJob struct {
	TaskName      string
	JobName       string
	IsRunning     bool
	TaskId        string
	TaskInputData string
	ResourceParam float64
	CreateTime    string
	UpdateTime    string
}


// read
func selectStock() {
    tasks := []*MonitorK8sJob{}
	//var monitorK8sJob MonitorK8sJob
    //err := db.Table("monitor_k8s_jobs").Where("task_name = ?", "xiangru_k8s").First(monitorK8sJob).Error
    err := db.Table("monitor_k8s_jobs").Where("task_name = ?", "xiangru_k8s").First(&tasks).Error
   if err!=nil{
    fmt.Println("error select", err)
}
    //fmt.Println(tasks[0])
}




func main() {
    var err error
	viper.SetConfigName("conf.local")
	viper.AddConfigPath("./")
	viper.ReadInConfig()
    // host=host.com user=u1 dbname=dev sslmode=disable password=pass
    dsn := viper.GetString("mo_pg_dsn")
    if dsn==""{
        panic("bad dsn:"+dsn)
    }
	for i:=0; i<15; i++{
        db, err = gorm.Open("postgres", dsn)
		if err == nil {
			break
		}
	}
	if err != nil {
        println("open error:",err.Error())
		fmt.Println(err)
		panic("连接数据库失败")
	}
	db.LogMode(false)

    // set pool
    if true{
        // SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
        db.DB().SetMaxIdleConns(3)

        // SetMaxOpenConns sets the maximum number of open connections to the database.
        db.DB().SetMaxOpenConns(3)

        // SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
        db.DB().SetConnMaxLifetime(2*time.Minute)
    }


	// 自动迁移模式
    multi_run(1000, selectStock)


	defer db.Close()
}
