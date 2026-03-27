package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"sync"
	"time"
)

const (
	targetURL      = "https://www.pgchoice.com/property/query"
	totalRequests  = 20
	maxConcurrency = 2
	requestTimeout = 10 * time.Second

	// Set this to your known daily limit per mobile/user
	dailyLimitPerMobile = 5
)

var mobileNumbers = []string{
	"6876543210",
	"7123456780",
	"9888877777",
	"2777766666",
	"3666655555",
	"5999944444",
	"6111122222",
	"7000033333",
	"9500011111",
	"4400099999",
}

var perMobileCount = make(map[string]int)
var mu sync.Mutex

func canSend(mobile string) bool {
	mu.Lock()
	defer mu.Unlock()
	if perMobileCount[mobile] >= dailyLimitPerMobile {
		return false
	}
	perMobileCount[mobile]++
	return true
}

func sendPost(id int, wg *sync.WaitGroup, sem chan struct{}) {
	defer wg.Done()
	defer func() { <-sem }()

	mobile := mobileNumbers[id%len(mobileNumbers)]

	if !canSend(mobile) {
		fmt.Printf("⛔ Skipping %s (daily cap reached)\n", mobile)
		return
	}

	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	writer.WriteField("recordId", "NzU2MzA1OTY0MDU0")
	writer.WriteField("page_url", "https://www.pgchoice.com/property/NzU2MzA1OTY0MDU0")
	writer.WriteField("user_name", fmt.Sprintf("User %d", id))
	writer.WriteField("user_mobile", mobile)
	writer.WriteField("plan_move", "More than 14 days")
	writer.WriteField("answer", "27")
	writer.WriteField("correct_answer", "27")

	writer.Close()

	req, err := http.NewRequest("POST", targetURL, &requestBody)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("X-Requested-With", "XMLHttpRequest")

	client := http.Client{Timeout: requestTimeout}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("❌ Failed:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("✅ Req %d | Mobile: %s | Status: %d\n", id, mobile, resp.StatusCode)
	fmt.Println(string(body))
}

func main() {
	var wg sync.WaitGroup
	sem := make(chan struct{}, maxConcurrency)

	for i := 0; i < totalRequests; i++ {
		wg.Add(1)
		sem <- struct{}{}
		go sendPost(i, &wg, sem)
		time.Sleep(2 * time.Second) // spacing to avoid burst
	}

	wg.Wait()
	fmt.Println("Finished without exceeding client-side quota")
}