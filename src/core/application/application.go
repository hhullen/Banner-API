package application

import (
	"main/core/model"
	"main/core/repository"
	"time"

	gen_api "main/go-server-generated/go"
)

type Application struct {
	storage repository.IRepository
}

func NewApplication(storage repository.IRepository) *Application {
	return &Application{storage: storage}
}

func (me *Application) CheckRoleByToken(token string) string {
	role, err := me.storage.GetUserRoleByToken(token)
	if err != nil {
		return ""
	}
	return role
}

func (me *Application) AddBanner(banner gen_api.BannerBody) (uint, error) {
	content := model.BannerContent{
		Title:     banner.Content.Title,
		Text:      banner.Content.Text,
		Url:       banner.Content.Url,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	id, err := me.storage.AddBannerContent(content)
	if err != nil {
		return 0, err
	}

	feature := model.Feature{
		ID:        uint(banner.FeatureId),
		ContentID: id,
	}
	me.storage.AddBannerFeature(feature)

	for _, tagID := range banner.TagIds {
		err = me.storage.AddBannerTag(model.Tag{
			ID:        uint(tagID),
			ContentID: id,
		})
		if err != nil {
			return 0, err
		}
	}

	return id, nil
}
