package main

import (
	sw "app/swagger"
	"encoding/base64"
	"testing"

	"github.com/antihax/optional"
	"golang.org/x/net/context"
)

func TestLogin(t *testing.T) {

	loginModel := sw.WctApiModelsTokenAuthAuthenticateModel{
		LoginType: 0,
		UserName:  base64.StdEncoding.EncodeToString([]byte("kmdxdk1@hotmail.com")),
		Password:  base64.StdEncoding.EncodeToString([]byte("1234qwer")),
		SmsCode:   "898017",
		ImageKey:  "2323",
		ImageCode: "8980"}

	loginModelpost := sw.TokenAuthApiApiTokenAuthAuthenticatePostOpts{
		Body: optional.NewInterface(loginModel),
	}

	data, r, err := client.TokenAuthApi.ApiTokenAuthAuthenticatePost(context.Background(), &loginModelpost)

	if err != nil {
		t.Errorf("Error while get Login")
		t.Log(err)
	}
	if r.StatusCode != 200 {
		t.Log(err)
	}
	if r.StatusCode == 200 {
		t.Log(data)
	}
}
