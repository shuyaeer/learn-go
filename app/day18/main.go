package main

import (
	"fmt"
	"net/http"
	"sync"
)

// sendRequest関数は、指定されたURLにHTTPリクエストを送信し、
// 成功か失敗かをchannelを通して返します。
func sendRequest(url string, wg *sync.WaitGroup, ch chan<- bool) {
	defer wg.Done()

	_, err := http.Get(url)
	if err != nil {
		ch <- false // 失敗
		return
	}

	ch <- true // 成功
}

func main() {
	var wg sync.WaitGroup
	requestCount := 1000 // 送信するリクエストの回数
	fmt.Printf("Request Count: %d\n", requestCount)
	url := "http://localhost:3000" // リクエストを送信するURL
	fmt.Printf("URL: %s\n", url)

	// 成功と失敗をカウントするための変数
	var successCount, failureCount int

	// 結果を受け取るためのchannelを作成
	ch := make(chan bool, requestCount)

	for i := 0; i < requestCount; i++ {
		wg.Add(1)
		// 各リクエストをgoroutineで実行
		go sendRequest(url, &wg, ch)
	}

	// すべてのgoroutineの完了を待機
	wg.Wait()
	close(ch)

	// 成功と失敗をカウント
	for result := range ch {
		if result {
			successCount++
		} else {
			failureCount++
		}
	}

	// 結果を表示
	fmt.Printf("Success: %d, Failure: %d\n", successCount, failureCount)
}
