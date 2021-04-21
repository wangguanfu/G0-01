### go训练营

#### 二.goroutine

1.Goroutine

2.Memory model

3.Package sync

4.chan

5.Package context

6.Referneces



###### 1.1

process and thread 

##### 进程是操作系统资源分配的基本单位，而线程是处理器任务调度和执行的基本单位

1.进程 ：资源运行的容器 。 资源分配的

2.线程：操作系统调度的一个执行路径。

##### Goroutines and Parallelism

Go 语言层面支持的 go 关键字，可以快速的让一个函数创建为 goroutine，我们可以认为 main 函数就是作为 goroutine 执行的。操作系统调度线程在可用处理器上运行，Go运行时调度 goroutines 在绑定到单个操作系统线程的逻辑处理器中运行(P)。即使使用这个单一的逻辑处理器和操作系统线程，也可以调度数十万 goroutine 以惊人的效率和性能并发运行。



###### 并发的关键是你有处理多个任务的能力，不一定要同时。

###### 并行的关键是你有同时处理多个任务的能力。

##### 1.2. 管理goroutine的生命周期

1.when will it terminate?  什么时候 

2.what could prevent it from terminating? 如何让他结束





































