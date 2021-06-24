package pinger

import (
	"time"
)

func Execute(project string, location string, pod string, ips []string) error {
	tickerC := map[string]<-chan time.Time{}
	for _, ip := range ips {
		ticker := time.NewTicker(60 * time.Second)
		defer ticker.Stop()
		tickerC[ip] = ticker.C
	}
	aggTickerC := make(chan string)
	for _, ip := range ips {
		go func(i string, c <-chan time.Time) {
			for range c {
				aggTickerC <- i
			}
		}(ip, tickerC[ip])
	}
	for {
		select {
		case ip := <-aggTickerC:
			rtt, err := singlePing(ip)
			if err != nil {
				return err
			}
			err = export(project, location, pod, ip, rtt)
			if err != nil {
				return err
			}
		}
	}
}
