package entity

type Book struct {
	ID          uint64 `gorm:"primary_key:auto_increment"`
	Title       string `gorm:"type:varchar(255)" json:"title"`
	Description string `gorm:"type:text" json:"description"`
	UserID      uint64 `gorm:"not null" json:"-"`
	User        User   `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE, onDelete:CASCADE" json:"user"`
}
