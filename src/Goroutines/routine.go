package main

import (
	"fmt"
	"math/rand"

	"time"
)

// var m = sync.RWMutex{}
// var wg = sync.WaitGroup{}
// var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
// var results = []string{}

// func main() {
// 	t0 := time.Now()
// 	for i := range dbData {
// 		wg.Add(1)
// 		go dbCall(i)
// 	}
// 	wg.Wait()
// 	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
// 	fmt.Printf("\nThe results are %v", results)
// }

// func dbCall(i int) {
// 	// Simulate DB call delay
// 	var delay float32 = 2000
// 	time.Sleep(time.Duration(delay) * time.Millisecond)
// 	save(dbData[i])
// 	log()
// 	defer wg.Done()
// }

// func save(result string) {
// 	m.Lock()
// 	results = append(results, result)
// 	m.Unlock()
// }

// func log() {
// 	m.RLock()
// 	fmt.Println("The result from the database is:", results)
// 	m.RUnlock()
// }
// Channels
// func main() {
// 	var c = make(chan int, 5)
// 	go process(c)
// 	for i := range c {
// 		fmt.Println(i)
// 		time.Sleep(time.Second * 1)
// 	}
// }

// func process(c chan int) {
// 	defer close(c)
// 	for i := range 5 {
// 		c <- i
// 	}
// 	fmt.Println("Exiting process")
// }

var MAX_CHICKEN_PRICE float32 = 5
var MAX_T0FU_PRICE float32 = 3

func main() {
	var chickenChannel = make(chan string)
	var tofuChannel = make(chan string)
	var websites = []string{"walmart.com", "costco.com", "wholefoods.com"}
	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel)
		go checkTofuPrices(websites[i], tofuChannel)
	}
	sendMessage(chickenChannel, tofuChannel)
}

func checkChickenPrices(website string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32() * 20
		if chickenPrice <= MAX_CHICKEN_PRICE {
			chickenChannel <- website
			break
		}
	}
}

func checkTofuPrices(website string, c chan string) {
	for {
		time.Sleep(time.Second * 1)
		var tofu_price = rand.Float32() * 20
		if tofu_price <= MAX_T0FU_PRICE {
			c <- website
			break
		}
	}
}

func sendMessage(chickenChannel chan string, tofuChannel chan string) {
	select {
	case website := <-chickenChannel:
		fmt.Printf("\nText Sent: Found deal on chicken at %v.", website)
	case website := <-tofuChannel:
		fmt.Printf("\nEmail Sent: Found deal on tofu at %v.", website)
	}
}
