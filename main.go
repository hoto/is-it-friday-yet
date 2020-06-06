package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	isItFridayYet := theAnswer()
	letTheWorldKnow(isItFridayYet)
	fmt.Printf("Is it friday yet?\n%s", isItFridayYet)
}

func theAnswer() string {
	switch time.Now().Weekday() {
	case time.Monday:
		return "oh boy... it's a monday"
	case time.Tuesday:
		return "nope but closer than monday"
	case time.Wednesday:
		return "half way there"
	case time.Thursday:
		return "just one more day"
	case time.Friday:
		return "yes 🤤"
	case time.Saturday:
		return "hihi nope, but it's a better friday"
	case time.Sunday:
		return "nope, sunday, weekend is almost over :("
	default:
		return "Andyyyyy something went super wrong!"
	}
}

func letTheWorldKnow(isItFridayYet string) {
	payload := strings.NewReader(fmt.Sprintf(`{"description": "%s"}`, isItFridayYet))
	req, err := http.NewRequest("PATCH", "https://api.github.com/repos/hoto/is-it-friday-yet", payload)
	check(err)
	req.Header.Add("Authorization", "Bearer "+retrieveGithubToken())
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	check(err)
	if res.StatusCode != 200 {
		log.Panic("Github did not respond well. " + res.Status)
	}
}

func retrieveGithubToken() string {
	githubApiToken := os.Getenv("API_TOKEN_GITHUB")
	if len(githubApiToken) <= 0 {
		log.Panic("Provide API_TOKEN_GITHUB to spread the word.")
	}
	return githubApiToken
}

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}
