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
	//下划线表示 只执行保重init函数 并不引入全部文件,mysql/driver中有如下初始化操作，
	//func init() {
	//	sql.Register("mysql", &MySQLDriver{})//注册驱动
	//}
	_ "github.com/Go-SQL-Driver/MySQL"
	"database/sql"
	_ "github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm"
)

var xp Person
var persons []Person
var theCouple map[string]Person
var language = "java"

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
	//operateDB()
}

//____________本应放在其他文件的内容，由于编译问题暂放这里________________

//开放博客
func (p *Person) showBlog() {
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
		//如果要指定请求方式只能在这里 用 r.Method="get"...来判断吗？
		w.Write([]byte("Hey ****" + r.FormValue("name") + "**** The time is: " + tm))
	}
	return http.HandlerFunc(fn)

	//下面是接收json并反序列化的方式
	//body, err := ioutil.ReadAll(r.Body)
	//if err != nil {
	//	fmt.Printf("read body err, %v\n", err)
	//	return
	//}
	//println("json:", string(body))
	//
	//var a AutotaskRequest
	//if err = json.Unmarshal(body, &a); err != nil {
	//	fmt.Printf("Unmarshal err, %v\n", err)
	//	return
	//}
	//fmt.Printf("%+v", a)
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
func (p *Person) Learn(learnContent string) {
	p.skill = append(p.skill, learnContent)
}

//吃饭
func (p *Person) Eat(food *food.Food, avoid string) error {
	//处理 味道忌口
	for _, taste := range food.Taste {
		if strings.Contains(taste, avoid) {
			return errors.New("忌口：" + avoid)
		}
	}
	p.weight += 1
	return nil
}

//自我介绍
func (p *Person) autoIntroduce() {
	fmt.Println("Hey guys, my name is " + p.name)
	fmt.Println("I do well in ", p.skill)
}

//旅行
func (p *Person) travel(city string) {
	p.local = city
}

//读书
func (p *Person) readBook() {
	path := "E:\\test\\背影.txt"
	log.Print(path)
	file, err := os.Open(path)

	if err == nil {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	} else {
		fmt.Println("file error:", err)
	}

}

//做饭
func (p *Person) cook() food.Food {

	return food.Food{
		Name:  "小龙虾",
		Taste: []string{"麻辣"},
	}
}

//恋爱
func (p *Person) fallInLoveWith(who *Person) {
	if p.sex == who.sex {
		panic("对不起，我" + p.name + "是直的！", )
	}
	fmt.Println(p.name, " love ", who.name)
}

//____________________________

func testUpdateVar() {
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
func init_local() {
	xp = Person{
		name:  "xiexiangpeng",
		sex:   "男",
		age:   18,
		local: "吉林啊",
	}
	wy := Person{
		name:  "wangwenya",
		sex:   "女",
		age:   14,
		local: "廊坊",
	}

	persons = []Person{xp, wy}
	//theCouple = make(map[string]Person)
	theCouple = map[string]Person{"husband": xp, "wife": wy}

}
func introduceCP(cp map[string]Person) {
	for role, w := range cp {
		fmt.Println(role + ":" + getIntroduce(w))
	}
}

//介绍全员
func introduceAll(persons []Person) {
	for _, man := range persons {
		fmt.Println("this is " + man.name + "," + strconv.Itoa(man.age) + " years old. come from " + man.local)
	}
}

//获得一个人的简介
func getIntroduce(person Person) string {
	return "this is " + person.name + "," + strconv.Itoa(person.age) + " years old. come from " + person.local
}

type man struct {
	id   int
	name string
	city string
}

//数据库操作
func operateDB() {
	//db 类型为sql.DB
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/mytest?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	//preStmt, _ := db.Prepare("insert into test_tab(name,city) values (?,?)")
	//preStmt.Exec("tr", "beijing")
	//db.Exec("insert into test_tab(name,city) values (?,?)","tr", "beijing")

	transaction, _ := db.Begin()
	stmt, _ := transaction.Prepare("insert into test_tab(name,city) values (?,?)")
	//id,_ := stmt.Exec("tr", "beijing")
	stmt.Exec("yf", "衡水")
	transaction.Commit()

	fmt.Println("--------批量查询--------")
	rows, _ := db.Query("select * from test_tab")
	for rows.Next() {
		var id int
		var name, city string
		if err := rows.Scan(&id, &name, &city); err == nil {
			fmt.Println("id:", id, " name:", name, " city:", city)
		}
	}

	fmt.Println("--------单个查询--------")
	var id int
	var name, city string
	db.QueryRow("select * from test_tab where id = 4").Scan(&id, &name, &city)
	//var record man
	//db.QueryRow("select * from test_tab where id = 1").Scan(&record)
	//fmt.Println(record)
	db.Close()

}

type knowledge struct {
	Id       int
	UserId   int
	Language string
	Skill    string
	Level    string
	Contents string
}

type joinResult struct {
	//注意字段名必须大写，否则scan将赋值失败
	Name  string
	City  string
	Skill string
	Level string
}

func operateGORM() {
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:3306)/mytest?charset=utf8")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	//db.Table("tableName")用于指定表名；不指定时默认查询 "结构体名+s"
	// Migrate the schema 建表
	db.Table("test2").AutoMigrate(&knowledge{})

	// Create 插入
	//db.Table("test2").Create(&knowledge{Language: "go", Skill: "orm", Level: "init"})
	//
	//
	//// Read
	var record knowledge
	//一条
	db.Table("test2").First(&record, 1).Scan(&record) // find product with id 1
	//
	var records []knowledge
	//多条
	db.Table("test2").Where("user_id = ?", 1).Find(&records).Scan(&records)

	var records2 []joinResult
	var records3 []joinResult
	//连接
	db.Table("test2").Select("test_tab.name name,test_tab.city city,test2.skill skill,test2.level level").
		Joins("left join test_tab on test2.user_id=test_tab.id").Scan(&records2)
	db.Table("test2").Select("test_tab.name name,test_tab.city city,test2.skill skill,test2.level level").
		Joins("left join test_tab on test2.user_id=test_tab.id").Where("test2.user_id = ?", 1).
		Scan(&records3)
	db.Table("test2").First(&record, "Language = ?", "go") // find

	// Update
	//db.Model(&record).Update("begin", "init")

	// Delete
	//db.Delete(&product)
}
