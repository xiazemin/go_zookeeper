package main
import(
	"fmt"
	"time"
	"runtime"
)
var num=14  //定义一工并发多少数量
var cnum chan int
func main(){
	maxProcs := runtime.NumCPU()// 获取cpu个数
	runtime.GOMAXPROCS(maxProcs)//限制同时运行的goroutines数量
	cnum=make(chan int,num) //make一个chan,缓存为num
	for i:=0;i<num;i++{
		go Printer(i)
	}
	// 下面这个for循环的意义就是利用信道的阻塞，一直从信道里取数据，直到取得跟并发数一样的个数的数据，则视为所有goroutines完成。
	for i:=0;i<num;i++{
		<-cnum
	}
	fmt.Println("WE DONE!!!")
}



func Printer(a int)(){
	time.Sleep(2000 * time.Millisecond)
	fmt.Printf("i am %d\n",a)
	cnum <- 1  //goroutine结束时传送一个标示给信道。
}