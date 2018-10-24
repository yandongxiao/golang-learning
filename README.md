# golang-learning

Features added between releases, available in the source repository but not part of the numbered binary releases,
are under active development. No promise of compatibility is made for software using such features until they have been released.

1. [variable](./variable): 介绍多重赋值、类型转换、**比较**、变量名与类型同名、作用域、零值和常量.
2. [control](./control): 介绍for、 if、 switch 和 type switch.
3. [error](./error): 介绍panic的工作原理，如何与recover结合使用；panic和error使用的边界；repanic.
4. [func](./func): 介绍init函数以及初始化顺序, 可变参数、命名的返回值, closure, 高阶函数
5. [struct](./struct): 介绍继承、封装、匿名内部类等属性.
6. [type](./type): 介绍map, slice, array, string等数据结构
7. [interface](./interface): 介绍interface的工作原理，比较方法，多态，type assertion等.
8. [chan](./chan): 介绍chan的各方面属性，如何使用routine+chan实现协程间同步，以及常见的同步手段.
9. [标准库](./lib): 介绍golang各个模块的使用方法, 例如sync, io, os,strings等
