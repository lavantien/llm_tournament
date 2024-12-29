package main

// import (
// 	"io"
// 	"net/http"
// 	"os"
// )
//
// func main() {
// 	url := "https://pkg.go.dev/fyne.io/fyne/v2"
//
// 	// Create HTTP client with custom user agent
// 	client := &http.Client{}
// 	req, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	// Set a user agent to avoid being blocked
// 	req.Header.Set("User-Agent", "Mozilla/5.0")
//
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer resp.Body.Close()
//
// 	// Create output file
// 	file, err := os.Create("fyne_api.html")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()
//
// 	// Copy response body to file
// 	_, err = io.Copy(file, resp.Body)
// 	if err != nil {
// 		panic(err)
// 	}
// }
