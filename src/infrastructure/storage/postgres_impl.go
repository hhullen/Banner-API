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

func (me *PostgreSQL) beginTransactionSafely() *gorm.DB {
	tx := me.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	return tx
}

func (me *PostgreSQL) GetUserRoleByToken(token string) (string, error) {
	tx := me.beginTransactionSafely()
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
	tx := me.beginTransactionSafely()

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
	tx := me.beginTransactionSafely()

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
	tx := me.beginTransactionSafely()

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
	tx := me.beginTransactionSafely()

	var banner model.BannerContent
	err := tx.Find(&banner, "ID = ?", id).Error
	if err != nil {
		return false
	}

	if err := tx.Commit().Error; err != nil {
		return false
	}
	return id == banner.ID
}

func (me *PostgreSQL) UpdateBannerContent(banner model.BannerContent) error {
	tx := me.beginTransactionSafely()

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
	tx := me.beginTransactionSafely()

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
	tx := me.beginTransactionSafely()

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
	return bannerContent, nil
}

func (me *PostgreSQL) GetBannerFeatere(id int32) (model.Feature, error) {
	return bannerFeature, nil
}

func (me *PostgreSQL) GetBannerTags(id int32) ([]model.Tag, error) {
	return []model.Tag{bannerTag1, bannerTag2}, nil
}

func (me *PostgreSQL) DeleteBannerContent(int32) error {
	return nil
}

func (me *PostgreSQL) DeleteBannerFeature(int32) error {
	return nil
}

func (me *PostgreSQL) DeleteBannerTags(int32) error {
	return nil
}

func (me *PostgreSQL) GetAllBannersByFilters(feature, tag, limit, offset int32, deactivated bool) ([]model.BannerContent, error) {
	if deactivated {
		return []model.BannerContent{bannerContent}, nil
	}
	return []model.BannerContent{}, nil
}

func (me *PostgreSQL) GetSpecificBanner(feature, tag int32, deactivated bool) (*model.BannerContent, error) {
	if deactivated {
		return &bannerContent, nil
	}
	return nil, nil
}
