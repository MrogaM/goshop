package main

import (
	"crypto/tls"
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"
	"shop/models"
	page "shop/models/view/pages"
	"text/template"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Context struct {
	response http.ResponseWriter
	request  *http.Request
}

var tpl = template.Must(template.ParseGlob("views/*"))

func main() {

	http.Handle("/statics/", http.StripPrefix("/", http.FileServer(http.Dir("."))))

	http.HandleFunc(`/`, homeDispatch)

	server := http.Server{
		Addr:              ":8080",
		Handler:           nil,
		TLSConfig:         &tls.Config{},
		ReadTimeout:       60,
		ReadHeaderTimeout: 60,
		WriteTimeout:      60,
		IdleTimeout:       60,
		MaxHeaderBytes:    0,
		TLSNextProto:      map[string]func( *http.Server,  *tls.Conn,  http.Handler){},
		ConnState: func( net.Conn,  http.ConnState) {
		},
		ErrorLog: &log.Logger{},
	}
	fmt.Println("Server started at: http://localhost:8080")
	log.Fatal(server.ListenAndServe())
}
func bench() {
	/* db, _ := getDb("shop_go")

	db.Migrator().AutoMigrate(&models.User{}, &models.Roles{}, &models.Permissions{})
	n := time.Now().UnixNano() / (int64(time.Millisecond)/int64(time.Nanosecond))
	
	var a models.User
	var p []models.Permissions

	db.First(&a, "id = ?", 1)
	db.Where("user_id = ?", a.Id).Find(&p)
	a.Permissions = p
	sort.SliceStable(a.Permissions, func(i,j int)bool{return a.Permissions[i].Id < a.Permissions[j].Id})
	r := strings.Split(a.Roles, ",")

	 for _, v := range r {
		fmt.Println(v)
	}
	
	a.Roles = strings.Join([]string{a.Roles, "superuser"}, ",")
	for _, v := range a.Permissions {
		var role models.Roles
		db.Where("id = ?", v.RolesID).First(&role)
		fmt.Println(role.Name)
	}

	fmt.Println(time.Duration(time.Now().UnixNano() / (int64(time.Millisecond)/int64(time.Nanosecond)) - n))
	db.Save(a)
	os.Exit(0) */
}
func homeDispatch(rw http.ResponseWriter, req *http.Request) {

	//rw.Header().Set("Content-Type", "application/html")

	ctx := Context{response: rw, request: req}

	if !ctx.validatePath("/") {
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

func (ctx *Context) validatePath(required string) bool {

	return ctx.request.URL.Path == required
}

func renderHome(ctx Context) {

	pageHeader := page.PageHeader{HeaderTitle: "FiraneczkiuOleczki"}
	pageHeader.Menu = page.Menu{}

	tpl.ExecuteTemplate(ctx.response, "index", pageHeader)
}

func getAllProducts(rw http.ResponseWriter, req *http.Request) {

	product := models.Product{Name: "Programowanie"}

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

func getDBContext(s string) (*gorm.DB, error) {

	dsn := "root:905187G.@tcp(127.0.0.1:3306)/" + s
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}
