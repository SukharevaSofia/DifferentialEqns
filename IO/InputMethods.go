package IO

import (
	"DifferentialEqns/utils"
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func UserInput(eqn1, eqn2, eqn3 utils.Equation) (float64, float64, float64, float64, float64, float64, int) {
	fmt.Println(utils.INFO)
	fmt.Println(utils.CHOOSE_FUNC)
	fmt.Println("1: ", eqn1.NameOfFunction)
	fmt.Println("2: ", eqn2.NameOfFunction)
	fmt.Println("3: ", eqn3.NameOfFunction)
	var inputString string
	var eqn int
	for {
		fmt.Scan(&eqn)
		if !(eqn == 1 || eqn == 2 || eqn == 3) {
			fmt.Println(utils.INPUT_ERR)
			continue
		}
		break
	}
	fmt.Println(utils.CHOOSE_INPUT)
	for {
		fmt.Scan(&inputString)
		if !(inputString == "T" || inputString == "t" || inputString == "F" || inputString == "f") {
			fmt.Println(utils.INPUT_ERR)
			continue
		}
		break
	}
	if inputString == "T" || inputString == "t" {
		return terminalInput(eqn)
	}
	return fileInput(eqn)
}

func terminalInput(eqn int) (float64, float64, float64, float64, float64, float64, int) {
	fmt.Println(utils.INPUT_X0Y0)
	var x0, y0 float64
	fmt.Scan(&x0, &y0)

	fmt.Println(utils.INPUT_INTERVAL)
	var xo, xn float64
	for {
		fmt.Scan(&xo, &xn)
		if xo >= xn {
			fmt.Println(utils.INPUT_ERR)
			continue
		}
		break
	}

	fmt.Println(utils.INPUT_STEP)
	var h float64
	for {
		fmt.Scan(&h)
		if xn-xo <= h {
			fmt.Println(utils.INPUT_ERR)
			continue
		}
		break
	}

	fmt.Println(utils.INPUT_ACC)
	var accuracyStr string
	var accuracy float64
	for {
		fmt.Scan(&accuracyStr)
		if _, err := strconv.ParseFloat(accuracyStr, 64); err == nil {
			break
		}
		fmt.Print(utils.INPUT_ERR)
	}
	acc, _ := strconv.ParseFloat(accuracyStr, 64)
	for math.Mod(acc, 1) != 0 {
		accuracy++
		acc *= 10
	}
	return x0, y0, xo, xn, h, accuracy, eqn
}

func fileInput(eqn int) (float64, float64, float64, float64, float64, float64, int) {
	f, _ := os.Open("./data/data1.txt")
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	data := strings.Split(scanner.Text(), " ")
	x0_, y0_, xo_, xn_, h_, accuracy_ := data[0], data[1], data[2], data[3], data[4], data[5]
	x0, _ := strconv.ParseFloat(x0_, 64)
	y0, _ := strconv.ParseFloat(y0_, 64)
	xo, _ := strconv.ParseFloat(xo_, 64)
	xn, _ := strconv.ParseFloat(xn_, 64)
	h, _ := strconv.ParseFloat(h_, 64)
	acc, _ := strconv.ParseFloat(accuracy_, 64)
	if xo > xn || (xn-xo <= h) {
		fmt.Print(utils.INPUT_ERR)
		os.Exit(1)
	}
	var accuracy float64
	for math.Mod(acc, 1) != 0 {
		accuracy++
		acc *= 10
	}
	return x0, y0, xo, xn, h, accuracy, eqn
}
