package mod

import (
	"fmt"
	"os"
	"strconv"
)

/*
Interpreter Translater
*/
func Inter(code []string) {
	var (
		save = make(map[string]int)
	)

	for i := 0; i < len(code); i++ {
		if i < len(code)-1 {
			switch code[i] {
			case "add":
				_, ok1 := save[code[i+1]]
				_, ok2 := save[code[i+2]]

				if ok1 && ok2 {
					save[code[i+1]] = save[code[i+1]] + save[code[i+2]]
				} else if ok1 {
					n, _ := strconv.Atoi(code[i+2])
					save[code[i+1]] = save[code[i+1]] + n
				} else {
					fmt.Println("Error: conteiner not found")
					os.Exit(1)
				}
				i += 2
			case "mul":
				_, ok1 := save[code[i+1]]
				_, ok2 := save[code[i+2]]

				if ok1 && ok2 {
					save[code[i+1]] = save[code[i+1]] * save[code[i+2]]
				} else if ok1 {
					n, _ := strconv.Atoi(code[i+2])
					save[code[i+1]] = save[code[i+1]] * n
				} else {
					fmt.Println("Error: conteiner not found")
					os.Exit(1)
				}
				i += 2
			case "div":
				_, ok1 := save[code[i+1]]
				_, ok2 := save[code[i+2]]

				if ok1 && ok2 {
					save[code[i+1]] = save[code[i+1]] / save[code[i+2]]
				} else if ok1 {
					n, _ := strconv.Atoi(code[i+2])
					save[code[i+1]] = save[code[i+1]] / n
				} else {
					fmt.Println("Error: conteiner not found")
					os.Exit(1)
				}
				i += 2
			case "out":
				val, ok := save[code[i+1]]
				if ok {
					fmt.Println(val)
				} else {
					fmt.Println("Error: value not found")
					os.Exit(1)
				}
				i++
			case "inp":
				val, ok := save[code[i+1]]
				if ok {
					fmt.Scan(&val)
					save[code[i+1]] = val
				} else {
					fmt.Println("Error: value not found")
					os.Exit(1)
				}
				i++
			default:
				if code[i+1] == "=" {
					save[code[i]], _ = strconv.Atoi(code[i+2])
				} else {
					fmt.Println("Error: syntax error")
					os.Exit(1)
				}
				i += 2
			}
		}
	}
}
