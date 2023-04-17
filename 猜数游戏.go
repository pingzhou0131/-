package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	maxNum := 100
	//rand.Seed(time.Now().UnixNano())
	rand.New(rand.NewSource(time.Now().UnixNano()))
	secretNumber := rand.Intn(maxNum)
	//fmt.Println("The secret number is ", secretNumber)
	fmt.Println("Please input your guess")
	var guess, cnt int
	for {
		cnt++
		// 输入我们猜的数字
		_, err := fmt.Scanf("%d", &guess)
		fmt.Scanln()
		// Go语言中处理错误的方法
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer value")
			continue
		}
		fmt.Println("You guess is", guess)
		if guess == secretNumber {
			fmt.Println("Correct, you Legend!")
			break
		} else if guess < secretNumber {
			fmt.Println("Your guess is smaller than the secret number. Please try again count=", cnt)
		} else if guess > secretNumber {
			fmt.Println("Your guess is bigger than the secret number. Please try again count=", cnt)
		}
	}
}
