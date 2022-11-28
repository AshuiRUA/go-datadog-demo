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
	http.HandleFunc("/callOneFunc", func(writer http.ResponseWriter, request *http.Request) {
		randomDoubleSort(1000000)
		writer.Write([]byte("call one func"))
		randomDoubleSort(500000)
	})

	http.HandleFunc("/callTwoFunc", func(writer http.ResponseWriter, request *http.Request) {
		randomIntSort(1000000)
		writer.Write([]byte("call Two func"))
		randomDoubleSort(500000)
	})
	http.ListenAndServe("localhost:8080", nil)
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
