package main

import "fmt"
import "net/http"
import "time"

// func Racer(slow, fast string) (winner string) {
// 	slowDuration := measureResponseTime(slow)
// 	fastDuration := measureResponseTime(fast)
//
// 	if slowDuration < fastDuration {
// 		return slow
// 	}
//
// 	return fast
// }

// func measureResponseTime(url string) time.Duration {
// 	start := time.Now()
// 	http.Get(url)
// 	return time.Since(start)
// }

const tenSecondTimeout = 10 * time.Second

func Racer(slow, fast string) (winner string, err error) {
	return ConfigurableRacer(slow, fast, tenSecondTimeout)
}

func ConfigurableRacer(slow, fast string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(slow):
		return slow, nil
	case <-ping(fast):
		return fast, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", slow, fast)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
