fmt := import("fmt")
time := import("time")
rand := import("rand")
text := import("text")
http := import("http")
crypto := import("crypto")
os := import("os")
arr := ["foo", "bar", [1, 2, 3], {i:"j"}]
map := {a: [1,2,3], b: {c: "foo", d: "bar"}, e: "f", g:{i:"j"}}
/*
sum := import("./sum") 
fmt.println(sum(0, [1, 2, 3, 4, 5]))   // "6"
myfunc := func(n,...v){
    v1 := 0
    v2 := 0
    if(v[0]!=nil){
        v1 = v[0]
    }
    if(v[1]!=nil){
        v2 = v[1]
    }
    return n+v1+v2
}
fmt.println(myfunc(1))
fmt.println(myfunc(1,2))
fmt.println(string(myfunc(1,2,3))+"s")
n := 5
for n>1 {
    n--
}
fmt.println(n)
fmt.println(time.time_format(time.now(),"2006-01-02 15:04:05"))
fmt.println(time.time_unix(time.now()))
fmt.println(time.time_unix_nano(time.now()))
rand.seed(time.time_unix_nano(time.now()))
fmt.println(rand.intn(5))
fmt.println(text.contains("hello,world",","))
fmt.println(text.split("hello,world",","))
fmt.println(os.read_file("test.txt"))
request := http.request("POST","http://localhost:6666/php.php")
request.set_timeout(30)
request.set_header("User-Agent","Mozilla/5.0 (Windows NT 10.0; WOW64; Trident/7.0; rv:11.0) like Gecko")
//request.set_header("Content-Type","application/x-www-form-urlencoded")
request.set_header("Content-Type","application/json")
request.set_body({"test":1})
response := request.get_response()
fmt.println(response.get_body())
fmt.println(crypto.md5_file("test.txt"))
file := os.stat("test.txt")
if is_error(file) {
    fmt.println(file.value)
}
fmt.println(arr)
fmt.println(delete(arr,1))
fmt.println(arr)
fmt.println(map)
fmt.println(delete(map,"b"))
fmt.println(array_sort([2,1,4,3,3,5]))
fmt.println(map)
fmt.println(exists(map,"f"))
fmt.println(exists(map,[1, 2, 3]))
fmt.println(exists(map,{i:"j"}))
fmt.println(arr)
fmt.println(exists(arr,"foo"))
fmt.println(exists(arr,[1, 2, 3]))
fmt.println(exists(arr,{i:"j"}))
*/
fmt.println(array_column([{a:"1",b:"2"},{a:"2",b:"3"},{a:"4",b:"5"}],"c"))
fmt.println(array_unique(["a", "b", "b", 1,["a"],{b:"c"},1,["a"]]))

