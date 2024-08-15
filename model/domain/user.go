package domain

type User struct {
	ID        int    	`gorm:"column:id;primary_key;autoIncrement"`
	Password  string    `gorm:"column:password"`
	Name      string    `gorm:"column:name"`
	Token     string    `gorm:"column:token"`
	CreatedAt int64     `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt int64     `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
	Note []Note 		`gorm:"foreignKey:user_id;references:id"`
}

func (u *User) TableName() string  {
	return "users"
}