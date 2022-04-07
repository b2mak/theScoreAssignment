package controllers

import (
	"b2mak/theScoreAssignemnt/models"
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"

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
		"root:mypassword@tcp(host.docker.internal:3306)/theScoreAssignment",
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

	var players []models.Player
	if err := query.ScanStructs(&players); err != nil {
		panic(err.Error())
	}

	return players
}

func buildFile(
	name string,
	orderCol string,
	orderDirection string,
) string {
	players := getPlayers(
		name,
		orderCol,
		orderDirection,
	)

	filePath := fmt.Sprintf(
		"/tmp/%d.csv",
		time.Now().Unix(),
	)

	f, err := os.Create(filePath)
	defer f.Close()

	if err != nil {
		panic(err.Error())
	}

	w := csv.NewWriter(f)
	defer w.Flush()

	st := reflect.TypeOf(models.Player{})
	numFields := st.NumField()

	var fieldNames []string
	for i := 0; i < numFields; i++ {
		fieldNames = append(fieldNames, st.Field(i).Name)
	}

	if err = w.Write(fieldNames); err != nil {
		panic(err.Error())
	}

	for _, p := range players {
		var values []string
		for i := 0; i < numFields; i++ {
			r := reflect.ValueOf(p)
			str := fmt.Sprintf("%v", reflect.Indirect(r).Field(i))
			values = append(values, str)
		}
		if err = w.Write(values); err != nil {
			panic(err.Error())
		}
	}

	return filePath
}

func (c *MainController) GetFile() {
	filePath := buildFile(
		c.GetString("name"),
		c.GetString("orderCol"),
		c.GetString("orderDirection"),
	)

	c.Ctx.Output.Download(filePath)
}

func (c *MainController) Get() {
	players := getPlayers(
		c.GetString("name"),
		c.GetString("orderCol"),
		c.GetString("orderDirection"),
	)

	c.Ctx.Output.Header(
		"Access-Control-Allow-Origin",
		"http://localhost",
	)
	c.Data["json"] = players
	c.ServeJSON()
}
