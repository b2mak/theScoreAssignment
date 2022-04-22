package models

type Player struct {
	Name             string  `db:"name"`
	Team             string  `db:"team"`
	Position         string  `db:"position"`
	Att_g            float64 `db:"att_g"`
	Att              int     `db:"att"`
	Yds              int     `db:"yds"`
	Avg              float64 `db:"avg"`
	Yds_g            float64 `db:"yds_g"`
	Td               int     `db:"td"`
	Lng              string  `db:"lng"`
	First            int     `db:"first"`
	First_percentage float64 `db:"first_percentage"`
	Twenty_plus      int     `db:"twenty_plus"`
	Forty_plus       int     `db:"forty_plus"`
	Fum              int     `db:"fum"`
}

type PlayerColumns string

const (
	Name             PlayerColumns = "name"
	Team             PlayerColumns = "team"
	Postion          PlayerColumns = "position"
	Att_g            PlayerColumns = "att_g"
	Att              PlayerColumns = "att"
	Yds              PlayerColumns = "yds"
	Avg              PlayerColumns = "avg"
	Yds_g            PlayerColumns = "yds_g"
	Td               PlayerColumns = "td"
	Lng              PlayerColumns = "lng"
	First            PlayerColumns = "first"
	First_percentage PlayerColumns = "first_percentage"
	Twenty_plus      PlayerColumns = "twenty_plus"
	Forty_plus       PlayerColumns = "forty_plus"
	Fum              PlayerColumns = "fum"
)
