package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
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
		}, []string{"exapmle", "gage"},
	)
)
