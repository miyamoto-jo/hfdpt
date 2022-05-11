package strategy

import "fmt"

type FlyBehavior interface {
	Fly()
}

type FlyWithWings struct{}

func (m *FlyWithWings) Fly() {
	fmt.Println("私は飛ぶ")
}

type FlyNoWay struct{}

func (m *FlyNoWay) Fly() {
	fmt.Println("私は飛べません")
}
