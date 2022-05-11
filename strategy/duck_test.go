package strategy_test

import (
	"fmt"
	"testing"

	"github.com/miyamoto-jo/hfdp/strategy"
)

// いろんな種類のカモの挙動を確認
func TestDucks(t *testing.T) {
	ducks := []strategy.Duck{
		*strategy.MallardDuck, *strategy.RedHeadDuck, *strategy.RubberDuck, *strategy.DecoyDuck,
	}

	for _, d := range ducks {
		d.DisplayName()
		d.Swim()
		d.PerformFly()
		d.PerformQuack()
		fmt.Println("-------------------------------------------------------")
	}
}

type FlyRocket struct{}

func (m *FlyRocket) Fly() {
	fmt.Println("ロケットで飛ぶよ")
}

type RoboQuack struct{}

func (m *RoboQuack) Quack() {
	fmt.Println("ｶﾞｰ...ｶﾞｰ...")
}

// 独自のカモインスタンスを生成
func TestNewDuck(t *testing.T) {
	roboDuck := strategy.New(`ロボットカモ`, &RoboQuack{}, &FlyRocket{})

	roboDuck.DisplayName()
	roboDuck.Swim()
	roboDuck.PerformFly()
	roboDuck.PerformQuack()
}
