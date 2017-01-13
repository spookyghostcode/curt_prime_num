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

    //Making sure that the argument is passed in
    param = req.URL.Query().Get("max")
    if param == "" {
        errorMsg := ErrorResp{Error: "No number provided"}
        writer.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(writer).Encode(errorMsg)
        return
    }

    //Making sure that the argument is an integer
    max, err := strconv.Atoi(param)
    if err != nil {
        errorMsg := ErrorResp{Error: "Must be integer"}
        writer.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(writer).Encode(errorMsg)
        return
    }

   if (max < 2) || (max > 1000) {
      errorMsg := ErrorResp{Error: "Number cannot be smaller than 2, or larger than 1000"}
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
