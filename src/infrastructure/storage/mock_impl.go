package storage

import (
	"main/core/model"
	"time"
)

var bannerContent = model.BannerContent{
	ID:        1,
	Title:     "MOCK TITLE",
	Text:      "MOCK TEXT",
	Url:       "MOCK://URL.COM",
	IsActive:  false,
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

var bannerFeature = model.Feature{
	ID:        1,
	ContentID: 1,
}

var bannerTag1 = model.Tag{
	ID:        1,
	ContentID: 1,
}

var bannerTag2 = model.Tag{
	ID:        2,
	ContentID: 1,
}

type MOCKDB struct {
}

func (me *MOCKDB) GetUserRoleByToken(token string) (string, error) {
	return string(token), nil
}

func (me *MOCKDB) AddBannerContent(model.BannerContent) (int32, error) {
	return bannerContent.ID, nil
}

func (me *MOCKDB) AddBannerFeature(model.Feature) error {
	return nil
}

func (me *MOCKDB) AddBannerTag(model.Tag) error {
	return nil
}

func (me *MOCKDB) IsBannerExists(id int32) bool {
	return id == bannerContent.ID
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
	return bannerContent, nil
}

func (me *MOCKDB) GetBannerFeatere(id int32) (model.Feature, error) {
	return bannerFeature, nil
}

func (me *MOCKDB) GetBannerTags(id int32) ([]model.Tag, error) {
	return []model.Tag{bannerTag1, bannerTag2}, nil
}

func (me *MOCKDB) DeleteBannerContent(int32) error {
	return nil
}

func (me *MOCKDB) DeleteBannerFeature(int32) error {
	return nil
}

func (me *MOCKDB) DeleteBannerTags(int32) error {
	return nil
}

func (me *MOCKDB) GetAllBannersByFilters(feature, tag, limit, offset int32, deactivated bool) ([]model.BannerContent, error) {
	if deactivated {
		return []model.BannerContent{bannerContent}, nil
	}
	return []model.BannerContent{}, nil
}

func (me *MOCKDB) GetSpecificBanner(feature, tag int32, deactivated bool) (*model.BannerContent, error) {
	if deactivated {
		return &bannerContent, nil
	}
	return nil, nil
}
