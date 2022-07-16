package bd

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gitlab.com/apernet/infra-devops/dba/rds-killmysql/models"
)

func KillMySQL(claves models.SecretRDSJson, parametros models.DatosEntrada) {
	var dbUser, authToken, dbEndpoint, dbName string

	dbUser = claves.Username
	dbEndpoint = fmt.Sprintf("%s:%d", claves.Host, claves.Port)
	dbName = ""
	authToken = claves.Password
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true", dbUser, authToken, dbEndpoint, dbName)

	/* Abro la base con las credenciales de root */
	db, err := sql.Open("mysql", dsn)
	defer db.Close()

	if err != nil {
		panic("[ERROR] sql.Open = " + err.Error())
	}

	fmt.Println("Probando la conexion a " + dbEndpoint + " con credenciales de " + dbUser)
	err = db.Ping()
	if err != nil {
		panic("[ERROR] al hacer el ping " + err.Error())
	}
	fmt.Println("Conexión exitosa a la base de datos")

	/* Traigo el Token del Supplier para conectarnos a la BD con sus credenciales */
	sentencia := "select id from information_schema.processlist where ((time > " + parametros.Kill_level1 + " and command = 'Sleep') OR (time > " + parametros.Kill_level2 + ")) and user = '" + dbUser + "';"
	fmt.Println(sentencia)
	rows, err2 := db.Query(sentencia)
	if err2 != nil {
		panic("[ERROR] en la consulta de datos " + err2.Error())
	}

	for rows.Next() {
		var id string
		rows.Scan(&id)
		killStr := "KILL " + id

		_, err = db.Query(killStr)
		if err != nil {
			fmt.Println("no se pudo matar el proceso " + id + " > " + err.Error())
		} else {
			fmt.Println("Proceso " + id + " borrado satisfactoriamente")
		}
	}
	rows.Close()

	fmt.Println("Proceso Finalizado")

}
