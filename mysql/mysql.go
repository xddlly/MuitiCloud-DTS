package mysql

import (
	"database/sql"
	config "MultiCloud-DTS/configparser"
)


type Mysql struct {
	Ismaster bool
	Config string
	db *sql.DB
    masterInfo  * config.Config
}

func InitMysql(conf map[string]interface{}) * Mysql {
    host := MapGet(conf,"host","127.0.0.1").(string)
	port := MapGet(conf,"port","3306").(string)
	username := MapGet(conf,"username","root").(string)
	password := MapGet(conf,"password","123456").(string)
	database := MapGet(conf,"database","test").(string)
	connection := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + database
	con, err := sql.Open("mysql", connection)
	if err != nil {
		panic(err)
	}
	mysql := &Mysql{}
	mysql.Config = connection
	mysql.Ismaster = false
	mysql.db = con
	return mysql
}

