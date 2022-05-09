package duck

import "fmt"

type Duck struct {
	name          string
	quackBehavior QuackBehavior // 鳴く振る舞い
	flyBehavior   FlyBehavior   // 飛ぶ振る舞い
}

func New(name string, q QuackBehavior, f FlyBehavior) *Duck {
	return &Duck{
		name:          name,
		quackBehavior: q,
		flyBehavior:   f,
	}
}

// 名前を表示
func (m *Duck) DisplayName() {
	fmt.Println(m.name)
}

// 泳ぐ
func (m *Duck) Swim() {
	fmt.Println("泳ぐよ〜")
}

// 鳴く
func (m *Duck) PerformQuack() {
	m.quackBehavior.Quack()
}

// 飛ぶ
func (m *Duck) PerformFly() {
	m.flyBehavior.Fly()
}

// 鳴く振る舞いをセット
func (m *Duck) SetQuackBehavior(q QuackBehavior) {
	m.quackBehavior = q
}

// 飛ぶ振る舞いをセット
func (m *Duck) SetFlyBehavior(f FlyBehavior) {
	m.flyBehavior = f
}
