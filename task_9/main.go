/*
 *	  Большой ручей преграждает вам путь. По словам местных жителей, сейчас переходить ручей небезопасно,
 *	потому что он завален мусором. Вы смотрите вниз на ручей, а там не вода, а большой поток символов.
 *	  Вы сидите какое-то время и записываете часть потока (они находятся по кнопке Данные внизу задачи).
 *	Символы представляют собой группы — некие последовательности, которые начинаются с символа {
 *	и заканчиваются на символ }. Внутри каждой группы есть ноль или более других вещей,
 *	разделенных запятыми: это либо другая группа, либо мусор. Поскольку группы могут содержать
 *	другие группы, символ } закрывает только самую последнюю незакрытую группу, то есть группы
 *	могут быть вложенными. Данные этой задачи в файле (по кнопке Данные) представляет собой одну
 *	большую группу, которая сама содержит множество маленьких групп.
 *	  Иногда вместо группы вы найдете мусор. Мусор начинается с < и заканчивается >. Между этими угло-
 *	выми скобками может стоять практически любой символ, включая { и }. Внутри мусора < не имеет осо-
 *	богозначения.
 *	  В тщетной попытке очистить мусор, какая-то программа использовала ! и отменила некоторые из сим-
 *	волов внутри себя. Внутри мусора любой символ, который идет после ! следует игнорировать, в том
 *	числе <, > и даже еще один !.
 *	  В файле нет больше никаких других символов, которые не соответствуют этим правилам. Вне мусора вы
 *	найдете только правильно сформированные группы, а мусор всегда завершается в соответствии с
 *	приведенными выше правилами.
 *	  Ваша цель — найти общий балл для всех групп. Каждой группе присваивается балл, который на единицу
 *	больше, чем балл группы, в которую она непосредственно входит. (То есть самая дальняя от середины
 *	группа получает 1 балл, а самая вложенная - наибольший балл.)
 *	  Каков общий балл для всех групп в вашем вводе?
 */

package main

import (
	"fmt"
	"io/ioutil"
)

// Значение для проверки на то, находимся ли мы в "мусоре"
const (
	isCheck  = true
	isnCheck = false
)

func main() {
	// Считываем данный из файла в слайс битов, и если ошибки нет, то работаем с полученными данными
	data, err := ioutil.ReadFile("09.txt")
	if err != nil {
		fmt.Println(err)
	}

	searchData := ""   // Мы ищем все данные, кроме мусора
	myFlag := isnCheck // Флаг для понимания, зашли мы в мусор или нет

	for i := 0; i < len(data); i++ {
		//Если мы до этого еще не заходили в "мусор" и если мы в позиции символа "<",
		//то обновляем флаг, указывая на то, что мы зашли в "мусор"
		if !myFlag {
			if string(data[i]) == "<" {
				myFlag = isCheck
			} else {
				searchData += string(data[i])
			}
		} else {
			if string(data[i]) == "!" { //Если мы в мусоре, то проверяем в позиции знак "!", после него знаки не действительны
				i++
			} else if string(data[i]) == ">" { // Если нашли знак ">", то выходим из "мусора"
				myFlag = isnCheck
			}
		}
	}
	data = []byte(searchData)
	searchData = ""

	//После того, как все действительные данные были найдены, выделяем символы "{", "}" от запятых
	for i := 0; i < len(data); i++ {
		if (string(data[i]) == "{") || (string(data[i]) == "}") {
			searchData += string(data[i])
		}
	}
	data = []byte(searchData)
	searchData = ""

	//Считаем сумму очков для всех групп
	sum, j := 0, 0
	for i := 0; i < len(data); i++ {
		if string(data[i]) == "{" {
			j++
			sum += j
		} else if string(data[i]) == "}" {
			j--
		}
	}
	fmt.Println(sum)
}
