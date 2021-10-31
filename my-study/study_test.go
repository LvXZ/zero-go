package my_study

import (
	"context"
	"fmt"
	"golang.org/x/time/rate"
	"log"
	"testing"
	"time"
)

// @Author: lvxiaozheng
// @Date: 2021/7/15 10:26
// @Description:

func TestMyStudy(t *testing.T) {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}

//limit表示每秒产生token数，Bucket最多存token数
//Allow判断当前是否可以取到token
//Wait阻塞等待知道取到token
//Reserve返回等待时间，再去取token
func TestMyStudy2(t *testing.T) {
	l := rate.NewLimiter(1, 5)
	log.Println(l.Limit(), l.Burst())
	for i := 0; i < 100; i++ {
		//阻塞等待直到，取到一个token
		log.Println("before Wait")
		c, _ := context.WithTimeout(context.Background(), time.Second*2)
		if err := l.Wait(c); err != nil {
			log.Println("limiter wait err:" + err.Error())
		}
		log.Println("after Wait")

		//返回需要等待多久才有新的token,这样就可以等待指定时间执行任务
		r := l.Reserve()
		log.Println("reserve Delay:", r.Delay())

		//判断当前是否可以取到token
		a := l.Allow()
		log.Println("Allow:", a)
	}
}

type dog struct {
	name string
}

type cat struct {
	name string
}

type duck struct {
	name string
}

// 定义接口
type speaker interface {
	speak() //只要实现了speak()方法的变量，都是speaker的类型
}

// dog叫
func (d dog) speak() {
	fmt.Printf("%s:汪\n", d.name)
}

// cat叫
func (c cat) speak() {
	fmt.Printf("%s:喵\n", c.name)
}

// duck叫
func (d duck) speak() {
	fmt.Printf("%s:嗄\n", d.name)
}

// 打
func da(s speaker) {
	//接受参数，打动物
	s.speak()
}

// 结构体
func TestMyStruct(t *testing.T) {

	d := dog{
		name: "小狗",
	}
	c := cat{
		name: "小猫",
	}
	du := duck{
		name: "小鸭",
	}

	da(d)
	da(c)
	da(du)
}
