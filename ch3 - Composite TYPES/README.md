# Composite Types

## Arrays
* rarely used
  * go thinks that the type of the array includes its length, so `[3]int` is different than `[4]int`
  * cannot use a variable to specify the array size, it needs to be resolved at compile time
* it has `==` and `!=`
* it is only used when you know the size ahead of time, and it does not changes.


## Slices
* length is not a part of the type
  * you can create a function that processes slices of any size
* it starts with the value of `nil`
  * it represents the lack of a value for some types.
* Cannot compare slices with `==`, only with `nil`
* `append` is used to grow slices
* you expand the slice using `...`
  * you use it to concat two slices
  
### Capacity
* number of consecutive memory locations reserved. It can be larger than the length.

> the Go runtime usually increases a slice by more than one each time it runs out of capacity. The rules as of Go 1.14 are to double the size of the slice when the capacity is less than 1,024 and then grow by at least 25% afterward.

### Make
* create a slice with an already defined size
* you can also pass the capacity
* it starts with the default zero value
* 

### Slicing
```golang
y := x[2:5]
```

* Slices are not copies, they are sharing the same storage (sometimes)
  * because of that, we need to set the third argument of the slice, which means the slice's capacity


Stopped: https://learning.oreilly.com/library/view/learning-go/9781492077206/ch03.html#:-:text=copy
