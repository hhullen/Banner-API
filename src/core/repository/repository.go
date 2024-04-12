package repository

import (
	"main/core/model"
)

type IRepository interface {
	GetUserRoleByToken(string) (string, error)
	AddBannerContent(model.BannerContent) (uint, error)
	AddBannerFeature(model.Feature) error
	AddBannerTag(model.Tag) error
}
