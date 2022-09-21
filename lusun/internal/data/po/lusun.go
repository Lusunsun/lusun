package po

import (
	"time"
)

// Article [...]
type Article struct {
	ID           int       `gorm:"primaryKey;column:id"`
	CategoryID   int       `gorm:"column:category_id"`
	ArticleTagID string    `gorm:"column:article_tag_id"`
	Status       int8      `gorm:"column:status"`
	Title        string    `gorm:"column:title"`
	Content      string    `gorm:"column:content"`
	ViewCount    int       `gorm:"column:view_count"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
}

// TableName get sql table name.获取数据库表名
func (m *Article) TableName() string {
	return "article"
}

// ArticleColumns get sql column name.获取数据库列名
var ArticleColumns = struct {
	ID           string
	CategoryID   string
	ArticleTagID string
	Status       string
	Title        string
	Content      string
	ViewCount    string
	CreatedAt    string
	UpdatedAt    string
}{
	ID:           "id",
	CategoryID:   "category_id",
	ArticleTagID: "article_tag_id",
	Status:       "status",
	Title:        "title",
	Content:      "content",
	ViewCount:    "view_count",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

// Category [...]
type Category struct {
	ID        int       `gorm:"primaryKey;column:id"`
	Name      string    `gorm:"column:name"`
	Order     int       `gorm:"column:order"`
	Count     int       `gorm:"column:count"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

// TableName get sql table name.获取数据库表名
func (m *Category) TableName() string {
	return "category"
}

// CategoryColumns get sql column name.获取数据库列名
var CategoryColumns = struct {
	ID        string
	Name      string
	Order     string
	Count     string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	Name:      "name",
	Order:     "order",
	Count:     "count",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// Comment [...]
type Comment struct {
	ID        int       `gorm:"primaryKey;column:id"`
	ArticleID int       `gorm:"column:article_id"`
	Email     string    `gorm:"column:email"`
	Content   string    `gorm:"column:content"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

// TableName get sql table name.获取数据库表名
func (m *Comment) TableName() string {
	return "comment"
}

// CommentColumns get sql column name.获取数据库列名
var CommentColumns = struct {
	ID        string
	ArticleID string
	Email     string
	Content   string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	ArticleID: "article_id",
	Email:     "email",
	Content:   "content",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// Tag [...]
type Tag struct {
	ID        int       `gorm:"primaryKey;column:id"`
	Name      string    `gorm:"column:name"`
	Order     int       `gorm:"column:order"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

// TableName get sql table name.获取数据库表名
func (m *Tag) TableName() string {
	return "tag"
}

// TagColumns get sql column name.获取数据库列名
var TagColumns = struct {
	ID        string
	Name      string
	Order     string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	Name:      "name",
	Order:     "order",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// TagArticleID [...]
type TagArticleID struct {
	ID        int       `gorm:"primaryKey;column:id"`
	ArticleID int       `gorm:"column:article_id"`
	TagID     int       `gorm:"column:tag_id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

// TableName get sql table name.获取数据库表名
func (m *TagArticleID) TableName() string {
	return "tag_article_id"
}

// TagArticleIDColumns get sql column name.获取数据库列名
var TagArticleIDColumns = struct {
	ID        string
	ArticleID string
	TagID     string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	ArticleID: "article_id",
	TagID:     "tag_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}
