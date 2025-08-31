package tests

import (
	"shortener/internal/dbrepo"
	"shortener/internal/models"
	"shortener/pkg/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func createRandomUser(t *testing.T) models.User {
	hashPwd, err := utils.GenHashedPassword(utils.RandomString(8))
	require.NoError(t, err)
	require.NotEqual(t, "", hashPwd)

	username := utils.RandomString(6)
	email := utils.RandomEmail()

	user := models.User{
		Username:     username,
		Email:        email,
		PasswordHash: hashPwd,
		Avatar:       "http://" + utils.RandomString(10) + ".com/files/avatar.png",
	}

	err = tRepo.UserRepo.Create(&user)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, username, user.Username)
	require.Equal(t, email, user.Email)
	require.True(t, user.ID > 0)
	require.True(t, user.Role == false)
	return user
}

func TestCreateUser(t *testing.T) {
	_ = createRandomUser(t)
}

func TestGetUserById(t *testing.T) {
	user := createRandomUser(t)
	var findUser models.User
	err := tRepo.UserRepo.GetById(user.ID, &findUser)
	require.NoError(t, err)
	require.Equal(t, user.ID, findUser.ID)
	require.Equal(t, user.Username, findUser.Username)
	require.Equal(t, user.Email, findUser.Email)
	require.Equal(t, user.Role, findUser.Role)
	require.Equal(t, user.Avatar, findUser.Avatar)
	require.Equal(t, user.DeletedAt, findUser.DeletedAt)
	require.Equal(t, user.PasswordHash, findUser.PasswordHash)

	require.WithinDuration(t, user.CreatedAt.Time, findUser.CreatedAt.Time, time.Second)
	require.WithinDuration(t, user.UpdatedAt.Time, findUser.UpdatedAt.Time, time.Second)
}

func TestGetUserByEmail(t *testing.T) {
	user := createRandomUser(t)
	var findUser models.User
	err := tRepo.UserRepo.GetByEmail(user.Email, &findUser)
	require.NoError(t, err)
	require.Equal(t, user.ID, findUser.ID)
	require.Equal(t, user.Username, findUser.Username)
	require.Equal(t, user.Email, findUser.Email)
	require.Equal(t, user.Role, findUser.Role)
	require.Equal(t, user.Avatar, findUser.Avatar)
	require.Equal(t, user.DeletedAt, findUser.DeletedAt)
	require.Equal(t, user.PasswordHash, findUser.PasswordHash)

	require.WithinDuration(t, user.CreatedAt.Time, findUser.CreatedAt.Time, time.Second)
	require.WithinDuration(t, user.UpdatedAt.Time, findUser.UpdatedAt.Time, time.Second)
}

func TestUpdateUser(t *testing.T) {
	user := createRandomUser(t)

	username := utils.RandomString(10)
	avatar := "http://" + utils.RandomString(12)

	user.Username = username
	user.Avatar = avatar

	err := tRepo.UserRepo.Update(&user)
	require.NoError(t, err)

	var user2 models.User
	tRepo.UserRepo.GetById(user.ID, &user2)
	require.Equal(t, username, user2.Username)
	require.Equal(t, avatar, user2.Avatar)
	require.Equal(t, user.ID, user2.ID)
}

func TestGetUserList(t *testing.T) {
	var rows = 10
	for i := 0; i < rows; i++ {
		_ = createRandomUser(t)
	}

	f := dbrepo.Filter{
		PageNum:  1,
		PageSize: 10,
	}
	users, filter, err := tRepo.UserRepo.List(f)
	require.NoError(t, err)
	require.True(t, len(users) == 10)
	require.True(t, f.PageNum == filter.PageNum)
	require.True(t, f.PageSize == filter.PageSize)
	require.True(t, filter.TotalRows >= 10)
}

func TestDeleteUser(t *testing.T) {
	user := createRandomUser(t)
	err := tRepo.UserRepo.Delete(user.ID)
	require.NoError(t, err)

	var user2 models.User
	err = tRepo.UserRepo.GetById(user.ID, &user2)
	require.ErrorIs(t, err, gorm.ErrRecordNotFound)
}
