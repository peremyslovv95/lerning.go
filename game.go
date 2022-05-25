//quess-игра,в которой игрок должен угадать случайное числою
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix() //Получаем текущую дату и время в виде целого числа
	rand.Seed(seconds)           //Иницилизируем генератор случайных чисел
	target := rand.Intn(100) + 1 //Генерируем число от 1 до 100
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you qess it?")
	//fmt.Println(target)

	reader := bufio.NewReader(os.Stdin) //для чтения ввода с клавиаутры
	success := false                    //вывод сообщение о проигрыше               //Обьявляем переменную до цикла чтобы она оставалась в области видимости после выода из цикла

	for quesses := 0; quesses < 10; quesses++ {
		fmt.Println("You have", 10-quesses, "quesses left.") // обозначаем кол-во попыток ввода

		fmt.Println("Make a quess: ")         //запрашиваем число
		input, err := reader.ReadString('\n') // чтение данных введеных с клавиатуры до enter
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)  //удаление символа новой строки
		quess, err := strconv.Atoi(input) // входная строка преобразуется в целое число
		if err != nil {
			log.Fatal(err)
		}
		if quess < target { //если введенное значение меньше загаданного сообщить об этом
			fmt.Println("Oops.Your qess was LOW.")
		} else if quess > target { ////если введенное значение меньше загаданного сообщить об этом
			fmt.Println("Oops.Your quess was HIGH.")
		} else {
			success = true // если угадал выводить о проигрыше не нужно
			fmt.Println("Good job!You qessed it!")
			break //прерывание цикла в случае успешного выбора1
		}
	} //конец цикла for.
	if !success {
		fmt.Println("Sorry,you didn't quess my number.It was:", target) // сообщение о проигрыше
	}
}
