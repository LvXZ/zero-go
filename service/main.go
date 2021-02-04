package main

import "fmt"

// Sayer 说话接口
type Sayer interface {
	say()
}

// Player 玩耍接口
type Player interface {
	play()
}

// Student 学生结构体
type Student struct {
	name string
}

// say 学生实现 Sayer说话接口
func (stu Student) say() {
	fmt.Printf("%s 正在说话\n", stu.name)
}

// say 学生实现 Player玩耍接口
func (stu Student) play() {
	fmt.Printf("%s 正在玩耍\n", stu.name)
}


func main() {

	student := Student{name:"小明"}
	student.play()
	student.say()

	student2 := Student{name:"小红"}

	var x Sayer
	var y Player

	x = student
	y = student2

	x.say()
	y.play()
}
