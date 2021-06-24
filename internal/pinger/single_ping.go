package pinger

import (
	"fmt"
	"time"

	"github.com/go-ping/ping"
)

func singlePing(ip string) (int64, error) {
	pinger, err := ping.NewPinger(ip)
	if err != nil {
		return -1, err
	}
	pinger.Count = 1
	pinger.Timeout = time.Second
	pinger.SetPrivileged(true)
	err = pinger.Run()
	if err != nil {
		fmt.Println(err)
		return -1, nil
	}
	stats := pinger.Statistics()
	rtt := int64(stats.MaxRtt / time.Millisecond)
	return rtt, nil
}
