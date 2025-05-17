package models

type User struct {
	BaseModel
	Phone     string `gorm:"type:string;size:11;not null,unique;"`
	Fullname  string `gorm:"type:string;size:32;null"`
	UserRoles []UserRole
}

type Role struct {
	BaseModel
	Name      string `gorm:"type:string;size:15;not null,unique;"`
	UserRoles []UserRole
}

type UserRole struct {
	BaseModel
	UserId int
	User   User `gorm:"foreignKey:UserId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	RoleId int
	Role   Role `gorm:"foreignKey:RoleId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
}
