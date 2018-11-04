package main

import ("fmt"
"time"
"sync"
"math/rand")

func reallyxpensiveFunc(respond chan<- int,wg *sync.WaitGroup, number int) {
  defer wg.Done()
  timeConsumed := rand.Intn(7)
  time.Sleep(time.Duration(timeConsumed) * time.Second)
  fmt.Printf("finished %d, %d \n", number, timeConsumed)
  respond <- number
}

func main() {
  a := 1
  b := 2

  operationDone := make(chan bool)

  go func() {
    time.Sleep(time.Duration(1) * time.Second)
    b = a * b

    operationDone <- true
  } ()

  is_true := <-operationDone
  fmt.Println(is_true)

  a = b * b

  fmt.Println("end")
  fmt.Printf("a = %d, b = %d\n", a, b)

  operations := make(chan int, 5)

  var wg sync.WaitGroup

wg.Add(5)
  go reallyxpensiveFunc(operations, &wg, 1)
  go reallyxpensiveFunc(operations, &wg,2)
  go reallyxpensiveFunc(operations, &wg,3)
  go reallyxpensiveFunc(operations, &wg,4)
  go reallyxpensiveFunc(operations, &wg,5)

  wg.Wait()
  close(operations)
  for resp := range operations {
    fmt.Println(resp)
  }
}
