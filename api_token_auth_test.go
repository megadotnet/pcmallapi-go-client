package main

import (
	sw "app/swagger"
	"encoding/base64"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
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

	resp, r, err := CreateRegisterProc(t)
	assert := assert.New(t)
	if err != nil {
		t.Errorf("Error while get TestRegister")
		t.Log(err)
	} else {
		assert.NotNil(resp, "Result should  not be null")
	}
	if r.StatusCode != 200 {
		t.Log(err)
	}
}

func TestRegisterConcurrency(t *testing.T) {
	errc := make(chan error)
	sum := 0
	for i := 0; i <= 100; i++ {
		go func() {
			resp, r, err := CreateRegisterProc(t)
			if err != nil {
				t.Errorf("Error while get TestRegister")
				t.Log(err, r, resp)
			}
			errc <- err
		}()
		sum++

	}
	waitOnFunctions(t, errc, sum)
}

//create random customer accout by phonenumber
func CreateRegisterProc(t *testing.T) (bool, *http.Response, error) {

	registerModel := sw.WctApiApplicationAuthorizationDtoRegisterModel{
		BindAccount: RandStringBytesMaskImprSrcSB(10) + "@gmail.com",
		SmsCode:     SMS_CODE,
		Password:    "@234qwer"}

	registerModelpost := sw.TokenAuthApiApiTokenAuthRegisterPostOpts{
		Body: optional.NewInterface(registerModel),
	}
	t.Log(registerModel.BindAccount)

	return client.TokenAuthApi.ApiTokenAuthRegisterPost(context.Background(), &registerModelpost)
}

func GeneratePhoneNumber() string {
	return ACCOUNT_PREFIX + strconv.Itoa(GenerateRandInt(1, 9999999))
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func RandStringBytesMaskImprSrcSB(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			sb.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return sb.String()
}

func GenerateRandInt(min, max int) int {
	rand.Seed(time.Now().UnixNano()) //随机种子
	return rand.Intn(max-min) + min
}
