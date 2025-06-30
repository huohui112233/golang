package main

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	//1.1-编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，
	// 在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
	// 考察点 ：指针的使用、值传递与引用传递的区别。
	num := 12
	ptr := &num
	num = addNum(ptr)
	// 1.2-实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
	//考察点 ：指针运算、切片操作
	numbers := []int{1, 2, 3, 4, 5}
	sliceNumber(&numbers)
	fmt.Println("切片乘以2后结果为：", numbers)

	// 1.3-编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
	//考察点 ： go 关键字的使用、协程的并发执行。
	var wg sync.WaitGroup
	wg.Add(2)
	go even(&wg)
	go odd(&wg)
	wg.Wait()

	// 1.4-设计一个任务调度器，接收一组任务（可以用函数表示），
	// 并使用协程并发执行这些任务，同时统计每个任务的执行时间。
	//考察点 ：协程原理、并发任务调度。
	tasks := []Task{
		func() error {
			time.Sleep(500 * time.Millisecond)
			return nil
		},
		func() error {
			time.Sleep(1 * time.Second)
			return nil
		},
		func() error {
			time.Sleep(300 * time.Millisecond)
			return nil
		},
	}

	scheduler := NewScheduler(tasks, 2)
	results := scheduler.Run()
	// 打印统计结果
	for i, result := range results {
		fmt.Printf("Task %d took %v\n", i+1, result.Durtion)
	}

	// 1.5-定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。
	// 然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，
	// 并调用它们的 Area() 和 Perimeter() 方法
	r := Rectangle{a: 3, b: 2}
	fmt.Println("矩形面积为：", r.Area())
	fmt.Println("矩形周长为：", r.Perimeter())
	c := Circle{radius: 3.0}
	fmt.Println("圆形面积为：", c.Area())
	fmt.Println("圆形周长为：", c.Perimeter())

	// 1.6-使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，
	// 组合 Person 结构体并添加 EmployeeID 字段。
	// 为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
	e := Employee{Person{Name: "zhangsan", Age: 18}, 123}
	e.PrintInfo()

	//1.7-编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，
	// 并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来
	ch := make(chan int)
	sendReceive(ch)

	// 1.8-实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
	buffer := make(chan int, 10)
	sendReceiveBuffer(buffer)

	// 1.9-编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。
	// 启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值
	incrementMutex()

	// 1.10-使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，
	// 每个协程对计数器进行1000次递增操作，最后输出计数器的值。
	incrementAtomic()
}

type Task func() error

type TaskResult struct {
	ID      int
	Durtion time.Duration
}

type Scheduler struct {
	taskList   []Task
	concurrent int
}

func NewScheduler(taskList []Task, concurrent int) *Scheduler {
	return &Scheduler{taskList: taskList, concurrent: concurrent}
}

func (s *Scheduler) Run() []TaskResult {
	results := make([]TaskResult, len(s.taskList))
	resultChan := make(chan TaskResult, len(s.taskList))
	var wg sync.WaitGroup
	wg.Add(len(s.taskList))
	for i, task := range s.taskList {
		go func(id int, t Task) {
			defer wg.Done()
			start := time.Now()
			t()
			duration := time.Since(start)
			resultChan <- TaskResult{
				ID:      id,
				Durtion: duration,
			}
		}(i, task)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for result := range resultChan {
		results[result.ID] = result
	}

	return results

}

func addNum(ptr *int) int {
	return *ptr + 10
}

func sliceNumber(sliceNum *[]int) bool {
	for i := range *sliceNum {
		(*sliceNum)[i] *= 2
	}
	return true
}

func even(waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	for i := 2; i <= 10; i += 2 {
		fmt.Println("偶数：", i)
	}
}

func odd(waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	for i := 1; i <= 10; i += 2 {
		fmt.Println("奇数：", i)
	}
}

type Shape interface {
	Area() int
	Perimeter() int
}

type Rectangle struct {
	a int
	b int
}

type Circle struct {
	radius float64
}

func (r *Rectangle) Area() int {
	return r.a * r.b
}

func (c *Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r *Rectangle) Perimeter() int {
	return 2 * (r.a + r.b)
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID int
}

func (e Employee) PrintInfo() {
	fmt.Println("员工姓名：", e.Name)
	fmt.Println("员工年龄：", e.Age)
	fmt.Println("员工ID：", e.EmployeeID)
}

func sendReceive(ch chan int) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		defer close(ch)
		for i := 1; i <= 10; i++ {
			ch <- i
			time.Sleep(100 * time.Millisecond)
		}
	}()
	go func() {
		defer wg.Done()
		for i := range ch {
			fmt.Println("通道消费消息:", i)
			// time.Sleep(100 * time.Millisecond)
		}
	}()
	wg.Wait()
	fmt.Println("通道结束")
}

func sendReceiveBuffer(buffer chan int) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		defer close(buffer)
		for i := 1; i < 20; i++ {
			buffer <- i
			time.Sleep(1000 * time.Millisecond)
		}
	}()

	go func() {
		defer wg.Done()
		for i := range buffer {
			fmt.Println("带缓冲通道消息：", i)
		}
	}()

	wg.Wait()
	fmt.Println("带缓冲执行结束")

}

type Counter struct {
	mutex sync.Mutex
	count int
}

func (c *Counter) increment() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.count++
}

func incrementMutex() {
	c := Counter{}
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				c.increment()
			}
		}()
	}

	// 等待所有协程完成
	wg.Wait()
	// 输出最终结果
	fmt.Printf("最终计数器值: %d\n", c.count)
}

type CountAtomic struct {
	count int64
}

func (ca *CountAtomic) increment() {
	atomic.AddInt64(&ca.count, 1)
}

func (ca *CountAtomic) getValue() int64 {
	return atomic.LoadInt64(&ca.count)
}

func incrementAtomic() {
	ca := CountAtomic{}
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				ca.increment()
			}
		}()
	}
	wg.Wait()
	fmt.Println("打印最终结果为：", ca.getValue())
}
