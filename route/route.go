package route

import (
	"net/http"
)

func CheckHealthAPI(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API is running"))
}
