package model

import "time"

type Metric struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Actual      float64   `json:"actual"`
	Target      float64   `json:"target"`
	LastCheckin time.Time `json:"lastCheckInDate"`
}

type Metrics []Metric

type FulLMetricResponse struct {
	Items Metrics `json:"items"`
}

func (m Metrics) FilterMetrics(date time.Time) Metrics {
	result := make([]Metric, 0)
	for _, metric := range m {
		if metric.LastCheckin.Before(date) {
			result = append(result, metric)
		}

	}

	return result
}

type CheckInMetricRequest struct {
	Actual       float64                 `json:"actual"`
	CheckInDate  *time.Time              `json:"checkInDate"`
	Comment      string                  `json:"comment"`
	Confidence   *float64                `json:"confidence"`
	CustomFields *map[string]interface{} `json:"customFields,omitempty"`
}
