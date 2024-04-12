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

func (me *Application) AddBanner(banner gen_api.BannerBody) (int32, error) {
	content := model.BannerContent{
		Title:     banner.Content.Title,
		Text:      banner.Content.Text,
		Url:       banner.Content.Url,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		IsActive:  banner.IsActive,
	}
	id, err := me.storage.AddBannerContent(content)
	if err != nil {
		return 0, err
	}

	feature := model.Feature{
		ID:        banner.FeatureId,
		ContentID: id,
	}
	me.storage.AddBannerFeature(feature)

	for _, tagID := range banner.TagIds {
		err = me.storage.AddBannerTag(model.Tag{
			ID:        tagID,
			ContentID: id,
		})
		if err != nil {
			return 0, err
		}
	}

	return id, nil
}

func (me *Application) IsBannerExists(id int32) bool {
	return me.storage.IsBannerExists(id)
}

func (me *Application) UpdatedBanner(id int32, banner gen_api.BannerBody) error {
	content, err := me.storage.GetBannerContent(id)
	if err != nil {
		return err
	}
	content.UpdatedAt = time.Now()
	content.Title = banner.Content.Title
	content.Text = banner.Content.Text
	content.Url = banner.Content.Url
	content.IsActive = banner.IsActive
	err = me.storage.UpdateBannerContent(content)
	if err != nil {
		return err
	}

	err = me.storage.DeleteBannerFeature(id)
	if err != nil {
		return err
	}

	err = me.storage.DeleteBannerTags(id)
	if err != nil {
		return err
	}

	feature := model.Feature{
		ID:        banner.FeatureId,
		ContentID: id,
	}
	me.storage.AddBannerFeature(feature)

	for _, tagID := range banner.TagIds {
		err = me.storage.AddBannerTag(model.Tag{
			ID:        tagID,
			ContentID: id,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
