package input

import (
	"fmt"
	"net/http"
	"os"
	"utils"
)

func CheckExists(dayInfo int) bool {
	filepath := fmt.Sprintf("./day%d/input.txt", dayInfo)
	if _, err := os.Stat(filepath); err == nil {
		println("exists:", filepath)
		return true
	}
	return false
}

func Fetch(dayInfo int) {
	fmt.Println("Fetching input")
	client := &http.Client{}
	fmt.Println("cookie: ")
	cookieValue := os.Getenv("SESSION_COOKIE")
	fmt.Println(cookieValue)
	if cookieValue == "" {
		panic("No session cookie value")
	}
	cookie := http.Cookie{Value: os.Getenv("SESSION_COOKIE"), Name: "session"}
	request, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/2022/day/%d/input", dayInfo), nil)
	utils.Check(err)
	request.AddCookie(&cookie)
	resp, err := client.Do(request)
	utils.Check(err)
	data := utils.ReadResponseBody(resp.Body)
	// fmt.Println(string(data))
	filepath := fmt.Sprintf("./day%d/input.txt", dayInfo)
	println(filepath)
	f, err := os.Create(filepath)
	utils.Check(err)
	defer f.Close()

	_, err = f.Write(data)
	utils.Check(err)
}
