package main

//需要把项目目录加到GOPATH中，才能引用到项目中的自定义包。GOPATH可以包括多个目录
import (
	"strconv"
	//点 . 点操作的含义就是这个包导入之后在调用这个包的函数时，可以省略前缀的包名，
	// 也就是前面你调用的Println(“hello world”)  可以省略的写成Println(“hello world”)
	. "fmt"
	"human"
	"log"
	"net/http"
	"time"

	//下划线表示 只执行包中init函数 并不引入全部文件,mysql/driver中有如下初始化操作，
	//func init() {
	//	sql.Register("mysql", &MySQLDriver{})//注册驱动
	//}
	"database/sql"
	//"util"
	//_ "util"

	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
	//restful 框架
	//"github.com/julienschmidt/httprouter"
)

var xp human.Person
var persons []human.Person
var theCouple map[string]human.Person
var language = "java"

func main() {
	// util.RedisPut("redisGoStr", "hello rediogo 3")
	//init_local()
	//
	//introduceAll(persons)
	//introduceCP(theCouple)
	//testUpdateVar()
	//
	//xp.autoIntroduce()
	//
	//xp.travel("北京")
	//Println(getIntroduce(xp))
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
	//	Println("拒绝吃饭！", err)
	//}
	//
	//if err := recover();err != nil {
	//	Println(err)
	//}
	//xp.fallInLoveWith(&wife)
	//
	//operateDB()

	c := make(chan int, 2)
	c <- 1
	c <- 2
	Println(<-c)
	Println(<-c)
	//close(c) 关闭channel
	// select 语句使得一个 goroutine 在多个通讯操作上等待。
	//select 会阻塞，直到条件分支中的某个可以继续执行，这时就会执行那个条件分支。当多个都准备好的时候，会随机选择一个
	//当 select 中的其他条件分支都没有准备好的时候，`default` 分支会被执行。
	//为了非阻塞的发送或者接收，可使用 default 分支
	//for {
	//	select {
	//	case v := <-c:
	//		fmt.Println("c,%v", v)
	//		return
	//	case v := <-quit:
	//		fmt.Println("quit,%v", v)
	//		return
	//	case <- time.After(5 * time.Second): //设置超时
	//	default:、
	//		fmt.Println("    .")
	//		time.Sleep(50 * time.Millisecond)
	//	}
	//}

		//operateGORM()
	//operateDB()
}

//____________本应放在其他文件的内容，由于编译问题暂放这里________________
type hperson human.Person

//开放博客
func (p *hperson) showBlog() {
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
	//	Printf("read body err, %v\n", err)
	//	return
	//}
	//println("json:", string(body))
	//
	//var a AutotaskRequest
	//if err = json.Unmarshal(body, &a); err != nil {
	//	Printf("Unmarshal err, %v\n", err)
	//	return
	//}
	//Printf("%+v", a)
}

//____________________________

func testUpdateVar() {
	//直接赋值
	person2 := xp
	person2.Local = "吉林"
	Println(getIntroduce(xp))
	//通过指针
	language3 := &language
	*language3 = "go"
	person3 := &xp
	person3.Local = "吉林"
	Println(getIntroduce(xp))

}

//初始化变量
func init_local() {
	xp = human.Person{
		Name:  "xiexiangpeng",
		Sex:   "男",
		Age:   18,
		Local: "吉林啊",
	}
	wy := human.Person{
		Name:  "wangwenya",
		Sex:   "女",
		Age:   14,
		Local: "廊坊",
	}

	persons = []human.Person{xp, wy}
	//theCouple = make(map[string]Person)
	theCouple = map[string]human.Person{"husband": xp, "wife": wy}

}
func introduceCP(cp map[string]human.Person) {
	for role, w := range cp {
		Println(role + ":" + getIntroduce(w))
	}
}

//介绍全员
func introduceAll(persons []human.Person) {
	for _, man := range persons {
		Println("this is " + man.Name + "," + strconv.Itoa(man.Age) + " years old. come from " + man.Local)
	}
}

//获得一个人的简介
func getIntroduce(person human.Person) string {
	return "this is " + person.Name + "," + strconv.Itoa(person.Age) + " years old. come from " + person.Local
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
		Println(err)
	}
	//preStmt, _ := db.Prepare("insert into test_tab(name,city) values (?,?)")
	//preStmt.Exec("tr", "beijing")
	//db.Exec("insert into test_tab(name,city) values (?,?)","tr", "beijing")

	transaction, _ := db.Begin()
	stmt, _ := transaction.Prepare("insert into test_tab(name,city) values (?,?)")
	//id,_ := stmt.Exec("tr", "beijing")
	stmt.Exec("yf", "衡水")
	transaction.Commit()

	Println("--------批量查询--------")
	rows, _ := db.Query("select * from test_tab")
	for rows.Next() {
		var id int
		var name, city string
		if err := rows.Scan(&id, &name, &city); err == nil {
			Println("id:", id, " name:", name, " city:", city)
		}
	}

	Println("--------单个查询--------")
	var id int
	var name, city string
	db.QueryRow("select * from test_tab where id = 4").Scan(&id, &name, &city)
	//var record man
	//db.QueryRow("select * from test_tab where id = 1").Scan(&record)
	//Println(record)
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
