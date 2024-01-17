package index

// https://gorm.io/docs/indexes.html
type User struct {
	Name  string `gorm:"index;not null;default:inited"`
	Name4 string `gorm:"uniqueIndex:idx_board_cell;not null"`
	Name5 string `gorm:"uniqueIndex:idx_board_cell;not null"`

	Name2 string `gorm:"index:idx_name,unique"`
	Name3 string `gorm:"index:,sort:desc,collate:utf8,type:btree,length:10,where:name3 != 'jinzhu'"`
	Age   int64  `gorm:"index:,class:FULLTEXT,comment:hello \\, world,where:age > 10"`
	Age2  int64  `gorm:"index:,expression:ABS(age)"`
}

// Composite Indexes, for example:
// create composite index `idx_member` with columns `name`, `number`
type UserCompositeIndexes struct {
	Name   string `gorm:"index:idx_member"`
	Number string `gorm:"index:idx_member"`
}
