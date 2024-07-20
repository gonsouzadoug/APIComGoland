package main

import {
	"log"
	"net/http"
}

func HelloAPI(httpResponse http.ResponseWriter, httpRequest *http.Request){
	var message string = "Hello API !"
	httpResponse.Write([]byte("{\"message\": \"" + message + "\" }"))
}

func main(){
	http.HandleFunc("/hello", HelloAPI)
	log.Panic(http.ListenerAndServer(":8080", nill))
}