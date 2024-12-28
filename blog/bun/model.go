package blob_bun

import (
	"database/sql"
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users"`

	ID        int64     `bun:"id,pk,autoincrement"`
	Name      string    `bun:"name,notnull,type:varchar(255)"`
	Email     string    `bun:"email,notnull,unique,type:varchar(255)"`
	CreatedAt time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:"updated_at,notnull,default:current_timestamp"`
}

type PostStatus string

const (
	PostStatusDraft     PostStatus = "draft"
	PostStatusPublished PostStatus = "published"
	PostStatusArchived  PostStatus = "archived"
)

type Post struct {
	bun.BaseModel `bun:"table:users"`

	ID          int64        `bun:"id,pk,autoincrement"`
	AuthorID    int64        `bun:"author_id,notnull"`
	Author      *User        `bun:"rel:belongs-to,join:author_id=id"`
	Title       string       `bun:"title,notnull,type:varchar(255)"`
	Content     string       `bun:"content,notnull,type:text"`
	Status      PostStatus   `bun:"status,notnull,type:varchar(255)"`
	PublishedAt sql.NullTime `bun:"published_at"`
	CreatedAt   time.Time    `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt   time.Time    `bun:"updated_at,notnull,default:current_timestamp"`
}
