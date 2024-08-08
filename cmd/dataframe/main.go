package main

import (
	"context"
	"flag"
	"os"
	"strings"
	"time"

	"github.com/jomei/notionapi"
	"github.com/rocketlaunchr/dataframe-go"
	"github.com/rocketlaunchr/dataframe-go/forecast/interpolation"
	"github.com/rocketlaunchr/dataframe-go/plot"
	"github.com/rocketlaunchr/dataframe-go/utils/utime"
	"github.com/samber/lo"
	"github.com/wcharczuk/go-chart/v2"
)

func main() {
	secret := flag.String("secret", "", "secret")
	database := flag.String("database", "", "database")
	show := flag.Bool("show", false, "show")
	output := flag.String("output", "", "output")

	flag.Parse()

	client := notionapi.NewClient(notionapi.Token(*secret))

	pages := lo.Must1(FetchAllRows(context.Background(), client, notionapi.DatabaseID(*database), &notionapi.DatabaseQueryRequest{Sorts: []notionapi.SortObject{{
		Property:  "Date",
		Direction: notionapi.SortOrderASC,
	}}}))

	var beginTime time.Time
	data := make(map[time.Time]float64)
	for _, page := range pages {
		dateStr := strings.Split(page.Properties["Date"].(*notionapi.DateProperty).Date.Start.String(), "T")[0]
		date := lo.Must1(time.ParseInLocation("2006-01-02", dateStr, time.Local))
		if beginTime.IsZero() {
			beginTime = date
		}
		weight := page.Properties["Weight"].(*notionapi.NumberProperty).Number / 1000
		data[date] = weight
	}
	now := time.Now()
	endTime := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, time.Local)
	dateSeries, _ := utime.NewSeriesTime(context.Background(), "date", "1D", beginTime, false, utime.NewSeriesTimeOptions{
		Until: &endTime,
	})
	weightVals := make([]interface{}, 0, len(dateSeries.Values))
	for _, value := range dateSeries.Values {
		val, ok := data[*value]
		if ok {
			weightVals = append(weightVals, val)
		} else {
			weightVals = append(weightVals, nil)
		}
	}
	weightSeries := dataframe.NewSeriesFloat64("weight", nil, weightVals...)
	df := dataframe.NewDataFrame(dateSeries, weightSeries)
	lo.Must1(interpolation.Interpolate(context.Background(), df, interpolation.InterpolateOptions{
		Method:  interpolation.Spline{Order: 3},
		InPlace: true,
	}))
	cs := chart.TimeSeries{
		Name: "Weights",
		XValues: lo.Map(df.Series[0].(*dataframe.SeriesTime).Values, func(item *time.Time, index int) time.Time {
			return *item
		}),
		YValues: df.Series[1].(*dataframe.SeriesFloat64).Values,
	}
	graph := chart.Chart{Series: []chart.Series{cs}}

	if *output != "" {
		type render = func(int, int) (chart.Renderer, error)
		for _, t := range []lo.Tuple2[string, render]{lo.T2("svg", chart.SVG), lo.T2("png", chart.PNG)} {
			f := lo.Must1(os.Create(*output + "." + t.A))
			lo.Must0(graph.Render(t.B, f))
			lo.Must0(f.Close())
		}
	}

	if *show {
		plt := lo.Must1(plot.Open("Weights", 450, 300))
		graph.Render(chart.SVG, plt)
		plt.Display(plot.None)
		<-plt.Closed
	}
}

func FetchAllRows(ctx context.Context, client *notionapi.Client, id notionapi.DatabaseID, requestBody *notionapi.DatabaseQueryRequest) ([]notionapi.Page, error) {
	if requestBody != nil {
		requestBody.PageSize = 2
	}
	resp, err := client.Database.Query(ctx, id, requestBody)
	if err != nil {
		return nil, err
	}
	var results []notionapi.Page
	results = append(results, resp.Results...)
	for resp.HasMore {
		requestBody.StartCursor = resp.NextCursor
		resp, err = client.Database.Query(ctx, id, requestBody)
		if err != nil {
			return nil, err
		}
		results = append(results, resp.Results...)
	}
	return results, nil
}
