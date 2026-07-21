package main

import ("fmt"

        "time"

)

 

func ShowGoroutine(id int) {

    fmt.Print("Goroutine #%d\n", id)

}

func main(){

     ShowGoroutine(1)

}

func main(){

    go ShowGoroutine(1)

}

func main(){

    go ShowGoroutine(1)

    time.Sleep(1 * time.Second)

}

package main

Import ("fmt"

        "math/rand"

        "time"

)

 

func ShowGoroutine(id int) {

    delay := rand.Intn(500)

    fmt.Print("Goroutine #%d\n with %dms\n", id, delay)

    time.Sleep(time.Milliseconds * time.Duration(delay))

    

}

func main(){

    for i :=0; i < 10; i++ {

        go ShowGoroutine(1)

    }

    time.Sleep(time.Minute)

}

 