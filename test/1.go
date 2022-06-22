package main

import "fmt"

//判断1-8000中哪些数是素数
func put(intChan chan int){
	for i:=2;i<=1000;i++{
		intChan<-i
	}
	close(intChan) //放完8000个元素后关闭通道
}
//一个是元数据通道   一个是存放素数通道  一个是判断是否结束通道
func primNum(intChan chan int ,primChan chan int , exitChan chan bool){
	for{
		num,ok:=<-intChan
		if !ok{
			break
		}
		flag:=true //是素数      简短声明的变量只能在函数内部使用
		for i:=2;i*i<=num;i++{
			if num%i==0{
				flag=false //不是素数
				break
			}
		}
		if flag{
			//如果是素数就将结果放入到 通道
			primChan<-num
		}
	}
	//已经读取完原始数据
	exitChan<-true  //向里面写一个true
}
func main(){
	intChan :=make(chan int ,100)
	exitChain:=make(chan bool ,4)
	primChain:=make(chan int ,2000)
	go put(intChan)
	for i:=0;i<4;i++{
		go primNum(intChan,primChain,exitChain) //开四个协程，但是素数队列里的顺序肯定不是从小到大 
	}
	go func(){
		for i:=0;i<4;i++{
			<-exitChain   //为什么 exitchain 不关闭没有问题  是因为不从里面取值吗
		}
		close(primChain)
	}()
	for {
		res,ok:=<-primChain
		if !ok{
			break
		}
		fmt.Println(res)
	}
	fmt.Println("结束")

}