package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	firstCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "first_total",
			Help: "first Counter",
		}, []string{"example", "count"},
	)

	firstGage = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "first_number",
			Help: "first Number",
		}, []string{"example", "gage"},
	)
)

func count() {
	// Labelsのmap[string]stringはvarで定義したものと一致する必要がある
	for {
		firstCounter.With(prometheus.Labels{"count": "increment", "example": "true"}).Inc()
		time.Sleep(10 * time.Second)
	}
}

func setRandomValue() {
	for {
		rand.Seed(time.Now().UnixNano())
		n := -1 + rand.Float64()*2
		firstGage.With(prometheus.Labels{"gage": "integer", "example": "true"}).Set(n)
		time.Sleep(10 * time.Second)
	}
}

func main() {
	go count()
	go setRandomValue()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
