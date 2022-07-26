iter
===
迭代器及相关操作工具。

迭代器模式的线程安全性由两部分影响：适配器/终结器/收集器的线程安全性和被迭代集合类型的线程安全性。
绝大多数场景下，迭代器不应在 goroutine 间共享，但迭代的集合对象可能会被共享。
当前模块只提供迭代器接口及操作函数，不会约束被迭代集合类型及其迭代器实现的线程安全性。
无法保证迭代集合对象及其迭代器实现的线程安全性的前提下，仅保证操作函数的线程安全性是徒劳的。
因此当前模块提供的所有函数均不对线程安全性做任何保证，并发使用时需要调用方自行进行线程安全性保证。

迭代器一旦被迭代，之后就不应再被使用。因 Go 1.18 泛型机制仅支持泛型类型和泛型函数，不支持泛型方法，
本模块无法提供如 Java、Rust 语言中Stream、迭代器相同的链式调用，仅能对同一迭代器变量进行重复赋值：
```go
it := CountTo(0, 1, 9)
it = Map(it, func(v int) int { return v * 2 })
sum, _ := Sum(it)
fmt.Printf("sum: %d\n", sum) // 90
_, ok := Sum(it) // Sum on an iterated iterator, DO NOT do this in practice.
fmt.Printf("ok? %t", ok) // ok? false
```
最后使用后的迭代器其实已经被"耗尽"，或者状态已经发生变化。
如果再次迭代该迭代器，迭代结果是无意义或不确定的。

# 生成器
* [x] Counter
* [x] Slice
* [x] FromFunc
# 适配器
* [x] Chain
* [x] Filter
* [x] FlatMap
* [x] Inspect
* [x] Map
* [x] Distinct
* [x] DistinctBy
# 终结器
* [x] All
* [x] Any
* [x] Count
* [x] Each
* [x] Last
* [x] Max, MaxBy
* [x] Min, MinBy
* [x] Mean
* [x] Reduce
* [x] Sum
# 收集器（其实是终结器的一种）
* [x] ToSlice
* ...