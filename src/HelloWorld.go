package main

import (
	"strconv"
	"fmt"
	"net/http"
	"time"
	"log"
	"food"
	"errors"
	"os"
	"bufio"
	"strings"
	_ "github.com/Go-SQL-Driver/MySQL"
	"database/sql"
	_ "github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm"
)

var xp Person
var persons []Person
var theCouple map[string]Person
var language ="java"

func main() {
	//init_local()
	//
	//introduceAll(persons)
	//introduceCP(theCouple)
	//testUpdateVar()
	//
	//xp.autoIntroduce()
	//
	//xp.travel("北京")
	//fmt.Println(getIntroduce(xp))
	//
	//var man human.Ability = &xp
	//man.Learn("java")
	//
	//xp.autoIntroduce()
	//
	////xp.showBlog()
	//
	//xp.readBook()
	//
	//xiaolongxia := xp.cook()
	//wife := theCouple["wife"]
	//err := wife.Eat(&xiaolongxia, "辣")
	//if err!=nil {
	//	fmt.Println("拒绝吃饭！", err)
	//}
	//
	//if err := recover();err != nil {
	//	fmt.Println(err)
	//}
	//xp.fallInLoveWith(&wife)
	//
	//operateDB()

	//c := make(chan int, 2)
	//c <- 1
	//c <- 2
	//fmt.Println(<-c)
	//fmt.Println(<-c)

	operateGORM()

}
//____________本应放在其他文件的内容，由于编译问题暂放这里________________

//开放博客
func (p *Person)showBlog(){
	mux := http.NewServeMux()
	rh := http.RedirectHandler("http://744722813.iteye.com/", 307)
	mux.Handle("/blog", rh)

	mux.Handle("/hello", timeHandler(time.RFC1123))
	log.Println("Listening...")

	http.ListenAndServe(":3000", mux)
}
func timeHandler(format string) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now().Format(format)
		w.Write([]byte("Hey ****"+ r.FormValue("name") +"**** The time is: " + tm))
	}
	return http.HandlerFunc(fn)
}
type Person struct {
	//姓名
	name string
	//性别
	sex string
	//年龄
	age int
	//所在地
	local string
	//技能
	skill []string
	//体重
	weight int
}

//学习
func (p *Person) Learn(learnContent string){
	p.skill = append(p.skill, learnContent)
}
//吃饭
func (p *Person) Eat(food *food.Food, avoid string) error {
	//处理 味道忌口
	for _,taste := range food.Taste  {
		if strings.Contains(taste, avoid) {
			return errors.New("忌口："+avoid)
		}
	}
	p.weight += 1
	return nil
}
//自我介绍
func (p *Person)autoIntroduce(){
	fmt.Println("Hey guys, my name is "+p.name)
	fmt.Println("I do well in ",p.skill)
}
//旅行
func (p *Person)travel(city string){
	p.local = city
}
//读书
func (p *Person)readBook(){
	path := "E:\\test\\背影.txt"
	log.Print(path)
	file, err := os.Open(path)

	if err==nil {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	} else {
		fmt.Println("file error:",err)
	}

}
//做饭
func (p *Person)cook() food.Food{

	return food.Food{
		Name:"小龙虾",
		Taste:[]string{"麻辣"},
	}
}
//恋爱
func (p *Person)fallInLoveWith(who *Person){
	if p.sex == who.sex {
		panic("对不起，我"+p.name+"是直的！", )
	}
	fmt.Println(p.name, " love ", who.name)
}

//____________________________


func testUpdateVar(){
	//直接赋值
	person2 := xp
	person2.local = "吉林"
	fmt.Println(getIntroduce(xp))
	//通过指针
	language3 := &language
	*language3 = "go"
	person3 := &xp
	person3.local = "吉林"
	fmt.Println(getIntroduce(xp))

}

//初始化变量
func init_local(){
	xp = Person{
		name:"xiexiangpeng",
		sex:"男",
		age:18,
		local:"吉林啊",
	}
	wy := Person{
		name:"wangwenya",
		sex:"女",
		age:14,
		local:"廊坊",
	}

	persons = []Person{xp, wy}
	//theCouple = make(map[string]Person)
	theCouple = map[string]Person{"husband":xp,"wife":wy}

}
func introduceCP(cp map[string]Person){
	for role, w := range cp  {
		fmt.Println(role+ ":" + getIntroduce(w))
	}
}

//介绍全员
func introduceAll(persons []Person){
	for _, man := range persons  {
		fmt.Println("this is " + man.name +","+strconv.Itoa(man.age) +" years old. come from " + man.local)
	}
}

//获得一个人的简介
func getIntroduce(person Person) string{
	return "this is " + person.name +","+strconv.Itoa(person.age) +" years old. come from " + person.local
}

//数据库操作
func operateDB(){
	//db 类型为sql.DB
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/mytest?charset=utf8")
	if err!=nil {
		fmt.Println(err)
	}
	//preStmt, _ := db.Prepare("insert into test_tab(name,city) values (?,?)")
	//preStmt.Exec("tr", "beijing")
	//db.Exec("insert into test_tab(name,city) values (?,?)","tr", "beijing")

	transaction,_ := db.Begin()
	stmt,_ := transaction.Prepare("insert into test_tab(name,city) values (?,?)")
	//id,_ := stmt.Exec("tr", "beijing")
	stmt.Exec("yf", "衡水")
	transaction.Commit()

	fmt.Println("--------批量查询--------")
	rows,_ := db.Query("select * from test_tab")
	for rows.Next()  {
		var id int
		var name,city string
		if err := rows.Scan(&id, &name, &city);err==nil{
			fmt.Println("id:",id," name:",name, " city:",city)
		}
	}

	fmt.Println("--------单个查询--------")
	var id int
	var name,city string
	db.QueryRow("select * from test_tab where id = 4").Scan(&id,&name,&city)
	fmt.Println("id:",id," name:",name, " city:",city)
	db.Close()


}

type knowledge struct {
	Id int
	Language string
	Skill string
	Level string
	Contents string
}

func operateGORM(){
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:3306)/mytest?charset=utf8")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&knowledge{})

	// Create
	db.Create(&knowledge{Language: "go", Skill: "orm", Level: "begin"})

	// Read
	var product knowledge
	db.First(&product, 1) // find product with id 1
	db.First(&product, "Language = ?", "go") // find product with code l1212

	// Update - update product's price to 2000
	db.Model(&product).Update("begin", "init")

	// Delete - delete product
	//db.Delete(&product)
}