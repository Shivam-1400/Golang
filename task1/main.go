package main

import (
    "fmt"
    "sync"
)

func main() {
    list := []int{1, 2, 3, 4, 5}
    fmt.Println("List:", list)

    // go map
    m := make(map[string]int)
    m["a"] = 1
    m["b"] = 2
    m["c"] = 3
    fmt.Println("Map:", m)

    // go set using map
    set := make(map[int]bool)
    set[1] = true
    set[2] = true
    set[3] = true
    fmt.Println("Set:", set)

    // go concurrent map
    var concurrentMap sync.Map
    concurrentMap.Store("a", 1)
    concurrentMap.Store("b", 2)
    concurrentMap.Store("c", 3)

    value, _ := concurrentMap.Load("a")
    fmt.Println("Concurrent Map Value for key 'a':", value)

    // Thread pool implementation
    poolSize := 3
    tasks := []func() int{
        func() int { return 1 },
        func() int { return 2 },
        func() int { return 3 },
        func() int { return 4 },
        func() int { return 5 },
    }

    resultChan := make(chan int)
    var wg sync.WaitGroup
    for i := 0; i < poolSize; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for task := range tasks {
                result := tasks[task]()
                resultChan <- result
            }
        }()
    }

    go func() {
        wg.Wait()
        close(resultChan)
    }()

    var results []int
    for result := range resultChan {
        results = append(results, result)
    }

    fmt.Println("Results:", results)
}
