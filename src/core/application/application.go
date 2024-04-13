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

func (me *Application) DeleteBanner(id int32) error {
	return me.storage.DeleteBannerContent(id)
}

func (me *Application) IsBannerExists(id int32) bool {
	return me.storage.IsBannerExists(id)
}

func (me *Application) GetAllBannersByFilters(feature, tag, limit, offset int32, deactivated bool) ([]gen_api.InlineResponse200, error) {
	contents, err := me.storage.GetAllBannersByFilters(feature, tag, limit, offset, deactivated)
	if err != nil {
		return nil, err
	}
	var returnable []gen_api.InlineResponse200
	for _, content := range contents {
		feature, err := me.storage.GetBannerFeatere(content.ID)
		if err != nil {
			return nil, err
		}
		tags, err := me.storage.GetBannerTags(content.ID)
		if err != nil {
			return nil, err
		}
		var tagsIds []int32
		for i := range tags {
			tagsIds = append(tagsIds, tags[i].ID)
		}
		returnable = append(returnable, gen_api.InlineResponse200{
			BannerId:  content.ID,
			TagIds:    tagsIds,
			FeatureId: feature.ID,
			Content: gen_api.ModelMap{
				Title: content.Title,
				Text:  content.Text,
				Url:   content.Url,
			},
			IsActive:  content.IsActive,
			CreatedAt: content.CreatedAt,
			UpdatedAt: content.UpdatedAt,
		})
	}
	return returnable, nil
}
