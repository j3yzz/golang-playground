package main

import (
	"context"
	"fmt"
	"time"
)

//
//func hello(w http.ResponseWriter, req *http.Request) {
//	ctx := req.Context()
//
//	childCtx := context.WithValue(ctx, "rand_token", rand.Uint32())
//	manager(childCtx)
//
//	fmt.Println("server: hello handler started")
//	defer fmt.Println("server: hello handler ended")
//
//	select {
//	case <-time.After(10 * time.Second):
//		fmt.Fprintf(w, "hello\n")
//	case <-ctx.Done():
//		err := ctx.Err()
//		fmt.Println("server:", err)
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//}
//
//func manager(ctx context.Context) {
//	fmt.Println(ctx.Value("rand_token"))
//}
//
//func main() {
//	ctx := context.Background()
//	cancelCtx, cancelFunc := context.WithCancel(ctx)
//	go task(cancelCtx)
//	time.Sleep(time.Second * 3)
//	cancelFunc()
//	time.Sleep(time.Second * 1)
//
//	http.HandleFunc("/hello", hello)
//	http.ListenAndServe(":8080", nil)
//}
//
//func task(ctx context.Context) {
//	i := 1
//	for {
//		select {
//		case <-ctx.Done():
//			fmt.Println("Gracefully exit")
//			fmt.Println(ctx.Err())
//			return
//		default:
//			fmt.Println(i)
//			time.Sleep(time.Second * 1)
//			i++
//		}
//	}
//}

func main() {
	ctx := context.Background()
	cancelCtx, cancelFunc := context.WithTimeout(ctx, time.Second*3)
	defer cancelFunc()
	go task1(cancelCtx)
	time.Sleep(time.Second * 4)
}

func task1(ctx context.Context) {
	i := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Gracefully exit")
			fmt.Println(ctx.Err())
		default:
			fmt.Println(i)
			time.Sleep(time.Second * 1)
			i++
		}
	}
}
