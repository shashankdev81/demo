package main

func main() {
	//ans := mathlib.Divide(10, 5)
	//fmt.Println(ans)
}

//
//import (
//	"demo/mathlib"
//	"fmt"
//	"io/ioutil"
//	"log"
//	"net/http"
//	"sync"
//	"sync/atomic"
//	"time"
//
//	"github.com/sony/gobreaker"
//)
//
//func main() {
//	ans := mathlib.Divide(10, 5)
//	fmt.Println(ans)
//}
//
//// MyClient is a dummy struct to hold the circuit breaker
//type MyClient struct {
//	cb *gobreaker.CircuitBreaker
//}
//type MyInterface interface {
//	doSomething(value int)
//}
//type MyInterfaceImpl1 struct {
//}
//
//func (i *MyInterfaceImpl1) doSomething(val int) {
//	fmt.Printf("Printing value=%v", val*2)
//}
//
//type MyInterfaceImpl2 struct {
//}
//
//func (i *MyInterfaceImpl2) doSomething(val int) {
//	fmt.Printf("Printing value=%v", val*3)
//
//}
//
//type MyStruct struct {
//	name string
//}
//
//type Vertex struct {
//	x int
//	y int
//}
//
//func main2() {
//	var myStructRef *Vertex
//	var myStruct Vertex
//	myStructRef = &Vertex{2, 3}
//	myStruct = Vertex{2, 3}
//
//	fmt.Println(myStructRef)
//	fmt.Println(myStruct)
//	changeByReferenceStruct(myStructRef)
//	changeByReferenceStruct(&myStruct)
//	//changeByValueStruct(myStruct)
//	fmt.Println(myStructRef)
//	fmt.Println(myStruct)
//}
//
//func changeByValueStruct(myStruct Vertex) {
//	myStruct.x = 5
//	fmt.Println(myStruct)
//}
//
//func changeByReferenceStruct(myStruct *Vertex) {
//	myStruct.x = 7
//	fmt.Println(myStruct)
//}
//
//func getTask() MyInterface {
//	impl := &MyInterfaceImpl1{}
//	return impl
//}
//
//// requester is the signature of the http.Client Do() method
//type requester func(*http.Request) (*http.Response, error)
//
//func main1() {
//
//	// - "example" is the name of the circuit breaker
//	// - 2 successful requests reset the state to closed
//	// - 1 failure sets the state to open
//	// - 1 second timeout
//	cb := NewCircuitBreaker("example", 2, 1, 1)
//	c := &MyClient{
//		cb: cb,
//	}
//
//	clientDo := c.addCircuitBreaker(wrapClient(&http.Client{}))
//
//	urls := []string{
//		"https://www.google.com",
//		"https://404.google.com",
//		"https://404.google.com",
//		"https://404.google.com",
//		"https://404.google.com",
//		"https://404.google.com",
//		"https://404.google.com",
//		"https://404.google.com",
//		"https://www.google.com",
//		"https://www.google.com",
//		"https://www.google.com",
//	}
//
//	for i, u := range urls {
//		log.Printf("%d %s", i, cb.State())
//
//		// Give some time to let Open state go to Half-Open (see above 1 second timeout)
//		time.Sleep(time.Millisecond * 700)
//
//		req, err := http.NewRequest(http.MethodGet, u, nil)
//
//		if err != nil {
//			log.Print(err)
//		}
//
//		resp, err := clientDo(req)
//		if err != nil {
//			log.Print(err)
//		}
//
//		if resp == nil {
//			continue
//		}
//		_, err = ioutil.ReadAll(resp.Body)
//		if err != nil {
//			log.Print(err)
//		}
//
//		err = resp.Body.Close()
//		if err != nil {
//			log.Print(err)
//		}
//	}
//}
//
//func NewCircuitBreaker(name string, maxRequests, consecutiveFailures uint32, timeout time.Duration) *gobreaker.CircuitBreaker {
//	return gobreaker.NewCircuitBreaker(gobreaker.Settings{
//		Name:        name,
//		MaxRequests: maxRequests,
//		Timeout:     timeout * time.Second,
//		ReadyToTrip: func(counts gobreaker.Counts) bool {
//			return counts.ConsecutiveFailures > consecutiveFailures
//		},
//		OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
//			// Fill this callback method with what you want
//			// Suggestions :
//			// 1. log state change
//			// 2. emit event
//		},
//	})
//}
//
//// addCircuitBreaker will either :
//// - allow requests and keep tracks of errors
//// - reject requests
//func (c *MyClient) addCircuitBreaker(requestExecuter requester) requester {
//	return func(req *http.Request) (*http.Response, error) {
//		if c.cb != nil {
//			r, err := c.cb.Execute(func() (interface{}, error) {
//				return requestExecuter(req)
//			})
//
//			if err != nil {
//				return nil, err
//			}
//			return r.(*http.Response), nil
//		}
//		return requestExecuter(req)
//	}
//}
//
//// wrapClient returns a closure that converts a http.Client to a requester type
//func wrapClient(client *http.Client) requester {
//	return func(req *http.Request) (*http.Response, error) {
//		return client.Do(req)
//	}
//}
//
//func waitGroupTest() {
//	//fmt.Println("!... Hello World ...!")
//	var a uint32 = 0
//	var wg sync.WaitGroup
//
//	for i := 0; i < 100; i++ {
//		wg.Add(1)
//		go func() {
//			a = atomic.AddUint32(&a, 1)
//			wg.Done()
//		}()
//	}
//	wg.Wait()
//
//	var lock sync.Mutex
//	l := lock
//	l.Lock()
//	l.Unlock()
//
//	f := Foo{lock: sync.Mutex{}}
//	f.Lock()
//}
//
//type Foo struct {
//	lock sync.Mutex
//}
//
//func (f Foo) Lock() {
//	f.lock.Lock()
//}
