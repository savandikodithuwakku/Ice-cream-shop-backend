package models

import "time"

type User struct {
	ID           uint   `gorm:"primaryKey"`
	Name         string `gorm:"not null"`
	Email        string `gorm:"not null;uniqueIndex"`
	PasswordHash string `gorm:"not null" json:"-"`
	Role         string `gorm:"not null;check:valid_role,role IN ('owner','manager','cashier')"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Category struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null;uniqueIndex"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Item struct {
	ID          uint   `gorm:"primaryKey"`
	CategoryID  uint   `gorm:"not null;index"`
	CreatedBy   uint   `gorm:"not null;index"`
	Name        string `gorm:"not null"`
	Description *string
	Quantity    int64 `gorm:"not null;check:valid_quantity,quantity >= 0"`
	PriceCents  int64 `gorm:"not null;check:valid_price,price_cents >= 0"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Category Category `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"-"`
	Creator  User     `gorm:"foreignKey:CreatedBy;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"-"`
}

type Customer struct {
	ID        uint   `gorm:"primaryKey"`
	CreatedBy uint   `gorm:"not null;index"`
	Name      string `gorm:"not null"`
	Email     *string
	Phone     *string
	Address   *string
	CreatedAt time.Time
	UpdatedAt time.Time

	Creator User `gorm:"foreignKey:CreatedBy;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"-"`
}
