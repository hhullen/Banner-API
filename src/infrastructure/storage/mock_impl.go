package storage

import (
	"main/core/model"
	"time"
)

type MOCKDB struct {
}

func (me *MOCKDB) GetUserRoleByToken(token string) (string, error) {
	return string(token), nil
}

func (me *MOCKDB) AddBannerContent(model.BannerContent) (int32, error) {
	return 1, nil
}

func (me *MOCKDB) AddBannerFeature(model.Feature) error {
	return nil
}

func (me *MOCKDB) AddBannerTag(model.Tag) error {
	return nil
}

func (me *MOCKDB) IsBannerExists(id int32) bool {
	return id == 21 || id == 42
}

func (me *MOCKDB) UpdateBannerContent(model.BannerContent) error {
	return nil
}

func (me *MOCKDB) UpdateBannerFeature(model.Feature) error {
	return nil
}

func (me *MOCKDB) UpdateBannerTag(model.Tag) error {
	return nil
}

func (me *MOCKDB) GetBannerContent(id int32) (model.BannerContent, error) {
	return model.BannerContent{
		ID:        99,
		Title:     "MOCK TITLE",
		Text:      "MOCK TEXT",
		Url:       "MOCK://URL.COM",
		IsActive:  false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (me *MOCKDB) DeleteBannerFeature(int32) error {
	return nil
}

func (me *MOCKDB) DeleteBannerTags(int32) error {
	return nil
}
