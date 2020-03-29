package corona

import "fmt"

// PrintOutput function is cool
func PrintOutput(choice int) int {
	if choice == 1 {
		fmt.Println("Covid19 cases in Pakistan are: 1,526 ")

	} else if choice == 2 {
		fmt.Println("Covid19 cases in South Korea are: 9,583 ")

	} else if choice == 3 {
		fmt.Println("Covid19 cases in France are: 37,575")

	} else if choice == 4 {
		var message string
		fmt.Println("Enter your personalised message for corona")
		fmt.Scanf("%s", &message)
		fmt.Println("Your message to corona is: ", message)
	} else if choice == 0 {
		fmt.Println("You have exited")
		return 0
	} else {
		fmt.Println("invalid input")
	}
	return 1
}
