package storage

import (
	"log"
	"main/core/model"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PostgreSQL struct {
	db *gorm.DB
}

func NewPostgreSQL(dsn string) *PostgreSQL {
	gormLogger := logger.New(
		log.New(os.Stdout, "\r\n[POSTGRES]", log.LstdFlags),
		logger.Config{
			IgnoreRecordNotFoundError: false,
		},
	)

	new_db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: gormLogger})
	if err != nil {
		log.Fatalln(err.Error())
		return nil
	}

	psql := &PostgreSQL{db: new_db}
	err = psql.Migrate()
	if err != nil {
		log.Fatalln(err.Error())
		return nil
	}
	return psql
}

func (me *PostgreSQL) Migrate() error {
	err := me.db.AutoMigrate(&model.BannerContent{})
	if err != nil {
		return err
	}

	err = me.db.AutoMigrate(&model.Feature{})
	if err != nil {
		return err
	}

	err = me.db.AutoMigrate(&model.Tag{})
	if err != nil {
		return err
	}

	err = me.db.AutoMigrate(&model.User{})
	if err != nil {
		return err
	}
	return nil
}

func (me *PostgreSQL) AddDefaultUser(user model.User) error {
	tx := me.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var user_old model.User
	err := tx.Where("role = ?", user.Role).Where("token = ?", user.Token).Find(&user_old).Error
	if err != nil {
		return err
	}
	if user_old.Role != user.Role || user_old.Token != user.Token {
		err = tx.Create(&user).Error
		if err != nil {
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}

func (me *PostgreSQL) GetUserRoleByToken(token string) (string, error) {
	tx := me.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var user model.User
	err := tx.First(&user, "token = ?", token).Error
	if err != nil {
		return "", err
	}

	if err := tx.Commit().Error; err != nil {
		return "", err
	}
	return user.Role, nil
}

func (me *PostgreSQL) AddBannerContent(banner model.BannerContent) (int32, error) {
	tx := me.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err := tx.Create(&banner).Error
	if err != nil {
		return 0, err
	}

	if err := tx.Commit().Error; err != nil {
		return 0, err
	}
	return banner.ID, nil
}

func (me *PostgreSQL) AddBannerFeature(feature model.Feature) error {
	tx := me.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err := tx.Create(&feature).Error
	if err != nil {
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}

func (me *PostgreSQL) AddBannerTag(tag model.Tag) error {
	tx := me.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err := tx.Create(&tag).Error
	if err != nil {
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}

func (me *PostgreSQL) IsBannerExists(id int32) bool {
	tx := me.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var banner model.BannerContent
	err := tx.First(&banner, "id = ?", id).Error
	if err != nil {
		return false
	}

	if err := tx.Commit().Error; err != nil {
		return false
	}
	return id == banner.ID
}

func (me *PostgreSQL) UpdateBannerContent(banner model.BannerContent) error {
	tx := me.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var banner_old model.BannerContent
	err := tx.First(&banner_old).Error
	if err != nil {
		return err
	}

	banner_old.UpdatedAt = banner.UpdatedAt
	banner_old.IsActive = banner.IsActive
	banner_old.Title = banner.Title
	banner_old.Text = banner.Text
	banner_old.Url = banner.Url

	err = tx.Save(&banner_old).Error
	if err != nil {
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}

func (me *PostgreSQL) UpdateBannerFeature(feature model.Feature) error {
	tx := me.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var feature_old model.Feature
	err := tx.First(&feature_old).Error
	if err != nil {
		return err
	}

	feature_old.ContentID = feature.ContentID

	err = tx.Save(&feature_old).Error
	if err != nil {
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}

func (me *PostgreSQL) UpdateBannerTag(tag model.Tag) error {
	tx := me.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var tag_old model.Feature
	err := tx.First(&tag_old).Error
	if err != nil {
		return err
	}

	tag_old.ContentID = tag.ContentID

	err = tx.Save(&tag_old).Error
	if err != nil {
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}

func (me *PostgreSQL) GetBannerContent(id int32) (model.BannerContent, error) {
	tx := me.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var banner model.BannerContent
	err := tx.First(&banner, "id = ?", id).Error
	if err != nil {
		return model.BannerContent{}, err
	}

	if err := tx.Commit().Error; err != nil {
		return model.BannerContent{}, err
	}
	return banner, nil
}

func (me *PostgreSQL) GetBannerFeatere(id int32) (model.Feature, error) {
	tx := me.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var feature model.Feature
	err := tx.First(&feature, "content_id = ?", id).Error
	if err != nil {
		return model.Feature{}, err
	}

	if err := tx.Commit().Error; err != nil {
		return model.Feature{}, err
	}
	return feature, nil
}

func (me *PostgreSQL) GetBannerTags(id int32) ([]model.Tag, error) {
	tx := me.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var tags []model.Tag
	err := tx.Find(&tags, "content_id = ?", id).Error
	if err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (me *PostgreSQL) DeleteBannerContent(id int32) error {
	tx := me.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err := tx.Delete(&model.BannerContent{}, id).Error
	if err != nil {
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}

func (me *PostgreSQL) DeleteBannerFeature(id int32) error {
	tx := me.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err := tx.Where("content_id = ?", id).Delete(&model.Feature{}).Error
	if err != nil {
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}

func (me *PostgreSQL) DeleteBannerTags(id int32) error {
	tx := me.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err := tx.Where("content_id = ?", id).Delete(&model.Tag{}).Error
	if err != nil {
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}

func (me *PostgreSQL) GetAllBannersByFilters(feature, tag, limit, offset int32, includeDeactivated bool) ([]model.BannerContent, error) {
	tx := me.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if offset > 0 {
		tx = tx.Offset(int(offset))
	}
	if limit > 0 {
		tx = tx.Limit(int(limit))
	}
	if !includeDeactivated {
		tx = tx.Where("is_active = true")
	}

	var contents []model.BannerContent
	err := tx.Joins("JOIN tags ON tags.content_id = banner_contents.id").
		Joins("JOIN features ON features.content_id = banner_contents.id").
		Where("tags.tag_id = ?", tag).Where("features.feature_id = ?", feature).
		Find(&contents).Error
	if err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}
	return contents, nil
}

func (me *PostgreSQL) GetSpecificBanner(feature, tag int32, includeDeactivated bool) (*model.BannerContent, error) {
	tx := me.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if !includeDeactivated {
		tx = tx.Where("is_active = true")
	}

	var banner []model.BannerContent
	err := tx.Joins("JOIN tags ON tags.content_id = banner_contents.id").
		Joins("JOIN features ON features.content_id = banner_contents.id").
		Where("tags.tag_id = ?", tag).Where("features.feature_id = ?", feature).
		Find(&banner).Error
	if err != nil {
		return nil, err
	}
	if len(banner) == 0 {
		return nil, nil
	}

	return &banner[0], nil
}
