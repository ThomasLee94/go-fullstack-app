package model

type Trainer struct {
    gorm.Model
    Slug string `gorm:"unique_index;not null"`
    Name string `gorm:"not null"`
    Pokemon []Pokemon 
    
}
