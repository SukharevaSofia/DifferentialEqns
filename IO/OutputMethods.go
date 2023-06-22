package IO

import (
	"DifferentialEqns/methods"
	"DifferentialEqns/utils"
	"github.com/olekukonko/tablewriter"
	"github.com/wcharczuk/go-chart/v2"
	"math"
	"os"
	"sort"
	"strconv"
)

func unify(currentList, newList []float64, precision float64) []float64 {
	newSet := map[float64]bool{}
	for _, element := range newList {
		newSet[utils.Truncate(element, precision)] = true
	}
	for _, element := range currentList {
		newSet[utils.Truncate(element, precision)] = true
	}
	result := []float64{}
	for key := range newSet {
		result = append(result, key)
	}
	sort.Float64s(result)
	return result
}

func tableMapRetrieveHelper(givenMap map[float64]float64, key float64, accuracy int) string {
	if value, ok := givenMap[key]; ok {
		return strconv.FormatFloat(value, 'f', accuracy, 64)
	}
	return ""
}

func OutputResults(euler, runge, adams, precise utils.XY, acc float64) {
	accuracy := int(acc)
	accuracyFloat := math.Pow(10, -acc)
	data := [][]string{}
	if utils.TABLE_FULL {
		xList := []float64{}
		xList = unify(xList, euler.X, accuracyFloat)
		xList = unify(xList, runge.X, accuracyFloat)
		xList = unify(xList, adams.X, accuracyFloat)
		xList = unify(xList, precise.X, accuracyFloat)

		eulerMap := utils.XYtoMap(euler, accuracyFloat)
		rungeMap := utils.XYtoMap(runge, accuracyFloat)
		adamsMap := utils.XYtoMap(adams, accuracyFloat)
		preciseMap := utils.XYtoMap(precise, accuracyFloat)

		for i := 0; i < len(xList); i++ {
			iStr := strconv.Itoa(i)
			xiStr := strconv.FormatFloat(xList[i], 'f', accuracy, 64)
			eulerStr := tableMapRetrieveHelper(eulerMap, xList[i], accuracy)
			rungeStr := tableMapRetrieveHelper(rungeMap, xList[i], accuracy)
			adamsStr := tableMapRetrieveHelper(adamsMap, xList[i], accuracy)
			preciseStr := tableMapRetrieveHelper(preciseMap, xList[i], accuracy)
			if eulerStr != "" || rungeStr != "" || adamsStr != "" || preciseStr != "" {
				line := []string{iStr, xiStr, eulerStr, rungeStr, adamsStr, preciseStr}
				data = append(data, line)
			}
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"i", "x_i", "Метод Эйлера", "Рунге-Кутта", "Адамса", "Точное решение"})

		for _, v := range data {
			table.Append(v)
		}
		if utils.GRAPH_FULL {
			buildChart(euler, runge, adams, precise)
		} else {
			buildChart(methods.Cutout(precise, euler), methods.Cutout(precise, runge), methods.Cutout(precise, adams), precise)
		}
		table.Render()
	} else {
		adamsUncut := adams
		rungeUncut := runge
		eulerUncut := euler
		adams = methods.Cutout(precise, adams)
		runge = methods.Cutout(precise, runge)
		euler = methods.Cutout(precise, euler)
		for i := 0; i < len(precise.X); i++ {
			iStr := strconv.Itoa(i)
			xiStr := strconv.FormatFloat(euler.X[i], 'f', accuracy, 64)
			eulerStr := strconv.FormatFloat(euler.Y[i], 'f', accuracy, 64)
			rungeStr := strconv.FormatFloat(runge.Y[i], 'f', accuracy, 64)
			adamsStr := strconv.FormatFloat(adams.Y[i], 'f', accuracy, 64)
			preciseStr := strconv.FormatFloat(precise.Y[i], 'f', accuracy, 64)
			line := []string{iStr, xiStr, eulerStr, rungeStr, adamsStr, preciseStr}
			data = append(data, line)
		}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"i", "x_i", "Метод Эйлера", "Рунге-Кутта", "Адамса", "Точное решение"})

		for _, v := range data {
			table.Append(v)
		}
		if utils.GRAPH_FULL {
			buildChart(eulerUncut, rungeUncut, adamsUncut, precise)
		} else {
			buildChart(euler, runge, adams, precise)
		}
		table.Render()
	}
}

func buildChart(euler, runge, adams, precise utils.XY) {
	graph := chart.Chart{
		Series: []chart.Series{
			chart.ContinuousSeries{
				Style: chart.Style{
					StrokeWidth: chart.Disabled,
					DotWidth:    2,
				},
				Name:    "точные значения",
				XValues: precise.X,
				YValues: precise.Y,
			}, chart.ContinuousSeries{
				Style: chart.Style{
					DotWidth: 2,
				},
				Name:    "Эйлер",
				XValues: euler.X,
				YValues: euler.Y,
			},
			chart.ContinuousSeries{
				Style: chart.Style{
					DotWidth: 2,
				},
				Name:    "Рунге-Кутта",
				XValues: runge.X,
				YValues: runge.Y,
			},
			chart.ContinuousSeries{
				Style: chart.Style{
					DotWidth: 2,
				},
				Name:    "Адамс",
				XValues: adams.X,
				YValues: adams.Y,
			},
		},
	}
	graph.Elements = []chart.Renderable{
		chart.Legend(&graph),
	}
	picture, _ := os.Create("graph.png")
	graph.Render(chart.PNG, picture)
	picture.Close()
}
