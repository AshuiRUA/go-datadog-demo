package main

import (
	"fmt"
	httptrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"gopkg.in/DataDog/dd-trace-go.v1/profiler"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"sort"
)

func main() {
	tracer.Start(
		tracer.WithService("go-datadog-demo"),
		tracer.WithEnv("dev"),
	)
	defer tracer.Stop()

	err := profiler.Start(
		profiler.WithService("go-datadog-demo"),
		profiler.WithEnv("dev"),
		profiler.WithProfileTypes(
			profiler.CPUProfile,
			profiler.HeapProfile,

			// The profiles below are disabled by
			// default to keep overhead low, but
			// can be enabled as needed.
			profiler.BlockProfile,
			profiler.MutexProfile,
			profiler.GoroutineProfile,
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer profiler.Stop()

	// Create a traced mux router
	mux := httptrace.NewServeMux()
	// Continue using the router as you normally would.
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	mux.HandleFunc("/callOneFunc", func(writer http.ResponseWriter, request *http.Request) {
		randomIntSort(1000000)
		writer.Write([]byte("call one func"))
	})
	mux.HandleFunc("/callTwoFunc", func(writer http.ResponseWriter, request *http.Request) {
		randomIntSort(1000000)
		randomDoubleSort(500000)
		writer.Write([]byte("call two func"))
	})
	mux.HandleFunc("/alloc", func(writer http.ResponseWriter, request *http.Request) {
		s1 := make([]byte, 1024*1024*1024)
		s2 := make([]byte, 1024*1024*1024)
		s3 := s1[0]
		s4 := s2[0]
		fmt.Println(s3, s4)
	})
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	http.ListenAndServe(":8080", mux)
}

func randomIntSort(len int) {
	var numList = make([]int, 0)
	for i := 0; i < len; i++ {
		numList = append(numList, rand.Int())
	}
	sort.Ints(numList)
}

func randomDoubleSort(len int) {
	var numList = make([]float64, 0)
	for i := 0; i < len; i++ {
		numList = append(numList, rand.Float64())
	}
	sort.Float64s(numList)
}
