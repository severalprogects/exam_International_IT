/*
 * Последовательности "посмотри и скажи" генерируются итеративно,
 * используя предыдущее значение в качестве входных данных для следующего шага.
 * Для каждого шага возьмите предыдущее значение и замените каждый ряд цифр
 * (например, 111) на количество цифр (3), за которым следует сама цифра (1).
 * Например:
 * - 1 становится 11 (1 копия цифры 1).
 * - 11 становится 21 (2 копии цифры 1).
 * - 21 становится 1211 (одна 2, за которой следует одна 1).
 * - 1211 становится 111221 (одна 1, одна 2 и две 1).
 * - 111221 становится 312211 (три 1, две 2 и одна 1).
 * Начиная с цифр 1113222113 проделайте этот процесс 40 раз. Какова длина результата?
 */

package main

import "fmt"

/*	* Input:
	oldCount - часть значения, которую мы хотим дополнить
	count 	 - все значение
	i 		 - индекс
	valI 	 - значение, которое было найдено по индексу i
	* Output:
	newCount - часть значения, которую мы дополнили
	newI 	 - обновленное значение итератора
*/
func searchSum(oldCount string, count string, i int, valI string) (newCount string, newI int) {
	//Мы знаем, что в последовательности встречаются только цифры 1, 2 и 3.
	//Длина последовательностей также не превышает трех

	sum := "1"
	if (i+1 < len(count)) && (string(count[i]) == string(count[i+1])) {
		sum = "2"
		i++
		if (i+1 < len(count)) && (string(count[i]) == string(count[i+1])) {
			sum = "3"
			i++
		}
	}

	newI = i                         //Возвращаем индекс, для дальнейшего просмотра значения
	newCount = oldCount + sum + valI //Возвращаем часть следующего значения
	return newCount, newI
}

func main() {
	count := "1113222113"
	newCount := ""

	for j := 0; j < 40; j++ {
		for i := 0; i < len(count); i++ {
			if string(count[i]) == "1" {
				newCount, i = searchSum(newCount, count, i, "1")
			} else if string(count[i]) == "2" {
				newCount, i = searchSum(newCount, count, i, "2")
			} else {
				newCount, i = searchSum(newCount, count, i, "3")
			}
		}
		count = newCount
		newCount = ""
	}
	fmt.Println(len(count))
}
