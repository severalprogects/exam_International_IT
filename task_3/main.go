/*
 *	  В многоквартирном доме вам нужно доставить подарки во все квартиры, они уже
 *	упакованы в коробки и вам нужно обернуть все коробки лентами. Для этого нужно
 *	выяснить сколько футов ленты вам нужно.
 *	  Коробка имеет форму правильной прямоугольной призмы и имеет длину l, ширину
 *	w, и высоту h. Лента представляет собой наименьший периметр какой-либо одной
 *	грани. Каждой коробке так же требуется бантик из ленты, количество футов лен-
 *	ты необходимое для банта равно кубическому футу объема подарка.
 *	  Например:
 *	- Для подарка размером 2x3x4 требуется 2+2+3+3 = 10 футов ленты, чтобы обер-
 *	нуть подарок, плюс 2*3*4 = 24 фута ленты для банта, итого 34 фута.
 *	- Для подарка размером 1x1x10 требуется 1+1+1+1+1 = 4 фута ленты, чтобы обер-
 *	нуть подарок, плюс 1*1*10 = 10 футов ленты для банта, итого 14 футов.
 *	  Все числа в списке (кнопка данные) указаны в футах. Сколько всего футов
 *	ленты требуется для оборачивания подарков?
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

// Подсчет футов ленты для банта
func (s *sidePrism) searchV() int64 {
	V := s.l * s.w * s.h //Объем призмы
	return V
}

// Подсчет периметра минимальной стороны призмы
func (s *sidePrism) searchMinP() int64 {
	P := [3]int64{}
	P[0] = s.l + s.l + s.h + s.h
	P[1] = s.l + s.l + s.w + s.w
	P[2] = s.h + s.h + s.w + s.w

	minP := P[0] //Минимальный периметр одной из сторон призмы
	for i := 1; i < 3; i++ {
		if minP > P[i] {
			minP = P[i]
		}
	}
	return minP
}

// Итоговое значения необходимого кол-ва футов картона
func (s *sidePrism) searchVP() int64 {
	VP := s.searchV() + s.searchMinP()
	return VP
}

func main() {
	// Считываем данный из файла в слайс битов, и если ошибки нет, то работаем с полученными данными
	data, err := ioutil.ReadFile("03.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Выделяем строки из прочитанных данных
	dataLine := strings.Split(string(data), "\n")

	// Выделяем отдельные значения сторон из строк в слайс структур
	dataLinePart := []string{}
	prism := sidePrism{}
	sumVP := 0

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

		sumVP += int(prism.searchVP())
	}

	fmt.Println(sumVP)
}
