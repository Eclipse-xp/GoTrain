package main

import (
	"strconv"
	"fmt"
	"human"
	"net/http"
	"time"
	"log"
	"food"
	"errors"
	"os"
	"bufio"
)

var xp Person
var persons []Person
var theCouple map[string]Person
var language ="java"

func main() {
	init_local()

	introduceAll(persons)
	introduceCP(theCouple)
	testUpdateVar()

	xp.autoIntroduce()

	xp.travel("北京")
	fmt.Println(getIntroduce(xp))

	var man human.Ability = &xp
	man.Learn("java")

	xp.autoIntroduce()

	//xp.showBlog()

	xp.readBook()

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
		if taste == avoid {
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
		age:18,
		local:"吉林啊",
	}
	wy := Person{
		name:"wangwenya",
		age:14,
		local:"廊坊",
	}

	persons = []Person{xp, wy}

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
