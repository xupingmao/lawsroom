package main

import (
	"net/http"

	"github.com/txthinking/lawsroom/signal"
)

func getSignalHandle(origins []string) *signal.Signal {
	signal.ROOM_CAPACITY = 5
	return signal.New(func(r *http.Request) bool {
		return true
		/*
		origin := r.Header.Get("Origin")
		for _, v := range origins {
			if v == origin {
				return true
			}
		}
		return false
		*/
	}, nil)
}
