package main

import (
	"fmt"
)

func main() {
	fmt.Println("Расчёт суммы скидки")

	var price float64
	var discount float64

	fmt.Println("Введите цену товара")
	fmt.Scan(&price)
	fmt.Println("Введите размер скидки")
	fmt.Scan(&discount)

	if discount < 2000 && discount < (price*0.3) {
		fmt.Println("Размер скидки меньше 2000 и не больше 30% от цены товара, что удовлетворяет условию и состовляет:", discount)
	} else if discount > 2000 {
		discount = 2000
		fmt.Println("Размер скидки больше:", discount)
	} else if discount > (price * 0.3) {
		discountNew := (price * 0.3)
		fmt.Println("Размер скидки составил", discount, "что больше 30% от цены которая равна:", discountNew)
	}
}
