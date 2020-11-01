package database

import "github.com/jinzhu/gorm"

type SqlHandler interface {
	Exec(string, ...interface{}) *gorm.DB
	Find(interface{}, ...interface{}) *gorm.DB
	First(interface{}, ...interface{}) *gorm.DB
	Raw(string, ...interface{}) *gorm.DB
	Create(interface{}) *gorm.DB
	Save(interface{}) *gorm.DB
	Model(interface{}) *gorm.DB
	Update(...interface{}) *gorm.DB
	Delete(interface{}) *gorm.DB
	Where(interface{}, ...interface{}) *gorm.DB
	Set(string, interface{}) *gorm.DB
	Query() *gorm.DB
}
