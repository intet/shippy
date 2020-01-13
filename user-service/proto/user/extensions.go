package go_micro_srv_user

import (
	"github.com/jinzhu/gorm"
	uid "github.com/satori/go.uuid"
)

func (model *User) BeforeCreate(scope *gorm.Scope) error {
	id, _ := uid.NewV4()
	return scope.SetColumn("Id", id.String())
}