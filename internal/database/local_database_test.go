package database

import (
	"os"
	"testing"

	dbModels "github.com/Defl8/pasture/internal/database/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// setupTestDB creates a temporary test database
func setupTestDB(t *testing.T) (*LocalDatabase, func()) {
	t.Helper()

	// Create a temporary database file
	dbName := "test_pasture.db"

	// Clean up any existing test database
	os.Remove(dbName)

	db, err := NewLocalDatabase(dbName)
	require.NoError(t, err, "Failed to create test database")
	require.NotNil(t, db, "Database should not be nil")

	// Initialize the database with migrations
	err = db.Initialize()
	require.NoError(t, err, "Failed to initialize database")

	// Return cleanup function
	cleanup := func() {
		db.Close()
		os.Remove(dbName)
	}

	return db, cleanup
}

func TestNewLocalDatabase(t *testing.T) {
	dbName := "test_new.db"
	defer os.Remove(dbName)

	db, err := NewLocalDatabase(dbName)

	assert.NoError(t, err)
	assert.NotNil(t, db)
	assert.Equal(t, dbName, db.Name)
	assert.NotNil(t, db.DB)

	db.Close()
}

func TestNewLocalDatabase_InvalidPath(t *testing.T) {
	// Try to create a database in a non-existent directory
	dbName := "/nonexistent/directory/test.db"

	db, err := NewLocalDatabase(dbName)

	// Should fail gracefully
	assert.Error(t, err)
	assert.Nil(t, db)
}

func TestLocalDatabase_Connect(t *testing.T) {
	dbName := "test_connect.db"
	defer os.Remove(dbName)

	ld := &LocalDatabase{Name: dbName}
	err := ld.Connect()

	assert.NoError(t, err)
	assert.NotNil(t, ld.DB)

	ld.Close()
}

func TestLocalDatabase_Close(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	err := db.Close()
	assert.NoError(t, err)
}

func TestLocalDatabase_Initialize(t *testing.T) {
	dbName := "test_initialize.db"
	defer os.Remove(dbName)

	db, err := NewLocalDatabase(dbName)
	require.NoError(t, err)
	defer db.Close()

	err = db.Initialize()
	assert.NoError(t, err)

	// Verify tables were created by checking if we can interact with them
	var post dbModels.Post
	hasPostTable := db.DB.Migrator().HasTable(&post)
	assert.True(t, hasPostTable)

	var profile dbModels.Profile
	hasProfileTable := db.DB.Migrator().HasTable(&profile)
	assert.True(t, hasProfileTable)
}

func TestLocalDatabase_AutoMigrate(t *testing.T) {
	dbName := "test_automigrate.db"
	defer os.Remove(dbName)

	db, err := NewLocalDatabase(dbName)
	require.NoError(t, err)
	defer db.Close()

	err = db.AutoMigrate(&dbModels.Post{}, &dbModels.Profile{})
	assert.NoError(t, err)

	// Verify migrations
	assert.True(t, db.DB.Migrator().HasTable(&dbModels.Post{}))
	assert.True(t, db.DB.Migrator().HasTable(&dbModels.Profile{}))
}

func TestLocalDatabase_CreatePost(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	post := &dbModels.Post{
		Title:     "Test Post",
		Content:   "This is a test post content",
		Published: true,
	}

	err := db.Create(post)
	assert.NoError(t, err)
	assert.NotZero(t, post.ID, "Post ID should be set after creation")
}

func TestLocalDatabase_CreateProfile(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	profile := &dbModels.Profile{
		Name:      "Test User",
		Bio:       "This is a test bio",
		AvatarURL: "https://example.com/avatar.jpg",
	}

	err := db.Create(profile)
	assert.NoError(t, err)
}

func TestLocalDatabase_CreatePost_EmptyFields(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	// Create post with empty strings (SQLite allows this even with NOT NULL)
	post := &dbModels.Post{
		Title:     "",
		Content:   "",
		Published: true,
	}

	err := db.Create(post)
	// Note: SQLite doesn't enforce NOT NULL on empty strings by default
	assert.NoError(t, err)
	assert.NotZero(t, post.ID)
}

func TestLocalDatabase_GetPostByID(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	// Create a post first
	originalPost := &dbModels.Post{
		Title:     "Get Post Test",
		Content:   "Content for get test",
		Published: false,
	}
	err := db.Create(originalPost)
	require.NoError(t, err)

	// Retrieve the post
	retrievedPost, err := db.GetPostByID(originalPost.ID)

	assert.NoError(t, err)
	assert.NotNil(t, retrievedPost)
	assert.Equal(t, originalPost.ID, retrievedPost.ID)
	assert.Equal(t, originalPost.Title, retrievedPost.Title)
	assert.Equal(t, originalPost.Content, retrievedPost.Content)
	assert.Equal(t, originalPost.Published, retrievedPost.Published)
}

func TestLocalDatabase_GetPostByID_NotFound(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	// Try to get a non-existent post
	post, err := db.GetPostByID(99999)

	assert.Error(t, err)
	assert.Nil(t, post)
}

func TestLocalDatabase_GetProfile(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	// Create a profile first
	originalProfile := &dbModels.Profile{
		Name:      "John Doe",
		Bio:       "Software Developer",
		AvatarURL: "https://example.com/john.jpg",
	}
	err := db.Create(originalProfile)
	require.NoError(t, err)

	// Retrieve the profile
	retrievedProfile, err := db.GetProfile()

	assert.NoError(t, err)
	assert.NotNil(t, retrievedProfile)
	assert.Equal(t, originalProfile.Name, retrievedProfile.Name)
	assert.Equal(t, originalProfile.Bio, retrievedProfile.Bio)
	assert.Equal(t, originalProfile.AvatarURL, retrievedProfile.AvatarURL)
}

func TestLocalDatabase_GetProfile_NotFound(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	// Try to get profile when none exists
	profile, err := db.GetProfile()

	assert.Error(t, err)
	assert.Nil(t, profile)
}

func TestLocalDatabase_UpdatePost(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	// Create a post
	post := &dbModels.Post{
		Title:     "Original Title",
		Content:   "Original Content",
		Published: false,
	}
	err := db.Create(post)
	require.NoError(t, err)

	// Update the post
	post.Title = "Updated Title"
	post.Content = "Updated Content"
	post.Published = true

	err = db.Update(post)
	assert.NoError(t, err)

	// Verify the update
	retrievedPost, err := db.GetPostByID(post.ID)
	require.NoError(t, err)
	assert.Equal(t, "Updated Title", retrievedPost.Title)
	assert.Equal(t, "Updated Content", retrievedPost.Content)
	assert.True(t, retrievedPost.Published)
}

func TestLocalDatabase_UpdateProfile(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	// Create a profile
	profile := &dbModels.Profile{
		Name:      "Original Name",
		Bio:       "Original Bio",
		AvatarURL: "https://example.com/original.jpg",
	}
	err := db.Create(profile)
	require.NoError(t, err)

	// Update the profile
	profile.Name = "Updated Name"
	profile.Bio = "Updated Bio"
	profile.AvatarURL = "https://example.com/updated.jpg"

	err = db.Update(profile)
	assert.NoError(t, err)

	// Verify the update
	retrievedProfile, err := db.GetProfile()
	require.NoError(t, err)
	assert.Equal(t, "Updated Name", retrievedProfile.Name)
	assert.Equal(t, "Updated Bio", retrievedProfile.Bio)
	assert.Equal(t, "https://example.com/updated.jpg", retrievedProfile.AvatarURL)
}

func TestLocalDatabase_DeletePost(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	// Create a post
	post := &dbModels.Post{
		Title:     "Post to Delete",
		Content:   "This will be deleted",
		Published: true,
	}
	err := db.Create(post)
	require.NoError(t, err)
	postID := post.ID

	// Delete the post
	err = db.Delete(post)
	assert.NoError(t, err)

	// Verify deletion
	_, err = db.GetPostByID(postID)
	assert.Error(t, err, "Post should not be found after deletion")
}

func TestLocalDatabase_DeleteProfile(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	// Create a profile
	profile := &dbModels.Profile{
		Name:      "Profile to Delete",
		Bio:       "This will be deleted",
		AvatarURL: "https://example.com/delete.jpg",
	}
	err := db.Create(profile)
	require.NoError(t, err)

	// Delete the profile
	err = db.Delete(profile)
	assert.NoError(t, err)

	// Verify deletion
	_, err = db.GetProfile()
	assert.Error(t, err, "Profile should not be found after deletion")
}

func TestLocalDatabase_MultiplePostsCRUD(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	// Create multiple posts
	posts := []*dbModels.Post{
		{Title: "Post 1", Content: "Content 1", Published: true},
		{Title: "Post 2", Content: "Content 2", Published: false},
		{Title: "Post 3", Content: "Content 3", Published: true},
	}

	for _, post := range posts {
		err := db.Create(post)
		require.NoError(t, err)
	}

	// Retrieve each post
	for _, originalPost := range posts {
		retrievedPost, err := db.GetPostByID(originalPost.ID)
		assert.NoError(t, err)
		assert.Equal(t, originalPost.Title, retrievedPost.Title)
	}

	// Update a post
	posts[1].Title = "Updated Post 2"
	err := db.Update(posts[1])
	assert.NoError(t, err)

	// Verify update
	updatedPost, err := db.GetPostByID(posts[1].ID)
	assert.NoError(t, err)
	assert.Equal(t, "Updated Post 2", updatedPost.Title)

	// Delete a post
	err = db.Delete(posts[0])
	assert.NoError(t, err)

	// Verify deletion
	_, err = db.GetPostByID(posts[0].ID)
	assert.Error(t, err)
}
