package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

//需要注意的是 { 不能单独放在一行，否则出错
//如果声明了变量不使用，也会标红
func main() {
	fmt.Println("hello")
	str:="go 语言"
	index:=strings.Index(str,"语")
	fmt.Println(index)
	str1:=str[index:]
	fmt.Println(str1)
	bytes:=[]byte(str)
	for i:=0;i<3;i++ {
		bytes[i]='1'
	}
	fmt.Println(string(bytes))
	str2:=fmt.Sprintf("%s  是啥时候  %s",str,str1)
	fmt.Println(str2)
	//builder 和buffer都没东西输出啊
	var cc strings.Builder
	cc.WriteString(str)

	fmt.Println(10)
	fmt.Sprintf(cc.String())

	fmt.Println(10)
	const (
		a=iota
		c="hello word"
		d
		e=iota
	)
	fmt.Println(a,b,c,d,e)
	pyramid(9)
	//var nums=[...]int{3,4,5,6,7}
	nums:=[]int{2,4,6,8,10}
	//range也太强大了吧
	for k,v :=range nums{
		fmt.Println(k,v)
	}
	for index:=0;index< len(nums);index++ {
		fmt.Println(nums[index],"###")
	}
	//切片
	nums1:=nums[4:5]
	nums1[0]=1
	fmt.Println(nums[4])
	fmt.Println(&nums[3])
	fmt.Println(&nums[4])
	fmt.Println(&nums1[0])
	//在不超过切片容量的前提下，切片引用的还是原数组，当扩容之后，就会重新cope，创建新的数组
	//开始给切片扩容
	nums1=append(nums1,12)
	fmt.Println("切片长度",len(nums1))
	fmt.Println("容量", cap(nums1))
	fmt.Println(&nums[4])
	fmt.Println(&nums1[0])
	//声明切片
	var student []int
	student=make([]int ,2,10)
	fmt.Println(student)
	//map使用
	var mp =map[string]int{
		"tom":80,
		"ben":90,
	}
	fmt.Println(mp)
	var studentmap map[string]int
	studentmap=make(map[string]int)
	studentmap["ma"]=90
	studentmap["peng"]=45
	studentmap["hui"]=60
	for k,v:=range studentmap{
		fmt.Println(k,v)
	}
	//随机的
	for k,v:=range studentmap{
		fmt.Println(k,v)
	}
	fmt.Println(studentmap["penghui"])
	delete(studentmap,"ma")
	fmt.Println(studentmap)
	a1,_:=mutireturn()
	_,b1:=mutireturn()
	fmt.Println(a1,b1)
	ListDir("C:\\Users")
	ReadFile("C:\\Users\\Penghui Ma\\Desktop\\test.txt")
	idb :=&Redis{DBName: "mapenghui"}
	var  iredis IRediser
	iredis=idb
	iredis.Connect()
	idb.Disconnect()
}
func pyramid(n int){
	for i:=1;i<=n;i++ {
		for j:=1;j<=n-i;j++ {
			fmt.Print(" ")
		}
		for k:=1;k<=2*i-1;k++ {
			fmt.Print("*")
		}
		fmt.Println()

	}
}
func mutireturn() (int,int) {
	return 10,20
}
//小驼峰命名法    表示标识符仅在本包内可以使用 ，不可导出
func ListDir(dirPath string) error  {
	dir,err := ioutil.ReadDir(dirPath)//"C:\\Users"
	if err!=nil {
		return err
	}
	for _,file :=range dir{
		if file.IsDir() {
			fmt.Println("目录："+file.Name())
		} else{                   //else 只能紧贴这个右括号这一行，否则就报错，很离谱
			fmt.Println("文件："+file.Name())
		}


	}
	return nil
}
//C:\Users\Penghui Ma\Desktop\test.txt
func ReadFile(path string)  {
	file,err:=os.Open(path)
	if err!=nil {
		fmt.Println(err)
	}
	buf :=make([]byte,1024)
	fmt.Println("以下是文件内容：")
	for {
		len,_:=file.Read(buf)  //不明白为啥这里还要写个for循环，一次不就读取结束了吗
		if len==0{
			break
		}
		fmt.Println(string(buf))
	}
	file.Close()
}
//接口类型的实现操作
type IDtabaser interface {
	Connect() error
	Disconnect() error
}
type IRediser interface{
	Connect() error
}
type Redis struct {
	DBName string
}

func (redis *Redis ) Connect() error  {
	fmt.Println("Redis Connect DB=> "+redis.DBName)
	fmt.Println("Redis Connect Success!")
	return nil
}
func (redis *Redis) Disconnect() error{
	fmt.Println("Redis Disconnect Success!")
	fmt.Println(time.Now().Format("15:04:05"))
	return nil
}
//类型断言
//t,ok:X.(T)
//并发与通道    协程（coroutine）通道（channel）

