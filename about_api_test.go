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
	cfg.AddDefaultHeader("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJodHRwOi8vc2NoZW1hcy54bWxzb2FwLm9yZy93cy8yMDA1LzA1L2lkZW50aXR5L2NsYWltcy9uYW1laWRlbnRpZmllciI6IjU3Iiwic3ViIjoiNTciLCJqdGkiOiIwN2ZiOTlkNC00NzgxLTQ5YmUtOTY0MC02MGNlOGUxM2EzZWUiLCJpYXQiOjE2NDczMjU5NDEsInRva2VuX3ZhbGlkaXR5X2tleTphcGkiOiJhZDhlMWFlYy02ZDBhLTQ5ZmEtYmY4NS02MDhmYTk0ZjFmY2QiLCJ1c2VyX2lkZW50aWZpZXIiOiI1NyIsIm5iZiI6MTY0NzMyNTk0MSwiZXhwIjoxNjc4ODYxOTQxLCJpc3MiOiJWUEhvcm5vciIsImF1ZCI6IlZQSG9ybm9yIn0.tgoCFWZVqanjBeHatVt0gqkpU7ptpHON6PmJP1_SFPQ")
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

func waitOnFunctions(t *testing.T, errc chan error, n int) {
	for i := 0; i < n; i++ {
		err := <-errc
		if err != nil {
			t.Errorf("Error performing concurrent test")
			t.Log(err)
		}
	}
}
