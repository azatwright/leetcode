package main

import (
	"encoding/json"
	"os"
	"testing"
)

func TestCriticalConnections(t *testing.T) {
	file, err := os.Open("testdata/1192.json")
	if err != nil {
		panic(err)
	}
	dec := json.NewDecoder(file)
	for dec.More() {
		var label string
		var args struct {
			N           int
			Connections [][]int
		}
		var expect [][]int
		arr := &[]interface{}{&label, &[]interface{}{&args.N, &args.Connections}, &[]interface{}{&expect}}
		if err := dec.Decode(arr); err != nil {
			panic(err)
		}

		t.Run(label, func(t *testing.T) {
			got := criticalConnections(args.N, args.Connections)
			if !intArrArrEqualNoOrder(got, expect) {
				t.Errorf("expected %v got %v", expect, got)
			}
		})
	}
}
