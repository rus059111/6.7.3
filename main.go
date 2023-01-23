package main

//Задание 3. Расчёт суммы скидки

import (
	"fmt"
	//"math"
)

func main() {

	var price int
	var discount int
	var discountSumm int

	maxDiscount := 30
	maxDiscountSumm := 2000

	fmt.Println("Введите стоимость товара")
	fmt.Scan(&price)
	fmt.Println("Введите скидку")
	fmt.Scan(&discount)

	if discount > maxDiscount {
		discountSumm = price * maxDiscount / 100
		if discountSumm > maxDiscountSumm {
			discountSumm = maxDiscountSumm
		}
	} else {
		discountSumm = price * discount / 100
		if discountSumm > maxDiscountSumm {
			discountSumm = maxDiscountSumm
		}

	}
	fmt.Println("Cумма скидки, составляет", discountSumm, "рублей")

}
