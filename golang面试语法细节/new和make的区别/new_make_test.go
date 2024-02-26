package new和make的区别

// Go 语⾔当中 new 和 make 有什么区别吗？

// new的作⽤是初始化⼀个纸箱类型的指针 new函数是内建函数，

// 使⽤ new函数来分配空间
// 传递给 new函数的是⼀个类型，⽽不是⼀个值
// 返回值是指向这个新分配的地址的【指】针

// func new(Type) *Type

// make的作⽤是为 slice, map or chan 的初始化然后返回【引⽤】 make函数是内建函数，

// func make(Type, size IntegerType) Type
