package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"
)

const (
	baseURL       = "https://short-link.me/"
	prefix        = "12"          // Can change to "ZD"
	totalLength   = 5
	numChecks     = 1000            // Number of URLs to generate
	requestTimeout = 5 * time.Second
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// Generate random code with given prefix
func generateRandomCode(prefix string, length int) string {
	suffixLength := length - len(prefix)
	suffix := make([]rune, suffixLength)
	for i := range suffix {
		suffix[i] = letters[rand.Intn(len(letters))]
	}
	return prefix + string(suffix)
}

// Check if the HTML contains phishing/interstitial keywords
func isRealLink(body string) bool {
	fakeMarkers := []string{
		"Antiphishing.biz", "not a robot", "cyber security",
		"landing page address will be shown", "confirm that you are not a robot",
	}
	for _, marker := range fakeMarkers {
		if strings.Contains(strings.ToLower(body), strings.ToLower(marker)) {
			return false
		}
	}
	return true
}

func checkURL(code string, wg *sync.WaitGroup, mu *sync.Mutex, results *[]string) {
	defer wg.Done()

	client := http.Client{Timeout: requestTimeout}
	fullURL := baseURL + code

	resp, err := client.Get(fullURL)
	if err != nil {
		// fmt.Printf("❌ Error: %s - %s\n", fullURL, err)
		return
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	body := string(bodyBytes)

	if resp.StatusCode == 200 && isRealLink(body) {
		fmt.Printf("✅ Valid: %s\n", fullURL)
		mu.Lock()
		*results = append(*results, fullURL)
		mu.Unlock()
	} else {
		// fmt.Printf("⚠️ Skipped: %s\n", fullURL)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup
	var mu sync.Mutex
	var workingLinks []string

	for i := 0; i < numChecks; i++ {
		code := generateRandomCode(prefix, totalLength)
		wg.Add(1)
		go checkURL(code, &wg, &mu, &workingLinks)
	}

	wg.Wait()

	fmt.Println("\n=== ✅ WORKING SHORT LINKS ===")
	for _, link := range workingLinks {
		fmt.Println(link)
	}
}
