package routers

import (
	"encoding/json"
	"fmt"
	"strconv"

	//"strings"

	//"github.com/aws/aws-lambda-go/events"
	"github.com/gambit/bd"
	"github.com/gambit/models"
)

func InsertProduct(body string, User string) (int, string) {
	fmt.Println("Inicializando funcion  router.InsertProduct")

	var t models.Product

	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "Error en los datos recibidos" + err.Error()
	}

	if len(t.ProdTitle) == 0 {
		return 400, "Debe especificar el Nombre (title) del producto "
	}

	isAdmin, msg := bd.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	result, err2 := bd.InsertProduct(t)
	if err2 != nil {
		return 400, "Ocurrio un error al intentar registrar el producto" + t.ProdTitle + " > " + err2.Error()
	}

	return 200, "{ProdID: " + strconv.Itoa(int(result)) + "}"

}

func UpdateProduct(body string, User string, id int) (int, string) {
	fmt.Println("Inicializando funcion  router.UpdateProduct")

	var t models.Product

	fmt.Println(t.ProdTitle)

	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "Error en los datos recibidos" + err.Error()
	}

	if len(t.ProdTitle) == 0 {
		fmt.Println(t.ProdTitle)
		return 400, "Debe especificar el nombre del producto DALE "
	}

	isAdmin, msg := bd.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	t.ProdId = id
	err2 := bd.UpdateProduct(t)
	if err2 != nil {
		return 400, "Ocurrio un error al intentar actualizar el registro de producto" + strconv.Itoa(id) + " > " + err2.Error()
	}

	return 200, " Update OK"

}
