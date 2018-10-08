package main
import(
	"fmt"
	"sync"
	"time"
	"runtime"
)
var wg sync.WaitGroup  //定义一个同步等待的组
func main() {
	maxProcs := runtime.NumCPU()   //获取cpu个数
	runtime.GOMAXPROCS(maxProcs)  //限制同时运行的goroutines数量
	for i:=0;i<10;i++{
		wg.Add(1)//为同步等待组增加一个成员
		go Printer(i)//并发一个goroutine
	}
	wg.Wait() //阻塞等待所有组内成员都执行完毕退栈
	fmt.Println("WE DONE!!!")
}
//定义一个Printer函数用于并发
func Printer(a int)(){
	time.Sleep(2000 * time.Millisecond)
	fmt.Printf("i am %d\n",a)
	defer wg.Done()
}