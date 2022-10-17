package utils

import (
	"fmt"
	"net/http"
)

func SendError(err *CustomError, res http.ResponseWriter) {
	res.WriteHeader(err.Code)
	fmt.Println(err.Err.Error(), err.Op)
	res.Write([]byte(`{"message": "` + err.Error() + `"}`))
}
