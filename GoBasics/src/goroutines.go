package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
var m = sync.Mutex{}
var results = []string{}
var dbCalls = []string{"id1", "id2", "id3", "id4", "id5"}

func concurrency() {
	t0 := time.Now()
	for i := 0; i < len(dbCalls); i++ {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Printf("\n Total execution time: %v | Results are %v", time.Since(t0), results)
}

func dbCall(i int) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	m.Lock() // Otherwise memory corruption
	results = append(results, dbCalls[i])
	m.Unlock()
	wg.Done()
}
