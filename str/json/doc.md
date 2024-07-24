# 不同语言json 对[]byte的处理
json　本身无法区分string 和 `[]byte`，只有string。

1. php 自己本身就没有`[]byte`，甚至连dict/array 都不区分。直接回避了问题
2. python　有`[]byte`, 但是不能json序列化, 如果是utf8就需要用 data.decode('utf-8') 先转成str
3. golang　序列化`[]byte` 时用base64 将它转成string, json反序列时则用相应的类型去关联`[]byte`类型
4. java , c sharp 跟golang 处理方式一样的。

# golang 对json 的支持
json.Marshal 方法:
1. json.Marshal 本身是使用了内置的json encoder 默认使用base64 处理 bytes, eacapHTML默认为true
1. 但是 json.RawMessage  本身带的json encoder，不会对bytes做任何处理，如果bytes不是json格式，就会报编码出错

某些golang　库实现，可能出于减少值复制，喜欢用bytes去表达string，序列化时就只能是base64　string了。
如果接收者不要base64 string json, 解决办法有：
1. 自己先定义一个相应string struct，再将bytes struct转成string struct，最后json　序列化自然就不是base64了
1. 通过反射将 string struct/map 转成　string map 再json. 参考：https://github.com/ahuigo/golib/tree/main/spec/object/convert/objbytes2string_test.go

