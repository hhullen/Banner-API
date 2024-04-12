package repository

import (
	"main/core/model"
)

type IRepository interface {
	GetUserRoleByToken(string) (string, error)
	AddBannerContent(model.BannerContent) (int32, error)
	AddBannerFeature(model.Feature) error
	AddBannerTag(model.Tag) error
	UpdateBannerContent(model.BannerContent) error
	UpdateBannerFeature(model.Feature) error
	UpdateBannerTag(model.Tag) error
	GetBannerContent(int32) (model.BannerContent, error)
	DeleteBannerContent(int32) error
	DeleteBannerFeature(int32) error
	DeleteBannerTags(int32) error
	IsBannerExists(int32) bool
}
