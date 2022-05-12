package beverage

type Topping struct {
	Name  string
	Price int
}

// トッピング一覧
var (
	Milk  = Topping{Name: "ミルク", Price: 100}
	Mocha = Topping{Name: "モカ", Price: 150}
	Soy   = Topping{Name: "豆乳", Price: 80}
	Whip  = Topping{Name: "ホイップ", Price: 120}
)
