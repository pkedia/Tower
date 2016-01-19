// WebServer.go
package main

import (
    "fmt"
	"time"
    "net/http"
	"net/url"
	"strconv"
	"reflect"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func runtower(w http.ResponseWriter, r *http.Request) {
	var moves [200]int
	var nr int
	var resp string
	fmt.Println(r.Method)
	fmt.Println(r.Host)
	fmt.Println(r.PostForm)
	fmt.Println(r.URL)
	u := r.URL.RawQuery
	fmt.Println("u=",u)
	m, _ := url.ParseQuery(u)
    fmt.Println("m",m)
    v := m["rings"]
	fmt.Println(v[0])
	sv := v[0]
	fmt.Println(sv,reflect.TypeOf(sv),len(sv),sv[0])
	nr,_ = strconv.Atoi(sv)
	nm := TowerOfBramha(nr,&moves)
	fmt.Println("No. of moves: ",nm)
	for i := 0; i < nm;i += 10 {
		for j := i; j < i + 10 && j < nm; j++ {
		   fmt.Print(j,":",moves[j],";")
		}
	}
	fmt.Println()
	resp += "<?xml version=\"1.0\" encoding=\"UTF-8\"?>"
	resp += "<moves>"
	for i := 0; i < nm ; i++ {
		mv := strconv.Itoa(moves[i])
		resp += "<move>"+mv+"</move>"
	}
	resp += "</moves>"
	fmt.Println(resp)
	fmt.Fprintf(w, resp) // write data to response
}

func main() {
	fmt.Println("We are up and listening"+time.Now().Format(time.Kitchen))
    http.Handle("/", http.FileServer(http.Dir("./data")))
    http.HandleFunc("/Tower", runtower)
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
		fmt.Println("Error in ListenAndServe",err)
	}
	fmt.Println("served at"+time.Now().Format(time.Kitchen))
}