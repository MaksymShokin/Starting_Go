package main

import "fmt"

func main() {
	// var greetingText string
	// greetingText = "Hello World"

	var greetingText string = "Hello World"

	// var greetingText = "Hello World"

	// greetingText := "Hello World"
	luckyNumber := 17

	fmt.Println(greetingText)
	fmt.Println(luckyNumber)

	var superLuckyNumber float64 = float64(luckyNumber) / 4

	fmt.Println(superLuckyNumber)

	var defaultFloat float64 = 1.12345678901234567890 // 1.1234567890123457
	var smallFloat float32 = 1.12345678901234567890   // 1.1234568

	fmt.Println(defaultFloat)
	fmt.Println(smallFloat)

	var firstRune rune = '$'
	fmt.Println(firstRune)
	fmt.Println(string(firstRune))

	var firstByte byte = 'a'
	fmt.Println(firstByte)
	fmt.Println(string(firstByte))

	// var firstBool = false

	firstName := "Max"
	lastName := "Shokin"

	fmt.Println(firstName + " " + lastName)

	var fl float32 = 3.3
	var fl2 float64 = 3.3

	fmt.Println(fl + float32(fl2))

	fmt.Println("8" + string(9))

	mline := `This is a multiline string!	
	     It is quite long.
	
	And spans multiple lines!`

	fmt.Println(mline)

}

// run file go run first_app.go
