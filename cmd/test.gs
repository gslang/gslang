fmt := import("fmt")
sum := import("./sum") 

fmt.println(sum(0, [1, 2, 3, 4, 5]))   // "6"
arr := ["foo", "bar", [1, 2, 3]]
map := {a: [1,2,3], b: {c: "foo", d: "bar"}}
for k,v in arr{
    fmt.println(k,v)
}
for k,v in map{
    fmt.println(k,v)
}
if arr[4]==nil {
    fmt.println("nil")
}
fmt.println(type(arr))
for k,v in range(10,100) {
    fmt.println(k,v)
}