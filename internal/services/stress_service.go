package services

import (
	"log"
	"net/http"
	"sync"
	"time"
)

type StressServiceInterface interface {
	Run(url string, requests int, concurrency int) error
}

type StressService struct {
	results      map[int]int
	errors       int
	mutex        sync.Mutex
	requestsSent int
}

func NewStressService() *StressService {
	return &StressService{
		results:      make(map[int]int),
		errors:       0,
		mutex:        sync.Mutex{},
		requestsSent: 0,
	}
}

func (s *StressService) Run(url string, requests int, concurrency int) error {
	log.Printf("Stress testing %s with %d requests and %d concurrent workers\n", url, requests, concurrency)

	requestsPerWorker := requests / concurrency
	extraRequests := requests % concurrency

	startTime := time.Now()

	wg := sync.WaitGroup{}
	wg.Add(concurrency)

	for i := 0; i < concurrency; i++ {
		go func(workerId int) {
			defer wg.Done()

			totalRequests := requestsPerWorker

			if workerId == concurrency-1 {
				totalRequests += extraRequests
			}

			for j := 0; j < totalRequests; j++ {
				res, err := http.Get(url)
				if err != nil {
					s.mutex.Lock()
					s.requestsSent++
					s.errors++
					s.mutex.Unlock()
					continue
				}

				defer res.Body.Close()

				s.mutex.Lock()
				s.requestsSent++
				s.results[res.StatusCode]++
				s.mutex.Unlock()
			}
		}(i)
	}

	wg.Wait()

	log.Println("Stress test completed in", time.Since(startTime))
	log.Println("Total requests:", s.requestsSent)
	log.Println("Total requests with errors:", s.errors)

	for code, count := range s.results {
		log.Printf("Status code %d received %d times\n", code, count)
	}

	return nil
}
