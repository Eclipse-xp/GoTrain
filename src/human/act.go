package human

import "food"

//本能
type Ability interface {

	//学习 需要导出，则首字母必须大写
	Learn(learnContent string)
	//吃饭,针对忌口返回异常信息
	Eat(food *food.Food, avoid string) error
}
