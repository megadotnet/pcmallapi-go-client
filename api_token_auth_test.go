package main

import (
	sw "app/swagger"
	"encoding/base64"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/antihax/optional"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

const SMS_CODE string = "898017"
const ACCOUNT_PREFIX = "+218 91"

func TestLogin(t *testing.T) {

	loginModel := sw.WctApiModelsTokenAuthAuthenticateModel{
		LoginType: 0,
		UserName:  base64.StdEncoding.EncodeToString([]byte("kmdxdk1@hotmail.com")),
		Password:  base64.StdEncoding.EncodeToString([]byte("1234qwer")),
		SmsCode:   SMS_CODE,
		ImageKey:  "2323",
		ImageCode: "8980"}

	loginModelpost := sw.TokenAuthApiApiTokenAuthAuthenticatePostOpts{
		Body: optional.NewInterface(loginModel),
	}

	resp, r, err := client.TokenAuthApi.ApiTokenAuthAuthenticatePost(context.Background(), &loginModelpost)
	assert := assert.New(t)
	if err != nil {
		t.Errorf("Error while get Login")
		t.Log(err)
	} else {
		assert.NotNil(resp.Result.AccessToken, "AccessToken should  not be null")
		//t.Log(resp.Result)
	}
	if r.StatusCode != 200 {
		t.Log(err)
	}

}

func TestRegister(t *testing.T) {

	registerModel := sw.WctApiApplicationAuthorizationDtoRegisterModel{
		BindAccount: ACCOUNT_PREFIX + strconv.Itoa(GenerateRandInt(1, 9999999)),
		SmsCode:     SMS_CODE,
		Password:    "@234qwer"}

	registerModelpost := sw.TokenAuthApiApiTokenAuthRegisterPostOpts{
		Body: optional.NewInterface(registerModel),
	}

	resp, r, err := client.TokenAuthApi.ApiTokenAuthRegisterPost(context.Background(), &registerModelpost)
	assert := assert.New(t)
	if err != nil {
		t.Errorf("Error while get TestRegister")
		t.Log(err)
	} else {
		t.Log(registerModel.BindAccount)
		assert.NotNil(resp, "Result should  not be null")
	}
	if r.StatusCode != 200 {
		t.Log(err)
	}

}

func GenerateRandInt(min, max int) int {
	rand.Seed(time.Now().UnixNano()) //随机种子
	return rand.Intn(max-min) + min
}
