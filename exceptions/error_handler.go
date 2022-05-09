package exceptions

import (
	"encoding/json"
	"fmt"
)

func ErrorHandler(err error) {
	errS := ErrorStruct{}
	json.Unmarshal([]byte(err.Error()), &errS)
	fmt.Println("Eror Code = ", errS.Code)
}
