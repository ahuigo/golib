 package main

 import (
         "fmt"
         "path/filepath"
 )

 // like dir.Walk
 func main() {

         pattern := "*input*"

         matches, err := filepath.Glob(pattern)

         if err != nil {
                 fmt.Println(err)
         }

         fmt.Println(matches)

         // search upper directory
         upperDirPattern := "../*input*"  

         // you can specify directly the directory you want to search as well with the pattern
         // for example, /usr/files/*input*

         matches, err = filepath.Glob(upperDirPattern)

         if err != nil {
                 fmt.Println(err)
         }

         fmt.Println(matches)
 }
