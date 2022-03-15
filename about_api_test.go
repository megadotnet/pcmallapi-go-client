package main

import (
	sw "app/swagger"
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"golang.org/x/net/context"
)

var client *sw.APIClient

//const testHost = "lby-stage.flyhome.net/almadar-stage/libya-mall-backend-api"

func TestMain(m *testing.M) {
	cfg := sw.NewConfiguration()
	//cfg.Host = testHost

	//login
	loginModel := CreateLoginModel()

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	myclient := &http.Client{Transport: tr}
	payloadBytes, err := json.Marshal(loginModel)
	if err != nil {
		// handle err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", cfg.BasePath+"api/TokenAuth/Authenticate", body)
	if err != nil {
		// handle err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := myclient.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()
	localVarBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err)
	}
	//fmt.Printf("read resp.Body successfully:\n%v\n", string(localVarBody))
	var respObj sw.WctApiModelsTokenAuthAuthenticateResultModel
	jsonerr := json.Unmarshal(localVarBody, &respObj)
	if jsonerr != nil {
		fmt.Println("error:", jsonerr)
		return
	}

	cfg.AddDefaultHeader("Authorization", "Bearer "+respObj.Result.AccessToken)
	client = sw.NewAPIClient(cfg)
	retCode := m.Run()
	os.Exit(retCode)
}

func CreateLoginModel() sw.WctApiModelsTokenAuthAuthenticateModel {
	return sw.WctApiModelsTokenAuthAuthenticateModel{
		LoginType: 0,
		UserName:  base64.StdEncoding.EncodeToString([]byte("ZtfyRZxPsG@gmail.com")),
		Password:  base64.StdEncoding.EncodeToString([]byte(PASSWORD)),
		SmsCode:   SMScode,
		ImageKey:  "2323",
		ImageCode: "8980"}
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
