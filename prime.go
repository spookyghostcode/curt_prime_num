package main

import (
    "encoding/json"
    "net/http"
    "log"
    "strconv"

    "github.com/gorilla/mux"
)

//models for JSON encoding
type SuccessResp struct {
    Success    []int
}

type ErrorResp struct {
    Error      string
}

//Function that will be doing the heavy lifting of actually calculating
//the prime numbers
func calcPrimes(writer http.ResponseWriter, req *http.Request) {

    var param string

    param = req.URL.Query().Get("max")
    if param == "" {
        errorMsg := ErrorResp{Error: "Argument 'max' not provided"}
        writer.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(writer).Encode(errorMsg)
        return
    }

    max, err := strconv.Atoi(param)
    if err != nil {
        errorMsg := ErrorResp{Error: "Argument 'max' must be an integer"}
        writer.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(writer).Encode(errorMsg)
        return
    }

    var primeList []int
    var prime bool

    for i := 2; i<=max; i++ {
        prime = true
        for j:= 2; j<i; j++ {
            if i == j {
                continue
            }
            if i % j == 0 {
                prime = false
                break
            }
        }
        if prime == true {
            primeList = append(primeList, i)
        }
    }

    successMsg := SuccessResp{Success: primeList}
    json.NewEncoder(writer).Encode(successMsg)

}

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/primes", calcPrimes).Methods("GET")
    router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
    http.Handle("/", router)

    err := http.ListenAndServe(":80", router)
    if err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}
