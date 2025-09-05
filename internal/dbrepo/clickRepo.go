package dbrepo

import "shortener/internal/models"

type ClickRepo interface {
	Create(click *models.Click) error
	ListByEntry(entryId uint, f Filter) (list []models.User, filter Filter, err error)
}

var _ ClickRepo = (*clickRepo)(nil)

type clickRepo struct{}

func NewClickRepo() *clickRepo {
	return &clickRepo{}
}

func (*clickRepo) Create(click *models.Click) error {
	return Db.Create(click).Error
}

func (*clickRepo) ListByEntry(entryId uint, f Filter) (list []models.User, filter Filter, err error) {
	db := Db.Where("entry_id = ?", entryId)
	if err = db.Count(&f.TotalRows).Error; err != nil {
		return
	}
	f.LastPageNum = f.GetLastPageNum()
	filter = f
	err = Db.Limit(f.Limit()).Offset(f.Offset()).Order("id DESC").Find(list).Error
	return
}
