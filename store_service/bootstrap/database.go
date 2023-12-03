package bootstrap

import (
	"fmt"
	"store_service/internal/models"

	log "github.com/sirupsen/logrus"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Gorm setup
func dbConnection(dsn string) (sqlClient models.SqlClient, ok bool) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Error(err.Error())
		return sqlClient, false

	} else {
		sqlClient.DB = db
		return sqlClient, true
	}
}

// Ping to DB
func PingToDB(sqlClient models.SqlClient) bool {
	if _, err := sqlClient.DB.DB(); err != nil {
		log.Error(err.Error())
		return false
	} else {
		log.Info("DB PONG!")
		return true
	}
}

// Connect to Mariadb to create database if not exists
func connectMariadb(username string, password string, host string, port string, dbName string) models.SqlClient {

	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/"
	sqlClient, ok := dbConnection(dsn)

	if !ok {
		log.Error("database connection is distruped!")
	}

	// Create the database if it does not exist
	sqlClient.DB.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)

	return sqlClient

}

// Get database instance
func ConnectDatabaseMariadb(username string, password string, host string, port string, dbName string) models.SqlClient {

	// Connect to mariadb if database not exists create with migration
	connectMariadb(username, password, host, port, dbName)

	// Connect to database
	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8"
	sqlClient, ok := dbConnection(dsn)
	if !ok {
		log.Error(fmt.Sprintf("connection to database %s is intruped", dbName))
	}
	return sqlClient
}
