package main

import (
	sw "app/swagger"
	"testing"

	"github.com/antihax/optional"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

//安全测试-LogisticsSearch接口不可以查询任意其他人的订单号
func TestOrderLogisticsSearchWithOthersId(t *testing.T) {

	ordersearchModel := sw.AbpApplicationServicesDtoEntityDto1SystemInt64{
		Id: 25228443550859264}

	ordersearchModelpost := sw.OrderApiApiServicesAppOrderLogisticsSearchPostOpts{
		Body: optional.NewInterface(ordersearchModel),
	}

	resp, r, err := client.OrderApi.ApiServicesAppOrderLogisticsSearchPost(context.Background(), &ordersearchModelpost)
	assert := assert.New(t)
	if err != nil {
		//t.Log(err)
	} else {
		assert.Equal(r.StatusCode, 500, "Shoule equal http 500")
		assert.False(resp.Success)
		//t.Log(resp.Result)
	}
	if r.StatusCode != 200 {

	}

}
