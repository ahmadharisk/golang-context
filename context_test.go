package golang_context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()
	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")
	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")
	contextF := context.WithValue(contextC, "f", "F")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)

	fmt.Println(contextD.Value("d"))
	fmt.Println(contextF.Value("f"))
	fmt.Println(contextF.Value("c")) // mengambil nilai dari parent, karena tidak ada key c
}

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
			}
		}
	}()

	return destination
}

func TestCounterWithCancel(t *testing.T) {
	fmt.Println("Total Goroutines:", runtime.NumGoroutine())
	ctx, cancel := context.WithCancel(context.Background())

	destination := CreateCounter(ctx)

	fmt.Println("Total Goroutines:", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("counter:", n)
		if n == 10 {
			break
		}
	}

	cancel()
	time.Sleep(2 * time.Second)
	fmt.Println("Total Goroutines:", runtime.NumGoroutine())
}
