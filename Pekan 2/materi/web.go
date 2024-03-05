//server -- struct

package main

server := http.Server{
	addr: "localhost:8081",
}

err:= server.ListenAndServe()
if err !=nil{
	panic(err)
}