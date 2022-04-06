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
