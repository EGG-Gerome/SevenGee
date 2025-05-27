package main

import (
	"fmt"
	"strconv"
	"time"
)

//func main() {
//var str1 string = "clever Y"
//str2 := "so so so clever"
//// 查看语言类型 reflect.TypeOf().Kind()
//fmt.Println(reflect.TypeOf(str1).Kind())    // string
//fmt.Println(reflect.TypeOf(str1[2]).Kind()) // unit8在go里面等价于byte,rune在go里面等价于int32
//fmt.Println(str2[1], string(str2[1]))       // 111 o
//fmt.Printf("%d %c\n", str2[1], str2[1])     // 111 	o
//fmt.Println("len(str2)", len(str2))         // len(str2) 15

// rune array
//str2 := "语言类型"
//runeArr := []rune(str2)
//fmt.Println(reflect.TypeOf(str2[0]).Kind())    // unit8
//fmt.Println(reflect.TypeOf(runeArr[0]).Kind()) // int32(rune)
//fmt.Println(runeArr[0], string(runeArr[0]))    // 语	语
//fmt.Println("len(runeAr)r：", len(runeArr))

// slice
//slice1 := make([]float32, 3) // 长度为3的切片
//slice2 := make([]float32, 3, 5)       // [0 0 0] 长度为3，容量为5的切片
//fmt.Println(len(slice2), cap(slice2)) // 3 5
//var slice3 []float32 = make([]float32, 6, 10)
//fmt.Println(len(slice3), cap(slice3))
// append elements
//slice1 = append(slice1, 1, 2, 3, 4)
////fmt.Println(len(slice1), cap(slice1)) // 7(3+4)
//// subSlice [start, end)
//sub1 := slice1[3:]
//sub2 := slice1[:3]
//sub3 := slice1[1:4] // [0 0 1]
//combined := append(sub1, sub2...)
//fmt.Println(sub3)
//fmt.Println(combined) // [1 2 3 4 0 0 0]

// map
// 仅声明
//m1 := make(map[string]int)
//// 声明时初始化
//m2 := map[string]string{
//	"Emily":  "Girl",
//	"Gerome": "Handsome Boy",
//}
//// 赋值/修改
//m1["Egg"] = 3
//m2["Emily"] = "Pretty Girl"
//fmt.Println(m1["Egg"])
//fmt.Println("little Y is a" + m2["Emily"])
//fmt.Println("Gerome is a" + m2["Gerome"])

// Pointer
//str1 := "Han han Y"
//var p *string = &str1
//fmt.Println(str1)
//fmt.Println(*p)
//*p = "Ha?"
//fmt.Println(*p)

//fmt.Println(get(5))
//fmt.Println("finished")

//}

//// defer and recover
//func get(index int) (ret int) {
//	// 在协程退出前，会执行完 defer 挂载的任务。因此如果触发了 panic，控制权就交给了 defer。
//	defer func() {
//		if r := recover(); r != nil {
//			// 在 defer 的处理逻辑中，使用 recover，使程序恢复正常，并且将 ret 设置为 -1，
//			// 在这里也可以不处理返回值，如果不处理返回值，返回值将被置为默认值 0。
//			fmt.Println("Some error happened!", r)
//			ret = -1
//		}
//	}()
//	arr := [3]int{2, 3, 4}
//	return arr[index]
//}

//// interface
//type Person interface {
//	getHobby()
//	getGender()
//	getName() string
//}
//
//type Egg struct {
//	name  string
//	hobby string
//	age   int
//}
//
//func (E *Egg) getHobby() {
//	fmt.Println(E.name + "'s hobby is " + E.hobby)
//}
//func (e *Egg) getName() string {
//	return e.name
//}
//func (e *Egg) getGender() {
//	fmt.Println("Ha?! I am a EGG!!!")
//}
//
//type Student struct {
//	name   string
//	gender string
//	id     int
//}
//
//func (s *Student) getName() string {
//	return s.name
//}
//func (s *Student) getHobby() {
//	fmt.Println(s.name + "'s hobby is studying") // 示例实现
//}
//func (S *Student) getGender() {
//	fmt.Println(S.name + "'s gender is " + S.gender)
//}
//
//func main() {
//	var p Person = &Egg{
//		name:  "littleEgg",
//		hobby: "eating",
//		age:   2,
//	}
//	p.getHobby()
//	p = &Student{
//		name:   "Gerome",
//		gender: "boy",
//		id:     01,
//	}
//	p.getGender()
//	p.getHobby()
//	// 我们如何确保某个类型实现了某个接口的所有方法呢？一般可以使用下面的方法进行检测， 如果实现不完整，
//	//编译期将会报错。 var _ Person = (*Student)(nil)     var _ Person = (*Worker)(nil)
//  // 将空值 nil 转换为 *Student 类型，再转换为 Person 接口，如果转换失败，说明 Student 并没有实现 Person 接口的所有方法。
//  实例可以强制类型转换为接口，接口也可以强制类型转换为实例。
//
//	var p Person = &Student{
//		name: "Tom",
//		age:  18,
//  }
//	stu := p.(*Student) // 接口转为实例
//	fmt.Println(stu.getAge())
//}

// empty interface can present anyType
// type anyType interface {
// }
//
//	func main() {
//		//m := make(map[string]interface{})
//		m := make(map[string]anyType)
//		m["name"] = "Tom"
//		m["age"] = 18
//		m["scores"] = [3]int{98, 99, 85}
//		fmt.Println(m) // map[age:18 name:Tom scores:[98 99 85]]
//	}

// 并发编程(goroutine)

// sync
//var wg sync.WaitGroup // 声明并初始化一个 WaitGroup 实例，用于协调多个 goroutine 的执行。
//// 内部维护一个计数器，记录尚未完成的 goroutine 数量。
//
//func download(url string) {
//	fmt.Println("start to download ", url)
//	//time.Sleep(time.Second) // 模拟耗时操作
//	fmt.Println("end to download ", url)
//	//time.Sleep(time.Second)
//	wg.Done()
//}
//
//func main() {
//	for i := 0; i < 10; i++ {
//		wg.Add(1)
//		url := "tangyuan.com/0" + strconv.Itoa(i+1)
//		go download(url)
//	}
//	wg.Wait()
//	fmt.Println("Done!")
//}

// channel

var ch = make(chan string, 10) //// 创建大小为 10 的缓冲信道
func download(url string) {
	fmt.Println("start to download" + url)
	time.Sleep(time.Second)
	ch <- url
}
func main() {
	for i := 0; i < 10; i++ {
		go download("tangyuan.com/0" + strconv.Itoa(i+0))
	}
	for i := 0; i < 10; i++ {
		msg := <-ch
		fmt.Println("finish", msg)
	}
	fmt.Println("Done!")
}
