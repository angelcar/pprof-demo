package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)

func main() {
	// we need a webserver to get the pprof webserver
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	fmt.Println("listening on localhost:6060")
	var wg sync.WaitGroup
	wg.Add(1)
	//go leakyMemFunction(wg)
	go leakyGoRoutineFunction(wg)
	go otherGoRoutineFunction(wg)
	wg.Wait()
}

type foo struct{
	bar string
}

func leakyMemFunction(wg sync.WaitGroup) {
	//defer wg.Done()
	s := make([]*foo, 3)
	for i:= 0; i < 10000000; i++{
		s = append(s, &foo{
			bar: "leak",
		})
		if (i % 100000) == 0 {
			time.Sleep(500 * time.Millisecond)
		}
	}
}


func leakyGoRoutineFunction(wg sync.WaitGroup) chan bool {
	//defer wg.Done()
	c := make(chan bool, 1)
	for i := 0; i < 10000; i++ {
		go func (){
			s := make([]string, 3)
			for i := 0; i < 1000; i++ {
				s = append(s, "leak")
			}
			c <- true
		}()
		time.Sleep(time.Millisecond * 100)
	}

	return c
}

func otherGoRoutineFunction(wg sync.WaitGroup) chan bool {
	//defer wg.Done()
	c := make(chan bool, 1)
	for i := 0; i < 10000; i++ {
		go func (){
			for i := 0; i < 1000; i++ {
				go func() {
					f := &foo{
						bar: "leak",
					}
					fmt.Fprint(io.Discard, f.bar)
				}()
			}
			c <- true
		}()
		time.Sleep(time.Millisecond * 100)
	}
	return c
}
