package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

var adsblolInstance *adsblol

type adsblol struct {
	t    *time.Ticker
	done chan bool
}

func newAdsblol() {
	adsblolInstance = &adsblol{
		t:    time.NewTicker(10 * time.Second),
		done: make(chan bool, 1),
	}
}

func (s *adsblol) start() {
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

func (s *adsblol) stop() {
	s.t.Stop()
	s.done <- true
}

func (s *adsblol) collectData() {
	// LA coords.
	resp, err := http.Get("https://api.adsb.lol/v2/point/34.05/-118.24/250")
	if err != nil {
		log.Println("cannot get data from adsblol:", err)
		return
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("cannot read response body from adsblol:", err)
		return
	}

	createAdsblolRecord(data)
}
