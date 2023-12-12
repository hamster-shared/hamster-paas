package zan

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func GetClient() *ZanClient {
	return NewZanClient("http://webtcapi.unchartedw3s.com", "478f53734d284889a6a0fbfc648cc061", "2def8d1826884fdd896508d078b26a91", "/Users/mohaijiang/IdeaProjects/blockchain/hamster-paas/rsa_private_key_pkcs8.pem")
}

func TestAuthURL(t *testing.T) {
	client := GetClient()
	response, err := client.AuthUrl()
	if err == nil {
		fmt.Println(response)
	}
}

func TestAccessToken(t *testing.T) {
	client := GetClient()
	response, err := client.AccessToken("13463c11f44d46b59272a56a15e9e577")
	if err == nil {
		fmt.Println(response)
	}
}

func TestCreateApiKey(t *testing.T) {
	client := GetClient()
	response, err := client.ApiKeyCreate("test3", client.baseAccessToken)
	if err == nil {
		fmt.Println(response)
	}

}

func TestApiKeyList(t *testing.T) {
	client := GetClient()
	response, err := client.ApiKeyList(1, 10, client.baseAccessToken)
	if err == nil {
		fmt.Println(response)
	}
}

func TestApiKeyDetail(t *testing.T) {
	client := GetClient()
	response, err := client.ApiKeyDetail("f3d67401-d054-40f8-a332-be17e48a07e8", client.baseAccessToken)
	if err == nil {
		fmt.Println(response)
	}
}

func TestApiKeyCost(t *testing.T) {
	client := GetClient()
	response, err := client.ApiKeyCreditCost("f3d67401-d054-40f8-a332-be17e48a07e8", client.baseAccessToken)
	if err == nil {
		fmt.Println(response)
	}
}

func TestApiKeyRequestStats(t *testing.T) {
	client := GetClient()
	response, err := client.ApiKeyRequestActivityStats("f3d67401-d054-40f8-a332-be17e48a07e8", "STAT_1_MONTH", "ethereum", client.baseAccessToken)
	if err == nil {
		fmt.Println(response)
	}
}

func TestEcosystemsDigest(t *testing.T) {
	client := GetClient()
	response, err := client.EcosystemsDigest()
	if err == nil {
		fmt.Println(response)
	}
}

func TestPlan(t *testing.T) {
	accessToken := "cc9f76447eaf492199f5f1e81311dc97"
	client := GetClient()
	response, err := client.Plan(accessToken)
	assert.NoError(t, err)
	fmt.Println(response.Data)

}

func TestInitPlan(t *testing.T) {
	accessToken := "v29f76447eaf492199f5f1e542ccqevd"
	client := GetClient()

	err := client.InitFreeSpec(accessToken)
	assert.NoError(t, err)
	response, err := client.Plan(accessToken)
	assert.NoError(t, err)
	fmt.Println(response.Data)
}
