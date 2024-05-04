/**
 * @Author: Cc
 * @Description: prometheus Client
 * @File: prometheus
 * @Version: 1.0.0
 * @Date: 2022/8/2 10:21
 * @Software : GoLand
 */

package prometheus

import (
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})
)

func InitStart(address string) {
	recordMetrics()
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
