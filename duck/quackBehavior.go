package duck

import "fmt"

type QuackBehavior interface {
	Quack()
}

type Quack struct{}

func (m *Quack) Quack() {
	fmt.Println("ガーガー")
}

type Squeak struct{}

func (m *Squeak) Quack() {
	fmt.Println("キューキュー")
}

type MuteQuack struct{}

func (m *MuteQuack) Quack() {
	fmt.Println("...(沈黙)")
}
