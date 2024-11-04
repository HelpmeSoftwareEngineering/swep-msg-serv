package repository_test

import (
	"testing"
	"time"

	"github.com/Ateto1204/swep-msg-serv/entity"
	"github.com/Ateto1204/swep-msg-serv/internal/repository"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var testDB *gorm.DB

func setupTestDB() {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to test database")
	}

	db.AutoMigrate(&entity.Message{})
	testDB = db
}

func TestSave(t *testing.T) {
	setupTestDB()
	repo := repository.NewMsgRepository(testDB)

	msgID := "user123"
	content := "this is a demo msg content"
	sender := "Test User Sender"
	now := time.Now()

	msg, err := repo.Save(msgID, sender, content, now)
	assert.NoError(t, err)
	assert.Equal(t, msgID, msg.ID)
	assert.Equal(t, content, msg.Content)
	assert.Equal(t, sender, msg.Sender)
	assert.Equal(t, now, msg.CreateAt)
	assert.Equal(t, false, msg.Read)
}

func TestGetByID(t *testing.T) {
	setupTestDB()
	repo := repository.NewMsgRepository(testDB)

	msgID := "user123"
	content := "this is a demo msg content"
	sender := "Test User Sender"
	now := time.Now()
	repo.Save(msgID, sender, content, now)

	msg, err := repo.GetByID(msgID)
	assert.NoError(t, err)
	assert.Equal(t, msgID, msg.ID)
	assert.Equal(t, content, msg.Content)
	assert.Equal(t, sender, msg.Sender)
	assert.Equal(t, false, msg.Read)

	assert.True(t, msg.CreateAt.Equal(now), "CreateAt should match")
}

func TestUpdByID(t *testing.T) {
	setupTestDB()
	repo := repository.NewMsgRepository(testDB)

	msgID := "user123"
	content := "this is a demo msg content"
	sender := "Test User Sender"
	now := time.Now()
	repo.Save(msgID, sender, content, now)

	err := repo.UpdByID(msgID)
	assert.NoError(t, err)
}
