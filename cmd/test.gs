fmt := import("fmt")
sum := import("./sum") 

fmt.println(sum(0, [1, 2, 3, 4, 5]))   // "6"
fmt.println(sum("", [1, 2, 3, 4]))  // "123"