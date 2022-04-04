package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	_ "github.com/go-sql-driver/mysql"
)

// CREATE TABLE IF NOT EXISTS nfl_players (
// 	id INT AUTO_INCREMENT PRIMARY KEY,
// 	name VARCHAR(255) NOT NULL,
// 	team VARCHAR(255) NOT NULL,
// 	position VARCHAR(255) NOT NULL,
// 	att_g FLOAT(23) NOT NULL,
// 	att SMALLINT UNSIGNED NOT NULL,
// 	yds SMALLINT NOT NULL,
// 	avg FLOAT(23) NOT NULL,
// 	yds_g FLOAT(23) NOT NULL,
// 	td SMALLINT UNSIGNED NOT NULL,
// 	lng VARCHAR(255) NOT NULL,
// 	first SMALLINT UNSIGNED NOT NULL,
// 	first_percentage FLOAT(23) NOT NULL,
// 	twenty_plus SMALLINT UNSIGNED NOT NULL,
// 	forty_plus SMALLINT UNSIGNED NOT NULL,
// 	fum SMALLINT UNSIGNED NOT NULL,
// 	INDEX name_idx (name),
// )

type player struct {
	name             string
	team             string
	position         string
	att_g            float64
	att              int
	yds              int
	avg              float64
	yds_g            float64
	td               int
	lng              string
	first            int
	first_percentage float64
	twenty_plus      int
	forty_plus       int
	fum              int
}

func interfaceToInt(val interface{}) int {
	switch val.(type) {
	case int:
		return val.(int)
	case float64:
		return int(val.(float64))
	case string:
		reg, err := regexp.Compile("[^0-9]+")
		i, err := strconv.Atoi(reg.ReplaceAllLiteralString(val.(string), ""))
		if err != nil {
			panic(err)
		}
		return i
	default:
		panic("unknown type")
	}
}

func interfaceToString(val interface{}) string {
	switch val.(type) {
	case int:
		return strconv.Itoa(val.(int))
	case float64:
		return fmt.Sprintf("%v", val.(float64))
	case string:
		return val.(string)
	default:
		panic("unknown type")
	}
}

func getFormattedPlayers() []player {
	str, err := os.ReadFile("/Users/bmak/Code/theScoreAssignment/rushing.json")
	if err != nil {
		panic(err)
	}

	var data []map[string]interface{}
	json.Unmarshal([]byte(str), &data)

	var players []player

	for i, obj := range data {
		fmt.Println(i, obj)
		p := player{
			name:             obj["Player"].(string),
			team:             obj["Team"].(string),
			position:         obj["Pos"].(string),
			att_g:            obj["Att/G"].(float64),
			att:              interfaceToInt(obj["Att"]),
			yds:              interfaceToInt(obj["Yds"]),
			avg:              obj["Avg"].(float64),
			yds_g:            obj["Yds/G"].(float64),
			td:               interfaceToInt(obj["TD"]),
			lng:              interfaceToString(obj["Lng"]),
			first:            interfaceToInt(obj["1st"]),
			first_percentage: obj["1st%"].(float64),
			twenty_plus:      interfaceToInt(obj["20+"]),
			forty_plus:       interfaceToInt(obj["40+"]),
			fum:              interfaceToInt(obj["FUM"]),
		}
		players = append(players, p)
	}
	return players
}

func insertData(players []player) {
	fmt.Println("Go MySQL Tutorial")

	// Open up our database connection.
	db, err := sql.Open(
		"mysql",
		"root:mypassword@tcp(127.0.0.1:3306)/theScoreAssignment",
	)

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	var values [][]interface{}
	// twenty_plus, forty_plus, fum)
	for _, p := range players {
		// name := fmt.Sprintf("%s", p.name)
		// team := fmt.Sprintf("%s'", p.team)
		// position := fmt.Sprintf("'%s'", p.position)
		// att_g := fmt.Sprintf("%g", p.att_g)
		// att := fmt.Sprintf("%d", p.att)
		// yds := fmt.Sprintf("%d", p.yds)
		// avg := fmt.Sprintf("%g", p.avg)
		// yrds_g := fmt.Sprintf("%g", p.yds_g)
		// td := fmt.Sprintf("%d", p.td)
		// lng := fmt.Sprintf("'%s'", p.lng)
		// first := fmt.Sprintf("%d", p.first)
		// first_percentage := fmt.Sprintf("%g", p.first_percentage)
		// twenty_plus := fmt.Sprintf("%d", p.twenty_plus)
		// forty_plus := fmt.Sprintf("%d", p.forty_plus)
		// fum := fmt.Sprintf("%d", p.fum)

		val := goqu.Vals{
			p.name,
			p.team,
			p.position,
			p.att_g,
			p.att,
			p.yds,
			p.avg,
			p.yds_g,
			p.td,
			p.lng,
			p.first,
			p.first_percentage,
			p.twenty_plus,
			p.forty_plus,
			p.fum,
		}

		// str := fmt.Sprintf(
		// 	"(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)",
		// )
		values = append(values, val)
	}

	dialect := goqu.Dialect("mysql")
	ds := dialect.Insert("nfl_players").
		Cols("name", "team", "position", "att_g", "att", "yds", "avg", "yds_g", "td", "lng", "first", "first_percentage", "twenty_plus", "forty_plus", "fum").
		Vals(values...)

	// sql := fmt.Sprintf(
	// 	"INSERT INTO players (name, team, position, att_g, att, yds, avg, yrds_g, td, lng, first, first_percentage, twenty_plus, forty_plus, fum) VALUES %s",
	// 	strings.Join(valueStrings, ", "),
	// )

	sql, _, err_on_sql_build := ds.ToSQL()
	if err_on_sql_build != nil {
		panic(err_on_sql_build.Error())
	}

	fmt.Println(sql)

	insert, err_on_insert := db.Query(sql)
	if err_on_insert != nil {
		panic(err_on_insert.Error())
	}

	defer insert.Close()
}

func main() {
	players := getFormattedPlayers()
	insertData(players)
}
