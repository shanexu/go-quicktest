package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/rocketlaunchr/dataframe-go"
	"github.com/rocketlaunchr/dataframe-go/forecast/interpolation"
	"github.com/rocketlaunchr/dataframe-go/imports"
	"github.com/rocketlaunchr/dataframe-go/plot"
	"github.com/rocketlaunchr/dataframe-go/utils/utime"
	"github.com/samber/lo"
	"github.com/wcharczuk/go-chart/v2"
)

func main() {
	f, _ := os.Open("./cmd/dataframe/example.csv")
	defer f.Close()
	csvDf, _ := imports.LoadFromCSV(context.Background(), f)
	data := make(map[time.Time]float64)
	var beginTime time.Time
	iter := csvDf.ValuesIterator(dataframe.ValuesOptions{InitialRow: 0, Step: 1})
	for {
		row, vals, _ := iter()
		if row == nil {
			break
		}
		d, _ := vals["Date"].(string)
		day, _ := time.ParseInLocation("2006/01/02", d, time.Local)
		if beginTime.IsZero() {
			beginTime = day
		}
		w, _ := vals["Weight"].(string)
		weight, _ := strconv.ParseFloat(w, 64)
		weight /= 1000
		fmt.Println(day, weight)
		data[day] = weight
	}
	now := time.Now()
	endTime := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, time.Local)
	dates, _ := utime.NewSeriesTime(context.Background(), "date", "1D", beginTime, false, utime.NewSeriesTimeOptions{
		Until: &endTime,
	})
	weightVals := make([]interface{}, 0, len(dates.Values))
	for _, value := range dates.Values {
		val, ok := data[*value]
		if ok {
			weightVals = append(weightVals, val)
		} else {
			weightVals = append(weightVals, nil)
		}
	}

	weights := dataframe.NewSeriesFloat64("weight", nil, weightVals...)
	df := dataframe.NewDataFrame(dates, weights)
	fmt.Println(df.String())
	_, err := interpolation.Interpolate(context.Background(), df, interpolation.InterpolateOptions{
		Method:  interpolation.Spline{Order: 3},
		InPlace: true,
	})
	if err != nil {
		panic(err)
	}
	cs := chart.TimeSeries{
		Name: "Weight",
		XValues: lo.Map(df.Series[0].(*dataframe.SeriesTime).Values, func(item *time.Time, index int) time.Time {
			return *item
		}),
		YValues: df.Series[1].(*dataframe.SeriesFloat64).Values,
	}
	fmt.Println(df.String())
	graph := chart.Chart{Series: []chart.Series{cs}}
	plt, err := plot.Open("Weights", 450, 300)
	if err != nil {
		panic(err)
	}
	graph.Render(chart.SVG, plt)
	plt.Display(plot.None)
	<-plt.Closed
}
