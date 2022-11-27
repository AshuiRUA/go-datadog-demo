package main

import (
	"gopkg.in/DataDog/dd-trace-go.v1/profiler"
	"log"
	"math/rand"
	"sort"
)
import "net/http"

func main() {
	err := profiler.Start(
		profiler.WithService("go-datadog-demo"),
		profiler.WithEnv("dev"),
		profiler.WithVersion("0.0.1"),
		profiler.WithProfileTypes(
			profiler.CPUProfile,
			profiler.HeapProfile,
			// The profiles below are disabled by default to keep overhead
			// low, but can be enabled as needed.

			// profiler.BlockProfile,
			// profiler.MutexProfile,
			// profiler.GoroutineProfile,
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer profiler.Stop()
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		wait1()
		writer.Write([]byte("Hello"))
		wait2()
	})
	http.ListenAndServe("localhost:8080", nil)
}

func wait1() {
	var numList = make([]float64, 0)
	for i := 0; i < 10000000; i++ {
		numList = append(numList, rand.Float64())
	}
	sort.Float64s(numList)
}
func wait2() {
	var numList = make([]float64, 0)
	for i := 0; i < 5000000; i++ {
		numList = append(numList, rand.Float64())
	}
	sort.Float64s(numList)
}
