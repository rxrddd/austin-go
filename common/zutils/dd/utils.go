package dd

import (
	"encoding/json"
	"fmt"
)

func Print(val ...interface{}) {
	for _, v := range val {
		marshal, _ := json.Marshal(v)
		fmt.Println(string(marshal))
	}

}
