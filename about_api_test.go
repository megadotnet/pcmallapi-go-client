package main

import (
	sw "app/swagger"
	"os"
	"testing"

	"golang.org/x/net/context"
)

var client *sw.APIClient

const testHost = "lby-stage.flyhome.net/almadar-stage/libya-mall-backend-api"

func TestMain(m *testing.M) {
	cfg := sw.NewConfiguration()
	cfg.AddDefaultHeader("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJodHRwOi8vc2NoZW1hcy54bWxzb2FwLm9yZy93cy8yMDA1LzA1L2lkZW50aXR5L2NsYWltcy9uYW1laWRlbnRpZmllciI6IjgxIiwiaHR0cDovL3NjaGVtYXMueG1sc29hcC5vcmcvd3MvMjAwNS8wNS9pZGVudGl0eS9jbGFpbXMvbmFtZSI6InN1cGVyIiwiQXNwTmV0LklkZW50aXR5LlNlY3VyaXR5U3RhbXAiOiJZSlpJVVZDVFZKT0gyQkZQVjc3WkdCTVZXTllQVlkyVCIsImh0dHA6Ly9zY2hlbWFzLm1pY3Jvc29mdC5jb20vd3MvMjAwOC8wNi9pZGVudGl0eS9jbGFpbXMvcm9sZSI6IkFkbWluIiwic3ViIjoiODEiLCJqdGkiOiJiNmNlYWJiMi0wNDMwLTQ5YTQtODljNC1kMGY0MGYwOTBkZmYiLCJpYXQiOjE2NDY2MTU4NTQsInRva2VuX3ZhbGlkaXR5X2tleTphZG1pbiI6ImY3NzAwZjM3LWMzZmItNDU4MC1hMjM1LTk1ZDc0NDk4N2FhZCIsInVzZXJfaWRlbnRpZmllciI6IjgxIiwibmJmIjoxNjQ2NjE1ODU0LCJleHAiOjE2NDcyMjA2NTQsImlzcyI6IkFkbWluIiwiYXVkIjoiQWRtaW4ifQ.uGhQqxMwWT-0SZ5b-FAFkT0PAyIgS4jULzcVvoAQYC4; language=en; IOEB=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJodHRwOi8vc2NoZW1hcy54bWxzb2FwLm9yZy93cy8yMDA1LzA1L2lkZW50aXR5L2NsYWltcy9uYW1laWRlbnRpZmllciI6IjQ3Iiwic3ViIjoiNDciLCJqdGkiOiIzYjAzOTE1OS00YmQ4LTQ0YTctOTQ0Ny1jMmEwNWVkMjUyZWEiLCJpYXQiOjE2NDY2MTc3MjMsInRva2VuX3ZhbGlkaXR5X2tleTphcGkiOiJmOTU3YTAxMC1mMWRmLTQ1MWQtYWIyZS1mYjAzODQ5MmY4MjQiLCJ1c2VyX2lkZW50aWZpZXIiOiI0NyIsIm5iZiI6MTY0NjYxNzcyMywiZXhwIjoxNjc4MTUzNzIzLCJpc3MiOiJWUEhvcm5vciIsImF1ZCI6IlZQSG9ybm9yIn0.H6Gn1iVzoQdNlGKS9e4Mk493RswDc3Z0yy2NGlXuFFs")
	cfg.Host = testHost
	client = sw.NewAPIClient(cfg)
	retCode := m.Run()
	os.Exit(retCode)
}

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
	for i := 0; i <= 10; i++ {
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

func waitOnFunctions(t *testing.T, errc chan error, n int) {
	for i := 0; i < n; i++ {
		err := <-errc
		if err != nil {
			t.Errorf("Error performing concurrent test")
			t.Log(err)
		}
	}
}
