package dbrepo

import "shortener/internal/models"

type TagRepo interface {
}

var _ TagRepo = (*tagRepo)(nil)

type tagRepo struct{}

func (*tagRepo) Create(tag *models.Tag) error {
	return Db.Create(tag).Error
}

func (*tagRepo) GetByName(userId uint, name string, tag *models.Tag) error {
	return Db.Where(
		"tag_name = ?", name).Where(
		"owner_id = ?", userId).First(tag).Error
}

func (repo *tagRepo) Update(tag *models.Tag) error {
	var res models.Tag
	if err := Db.First(&res, tag.ID).Error; err != nil {
		return err
	}

	if tag.TagName != "" {
		res.TagName = tag.TagName
	}

	return Db.Save(res).Error
}

func (*tagRepo) Delete(id uint) error {
	return Db.Delete(&models.Tag{}, id).Error
}

func (*tagRepo) ListByOwner(ownerId uint) (list []models.Tag, err error) {
	err = Db.Where("owner_id = ?", ownerId).Order("id DESC").Find(list).Error
	return
}
