package utils

import (
	"fmt"
	"net/http"
)

func SendError(err *CustomError, res http.ResponseWriter) {
	fmt.Println("error: " + err.Err.Error() + " in " + err.Op)
	res.WriteHeader(err.Code)
	res.Write([]byte(`{"message": "` + err.Message + `"}`))
}
