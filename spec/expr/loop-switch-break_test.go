package demo
func TestLoopBreak(t *testing.T){
loop:
for {
        switch i {
        case foo:
                if condA {
                        doA()
                        break // like 'goto A'
                }

                if condB {
                        doB()
                        break loop // like 'goto B'                        
                }

                doC()
        case bar:
                // ...
        }
A:
    println("A")
        // ...
}

B:
println("b")
}
