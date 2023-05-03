package main

import "log"

func main() {
	svc := NewCatFactServer("http://catfact.ninja/fact")
	svc = NewLoggingService(svc)

	apiserver := NewApiServer(svc)
	log.Fatal(apiserver.Start(":3000"))

}
