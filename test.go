package main

import "fmt"

// First Test

func Person(age int, name string) string {
	// we can use this
	message1 := fmt.Sprintf("Hello my name is %v and im %v years old", name, age)
	fmt.Println(message1)
	// or this
	var message2 string
	message2 = fmt.Sprintf("if you dont know me; im %v and i have %v years old", name, age)
	return message2
}

// Prime Number Task

func GetNumber() int {
	var number int
	fmt.Println("put your number: ")
	fmt.Scanln(&number)
	return number
}

func CheckNumber() int {
	var count int
	number := GetNumber()
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			count++
		}
		if count > 2 {
			break
		}
	}
	return count
}

func PrimeNumber() string {
	if CheckNumber() == 2 {
		return "Prime Number!!!"
	} else {
		return "Not Prime Number :(((("
	}
}

// Months of the year Task

func GetDay() int {
	var day int
	fmt.Println("put your day: ")
	fmt.Scanln(&day)
	if day < 366 && day > 0 {
		return day
	} else {
		return 0
	}
}

func CheckMonth() []int {
	number := GetDay()
	var month_and_day []int
	if number <= 186 {
		day := number % 31
		month := number/31 + 1
		if day == 0 {
			month--
			day = 31
		}
		month_and_day = append(month_and_day, month, day)
		return month_and_day
	} else if number > 186 && number <= 336 {
		day := (number - 186) % 30
		month := number/30 + 1
		if day == 0 {
			month--
			day = 30
		}
		month_and_day = append(month_and_day, month, day)
		return month_and_day
	} else {
		day := (number - 336) % 29
		if day == 0 {
			day = 29
		}
		month_and_day = append(month_and_day, 12, day)
		return month_and_day
	}
}

func ReadMonth() string {
	day_and_month := CheckMonth()
	return fmt.Sprintf("Month: %v \nDay: %v", day_and_month[0], day_and_month[1])
}

// Main Func

func main() {
	//Person(23, "hosein")
	//fmt.Println(PrimeNumber())
	fmt.Println(ReadMonth())
}
