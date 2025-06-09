package domain

type Movie struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Title    string `json:"title"`
	Director string `json:"director"`
	Genre    string `json:"genre"`
	Year     int    `json:"year"`
}

func (Movie) TableName() string { return "movies" }
