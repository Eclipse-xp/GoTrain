//go的风格不建议过多建立文件夹，所以有依赖的结构体尽量在一个包中
package human

import (
	"bufio"
	"errors"
	"fmt"
	"food"
	"log"
	"os"
	"strings"
)

type Person struct {
	//姓名
	Name string
	//性别
	Sex string
	//年龄
	Age int
	//所在地
	Local string
	//技能
	Skill []string
	//体重
	Weight int
}

//学习
func (p *Person) Learn(learnContent string) {
	p.Skill = append(p.Skill, learnContent)
}

//吃饭
func (p *Person) Eat(food *food.Food, avoid string) error {
	//处理 味道忌口
	for _, taste := range food.Taste {
		if strings.Contains(taste, avoid) {
			return errors.New("忌口：" + avoid)
		}
	}
	p.Weight += 1
	return nil
}

//自我介绍
func (p *Person) AutoIntroduce() {
	fmt.Println("Hey guys, my name is " + p.Name)
	fmt.Println("I do well in ", p.Skill)
}

//旅行
func (p *Person) Travel(city string) {
	p.Local = city
}

//读书
func (p *Person) ReadBook() {
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
func (p *Person) Cook() food.Food {

	return food.Food{
		Name:  "小龙虾",
		Taste: []string{"麻辣"},
	}
}

//恋爱
func (p *Person) FallInLoveWith(who *Person) {
	if p.Sex == who.Sex {
		panic("对不起，我" + p.Name + "是直的！")
	}
	fmt.Println(p.Name, " love ", who.Name)
}
