package output

import (
	"fmt"
	"gtmhub-cli/model"
	"math"
	"os"
	"time"

	"github.com/jedib0t/go-pretty/table"
)

func PrintMetrics(metrics model.Metrics) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "Name", "Attainment", "Days since last checkin"})
	for _, metric := range metrics {
		t.AppendRows([]table.Row{
			{metric.ID, metric.Name, fmt.Sprintf("%v / %v", metric.Actual, metric.Target), GetDaysText(metric.LastCheckin)},
		})
	}
	t.Render()
}

func GetMetricSelectionTable(metrics model.Metrics, cursor int) string {
	t := table.NewWriter()
	//t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"✓", "ID", "Name", "Attainment", "Days since last checkin"})
	for i, metric := range metrics {
		selector := " "
		if cursor == i {
			selector = "> "
		}
		t.AppendRows([]table.Row{
			{selector, metric.ID, metric.Name, fmt.Sprintf("%v / %v", metric.Actual, metric.Target), GetDaysText(metric.LastCheckin)},
		})
	}
	return t.Render()
}

func PrintMetricUpdated() {
	fmt.Println(Red, "Sweet! Your metric has just been updated. Keep going!")
}

func GetDaysText(checkin time.Time) string {
	var defaultDate time.Time
	if checkin.Year() == defaultDate.Year() {
		return "never updated"
	}
	daysFloat := time.Now().Sub(checkin).Hours() / 24
	days := math.Round(daysFloat)

	if days == 0 {
		return "just updated"
	}

	daysStr := "days"
	if days == 1 {
		daysStr = "day"
	}

	return fmt.Sprintf("%v %s ago", days, daysStr)
}
