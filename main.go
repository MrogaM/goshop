package main

import (
	"database/sql"
	"fmt"
	"net/http"
	. "shop/models"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

type Context struct {
	response http.ResponseWriter
	request  *http.Request
}

var tpl = template.Must(template.ParseGlob("views/*"))

func main() {

	http.Handle("/statics/", http.StripPrefix("/", http.FileServer(http.Dir("."))))

	http.HandleFunc(`/`, homeDispatch)
	http.ListenAndServe(":8080", nil)
}

func homeDispatch(rw http.ResponseWriter, req *http.Request) {

	//rw.Header().Set("Content-Type", "application/html")

	ctx := Context{response: rw, request: req}

	if validatePath(ctx, "/") != true {
		//rw.Write([]byte("Not found!"))
		http.NotFound(rw, req)
		return
	}

	switch req.Method {

	case "GET":
		renderHome(ctx)
	default:
		rw.Write([]byte("Not allowed!"))
		return
	}
}

func validatePath(ctx Context, required string) bool {

	return ctx.request.URL.Path == required
}

func renderHome(ctx Context) {

	pageHeader := PageHeader{HeaderTitle: "FiraneczkiuOleczki"}
	pageHeader.Menu = Menu{Elements: []string{"Home", "Cart"}}

	tpl.ExecuteTemplate(ctx.response, "index", pageHeader)
}

func getAllProducts(rw http.ResponseWriter, req *http.Request) {

	product := Product{Name: "Programowanie"}

	db, err := sql.Open("mysql", "root:905187G.@/shop")

	if err != nil {
		fmt.Println(err)
	}
	query := `insert into products values(?,?,?,?)`
	db.Exec(query, nil, "name", "category", 90.0)

	fmt.Println(product)
	tpl, _ := template.ParseFiles("views/index.html")
	tpl.Execute(rw, product)
}
