package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 1}       // слайс a
	b := []int{4, 2, 5, 1, 1, 2} // слайс b

	result := []int{} // пустой слайс для результата

	for _, x := range a { // проходим по каждому числу из a
		for _, y := range b { // проходим по каждому числу из b
			if x == y { // если числа совпали
				already := false // переменная для проверки, есть ли число в result

				for _, z := range result { // проверяем, есть ли это число в result
					if z == x { // если есть, помечаем как добавленное
						already = true
						break
					}
				}

				if !already { // если числа нет в result, добавляем
					result = append(result, x)
				}

				break // выходим из второго цикла после первого совпадения
			}
		}
	}

	fmt.Println(result) // выводим результат
}









package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 1}       // слайс a
	b := []int{4, 2, 5, 1, 1, 2} // слайс b

	result := []int{} // пустой слайс для результата

	// Добавляем элементы из a в result
	for _, x := range a { // проходим по каждому элементу из a
		already := false // переменная для проверки, есть ли число в result
		for _, z := range result { // проверяем, есть ли это число в result
			if z == x { // если число уже есть в result, помечаем как добавленное
				already = true
				break // выходим из цикла
			}
		}
		if !already { // если число не найдено в result
			result = append(result, x) // добавляем число в result
		}
	}

	// Добавляем элементы из b в result, если их ещё нет
	for _, x := range b { // проходим по каждому элементу из b
		already := false // переменная для проверки, есть ли число в result
		for _, z := range result { // проверяем, есть ли это число в result
			if z == x { // если число уже есть в result, помечаем как добавленное
				already = true
				break // выходим из цикла
			}
		}
		if !already { // если число не найдено в result
			result = append(result, x) // добавляем число в result
		}
	}

	fmt.Println(result) // выводим результат
}
