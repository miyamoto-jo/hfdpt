package avator_test

import (
	"fmt"
	"testing"

	"github.com/miyamoto-jo/hfdp/decorator/avator"
)

func TestAvator(t *testing.T) {
	soldier := avator.New("戦士")

	wepons := []*avator.Weapon{
		&avator.WeaponA, &avator.WeaponC, &avator.WeaponD,
	}
	armors := []*avator.Armor{
		&avator.ArmorD, &avator.ArmorC, &avator.ArmorF,
	}

	// 武器を装備
	avator.AttachWeapons(soldier, wepons...)
	// 防具を装備
	avator.AttachArmor(soldier, armors...)

	display(soldier)
}

func display(a *avator.Avator) {
	fmt.Println("アバター情報")
	a.DispBasicInfo()
	a.DispDecorateInfo()
}
