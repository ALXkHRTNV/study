package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 1}
	b := []int{4, 2, 5, 1, 1, 2}

	intersection := intersections(a, b) // получаем пересечения
	unique := uniques(a, b)             // получаем уникальные

	fmt.Println("Пересечения:", intersection)
	fmt.Println("Уникальные:", unique)
}

func uniques(a, b []int) []int { // создаем функцию сравнения для уникальных чисел
	result := []int{}     // создаем пустой список для хранения результатов
	for _, x := range a { // проходим по каждому элементу из a
		already := false           // переменная для проверки, есть ли число в result
		for _, z := range result { // проверяем, есть ли это число в result
			if z == x { // если число уже есть в result, помечаем как добавленное
				already = true // проверяем, есть ли уже такое число в result.
				break          // останавливаем цикл
			}
		}
		if !already { // если число не найдено в result
			result = append(result, x) // добавляем число в result
		}
	}

	// Добавляем элементы из b в result, если их ещё нет
	for _, x := range b { // проходим по каждому элементу из b
		already := false           // переменная для проверки, есть ли число в result
		for _, z := range result { // проверяем, есть ли это число в result
			if z == x { // если число уже есть в result, помечаем как добавленное
				already = true
				break // останавливаем цикл
			}
		}
		if !already { // если число не найдено в result
			result = append(result, x) // добавляем число в result
		}
	}

	return result
}

func intersections(a, b []int) []int { // создаем функцию пересечения, которая берет 2 списка и взвращает список
	result := []int{} // создаем пустой список для хранения результатов

	for _, x := range a { // проходим по каждому числу из a
		for _, y := range b { // проходим по каждому числу из b
			if x == y { // если числа совпали
				already := false // создаём переменную, которая проверит, есть ли x уже в списке result

				for _, z := range result { // проверяем, есть ли это число в result
					if z == x { // если есть, помечаем как добавленное
						already = true
						break // останавливаем цикл
					}
				}

				if !already { // если числа нет в result, добавляем
					result = append(result, x)
				}

				break // останавливаем цикл
			}
		}
	}

	return result
}
