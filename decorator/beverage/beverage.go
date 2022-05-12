package beverage

type Beverage struct {
	Name      string
	Price     int
	FullPrice int
	Toppings  []*Topping
}

// menu
var (
	HouseBland = Beverage{Name: `ハウスブレンド`, Price: 500, FullPrice: 500}
	DarkRoast  = Beverage{Name: `ダークロースト`, Price: 550, FullPrice: 550}
	Espresso   = Beverage{Name: `エスプレッソ`, Price: 450, FullPrice: 450}
	Decaf      = Beverage{Name: `カフェイン抜き`, Price: 450, FullPrice: 450}
)
