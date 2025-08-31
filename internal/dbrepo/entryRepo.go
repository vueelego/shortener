package dbrepo

import (
	"shortener/internal/models"

	"gorm.io/gorm"
)

type EntryRepo interface {
	Create(entry *models.Entry) error
	GetById(id uint) (entry models.Entry, err error)
	Update(entry *models.Entry) error
	Delete(id uint) error
	List(f Filter) (list []models.Entry, filter Filter, err error)
	ListByUser(f Filter, userId uint) (list []models.Entry, filter Filter, err error)
}

var _ EntryRepo = (*entryRepo)(nil)

type entryRepo struct{}

func NewEntryRepo() *entryRepo {
	return &entryRepo{}
}

func (*entryRepo) Create(entry *models.Entry) error {
	return Db.Create(entry).Error
}

func (*entryRepo) GetById(id uint) (entry models.Entry, err error) {
	err = Db.Preload("clicks", "entry_id = ?", id).Limit(10).First(&entry, id).Error
	return entry, err
}

func (*entryRepo) Update(entry *models.Entry) error {
	var res models.Entry
	if err := Db.First(&res, entry.ID).Error; err != nil {
		return err
	}

	if entry.Title != "" {
		res.Title = entry.Title
	}

	return Db.Save(res).Error
}

func (*entryRepo) Delete(id uint) error {
	return Db.Delete(&models.Entry{}, id).Error
}

func (repo *entryRepo) list(f Filter, userIds ...uint) (list []models.Entry, filter Filter, err error) {
	var userId uint = 0
	if len(userIds) > 0 {
		userId = userIds[0]
	}

	curDb := Db.Table("entries")

	if userId > 0 {
		curDb = curDb.Where("user_id = ?", userId)
	}

	err = curDb.Limit(f.Limit()).Offset(f.Offset()).
		Preload("clicks", func(db *gorm.DB) *gorm.DB {
			return db.Order("id DESC").Limit(10)
		}).Find(&list).Error

	f.LastPageNum = f.GetLastPageNum()
	filter = f
	return list, filter, err
}

func (repo *entryRepo) List(f Filter) (list []models.Entry, filter Filter, err error) {
	return repo.list(f)
}

func (repo *entryRepo) ListByUser(f Filter, userId uint) (list []models.Entry, filter Filter, err error) {
	return repo.list(f, userId)
}
