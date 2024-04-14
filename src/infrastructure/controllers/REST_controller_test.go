package controllers_test

import (
	"bytes"
	"encoding/json"
	gen_api "main/go-server-generated/go"
	"net/http"
	"testing"
)

var testBanner1 = gen_api.BannerIdBody{
	TagIds:    []int32{1, 2, 3},
	FeatureId: 1,
	Content: gen_api.ModelMap{
		Title: "TEST 1",
		Text:  "TEST 1",
		Url:   "http://TEST_1.com",
	},
	IsActive: true,
}

var testBanner2 = gen_api.BannerIdBody{
	TagIds:    []int32{4, 5, 6},
	FeatureId: 2,
	Content: gen_api.ModelMap{
		Title: "TEST 2",
		Text:  "TEST 2",
		Url:   "http://TEST_2.com",
	},
	IsActive: true,
}

var testBanner3 = gen_api.BannerIdBody{
	TagIds:    []int32{3},
	FeatureId: 3,
	Content: gen_api.ModelMap{
		Title: "TEST 3",
		Text:  "TEST 3",
		Url:   "http://TEST_3.com",
	},
	IsActive: true,
}

var testBanner4 = gen_api.BannerIdBody{
	TagIds:    []int32{9},
	FeatureId: 4,
	Content: gen_api.ModelMap{
		Title: "TEST 4",
		Text:  "TEST 4",
		Url:   "http://TEST_4.com",
	},
	IsActive: false,
}

func TestAddActor(t *testing.T) {
	data, _ := json.Marshal(testBanner1)
	req, _ := http.NewRequest("POST", "localhost:8080/banner", bytes.NewBuffer(data))

	req.Header.Set("user_token", "TESTING_TOKEN_")
}
