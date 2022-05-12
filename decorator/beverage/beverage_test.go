package beverage_test

import (
	"fmt"
	"testing"

	"github.com/miyamoto-jo/hfdp/decorator/beverage"
)

func TestCoffeeShop(t *testing.T) {
	drink := beverage.HouseBland     // 注文されたドリンク
	toppings := []*beverage.Topping{ // 注文されたトッピング
		&beverage.Milk, &beverage.Mocha, &beverage.Soy,
	}
	//トッピングを追加
	beverage.AttachToppings(&drink, toppings...)

	display(&drink)
}

func display(drink *beverage.Beverage) {
	fmt.Printf("注文したドリンク\n\t%s/%d円\n", drink.Name, drink.Price)
	if len(drink.Toppings) > 0 {
		fmt.Println("トッピング一覧")
		for _, t := range drink.Toppings {
			fmt.Printf("\t・%s/%d円\n", t.Name, t.Price)
		}
	}
	fmt.Printf("料金：%d円\n", drink.FullPrice)
	fmt.Println("-------------------------------------------------------")
}
