package index

// https://gorm.io/docs/indexes.html
type User struct {
	Name  string `gorm:"index"`
	Name2 string `gorm:"index:idx_name,unique"`
	Name3 string `gorm:"index:,sort:desc,collate:utf8,type:btree,length:10,where:name3 != 'jinzhu'"`
	Name4 string `gorm:"uniqueIndex"`
	Age   int64  `gorm:"index:,class:FULLTEXT,comment:hello \\, world,where:age > 10"`
	Age2  int64  `gorm:"index:,expression:ABS(age)"`
}
