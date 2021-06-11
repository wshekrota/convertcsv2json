package main

import "fmt"
import "bufio"
import "os"
import "encoding/json"


// Convert a piped json stream to a csv stream you can direct to a file
//
func main() {

	obj := make([]map[string]interface{}, 2)
	r := bufio.NewReader(os.Stdin)
	s, err := r.ReadString('\n')
	if err != nil {
		fmt.Println("Error", err)
	}
	json.Unmarshal([]byte(s), &obj)
	for i := range obj {
		if i == 0 {
			for key, _ := range obj[0] {
				fmt.Print(key, ",")
			}
			fmt.Println()
		}
		for _, v := range obj[i] {
			fmt.Print(v, ",")
		}
		fmt.Println()

	}

}
