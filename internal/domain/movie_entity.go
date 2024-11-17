package domain

type Movie struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Director string `json:"director"`
	Genre    string `json:"genre"`
	Year     int    `json:"year"`
}
