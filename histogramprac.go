package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

)
func init() { 
	prometheus.MustRegister(random_val)
}
var (
	random_val= prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "random_val",
		Help:    "this is used as random",
		Buckets: prometheus.LinearBuckets(20, 2, 3),  
	})
)

func main(){
	
	for i := 0; i < 1000; i++ {
		random_val.Observe(float64(i))
	}

	fmt.Println("server started at port 9000")
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		panic(err)
	}

}