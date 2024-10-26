package main

import "github.com/mrngsht/realworld-goa-react/server"

func main() {
	if err := server.Run(); err != nil {
		panic(err)
	}
}
