package storage

import "main/core/model"

type MOCKDB struct {
}

func (me *MOCKDB) GetUserRoleByToken(token string) (string, error) {
	return string(token), nil
}

func (me *MOCKDB) AddBannerContent(model.BannerContent) (uint, error) {
	return 1, nil
}

func (me *MOCKDB) AddBannerFeature(model.Feature) error {
	return nil
}

func (me *MOCKDB) AddBannerTag(model.Tag) error {
	return nil
}
