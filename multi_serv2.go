package main

import (
    "fmt"
    "net/http"
    "math/rand"
    "strconv"
)


type Bar struct{}

func (f *Bar) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    fmt.Println("Invoking Bar response")
    numstr := strconv.Itoa(rand.Intn(1000))
    w.Write([]byte("node_bar_value{instance='myinstance.com' , job='myjob'} " + numstr + "\n"))
}


func start_http_server(iport int) {
    portstr := ":" + strconv.Itoa(iport) 
    go http.ListenAndServe(portstr, &Bar{})
    fmt.Println("Started serving on localhost port " + portstr)

}

func main() {

    baseportnum := 8080
    minport     := 8081
    maxport     := 9080
    
    for i := minport; i < maxport ; i++ {
        start_http_server (i)
    }

    fmt.Println("Started serving on localhost port " + ":8080")
    portstr := ":" + strconv.Itoa(baseportnum) 
    http.ListenAndServe(portstr, &Bar{})

    // go http.ListenAndServe(":8080", &Bar{})
    // go http.ListenAndServe(":8081", &Foo{})
    
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