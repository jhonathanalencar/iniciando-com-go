package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// func soma(x int, y int) (int, bool) {
// 	if x > 10 {
// 		return x + y, true
// 	}

// 	return x + y, false
// }

type Course struct {
	Name string `json:"course"`
	Description string `json:"description"`
	Price int `json:"price"`
}

func (c Course) GetFullInfo() string {
	return fmt.Sprintf("Name: %s, Description: %s, Price: %d", c.Name, c.Description, c.Price)
}

func counter(){
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func worker(workerID int, data chan int) {
	for x:= range data {
		fmt.Printf("Worker %d received %d\n", workerID, x)
		time.Sleep(time.Second)
	}
}

// Thread 1
func main() {
	// resultado, status := soma(10, 20)
	
	// course := Course{
	// 	Name: "Golang",
	// 	Description: "Golang Course",
	// 	Price: 100,
	// }
	
	// fmt.Println(course.GetFullInfo())

	// http.HandleFunc("/", home)
	// http.ListenAndServe(":8080", nil)

	// go counter()
	// go counter()
	// counter()

	// channel := make(chan string)

	// go func() {
	// 	channel <- "Hello World!"
	// }()

	// fmt.Println(<-channel)

	channel := make(chan int)
	go worker(1, channel) // Thread 2
	go worker(2, channel) // Thread 3
	go worker(3, channel) // Thread 4

	for i := 0; i < 100; i++ {
		channel <- i
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Hello World"))

	course := Course{
		Name: "Golang",
		Description: "Golang Course",
		Price: 100,
	}

	json.NewEncoder(w).Encode(course)

	res, _ := json.Marshal(course)
	var data Course
	err := json.Unmarshal(res, &data)
	if err != nil{
		log.Fatalln("error:", err)
	}
	json.NewEncoder(w).Encode(data)
}