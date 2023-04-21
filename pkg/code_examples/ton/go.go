package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "%s/api/v2/getAddressInformation?address=1"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
