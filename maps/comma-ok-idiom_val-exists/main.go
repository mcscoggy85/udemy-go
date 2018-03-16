package main
import "fmt"

func main() {
	myGreeting := map [string]string{
		"Tim": "Good Morning",
		"Jenny": "Bonjour",
		"Chris": "Hey-OH!",
	}

	myGreeting["Harleen"] = "Howdy"

	fmt.Println(myGreeting)
	fmt.Println(myGreeting["Tim"])
	fmt.Println(myGreeting["Jenny"])
	fmt.Println(myGreeting["Chris"])
	fmt.Println(myGreeting["Harleen"])

	myGreeting["Harleen"] = "Hola!"
	fmt.Println(myGreeting)
	fmt.Println(myGreeting["Harleen"])

	delete(myGreeting, "Tim")
	fmt.Println(myGreeting)

	if val, exists := myGreeting["Jenny"]; exists {
		delete(myGreeting, "Jenny")
		fmt.Println("val: ",val)
		fmt.Println("exists: ", exists)

		} else {
			fmt.Println("That value does not exist.")
			fmt.Println("val: ",val)
			fmt.Println("exists: ", exists)

	}
	fmt.Println(myGreeting)


}