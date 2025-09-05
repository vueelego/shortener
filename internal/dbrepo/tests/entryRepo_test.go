package tests

import (
	"encoding/json"
	"fmt"
	"shortener/internal/models"
	"shortener/internal/types"
	"shortener/pkg/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T) models.Entry {
	user := createRandomUser(t)
	expires := time.Now().Add(time.Hour)
	entry := models.Entry{
		UserID:      user.ID,
		ShortCode:   utils.RandomString(6),
		Title:       utils.RandomString(4),
		OriginalUrl: "https://google.com",
		ExpiresAt:   types.Time{Time: expires},
	}
	err := tRepo.EntryRepo.Create(&entry)
	require.NoError(t, err)
	require.True(t, entry.ID > 0)
	require.WithinDuration(t, entry.CreatedAt.Time, time.Now(), time.Second)
	require.WithinDuration(t, entry.UpdatedAt.Time, time.Now(), time.Second)
	addRandomClick(entry.ID, t)
	return entry
}

func addRandomClick(entryID uint, t *testing.T) models.Click {
	click := models.Click{
		EntryID:    entryID,
		IPAddress:  "127.0.0.1",
		UserAgent:  utils.RandomString(20),
		DeviceType: "IOS",
		CreatedAt:  types.Time{Time: time.Now()},
	}

	err := tRepo.ClickRepo.Create(&click)
	require.NoError(t, err)
	require.True(t, click.ID > 0)

	return click
}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestAddClick(t *testing.T) {
	addRandomClick(4, t)
}

func TestGetEntry(t *testing.T) {
	entry, err := tRepo.EntryRepo.GetById(4)
	require.NoError(t, err)
	jsonString, err := json.Marshal(entry)
	require.NoError(t, err)
	fmt.Println(string(jsonString))
}
