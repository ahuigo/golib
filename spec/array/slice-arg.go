package main

func main(){
    var a = []int{1,2,3}
    func (l []int){
        l[0]=100
    }(a)
    println(a[0])
}
