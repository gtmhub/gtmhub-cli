package model

import "time"

type FullListResponse struct {
	Items      []ListResponse `json:"items"`
	TotalCount int            `json:"totalCount"`
}

type LoadRequest struct {
	Columns []ListColumn `json:"columns"`
	Filter  Filter       `json:"filter"`
}

type ListResponse struct {
	ID      string       `json:"id"`
	Title   string       `json:"title"`
	Filter  Filter       `json:"filter"`
	Columns []ListColumn `json:"columns"`
}

type Filter struct {
	BooleanOperator string      `json:"booleanOperator"`
	RuleBounds      []RuleBound `json:"ruleBounds"`
}

type DynamicFilter struct {
	DynamicFieldValueType string   `json:"dynamicValueType"`
	DynamicFilterValue    []string `json:"value"`
}

type RuleBound struct {
	FieldName      string          `json:"fieldName"`
	Operator       string          `json:"operator"`
	Value          []string        `json:"value"`
	CustomField    bool            `json:"customField"`
	DynamicFilters []DynamicFilter `json:"dynamicFilters"`
}

type ListColumn struct {
	FieldName string `json:"fieldName"`
	FieldType string `json:"type"`
	Width     int    `json:"width"`
}

type KRListResponse struct {
	AccountID       string    `json:"accountId"`
	Attainment      float64   `json:"attainment"`
	GoalID          string    `json:"goalId"`
	LastCheckInDate time.Time `json:"lastCheckInDate"`
}
