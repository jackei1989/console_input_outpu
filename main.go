package main

import (
	"fmt"
	"strconv"

	"github.com/eiannone/keyboard"
)

func main() {
	coffees := make(map[int]string)
	coffees[1] = "Cappucino"
	coffees[2] = "Latte"
	coffees[3] = "Americano"
	coffees[4] = "Mocha"
	coffees[5] = "Macchiato"
	coffees[6] = "Espresso"

	fmt.Println("MENU")
	fmt.Println("-----")
	fmt.Println("1 - Cappucino")
	fmt.Println("2 - Latte")
	fmt.Println("3 - Americano")
	fmt.Println("4 - Mocha")
	fmt.Println("5 - Macchiato")
	fmt.Println("6 - Espresso")
	fmt.Println("Q - Quit the program")

	if err := keyboard.Open(); err != nil {
		panic(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Println("Press any key on the keyboard. Press ESC to quit.")

	for {
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			panic(err)
		}

		if char == 'q' || char == 'Q' {
			break
		}

		i, _ := strconv.Atoi(string(char))
		fmt.Println(fmt.Sprintf("You chose %s", coffees[i]))
	}

	fmt.Println("Program exiting...")
}
