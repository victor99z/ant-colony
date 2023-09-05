package utils

import "fmt"

var (
	Info  = Teal
	Warn  = Yellow
	Fatal = Red
)

var (
	Black   = Color("\033[1;30m%s\033[0m")
	Red     = Color("\033[1;31m%s\033[0m")
	Green   = Color("\033[1;32m%s\033[0m")
	Yellow  = Color("\033[1;33m%s\033[0m")
	Purple  = Color("\033[1;34m%s\033[0m")
	Magenta = Color("\033[1;35m%s\033[0m")
	Teal    = Color("\033[1;36m%s\033[0m")
	White   = Color("\033[1;37m%s\033[0m")
)

func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}

func PrettyPrint(environment *Enviroment) {
	for i := range environment.GetAll() {
		for j := range environment.GetAll()[i] {
			fodase := environment.GetCellValue(i, j)
			if fodase == 0 {
				fmt.Print(Info("0\t"))
			} else if fodase == 1 {
				fmt.Print(Fatal(fmt.Sprint(fodase) + "\t"))
			} else {
				fmt.Print(Warn(fmt.Sprint(fodase) + "\t"))
			}
			fmt.Print(";")
		}
		fmt.Println()
	}
}
