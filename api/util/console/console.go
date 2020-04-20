package console

import (
	"encoding/json"
	"fmt"
	"log"
)

func Pretty(data interface{}) {
	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println("error in console.Pretty():", err)
		return
	}

	fmt.Println(string(b))
}
