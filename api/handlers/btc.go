package handlers

import (
	"fmt"
	"net/http"
)

func GetBTCPrice(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Getting BTC price")
}
