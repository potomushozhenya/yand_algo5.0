package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n, year int
	var day_of_week string
	monthMap := map[string]int{
		"January":   1,
		"February":  2,
		"March":     3,
		"April":     4,
		"May":       5,
		"June":      6,
		"July":      7,
		"August":    8,
		"September": 9,
		"October":   10,
		"November":  11,
		"December":  12,
	}
	dayInMonth := map[int]int{
		1:  31,
		3:  31,
		4:  30,
		5:  31,
		6:  30,
		7:  31,
		8:  31,
		9:  30,
		10: 31,
		11: 30,
		12: 31,
	}
	days := [7]int{52, 52, 52, 52, 52, 52, 52}
	dayWeek := map[string]int{
		"Monday":    0,
		"Tuesday":   1,
		"Wednesday": 2,
		"Thursday":  3,
		"Friday":    4,
		"Saturday":  5,
		"Sunday":    6,
	}
	dayWeekRev := map[int]string{
		0: "Monday",
		1: "Tuesday",
		2: "Wednesday",
		3: "Thursday",
		4: "Friday",
		5: "Saturday",
		6: "Sunday",
	}
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	n, err = strconv.Atoi(scanner.Text())
	scanner.Scan()
	year, err = strconv.Atoi(scanner.Text())
	arr := make([]int, n*2)
	for i := 0; i < 2*n; i += 2 {
		var err error
		scanner.Scan()
		line := scanner.Text()
		line = strings.TrimSuffix(line, "\n")
		line = strings.TrimSuffix(line, "\r")
		curr := strings.Split(line, " ")
		arr[i], err = strconv.Atoi(curr[0])
		if err != nil {
			panic(err)
		}
		arr[i+1] = monthMap[curr[1]]
	}
	scanner.Scan()
	day_of_week = scanner.Text()
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		dayInMonth[2] = 29
		c := dayWeek[day_of_week]
		days[c] += 1
		c = (c + 1) % 7
		days[c] += 1
	} else {
		dayInMonth[2] = 28
		c := dayWeek[day_of_week]
		days[c] += 1
	}
	for i := 0; i < 2*n; i += 2 {
		dayOfYear := arr[i]
		i := arr[i+1]
		for {
			if i == 1 {
				break
			}
			i--
			dayOfYear += dayInMonth[i]
		}
		dayOfYear += dayWeek[day_of_week] - 1
		dayOfYear = dayOfYear % 7
		days[dayOfYear] -= 1
	}
	max := 0
	min := 0
	for i := 1; i < 7; i++ {
		if days[i] < days[min] {
			min = i
		}
		if days[i] > days[max] {
			max = i
		}
	}
	f, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString(dayWeekRev[max] + " ")
	f.WriteString(dayWeekRev[min])

}
