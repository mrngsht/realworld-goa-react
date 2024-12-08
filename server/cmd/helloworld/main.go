package main

import (
	"context"
	"fmt"

	"github.com/mrngsht/realworld-goa-react/mytime"
)

func main() {
	now := mytime.Now(context.Background())
	fmt.Println(now)
}
