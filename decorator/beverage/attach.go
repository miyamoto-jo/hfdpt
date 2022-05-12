package beverage

type AttachFunc func(*Beverage) *Beverage

// トッピングを付け加える関数を返す
func getAttachMethod(t *Topping) AttachFunc {
	return func(b *Beverage) *Beverage {
		b.Toppings = append(b.Toppings, t)
		b.FullPrice += t.Price
		return b
	}
}

// トッピングを付ける
func AttachToppings(b *Beverage, toppings ...*Topping) {
	funcs := []AttachFunc{}
	for _, t := range toppings {
		funcs = append(funcs, getAttachMethod(t))
	}
	for _, f := range funcs {
		b = f(b)
	}
}
