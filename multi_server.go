package main

import (
    "log"
    "net/http"
    "sync"
    "strconv"
)

type Foo struct{}

func (f *Foo) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    w.Write([]byte("Foo"))
}

type Bar struct{}

func (f *Bar) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    w.Write([]byte("Bar"))
}

func main() {
    portnum := 8080
    portstr := ""
    wg := &sync.WaitGroup{}

    for i := 0; i < 10 ; i++ {
        portstr = ":" + strconv.Itoa( (portnum + i) )
        print(portstr + "\n")
        wg.Add(1)
        go func() {
            log.Fatal(http.ListenAndServe(portstr, &Bar{}))
            wg.Done()
        }()

    }

    wg.Wait()
}



/*
package main

import (
    "net/http"
    "fmt"
    "log"
)

func main() {

    // Show on console the application stated
    log.Println("Server started on: http://localhost:9000")
    main_server := http.NewServeMux()

    //Creating sub-domain
    server1 := http.NewServeMux()
    server1.HandleFunc("/metrics", server1func)

    server2 := http.NewServeMux()
    server2.HandleFunc("/metrics", server2func)

    //Running First Server
    go func() {
        log.Println("Server started on: http://localhost:9001")
        http.ListenAndServe("localhost:9001", server1)
    }()

    //Running Second Server
    go func() {
        log.Println("Server started on: http://localhost:9002")
        http.ListenAndServe("localhost:9002", server2)
    }()

    //Running Main Server
    http.ListenAndServe("localhost:9000", main_server)
}


func server1func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Running First Server")
}

func server2func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Running Second Server")
}

*/
/*

package main

import (
    "fmt"
    //"html"
    "log"
    "net/http"
)

func main() {

    http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello Prometheus scraper!!")
    })
    


    log.Fatal(http.ListenAndServe(":8081", nil))
    log.Fatal(http.ListenAndServe(":8082", nil))
    log.Fatal(http.ListenAndServe(":8083", nil))

}

*/