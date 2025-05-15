package entity

type Menu struct {
    ID   uint   `gorm:"primaryKey"`
    Name string `gorm:"size:100"`
    Path string `gorm:"size:100"`
}

