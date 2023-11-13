package main

import "fmt"

type Usb interface {
	start()
	stop()
}

//手机
type Phone struct {
}

func (p Phone) start() {
	fmt.Println("手机开始工作了")
}
func (p Phone) stop() {
	fmt.Println("手机停止工作了")
}

//键盘
type Jianpan struct {
}

func (p Jianpan) start() {
	fmt.Println("键盘开始工作了")
}
func (p Jianpan) stop() {
	fmt.Println("键盘停止工作了")
}

type Computer struct {
}

func (C Computer) working(usb Usb) {
	usb.start()
	usb.stop()
}

func main() {
	c := Computer{}
	phone := Phone{}
	Jianpan := Jianpan{}
	c.working(phone)
	c.working(Jianpan)

}
