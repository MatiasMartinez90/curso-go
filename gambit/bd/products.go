package bd

import (
	"database/sql"
	//"errors"
	"fmt"
	"strconv"

	//"strings"

	"github.com/gambit/models"
	"github.com/gambit/tools"
	_ "github.com/go-sql-driver/mysql"
)

func InsertProduct(p models.Product) (int64, error) {
	fmt.Println("Inicializando funcion  db.InsertProduct")

	err := DbConnect()
	if err != nil {
		return 0, err
	}

	defer Db.Close()

	sentencia := "INSERT INTO products (Prod_Title"

	if len(p.ProdDescription) > 0 {
		sentencia += ", Prod_Description"
	}

	if p.ProdPrice > 0 {
		sentencia += ", Prod_Price"
	}

	if p.ProdCategId > 0 {
		sentencia += ", Prod_CategoryId"
	}

	if p.ProdStock > 0 {
		sentencia += ", Prod_Stock"
	}

	if len(p.ProdPath) > 0 {
		sentencia += ", Prod_Path"
	}

	sentencia += ") Values ('" + tools.EscapeString(p.ProdTitle) + "'"

	if len(p.ProdDescription) > 0 {
		sentencia += ", '" + tools.EscapeString(p.ProdDescription) + "'"
	}

	if p.ProdPrice > 0 {
		sentencia += ", " + strconv.FormatFloat(p.ProdPrice, 'e', -1, 64)
	}

	if p.ProdCategId > 0 {
		sentencia += ", " + strconv.Itoa(p.ProdCategId)
	}

	if p.ProdStock > 0 {
		sentencia += ", " + strconv.Itoa(p.ProdStock)
	}

	if len(p.ProdPath) > 0 {
		sentencia += ", '" + tools.EscapeString(p.ProdPath) + "'"
	}

	sentencia += ")"

	var result sql.Result

	fmt.Println("Vamos a ejecutar la sentancia: ")
	fmt.Println(sentencia)
	result, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	LastInsertId, err2 := result.LastInsertId()
	if err2 != nil {
		fmt.Println(err.Error())
		return 0, err2
	}

	fmt.Println("Insert Product > Ejecucion Exitosa")
	return LastInsertId, nil
}
