package main

func slice2interfce(){
b := make([]interface{}, len(a))
for i := range a {
    b[i] = a[i]
}

func main(){
    func InterfaceSlice(slice interface{}) []interface{} {
    s := reflect.ValueOf(slice)
    if s.Kind() != reflect.Slice {
        panic("InterfaceSlice() given a non-slice type")
    }

    // Keep the distinction between nil and empty slice input
    if s.IsNil() {
        return nil
    }

    ret := make([]interface{}, s.Len())

    for i:=0; i<s.Len(); i++ {
        ret[i] = s.Index(i).Interface()
    }

    return ret
}

}
