package IO

import (
	"DifferentialEqns/utils"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
)

func OutputResults(euler, runge, precise utils.XY, acc float64) {
	fmt.Println("Точное: ", precise.Y)
	fmt.Println("Euler: ", euler.Y)
	fmt.Println("runge: ", runge.Y)
	accuracy := int(acc)
	data := [][]string{}
	for i := 0; i < len(euler.X); i++ {
		iStr := strconv.Itoa(i)
		xiStr := strconv.FormatFloat(euler.X[i], 'f', accuracy, 64)
		eulerStr := strconv.FormatFloat(euler.Y[i], 'f', accuracy, 64)
		rungeStr := strconv.FormatFloat(runge.Y[i], 'f', accuracy, 64)
		preciseStr := strconv.FormatFloat(precise.Y[i], 'f', accuracy, 64)
		line := []string{iStr, xiStr, eulerStr, rungeStr, " ", preciseStr}
		data = append(data, line)
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"i", "x_i", "Метод Эйлера", "Рунге-Кутта", "Адамса", "Точное решение"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render() // Send output

}

func buildChart(euler, runge, adams, precise utils.XY) {

}
