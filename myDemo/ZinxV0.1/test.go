package main

type T1 interface {
	Cat()
}

type T2 struct {
	A int
}

func (t2 *T2) Cat() {

}

func (t2 *T2) Dog() {

}

func main() {
	var t1 T1 = &T2{
		A: 1,
	}
}
