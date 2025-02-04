package player

import "fmt"

func ExampleBot_GetName() {

	h := NewHuman("Alex")

	fmt.Println(h.GetName())

	// Output:
	// Alex
}

func ExampleBot_GetHand() {

}
