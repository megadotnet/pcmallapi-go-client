package main

import (
	"testing"

	"golang.org/x/net/context"
)

func TestGetAbout(t *testing.T) {
	data, r, err := client.AboutApi.ApiServicesAppAboutDetailPost(context.Background())

	if err != nil {
		t.Errorf("Error while get about detail")
		t.Log(err)
	}
	if r.StatusCode != 200 {
		t.Log(err)
	}
	if r.StatusCode == 200 {
		t.Log(data)
	}
}

// Test we can concurrently create, retrieve, update, and delete.
func TestConcurrency(t *testing.T) {
	errc := make(chan error)

	// get the about
	sum := 0
	for i := 0; i <= 1; i++ {
		go func() {
			data, r, err := client.AboutApi.ApiServicesAppAboutDetailPost(context.Background())
			if r.StatusCode != 200 {
				t.Log(r)
				t.Log(data)
			}
			errc <- err
		}()
		sum++
	}
	waitOnFunctions(t, errc, sum)
}
