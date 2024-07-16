package main

import "fmt"

const inf int = int(^uint(0) >> 1)

func sim(hp int, units int, enemyProd int, stop int) int {
	c := 0
	var enemy int = 0
	for {
		if hp-stop < 0 {
			break
		}
		if enemy >= units {
			return inf
		}
		hp -= (units - enemy)
		enemy = 0
		if hp > 0 {
			enemy += enemyProd
		}
		c++
	}
	for {
		if hp <= 0 {
			break
		}
		if hp >= units {
			hp -= units
		} else {
			enemy -= units - hp
			hp = 0
		}
		units -= enemy
		if hp > 0 {
			enemy += enemyProd
		}
		c++
		if units <= 0 {
			return inf
		}
	}
	for {
		if enemy <= 0 {
			break
		}
		enemy -= units
		if enemy > 0 {
			units -= enemy
		}
		c++
		if units <= 0 {
			return inf
		}
	}
	return c
}
func main() {
	var x, y, p, ans int
	fmt.Scan(&x, &y, &p)
	ans = inf
	for i := 0; i < y+2; i++ {
		ans = min(ans, sim(y, x, p, i))
	}
	if ans != inf {
		fmt.Print(ans)
	} else {
		fmt.Print(-1)
	}
}
