package main

import (
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// refer: https://www.v2ex.com/t/659233

type Blog struct {
	ID        uint
	Title     string
	Content   string
	Tags      pq.StringArray `json:"tags" gorm:"type:text[]"`
	CreatedAt time.Time
}
type Repository struct {
	db *gorm.DB
}

func (p *Repository) ListAll() ([]*Blog, error) {
	var l []*Blog
	err := p.db.Model(&Blog{}).Find(&l).Error
	// err = p.db.Table("blogs").Find(&Blog{}).Error
	return l, err
}

func (p *Repository) Find(id uint) (*Blog, error) {
	blog := &Blog{}
	err := p.db.Where(`id = ?`, id).First(blog).Error
	return blog, err
}

func mockBlog(t *testing.T) (*Repository, sqlmock.Sqlmock) {
	// sqldb
	// 用普通Match: sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	db, mock, _ := sqlmock.New()
	// defer db.Close()

	// gorm db
	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	gormDb, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Errorf("Failed to open gorm v2 db, got error: %v", err)
	}

	repository := &Repository{db: gormDb}
	return repository, mock
	//mock.ExpectationsWereMet()
}

func TestFindBlog(t *testing.T) {
	repository, mock := mockBlog(t)

	blog := &Blog{
		ID:    1,
		Title: "post",
	}

	rows := sqlmock.
		NewRows([]string{"id", "title", "content", "tags", "created_at"}).
		AddRow(blog.ID, blog.Title, blog.Content, blog.Tags, blog.CreatedAt)

	// found
	const sqlSelectOne = `SELECT * FROM "blogs" WHERE id = $1 ORDER BY "blogs"."id" LIMIT 1`
	mock.ExpectQuery(regexp.QuoteMeta(sqlSelectOne)).WithArgs(blog.ID).WillReturnRows(rows)

	dbBlog, err := repository.Find(blog.ID)
	if err != nil {
		t.Fatal(err)
	}
	if dbBlog.ID != blog.ID {
		t.Fatal("not found")
	}

	// not found
	mock.ExpectQuery(`.+`).WillReturnRows(sqlmock.NewRows(nil))
	_, err = repository.Find(1)
	if err == nil {
		t.Fatal("found")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Fatal(err)
	}
}

func TestListBlog(t *testing.T) {
	repository, mock := mockBlog(t)

	blog := &Blog{
		ID:    1,
		Title: "post",
	}

	rows := sqlmock.
		NewRows([]string{"id", "title", "content", "tags", "created_at"}).
		AddRow(blog.ID, blog.Title, blog.Content, blog.Tags, blog.CreatedAt)
	_ = rows

	// found
	// const sqlSelectOne = `SELECT * FROM "blogs" WHERE id = $1 ORDER BY "blogs"."id" LIMIT 1`
	const sqlSelectAll = `SELECT * FROM "blogs"`
	mock.ExpectQuery(regexp.QuoteMeta(sqlSelectAll)).WillReturnRows(rows)
	mock.ExpectQuery(regexp.QuoteMeta(sqlSelectAll)).WillReturnRows(sqlmock.NewRows(nil))

	l, err := repository.ListAll()
	if err != nil {
		t.Fatal(err)
	}
	if len(l) != 1 {
		t.Fatal("empty")
	}
}

func (p *Repository) Add(blog *Blog) error {
	return p.db.Save(blog).Error
}

func TestAddBlog(t *testing.T) {
	repository, mock := mockBlog(t)

	blog := &Blog{
		Title:     "post",
		Content:   "hello",
		Tags:      pq.StringArray{"a", "b"},
		CreatedAt: time.Now(),
	}

	// https://github.com/DATA-DOG/go-sqlmock/issues/118
	const sqlInsert = ` INSERT INTO "blogs" ("title","content","tags","created_at") 
	VALUES ($1,$2,$3,$4) RETURNING "id"`
	const newId = 1
	mock.ExpectBegin() // begin transaction
	mock.ExpectQuery(regexp.QuoteMeta(sqlInsert)).
		WithArgs(blog.Title, blog.Content, blog.Tags, blog.CreatedAt).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(newId))
	mock.ExpectCommit() // commit transaction

	if err := repository.Add(blog); err != nil {
		t.Fatal(err)
	}
	if newId != blog.ID {
		t.Fatal("failed insert")
	}

}
