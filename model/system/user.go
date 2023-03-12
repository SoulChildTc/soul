package system

import (
	"soul/model/common"
)

type User struct {
	common.ID
	Name     string `json:"name" gorm:"size:32;not null;comment:用户名"`
	Mobile   string `json:"mobile" gorm:"size:24;not null;uniqueIndex;comment:用户手机号"`
	Password string `json:"-" gorm:"not null;comment:用户密码"`
	common.Timestamps
}
