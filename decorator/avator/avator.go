// アバター
package avator

import "fmt"

type Avator struct {
	name    string
	hp      int       // 体力
	attack  int       // 攻撃力
	Weapons []*Weapon // 装備している武器
	Armors  []*Armor  // 装備している防具
}

func New(name string) *Avator {
	return &Avator{
		name:   name,
		hp:     100,
		attack: 10,
	}
}

// 基本情報を出力
func (m *Avator) DispBasicInfo() {
	fmt.Printf("\t名前: %s\n", m.name)
	fmt.Printf("\t攻撃力: %d\n", m.attack)
	fmt.Printf("\t体力: %d\n", m.hp)
}

// 装備品を出力
func (m *Avator) DispDecorateInfo() {
	m.DispWeaponsInfo()
	m.DispArmorsInfo()
}

// 装備武器を出力
func (m *Avator) DispWeaponsInfo() {
	if len(m.Weapons) > 0 {
		fmt.Println("装備武器")
		for _, w := range m.Weapons {
			fmt.Printf("%s ", w.Name)
		}
		fmt.Printf("\n")
	}
}

// 装備防具を出力
func (m *Avator) DispArmorsInfo() {
	if len(m.Armors) > 0 {
		fmt.Println("装備防具")
		for _, a := range m.Armors {
			fmt.Printf("%s ", a.Name)
		}
		fmt.Printf("\n")
	}
}
