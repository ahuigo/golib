//链接：https://www.zhihu.com/question/425625461/answer/2356045938

package main

func bar() (r int) {
 defer func() {
  r += 4
  if recover() != nil {
   r += 8
  }
 }()
 
 var f func()
 defer f()  // undefined nil
 f = func() {
  r += 2
 }

 return 1
}

func main() {
 println(bar())
}
