package main

import "fmt"
import "bufio"
import "os"
import "encoding/json"

// Convert a piped json stream to a headered csv stream you can direct to a file
// Assumption is keys are the same throughout. In other words column headers.
// Not all json will conform to this it is a specialized use.
// Tabular report <-> json
//
func main() {

	var maps = make([]map[string]interface{}, 3)

	r := bufio.NewReader(os.Stdin)
	s, err := r.ReadString('\n')
	if err != nil {
		fmt.Println("Error", err)
	}
	err = json.Unmarshal([]byte(s), &maps)
	fmt.Println(err)

	// Each array element map
	//
	for i := range maps {
		// Column headers
		//
		if i == 0 {
			for key, _ := range maps[i] {
				fmt.Print(key, ",")
			}
			fmt.Println()
		}
		// data values
		//
		for _, item := range maps[i] {
			fmt.Print(item, ",")
		}
		fmt.Println()

	}

}
