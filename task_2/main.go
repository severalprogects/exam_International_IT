/*
 *	  В многоквартирном доме вам нужно доставить товары во все квартиры, но прежде
 *	товары нужно запаковать в коробки. Для этого нужно выяснить сколько квадратных
 *	футов картона вам нужно.
 *	  Коробка имеет форму правильной прямоугольной призмы и имеет длину l, ширину
 *	w, и высоту h. На каждую коробку вам нужно:
 *	2*l*w + 2*w*h + 2*h*l + площадь наименьшей стороны.
 *	  Например:
 *	  - Для подарка размером 2x3x4 требуется 2*6 + 2*12 + 2*8 = 52
 *	квадратных футов оберточной бумаги плюс 6 квадратных футов по наименьшей сторо-
 *	не, итого итого 58 квадратных футов.
 *	  - Для подарка размером 1x1x10 требуется 2*1 + 2*10 + 2*10 = 42 квадратных фу-
 *	тов оберточной бумаги плюс 1 квадратный фут по наименьшей стороне, итого итого
 *	43 квадратных фута.
 *	  Все числа в списке (кнопка данные) указаны в футах. Сколько всего квадратных
 *	футов картона требуется для запаковки товаров?
 */

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type sidePrism struct {
	l int64
	w int64
	h int64
}

// Подсчет перовй части формулы
func (s *sidePrism) searchV() int64 {
	V := (2 * s.l * s.w) + (2 * s.w * s.h) + (2 * s.h * s.l) //Объем
	return V
}

// Подсчет площади минимальной стороны призмы (вторая часть уравнения)
func (s *sidePrism) searchMinS() int64 {
	S := [3]int64{}
	S[0] = s.l * s.h
	S[1] = s.l * s.w
	S[2] = s.h * s.w

	minS := S[0] //Минимальная площадь одной из сторон призмы
	for i := 1; i < 3; i++ {
		if minS > S[i] {
			minS = S[i]
		}
	}
	return minS
}

// Итоговое значения необходимого кол-ва футов картона
func (s *sidePrism) searchVS() int64 {
	VP := s.searchV() + s.searchMinS()
	return VP
}

func main() {
	// Считываем данный из файла в слайс битов, и если ошибки нет, то работаем с полученными данными
	data, err := ioutil.ReadFile("02.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Выделяем строки из прочитанных данных
	dataLine := strings.Split(string(data), "\n")

	// Выделяем отдельные значения сторон из строк в слайс структур
	dataLinePart := []string{}
	prism := sidePrism{}
	sumVS := 0

	// Считываем l, w, h для каждой призмы в структуру, и считыем необходимое кол-во футов картона
	for i := 0; i < len(dataLine); i++ {
		dataLinePart = strings.Split(dataLine[i], "x") //Выделяем значения h, w, l

		prism.l, err = strconv.ParseInt(dataLinePart[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		prism.w, err = strconv.ParseInt(dataLinePart[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		prism.h, err = strconv.ParseInt(dataLinePart[2], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		sumVS += int(prism.searchVS())
	}

	fmt.Println(sumVS)
}
