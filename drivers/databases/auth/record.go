package auth

type Auth struct {
	ID          uint   `gorm:"primaryKey; auto_increment"`
	UserID      uint   `gorm:"not null;"`
	AuthUUID    string `gorm:"size:255; not null;"`
	RefreshUUID string `gorm:"size:255; not null;"`
}
