package inputs

import (
	"bufio"
	"fmt"
	"os"
)

func WaitForEnter() {
	fmt.Print("Press Enter to continue...")
	reader := bufio.NewReader(os.Stdin)
	_, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("\nError reading input:", err)
	}
}
