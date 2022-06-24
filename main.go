package main

// fmt use for print
import (
	"fmt"
	"steward_api/api"
	"steward_api/calculator"
)

//---- Comment-----

//test

/*
Sample
*/

//-----------------

func main() {
	//! Print
	// fmt.Println("hello")

	//! Declaration
	// var firstname string
	// firstname = "Supanut"
	// fmt.Println(firstname)

	// lastname := "Milung"
	// fmt.Println(lastname)
	//!

	//! Constant
	// const nickname string = "Best"
	// fmt.Println(nickname)
	//!

	//! Value Type
	// name := "Best"
	// fmt.Printf("My name is %v\n", name)
	// fmt.Printf("My name is %T\n", name)
	//!

	//! Operator = + - * / %
	// var number1 int = 10
	// var number2 = 3
	// var number1, number2 = 10 , 3
	// number1, number2 := 10 , 3
	// fmt.Println("+ = ", number1+number2)
	// fmt.Println("- = ", number1-number2)
	// fmt.Println("* = ", number1*number2)
	// fmt.Println("/ = ", number1/number2)
	// fmt.Println("% = ", number1%number2)
	//!

	//! Operator == != > < >= <=
	// number3, number4 := 10, 5
	// fmt.Println("== ", number3 == number4)
	// fmt.Println("!= ", number3 != number4)
	// fmt.Println("> ", number3 > number4)
	// fmt.Println("< ", number3 < number4)
	// fmt.Println(">= ", number3 >= number4)
	// fmt.Println("<= ", number3 <= number4)
	// fmt.Println(number3, " > ", number4, " = ", number3 > number4)
	//!

	//! Scanf
	//string = %s
	//integer = %d
	//float = %f

	//keep value from keyborad enter

	// var nameScanf string
	// var score int
	// fmt.Print("enter student =")
	// fmt.Scanf("%s", &nameScanf)
	// fmt.Scanf("%d", &score)
	// fmt.Println("Hello = ", nameScanf)
	// fmt.Println("score = ", score)
	//!

	//! if else
	/// 100 => >=50 pass, not pass
	// var point int
	// fmt.Print("enter point = ")
	// fmt.Scanf("%d", &point)

	/// Process
	// if point >= 50 {
	// 	fmt.Println("Pass")
	// } else {
	// 	fmt.Println("Not pass")
	// }
	//!

	//! Swith case
	/// 100 => >=50 pass, not pass
	// var point int
	// fmt.Print("enter point = ")
	// fmt.Scanf("%d", &point)

	// switch point {
	// case 1:
	// 	fmt.Println("Create Account")
	// case 2:
	// 	fmt.Println("Deposit - withdraw")
	// case 3:
	// 	{
	// 		fmt.Println("Wrong")
	// 		fmt.Println("Infomation")
	// 	}
	// default:
	// 	fmt.Println("Wrong Infomation")
	// }
	//!

	//! Array
	// name := [3]string{"a", "b", "c"}
	//len หาจำนวนของ array

	// number1 := 100
	// number2 := 200
	// number3 := 300

	//init [0,0,0]
	// var numberList [3]int

	// var numberList [3]int = [3]int{100, 200, 300}
	// numberList := [3]int{100, 200, 300}
	// numberList := [3]int{}
	// numberList := [3]int{100}

	// numberList := [3]int{100, 200, 300} // 0 = 100 , 1 = 200 , 2 = 300
	// numberList := [...]int{}

	// numberList[0] = 100
	// numberList[1] = 200
	// numberList[2] = 300
	//setValue [0] = 100

	// fmt.Println(numberList)

	// size := len(numberList)
	// fmt.Println(size)

	// fmt.Println(number1)
	// fmt.Println(number2)
	// fmt.Println(number3)

	//!

	//! Slices
	//กำหนดช่วง
	//names[1:] //index ที่ 1 ถึงสุดท้าย
	//names[:1] //index ที่ 0 ถึง 1
	// numberList := []int{}
	// numberList = append(numberList, 100)
	// numberList = append(numberList, 200)
	// numberList = append(numberList, 300)
	// numberList = append(numberList, 400)
	// fmt.Println(numberList[:])
	// //pick index 2 to last
	// fmt.Println(numberList[2:])
	//!

	//! Map (Key,value)
	//						key  value
	//exp var country map[string]string
	//coins := map [string]string {"ETH":"Ether" , "BTC":"Bitcoin"}
	// country := [2]string{"thai", "japan"}
	// country := map[string]string{"th": "thai", "jp": "japan"}
	// country := map[string]string{}
	// country["th"] = "thai"
	// country["jp"] = "japan"

	// fmt.Println(country)
	// fmt.Println(country["th"])

	// value, check := country["th"]
	// value, isThai := country["th"]
	// if isThai {
	// 	fmt.Println(value)
	// } else {
	// 	fmt.Println("No data")
	// }
	//!

	//! Loop
	//for count :=1; count <=3; count++ { fmt.Println("Hello")}
	// for count := 1; count <= 2; count++ {
	// 	fmt.Println(count)
	// }

	// for count := 10; count >= 1; count-- {
	// 	fmt.Println(count)
	// }
	//!

	//!	loop #2 -> break , continue
	// for count := 1; count <= 10; count++ {
	// if count == 5 {
	// 	break
	// }
	// if count == 5 {
	// 	continue
	// }

	// fmt.Println(count)

	// }
	// fmt.Println("End")
	//!

	//! for loop + array + slice
	// numbers := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	// for number := 0; number < len(numbers); number++ {
	// 	fmt.Println("Index =", number, "Value =", numbers[number])
	// }

	// for index, value := range numbers {
	// 	fmt.Println("Index =", index, "Value =", value)
	// }

	// _ = ไม่สนใจ index

	// for _, v := range numbers {

	// }

	// languages := map[string]string{"TH": "Thai", "EN": "English", "JP": "Japan"}
	// for key, value := range languages {
	// 	fmt.Println("Key = ", key, ",Value = ", value)
	// }

	// for _, value := range languages {
	// 	fmt.Println("Value = ", value)
	// }

	//!

	// showMessage()
	// showMessage("test")
	// total(50, 10)

	// delivery := getDelivery()
	// fmt.Println("Delivery =", delivery)

	// cart := getTotalCart(100, 100)
	// fmt.Println("Total =", cart)

	// result, check := summation(100, 200)
	// fmt.Println("result =", result)
	// fmt.Println(check)

	// result := summation(100, 200, 300, 400, 500, 600)
	// fmt.Println(result)

	// product := Product{
	// 	name:     "Apple",
	// 	price:    10.5,
	// 	category: "Food",
	// 	discount: 0,
	// }
	// product.price = 100

	// fmt.Println(product)

	fmt.Println(calculator.Add(50, 100))
	fmt.Println(calculator.Subtract(50, 100))

	api.HandleRequest()

}

//! function
// func showMessage() {
// 	fmt.Println("Hello")
// }
// func showMessage(name string) {
// 	fmt.Println("Hello", name)
// }
// func total(number1 int, number2 int) {
// 	fmt.Println("Total = ", number1+number2)

// }
// func total(number1, number2 int) {
// 	fmt.Println("Total = ", number1+number2)
// }

// func getDelivery() int {
// 	return 50
// }

// func getTotalCart(number1, number2 int) int {
// 	return number1 + number2
// }

// func summation(number1, number2 int) (int, string) {
// 	total := number1 + number2
// 	status := ""
// 	if total%2 == 0 {
// 		status = "Even"
// 	} else {
// 		status = "Odd"
// 	}
// 	return total, status
// }
//!

//! Variadic function parametor ไม่จำกัดจำนวน
// func summation(number ...int) int {
// 	total := 0
// 	for _, v := range number {
// 		total += v
// 	}
// 	return total
// }
//!

//! structure
type Product struct {
	name     string
	price    float64
	category string
	discount int
}

//! constructure
// func AddProduct(name string, price float64, category string, discount int) Product {
// product := Product{
// 	name:     name,
// 	price:    price,
// 	category: category,
// 	discount: discount,
// }

// 	return product
// }
//!

//! sigleton
// nil & null หมายถึง ค่าว่าง
// sigleton
// var product *Product

// func Init(name string, price float64, category string, discount int) *Product {
// 	if product == nil {
// 		product = &Product{
// 			name:     name,
// 			price:    price,
// 			category: category,
// 			discount: discount,
// 		}
// 	}
// 	return product
// }

//!

//! package
//!
