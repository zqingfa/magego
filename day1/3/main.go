package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"

	//"os"
	//"strings"
	"time"
)

const MaxGuess = 5
const MaxGuessNum = 100

var guess int

func main() {
	guess1()

}
func guess1() {
	rand.Seed(time.Now().Unix())
	target := rand.Intn(MaxGuessNum)
	fmt.Printf("测试结果为：%d,\n", target)
	for i := 1; i <= MaxGuess; i++ {
		fmt.Println("输入你猜测的数字：")
		_, err := fmt.Scanln(&guess)
		if err != nil {
			fmt.Println("输入故障，请正确输入")
			i = 1
			continue
		}
		if guess == target {
			fmt.Println("猜测正确，非常棒！,是否继续再玩一把？【Y/N】")
			var oncemore string
			_, err = fmt.Scanln(&oncemore)
			if err != nil {
				fmt.Println("2输入故障，请正确输入")
				break
			}
			if strings.ToUpper(oncemore) == "Y" {
				fmt.Println("你重新开始了一局")
				guess1() //重新执行该方法
			} else if strings.ToUpper(oncemore) == "N" {
				fmt.Println("你选择了退出，游戏结束")
				//fmt.Println(strings.ToUpper(oncemore)) //debug信息
				os.Exit(0)
			}
		} else if guess > target {
			fmt.Println("猜错了！太大了，再小一点！")
			continue
		} else if guess < target {
			fmt.Println("猜错了！太小了，再大一点！")
			continue
		}
	}
}
