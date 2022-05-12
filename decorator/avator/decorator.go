// デコレーター
package avator

type DecorateMethod func(*Avator) *Avator

// 武器を装備する
func AttachWeapons(a *Avator, weapons ...*Weapon) {
	funcs := []DecorateMethod{}
	for _, w := range weapons {
		funcs = append(funcs, getAttachWeaponMethod(w))
	}
	for _, f := range funcs {
		a = f(a)
	}
}

// 防具を装備する
func AttachArmor(a *Avator, armors ...*Armor) {
	funcs := []DecorateMethod{}
	for _, a := range armors {
		funcs = append(funcs, getAttachArmorMethod(a))
	}
	for _, f := range funcs {
		a = f(a)
	}
}

// 防具を装着するメソッドを取得
func getAttachArmorMethod(a *Armor) DecorateMethod {
	return func(h *Avator) *Avator {
		h.Armors = append(h.Armors, a)
		h.hp += a.Value
		return h
	}
}

// 武器を装着するメソッドを取得
func getAttachWeaponMethod(w *Weapon) DecorateMethod {
	return func(h *Avator) *Avator {
		h.Weapons = append(h.Weapons, w)
		h.attack += w.Value
		return h
	}
}
