package main
import "fmt"

func f()(int,int){
    return 1,2
}

func main(){
    res, err:= func()(res int, err int){
        res1,err := f()
        if err>0{
            // err 返回局部变量2!
            return
        }
        return res1*100,err
        
    }()
    fmt.Println(res, err) 
    
}

