package main

import (
	sw "app/swagger"
	"testing"

	"github.com/antihax/optional"
	//"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

func TestOrderLogisticsSearch(t *testing.T) {

	ordersearchModel := sw.AbpApplicationServicesDtoEntityDto1SystemInt64{
		Id: 25029416043905024}

	ordersearchModelpost := sw.OrderApiApiServicesAppOrderLogisticsSearchPostOpts{
		Body: optional.NewInterface(ordersearchModel),
	}

	resp, r, err := client.OrderApi.ApiServicesAppOrderLogisticsSearchPost(context.Background(), &ordersearchModelpost)
	//	assert := assert.New(t)
	if err != nil {
		t.Errorf("Error while get Login")
		t.Log(err)
	} else {
		//assert.NotNil(resp.Result, "Result should  not be null")
		t.Log(resp)
	}
	if r.StatusCode != 200 {
		t.Log(err)
	}

}
