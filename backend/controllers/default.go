package controllers

import (
	"b2mak/theScoreAssignemnt/models"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type MainController struct {
	beego.Controller
}

func getDB() *sql.DB {
	mysqlDB, err := sql.Open(
		"mysql",
		"root:mypassword@tcp(mysql:3306)/theScoreAssignment",
	)

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	return mysqlDB
}

func getPlayers(
	name string,
	orderCol string,
	orderDirection string,
	limit string,
	offset string,
) []models.Player {
	mysqlDB := getDB()
	// defer the close till after the main function has finished
	// executing
	defer mysqlDB.Close()

	db := goqu.New("mysql", mysqlDB)
	query := db.From("nfl_players")

	if name != "" {
		query = query.Where(
			goqu.L("LOWER(name)").Like(
				fmt.Sprintf("%%%s%%", strings.ToLower(name)),
			),
		)
	}
	if orderCol != "" {
		if orderDirection == "asc" {
			query = query.Order(goqu.I(orderCol).Asc())
		} else {
			query = query.Order(goqu.I(orderCol).Desc())
		}
	}

	if offset != "" {
		v, _ := strconv.Atoi(offset)
		query = query.Offset(uint(v))
	}

	if limit != "" {
		v, _ := strconv.Atoi(limit)
		query = query.Limit(uint(v))
	}

	var players []models.Player
	if err := query.ScanStructs(&players); err != nil {
		panic(err.Error())
	}

	return players
}

func getTeamsImpl() []map[string]interface{} {
	mysqlDB := getDB()
	// defer the close till after the main function has finished
	// executing
	defer mysqlDB.Close()

	db := goqu.New("mysql", mysqlDB)
	query := db.From("nfl_players")
	query = query.Select("team", goqu.SUM("yds").As("total_yds")).GroupBy("team")

	sql, _, _ := query.ToSQL()
	rows, err := db.Query(sql)
	if err != nil {
		panic((err.Error()))
	}

	var teamYds []map[string]interface{}
	var (
		team     string
		totalYds int
	)
	for rows.Next() {
		err := rows.Scan(&team, &totalYds)
		teamMap := make(map[string]interface{})
		teamMap["Team"] = team
		teamMap["total_yds"] = totalYds
		teamYds = append(teamYds, teamMap)
		if err != nil {
			log.Fatal(err)
		}
	}

	return teamYds
}

func (c *MainController) GetFile() {
	test := make(map[string]string)
	test["ErrorMessage"] = "test error"
	test["ErrorCode"] = "1234"
	c.Data["json"] = test
	c.Abort("400")
}

func (c *MainController) Get() {
	name := c.GetString("name")
	orderCol := c.GetString("orderCol")
	orderDir := c.GetString("orderDirection")
	limit := c.GetString("limit")
	offset := c.GetString("offset")

	players := getPlayers(
		name,
		orderCol,
		orderDir,
		limit,
		offset,
	)

	c.Ctx.Output.Header(
		"Access-Control-Allow-Origin",
		"http://localhost",
	)
	c.Data["json"] = players
	c.ServeJSON()
}

func (c *MainController) GetTeams() {
	teamYds := getTeamsImpl()
	c.Ctx.Output.Header(
		"Access-Control-Allow-Origin",
		"http://localhost",
	)
	c.Data["json"] = teamYds
	c.ServeJSON()
}
