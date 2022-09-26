package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var result int
var resultrim = "no rim"
var err1 string = "no"
var boolprov bool
var namone, oper, namone2, summ string
var arr []string

func main() {
	fmt.Println("передайте данные ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	if len(input) < 4 {
		err1 = "Вывод ошибки, так как строка не является математической операцией."
	} else if len(input) > 4 {
		err1 = "Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)."
	} else {
		arr = strings.Split(input, " ")
		namone = arr[0]
		oper = arr[1]
		namone2 = arr[2]
	}
	g := &summ
	for _, v := range namone2 {
		if v == 10 {
			break
		}
		*g = summ + string(v)
	}
	namone2 = summ
	if (namone == "0" || namone == "1" || namone == "2" || namone == "3" || namone == "4" || namone == "5" || namone == "6" || namone == "7" || namone == "8" || namone == "9" || namone == "10") && (namone2 == "0" || namone2 == "1" || namone2 == "2" || namone2 == "3" || namone2 == "4" || namone2 == "5" || namone2 == "6" || namone2 == "7" || namone2 == "8" || namone2 == "9" || namone2 == "10") {
		a1, _ := strconv.Atoi(namone)
		a2, _ := strconv.Atoi(namone2)
		if oper == "-" {
			result = a1 - a2
		} else if oper == "+" {
			result = a1 + a2
		} else if oper == "*" {
			result = a1 * a2
		} else if oper == "/" {
			result = a1 / a2
		} else {
			err1 = "Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)."
		}
		boolprov = true
	}
	rimrar := []string{"I", "II", "III", "IV", "VI", "VII", "VIII", "IX", "X", "XI", "XII", "XIII", "XIV", "XVI", "XVII", "XVIII", "XIX", "XX", "XXI", "XXII", "XXIII", "XXIV", "XXVI", "XXVII", "XXVIII", "XXIX", "XXX", "XXXI", "XXXII", "XXXIII", "XXXIV", "XXXVI", "XXXVII", "XXXVIII", "XXXIX", "XL", "XLI", "XLII", "XLIII", "XLIV", "XLVI", "XLVII", "XLIII", "XLIX", "L", "LI", "LII", "LIII", "LIV", "LVI", "LVII", "LVIII", "LIX", "LX", "LXI", "LXII", "LXIII", "LXIV", "LXVI", "LXVII", "LXVIII", "LXIX", "LXX", "LXXI", "LXXII", "LXXIII", "LXXIV", "LXXVI", "LXXVII", "LXXVIII", "LXXIX", "LXXX", "LXXXI", "LXXXII", "LXXXIII", "LXXXIV", "LXXXVI", "LXXXVII", "LXXXVIII", "LXXXIX", "XC", "XCI", "XCII", "XCIII", "XCIV", "XCVI", "XCVII", "XCVIII", "XCIX", "C"}
	var znatsh1, znatsh2 = 0, 0
	for i, v := range rimrar {
		if namone == v {
			znatsh1 = i + 1
		}
		if namone2 == v {
			znatsh2 = i + 1
		}
	}
	if znatsh1 != 0 && znatsh2 != 0 {
		a1 := znatsh1
		a2 := znatsh2
		if oper == "-" {
			result = a1 - a2
			if result <= 0 {
				err1 = "Вывод ошибки, так как в римской системе нет отрицательных чисел."
			}
			resultrim = rimrar[result+1]
		} else if oper == "+" {
			result = a1 + a2
			if result <= 0 {
				err1 = "Вывод ошибки, так как в римской системе нет отрицательных чисел."
			}
			resultrim = rimrar[result+1]
		} else if oper == "*" {
			result = a1 * a2
			if result <= 0 {
				err1 = "Вывод ошибки, так как в римской системе нет отрицательных чисел."
			}
			resultrim = rimrar[result+1]
		} else if oper == "/" {
			result = a1 / a2
			if result <= 0 {
				err1 = "Вывод ошибки, так как в римской системе нет отрицательных чисел."
			}
			resultrim = rimrar[result+1]
		} else {
			err1 = "Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)."
		}
		boolprov = true
	}
	if err1 != "no" {
		fmt.Println(err1)
	} else if boolprov != true {
		fmt.Println("Вывод ошибки, так как используются одновременно разные системы счисления.")
	} else if resultrim != "no rim" {
		fmt.Println(resultrim)
	} else {
		fmt.Println(result)
	}
}
