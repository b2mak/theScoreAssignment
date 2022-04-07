package main

import (
	"b2mak/theScoreAssignemnt/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/doug-martin/goqu/v9"
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

func getFormattedPlayers() []models.Player {
	str, err := os.ReadFile("/Users/bmak/Code/theScoreAssignment/db/rushing.json")
	if err != nil {
		panic(err)
	}

	var data []map[string]interface{}
	json.Unmarshal([]byte(str), &data)

	var players []models.Player

	for i, obj := range data {
		fmt.Println(i, obj)
		p := models.Player{
			Name:             obj["Player"].(string),
			Team:             obj["Team"].(string),
			Position:         obj["Pos"].(string),
			Att_g:            obj["Att/G"].(float64),
			Att:              interfaceToInt(obj["Att"]),
			Yds:              interfaceToInt(obj["Yds"]),
			Avg:              obj["Avg"].(float64),
			Yds_g:            obj["Yds/G"].(float64),
			Td:               interfaceToInt(obj["TD"]),
			Lng:              interfaceToString(obj["Lng"]),
			First:            interfaceToInt(obj["1st"]),
			First_percentage: obj["1st%"].(float64),
			Twenty_plus:      interfaceToInt(obj["20+"]),
			Forty_plus:       interfaceToInt(obj["40+"]),
			Fum:              interfaceToInt(obj["FUM"]),
		}
		players = append(players, p)
	}
	return players
}

func insertData(players []models.Player) {
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
	for _, p := range players {
		val := goqu.Vals{
			p.Name,
			p.Team,
			p.Position,
			p.Att_g,
			p.Att,
			p.Yds,
			p.Avg,
			p.Yds_g,
			p.Td,
			p.Lng,
			p.First,
			p.First_percentage,
			p.Twenty_plus,
			p.Forty_plus,
			p.Fum,
		}

		values = append(values, val)
	}

	dialect := goqu.Dialect("mysql")
	ds := dialect.Insert("nfl_players").
		Cols("name", "team", "position", "att_g", "att", "yds", "avg", "yds_g", "td", "lng", "first", "first_percentage", "twenty_plus", "forty_plus", "fum").
		Vals(values...)

	sql, _, err_on_sql_build := ds.ToSQL()
	if err_on_sql_build != nil {
		panic(err_on_sql_build.Error())
	}

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
