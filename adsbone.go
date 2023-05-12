package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

var adsboneInstance *adsbone

type adsbone struct {
	t    *time.Ticker
	done chan bool
}

func newAdsbone() {
	adsboneInstance = &adsbone{
		t:    time.NewTicker(10 * time.Second),
		done: make(chan bool, 1),
	}
}

func (s *adsbone) start() {
	go func() {
		for {
			select {
			case <-s.t.C:
				go s.collectData()
			case <-s.done:
				return
			}
		}
	}()
}

func (s *adsbone) stop() {
	s.t.Stop()
	s.done <- true
}

func (s *adsbone) collectData() {
	// SLC coords.
	resp, err := http.Get("https://api.adsb.one/v2/point/40.76/-111.88/250")
	if err != nil {
		log.Println("cannot get data from adsbone:", err)
		return
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("cannot read response body from adsbone:", err)
		return
	}

	createAdsboneRecord(data)
}
