//go的风格不建议过多建立文件夹，所以有依赖的结构体尽量在一个包中
package main

//type Person struct {
//	//姓名
//	name string
//	//年龄
//	age int
//	//所在地
//	local string
//	//技能
//	skill []string
//	//体重
//	weight int
//}
//
////学习
//func (p *Person) Learn(learnContent string){
//	p.skill = append(p.skill, learnContent)
//}
////吃饭
//func (p *Person) Eat(food *food.Food, avoid string) error {
//	//处理 味道忌口
//	for _,taste := range food.Taste  {
//		if taste == avoid {
//			return errors.New("忌口："+avoid)
//		}
//	}
//	p.weight += 1
//	return nil
//}
////自我介绍
//func (p *Person)autoIntroduce(){
//	fmt.Println("Hey guys, my name is "+p.name)
//	fmt.Println("I do well in ",p.skill)
//}
////旅行
//func (p *Person)travel(city string){
//	p.local = city
//}