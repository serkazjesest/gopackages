package main

import (
	"fmt"
	"gopackages/citydigit"
	newcolor "gopackages/color"
	"gopackages/wordz"

	utilsV2 "github.com/dhnikolas/utils/v2"
	utilsV3 "github.com/dhnikolas/utils/v3"
	"github.com/fatih/color"
	"github.com/huandu/xstrings"
	"github.com/serkazjesest/utils"
)

func main() {
	isExistV2 := utilsV2.InSlice(wordz.Words, "Two")
	if isExistV2 {
		fmt.Println("Using utilsV2.InSlice and find value ")
		return
	}

	isExistInt := utilsV3.InSliceInt([]int{1, 2, 3, 4, 5}, 5)
	if isExistInt {
		fmt.Println("Slice Int contain finding value")
		return
	}

	//С помощью нашей новой библиотеки проверяем слайс на наличие определенного значения
	//Если его там находим, то сообщим это в stdout и закончим выполнение программы
	isExist := utils.Contains(wordz.Words, "Two")
	if isExist {
		fmt.Println("Slice Words contain finding value")
		return
	}

	newcolor.Greet()
	fmt.Println("Hello world")
	color.Red("Hello world again")

	fmt.Println(wordz.Hello)
	wordz.Words = []string{"Moscow", "New-York", "Amsterdam", "Barcelona", "Paris"}
	fmt.Println(wordz.Random())

	wordz.Prefix = ""
	fmt.Println(citydigit.City())
	fmt.Println(citydigit.Digit())

	fmt.Println(xstrings.Shuffle(citydigit.City()))
	fmt.Println(xstrings.Shuffle(citydigit.Digit()))

}
