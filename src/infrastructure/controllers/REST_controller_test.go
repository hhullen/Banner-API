package controllers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	gen_api "main/go-server-generated/go"
	"net/http"
	"testing"
)

var (
	client = &http.Client{}
)

const (
	host       = "http://127.0.0.1:8080"
	adminToken = "AVADAKEDABRA"
	userToken  = "REDUKTO"
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

func TestUnauthorizedRequest(t *testing.T) {
	req, _ := http.NewRequest("POST", host+"/banner", nil)

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusUnauthorized {
		t.Fatal("Status code:", resp.Status)
	}
}

func TestPostForbiddenUser(t *testing.T) {
	data, _ := json.Marshal(testBanner1)
	req, _ := http.NewRequest("POST", host+"/banner", bytes.NewBuffer(data))

	req.Header.Set("user_token", userToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusForbidden {
		t.Fatal("Status code:", resp.Status, "Expected: ", http.StatusForbidden)
	}
}

var existingID = "1"

func TestPostBanners(t *testing.T) {
	data1, _ := json.Marshal(testBanner1)
	req1, _ := http.NewRequest("POST", host+"/banner", bytes.NewBuffer(data1))
	data2, _ := json.Marshal(testBanner2)
	req2, _ := http.NewRequest("POST", host+"/banner", bytes.NewBuffer(data2))
	data3, _ := json.Marshal(testBanner3)
	req3, _ := http.NewRequest("POST", host+"/banner", bytes.NewBuffer(data3))
	data4, _ := json.Marshal(testBanner4)
	req4, _ := http.NewRequest("POST", host+"/banner", bytes.NewBuffer(data4))

	req1.Header.Set("user_token", adminToken)
	req1.Header.Set("Content-Type", "application/json")
	req2.Header.Set("user_token", adminToken)
	req2.Header.Set("Content-Type", "application/json")
	req3.Header.Set("user_token", adminToken)
	req3.Header.Set("Content-Type", "application/json")
	req4.Header.Set("user_token", adminToken)
	req4.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req1)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatal("Status code:", resp.Status, "Expected: ", http.StatusOK)
	}

	var resp_struct gen_api.InlineResponse201
	json.NewDecoder(resp.Body).Decode(&resp_struct)
	existingID = fmt.Sprint(resp_struct.BannerId)

	resp, err = client.Do(req2)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatal("Status code:", resp.Status, "Expected: ", http.StatusOK)
	}
	resp, err = client.Do(req3)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatal("Status code:", resp.Status, "Expected: ", http.StatusOK)
	}
	resp, err = client.Do(req4)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatal("Status code:", resp.Status, "Expected: ", http.StatusOK)
	}
}

func TestPatchBanner(t *testing.T) {
	patchedBanner := gen_api.BannerIdBody{
		TagIds:    []int32{1, 2, 3},
		FeatureId: 1,
		Content: gen_api.ModelMap{
			Title: "PATCHED",
			Text:  "PATCHED",
			Url:   "http://PATCHED.com",
		},
		IsActive: true,
	}
	data, _ := json.Marshal(patchedBanner)
	var resp *http.Response

	req, err := http.NewRequest("PATCH", host+"/banner/"+existingID, bytes.NewBuffer(data))
	req.Header.Set("user_token", adminToken)

	req.Header.Set("Content-Type", "application/json")
	resp, err = client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatal("Status code:", resp.Status, "Expected: ", http.StatusOK)
	}
}

func TestGetBanners(t *testing.T) {
	req, _ := http.NewRequest("GET", host+"/banner?feature_id=2&tag_id=4&limit=5&offset=0", nil)
	req.Header.Set("user_token", adminToken)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatal("Status code:", resp.Status, "Expected: ", http.StatusOK)
	}
}

func TestGetInnactiveBanners(t *testing.T) {
	req, _ := http.NewRequest("GET", host+"/banner?feature_id=4&tag_id=9&limit=5&offset=0", nil)
	req.Header.Set("user_token", userToken)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Fatal("Status code:", resp.Status, "Expected: ", http.StatusNotFound)
	}
}

func TestGetUserBanner(t *testing.T) {
	req1, _ := http.NewRequest("GET", host+"/user_banner?feature_id=3&tag_id=3&use_last_revision=true", nil)
	req2, _ := http.NewRequest("GET", host+"/user_banner?feature_id=3&tag_id=3&use_last_revision=false", nil)
	req1.Header.Set("user_token", userToken)
	req1.Header.Set("Content-Type", "application/json")
	req2.Header.Set("user_token", userToken)
	req2.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req1)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatal("Status code:", resp.Status, "Expected: ", http.StatusOK)
	}

	resp, err = client.Do(req2)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatal("Status code:", resp.Status, "Expected: ", http.StatusOK)
	}
}

func TestDeleteBanner(t *testing.T) {
	req, _ := http.NewRequest("DELETE", host+"/banner/"+existingID, nil)
	req.Header.Set("user_token", adminToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatal("Status code:", resp.Status, "Expected: ", http.StatusOK)
	}
}
