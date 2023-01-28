package input

import (
	"fmt"
	"net/http"
	"os"
	"utils"
)

func CheckExists() bool {
	if _, err := os.Stat("./day1/input"); err == nil {
		return false
	}
	return true
}

func Fetch() {
	fmt.Println("Fetching input")
	client := &http.Client{}
	fmt.Println("cookie: ")
	cookieValue := os.Getenv("SESSION_COOKIE")
	fmt.Println(cookieValue)
	if cookieValue == "" {
		panic("No session cookie value")
	}
	cookie := http.Cookie{Value: os.Getenv("SESSION_COOKIE"), Name: "session"}
	request, err := http.NewRequest("GET", "https://adventofcode.com/2022/day/1/input", nil)
	utils.Check(err)
	request.AddCookie(&cookie)
	resp, err := client.Do(request)
	utils.Check(err)
	data := utils.ReadResponseBody(resp.Body)
	// fmt.Println(string(data))

	f, err := os.Create("./day1/input.txt")
	utils.Check(err)
	defer f.Close()

	_, err = f.Write(data)
	utils.Check(err)
}
