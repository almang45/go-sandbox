package models

type MovieDetails struct {
	Id          int     `db:"id" json:"id"`
	ImdbId      string  `db:"imdb_id" json:"imdbId"`
	Title       string  `db:"title" json:"title"`
	Year        string  `db:"year" json:"year"`
	Quality     string  `db:"quality" json:"quality"`
	Type        string  `db:"type" json:"type"`
	Size        string  `db:"size" json:"size"`
	Rating      float32 `db:"rating" json:"rating"`
	Runtime     int     `db:"runtime" json:"runtime"`
	Genre       string  `db:"genre" json:"genre"`
	Actors      string  `db:"actors" json:"actors"`
	Director    string  `db:"director" json:"director"`
	Poster      string  `db:"poster" json:"poster"`
	NewRelease  bool    `db:"is_new_release" json:"isNewRelease"`
	NewDownload bool    `db:"is_new_download" json:"isNewDownload"`
}
