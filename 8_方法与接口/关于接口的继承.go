package main

type Binterface interface {
	test01()
}
type Cinterface interface {
	test02()
}
type Ainterface interface {
	Binterface
	Cinterface
	test03()
}

type Stu struct {
}

func (stu Stu) test01() {

}
func (stu Stu) test02() {

}
func (stu Stu) test03() {

}

func main() {
	var stu Stu
	var A Ainterface = stu
	A.test03()
}
