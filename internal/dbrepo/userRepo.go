package dbrepo

import "shortener/internal/models"

type UserRepo interface {
	Create(user *models.User) error
	GetById(id uint, user *models.User) error
	GetByEmail(email string, user *models.User) error
	Update(user *models.User) error
	Delete(id uint) error
	List(f Filter) (users []models.User, filter Filter, err error)
}

var _ UserRepo = (*userRepo)(nil)

type userRepo struct{}

func NewUserRepo() *userRepo {
	return &userRepo{}
}

func (*userRepo) Create(user *models.User) error {
	return Db.Create(user).Error
}

func (*userRepo) GetById(id uint, user *models.User) error {
	return Db.First(user, id).Error
}

func (*userRepo) GetByEmail(email string, user *models.User) error {
	return Db.Where("email = ?", email).First(user).Error
}

func (*userRepo) Update(user *models.User) error {
	var res models.User
	if err := Db.First(&res, user.ID).Error; err != nil {
		return err
	}

	if user.Username != "" {
		res.Username = user.Username
	}

	if user.Avatar != "" {
		res.Avatar = user.Avatar
	}

	return Db.Save(res).Error
}

func (*userRepo) Delete(id uint) error {
	return Db.Delete(&models.User{}, id).Error
}

func (*userRepo) List(f Filter) (users []models.User, filter Filter, err error) {
	curDb := Db.Model(models.User{})
	curDb = curDb.Count(&f.TotalRows)
	err = curDb.Limit(f.Limit()).Offset(f.Offset()).Order("id DESC").Find(&users).Error
	f.LastPageNum = f.GetLastPageNum()
	filter = f
	return
}
