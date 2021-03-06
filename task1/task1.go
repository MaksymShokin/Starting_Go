// Time to practice what you learned!

// 1) Create a new Go code file, set a package and add the main function
// 2) Create two new variables:
//   - For your first name
//   - and your last name

//   Create those variables in two different ways (with and without
//   explicit type assignment).

// 3) Output the two variables in the command line and execute the code file
//   to see the output in the command line.

// 4) Also add two other variables:
//   - The current year
//   - Your birth year

//   Again, create those variables in two different ways and set the correct
//   value type (choose an appropriate type).

// 5) Calculate the difference ("age") between the two years and
//   store it in a new variable. Output that result in the command line.

// 6) Overwrite the value stored in the "current year" variable with
//   the previous value + 1 (i.e. next year). Calculate the next year,
//   don't just change the number.
//   Repeat step 5) with that new value (without changing any of the previous code).

// Try all those steps on your own first, then have a look at my solution
// lecture to compare your solution to mine!

package task1

import "fmt"

func main() {
	var firstName string
	firstName = "Max"

	lastName := "Shokin"

	fmt.Println(firstName)
	fmt.Println(lastName)

	currentYear := 2021
	birthYear := 1994

	age := currentYear - birthYear

	fmt.Println(age)

	currentYear += 1
	age = currentYear - birthYear
	fmt.Println(age)
}
