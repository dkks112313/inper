package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Loop struct {
	check bool
	pozit int
}

func main() {
	format := os.Args[1]

	if format[(len(format)-4):] != ".bee" {
		fmt.Println("Error: format not found")
		os.Exit(1)
	}

	var file, err = os.ReadFile(os.Args[1])

	if err != nil {
		fmt.Println("Error: file not found")
		os.Exit(1)
	}

	Inter(strings.Fields(string(file)))
	//fmt.Println(strings.Fields(string(file)))
}

func Inter(code []string) {
	var (
		save  = make(map[string]int)
		point = make([]Loop, 0)
	)

	for i := 0; i < len(code); i++ {
		if i < len(code) {
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
			case "print":
				val, ok := save[code[i+1]]
				if ok {
					fmt.Println(val)
				} else {
					fmt.Println("Error: value not found")
					os.Exit(1)
				}
				i++
			case "input":
				val, ok := save[code[i+1]]
				if ok {
					fmt.Scan(&val)
					save[code[i+1]] = val
				} else {
					fmt.Println("Error: value not found")
					os.Exit(1)
				}
				i++
			case "loop":
				_, ok1 := save[code[i+1]]
				_, ok2 := save[code[i+3]]

				if ok1 && ok2 {
					switch code[i+2] {
					case ">":
						if save[code[i+1]] > save[code[i+3]] {
							point = append(point, Loop{check: save[code[i+1]] > save[code[i+3]], pozit: i})
							if len(point) >= 2 && point[len(point)-2] == point[len(point)-1] {
								point = point[:(len(point) - 1)]
							}
							i += 3
						} else {
							j := i
							for {
								j++
								if code[j] == "end" {
									i = j
									break
								}
							}
						}
					case "<":
						if save[code[i+1]] < save[code[i+3]] {
							point = append(point, Loop{check: save[code[i+1]] < save[code[i+3]], pozit: i})
							if len(point) >= 2 && point[len(point)-2] == point[len(point)-1] {
								point = point[:(len(point) - 1)]
							}
							i += 3
						} else {
							j := i
							for {
								j++
								if code[j] == "end" {
									i = j
									break
								}
							}
						}
					case ">=":
						if save[code[i+1]] >= save[code[i+3]] {
							point = append(point, Loop{check: save[code[i+1]] >= save[code[i+3]], pozit: i})
							if len(point) >= 2 && point[len(point)-2] == point[len(point)-1] {
								point = point[:(len(point) - 1)]
							}
							i += 3
						} else {
							j := i
							for {
								j++
								if code[j] == "end" {
									i = j
									break
								}
							}
						}
					case "<=":
						if save[code[i+1]] <= save[code[i+3]] {
							point = append(point, Loop{check: save[code[i+1]] <= save[code[i+3]], pozit: i})
							if len(point) >= 2 && point[len(point)-2] == point[len(point)-1] {
								point = point[:(len(point) - 1)]
							}
							i += 3
						} else {
							j := i
							for {
								j++
								if code[j] == "end" {
									i = j
									break
								}
							}
						}
					case "==":
						if save[code[i+1]] == save[code[i+3]] {
							point = append(point, Loop{check: save[code[i+1]] == save[code[i+3]], pozit: i})
							if len(point) >= 2 && point[len(point)-2] == point[len(point)-1] {
								point = point[:(len(point) - 1)]
							}
							i += 3
						} else {
							j := i
							for {
								j++
								if code[j] == "end" {
									i = j
									break
								}
							}
						}
					case "!=":
						if save[code[i+1]] != save[code[i+3]] {
							point = append(point, Loop{check: save[code[i+1]] != save[code[i+3]], pozit: i})
							if len(point) >= 2 && point[len(point)-2] == point[len(point)-1] {
								point = point[:(len(point) - 1)]
							}
							i += 3
						} else {
							j := i
							for {
								j++
								if code[j] == "end" {
									i = j
									break
								}
							}
						}
					}
				} else if ok1 {
					n, _ := strconv.Atoi(code[i+3])

					switch code[i+2] {
					case ">":
						if save[code[i+1]] > n {
							point = append(point, Loop{check: save[code[i+1]] > n, pozit: i})
							if len(point) >= 2 && point[len(point)-2] == point[len(point)-1] {
								point = point[:(len(point) - 1)]
							}
							i += 3
						} else {
							j := i
							for {
								j++
								if code[j] == "end" {
									i = j
									break
								}
							}
						}
					case "<":
						if save[code[i+1]] < n {
							point = append(point, Loop{check: save[code[i+1]] < n, pozit: i})
							if len(point) >= 2 && point[len(point)-2] == point[len(point)-1] {
								point = point[:(len(point) - 1)]
							}
							i += 3
						} else {
							j := i
							for {
								j++
								if code[j] == "end" {
									i = j
									break
								}
							}
						}
					case ">=":
						if save[code[i+1]] >= n {
							point = append(point, Loop{check: save[code[i+1]] >= n, pozit: i})
							if len(point) >= 2 && point[len(point)-2] == point[len(point)-1] {
								point = point[:(len(point) - 1)]
							}
							i += 3
						} else {
							j := i
							for {
								j++
								if code[j] == "end" {
									i = j
									break
								}
							}
						}
					case "<=":
						if save[code[i+1]] <= n {
							point = append(point, Loop{check: save[code[i+1]] <= n, pozit: i})
							if len(point) >= 2 && point[len(point)-2] == point[len(point)-1] {
								point = point[:(len(point) - 1)]
							}
							i += 3
						} else {
							j := i
							for {
								j++
								if code[j] == "end" {
									i = j
									break
								}
							}
						}
					case "==":
						if save[code[i+1]] == n {
							point = append(point, Loop{check: save[code[i+1]] == n, pozit: i})
							if len(point) >= 2 && point[len(point)-2] == point[len(point)-1] {
								point = point[:(len(point) - 1)]
							}
							i += 3
						} else {
							j := i
							for {
								j++
								if code[j] == "end" {
									i = j
									break
								}
							}
						}
					case "!=":
						if save[code[i+1]] != n {
							point = append(point, Loop{check: save[code[i+1]] != n, pozit: i})
							if len(point) >= 2 && point[len(point)-2] == point[len(point)-1] {
								point = point[:(len(point) - 1)]
							}
							i += 3
						} else {
							j := i
							for {
								j++
								if code[j] == "end" {
									i = j
									break
								}
							}
						}
					}
				} else if ok2 {
					n, _ := strconv.Atoi(code[i+1])

					switch code[i+2] {
					case ">":
						if n > save[code[i+3]] {
							point = append(point, Loop{check: n > save[code[i+3]], pozit: i})
							if len(point) >= 2 && point[len(point)-2] == point[len(point)-1] {
								point = point[:(len(point) - 1)]
							}
							i += 3
						} else {
							j := i
							for {
								j++
								if code[j] == "end" {
									i = j
									break
								}
							}
						}
					case "<":
						if n < save[code[i+3]] {
							point = append(point, Loop{check: n < save[code[i+3]], pozit: i})
							if len(point) >= 2 && point[len(point)-2] == point[len(point)-1] {
								point = point[:(len(point) - 1)]
							}
							i += 3
						} else {
							j := i
							for {
								j++
								if code[j] == "end" {
									i = j
									break
								}
							}
						}
					case ">=":
						if n >= save[code[i+3]] {
							point = append(point, Loop{check: n >= save[code[i+3]], pozit: i})
							if len(point) >= 2 && point[len(point)-2] == point[len(point)-1] {
								point = point[:(len(point) - 1)]
							}
							i += 3
						} else {
							j := i
							for {
								j++
								if code[j] == "end" {
									i = j
									break
								}
							}
						}
					case "<=":
						if n <= save[code[i+3]] {
							point = append(point, Loop{check: n <= save[code[i+3]], pozit: i})
							if len(point) >= 2 && point[len(point)-2] == point[len(point)-1] {
								point = point[:(len(point) - 1)]
							}
							i += 3
						} else {
							j := i
							for {
								j++
								if code[j] == "end" {
									i = j
									break
								}
							}
						}
					case "==":
						if n == save[code[i+3]] {
							point = append(point, Loop{check: n == save[code[i+3]], pozit: i})
							if len(point) >= 2 && point[len(point)-2] == point[len(point)-1] {
								point = point[:(len(point) - 1)]
							}
							i += 3
						} else {
							j := i
							for {
								j++
								if code[j] == "end" {
									i = j
									break
								}
							}
						}
					case "!=":
						if n != save[code[i+3]] {
							point = append(point, Loop{check: n != save[code[i+3]], pozit: i})
							if len(point) >= 2 && point[len(point)-2] == point[len(point)-1] {
								point = point[:(len(point) - 1)]
							}
							i += 3
						} else {
							j := i
							for {
								j++
								if code[j] == "end" {
									i = j
									break
								}
							}
						}
					}
				} else {
					fmt.Println("Error: conteiner not found")
					os.Exit(1)
				}
			case "end":
				if len(point) != 0 {
					if point[len(point)-1].check {
						i = point[len(point)-1].pozit - 1
					} else {
						point = point[:(len(point) - 1)]
					}
				}
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
