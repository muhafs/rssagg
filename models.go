package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/muhamadhaspin/rssagg/internal/database"
)

// * User Model
type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ApiKey    string    `json:"api_key"`
}

func UserFromDB(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		Name:      dbUser.Name,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		ApiKey:    dbUser.ApiKey,
	}
}

// * Feed Model
type Feed struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FeedFromDB(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		Name:      dbFeed.Name,
		URL:       dbFeed.Url,
		UserID:    dbFeed.UserID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
	}
}

func FeedsFromDB(dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}

	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, FeedFromDB(dbFeed))
	}

	return feeds
}

// * Feed Follows Model
type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FeedFollowFromDB(dbFeed database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        dbFeed.ID,
		UserID:    dbFeed.UserID,
		FeedID:    dbFeed.FeedID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
	}
}

func FeedFollowsFromDB(dbFeedFollows []database.FeedFollow) []FeedFollow {
	feedFollows := []FeedFollow{}

	for _, dbFeedFollow := range dbFeedFollows {
		feedFollows = append(feedFollows, FeedFollowFromDB(dbFeedFollow))
	}

	return feedFollows
}

// * Posts Model
type Post struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	URL         string    `json:"url"`
	FeedID      uuid.UUID `json:"feed_id"`
	PublishedAt time.Time `json:"published_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func PostFromDB(dbPost database.Post) Post {
	var desc *string
	if dbPost.Description.Valid {
		desc = &dbPost.Description.String
	}
	return Post{
		ID:          dbPost.ID,
		Title:       dbPost.Title,
		Description: desc,
		URL:         dbPost.Url,
		FeedID:      dbPost.FeedID,
		PublishedAt: dbPost.PublishedAt,
		CreatedAt:   dbPost.CreatedAt,
		UpdatedAt:   dbPost.UpdatedAt,
	}
}

func PostsFromDB(dbPosts []database.Post) []Post {
	posts := []Post{}

	for _, dbPost := range dbPosts {
		posts = append(posts, PostFromDB(dbPost))
	}

	return posts
}
