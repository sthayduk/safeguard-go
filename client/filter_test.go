package client

import (
	"reflect"
	"testing"
)

func TestAddField(t *testing.T) {
	tests := []struct {
		name     string
		initial  []string
		field    string
		expected []string
	}{
		{
			name:     "Add field to empty list",
			initial:  []string{},
			field:    "newField",
			expected: []string{"newField"},
		},
		{
			name:     "Add field to non-empty list",
			initial:  []string{"field1", "field2"},
			field:    "newField",
			expected: []string{"field1", "field2", "newField"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &Filter{Fields: tt.initial}
			filter.AddField(tt.field)
			if !reflect.DeepEqual(filter.Fields, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, filter.Fields)
			}
		})
	}
}

func TestRemoveField(t *testing.T) {
	tests := []struct {
		name     string
		initial  []string
		field    string
		expected []string
	}{
		{
			name:     "Remove field from list with one element",
			initial:  []string{"field1"},
			field:    "field1",
			expected: []string{},
		},
		{
			name:     "Remove field from list with multiple elements",
			initial:  []string{"field1", "field2", "field3"},
			field:    "field2",
			expected: []string{"field1", "field3"},
		},
		{
			name:     "Remove field that does not exist",
			initial:  []string{"field1", "field2"},
			field:    "field3",
			expected: []string{"field1", "field2"},
		},
		{
			name:     "Remove field from empty list",
			initial:  []string{},
			field:    "field1",
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &Filter{Fields: tt.initial}
			filter.RemoveField(tt.field)
			if !reflect.DeepEqual(filter.Fields, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, filter.Fields)
			}
		})
	}
}

func TestGetFields(t *testing.T) {
	tests := []struct {
		name     string
		initial  []string
		expected []string
	}{
		{
			name:     "Get fields from empty list",
			initial:  []string{},
			expected: []string{},
		},
		{
			name:     "Get fields from non-empty list",
			initial:  []string{"field1", "field2"},
			expected: []string{"field1", "field2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &Filter{Fields: tt.initial}
			got := filter.GetFields()
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, got)
			}
		})
	}
}

func TestAddFilter(t *testing.T) {
	tests := []struct {
		name     string
		initial  string
		field    string
		operator string
		value    string
		expected string
	}{
		{
			name:     "Add filter to empty filter",
			initial:  "",
			field:    "field1",
			operator: "eq",
			value:    "value1",
			expected: "field1 eq value1",
		},
		{
			name:     "Add filter to non-empty filter",
			initial:  "field1 eq value1",
			field:    "field2",
			operator: "ne",
			value:    "value2",
			expected: "field1 eq value1field2 ne value2",
		},
		{
			name:     "Add filter with special characters",
			initial:  "",
			field:    "field1",
			operator: "contains",
			value:    `value"with\special*chars`,
			expected: `field1 contains value\"with\\special\*chars`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &Filter{Filter: tt.initial}
			filter.AddFilter(tt.field, tt.operator, tt.value)
			if filter.Filter != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, filter.Filter)
			}
		})
	}
}

func TestAddOrderBy(t *testing.T) {
	tests := []struct {
		name     string
		initial  []string
		field    string
		expected []string
	}{
		{
			name:     "Add orderby to empty list",
			initial:  []string{},
			field:    "newOrderBy",
			expected: []string{"newOrderBy"},
		},
		{
			name:     "Add orderby to non-empty list",
			initial:  []string{"orderby1", "orderby2"},
			field:    "newOrderBy",
			expected: []string{"orderby1", "orderby2", "newOrderBy"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &Filter{Orderby: tt.initial}
			filter.AddOrderBy(tt.field)
			if !reflect.DeepEqual(filter.Orderby, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, filter.Orderby)
			}
		})
	}
}

func TestGetOrderBy(t *testing.T) {
	tests := []struct {
		name     string
		initial  []string
		expected []string
	}{
		{
			name:     "Get orderby from empty list",
			initial:  []string{},
			expected: []string{},
		},
		{
			name:     "Get orderby from non-empty list",
			initial:  []string{"orderby1", "orderby2"},
			expected: []string{"orderby1", "orderby2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &Filter{Orderby: tt.initial}
			got := filter.GetOrderBy()
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, got)
			}
		})
	}
}

func TestRemoveOrderBy(t *testing.T) {
	tests := []struct {
		name     string
		initial  []string
		field    string
		expected []string
	}{
		{
			name:     "Remove orderby from list with one element",
			initial:  []string{"orderby1"},
			field:    "orderby1",
			expected: []string{},
		},
		{
			name:     "Remove orderby from list with multiple elements",
			initial:  []string{"orderby1", "orderby2", "orderby3"},
			field:    "orderby2",
			expected: []string{"orderby1", "orderby3"},
		},
		{
			name:     "Remove orderby that does not exist",
			initial:  []string{"orderby1", "orderby2"},
			field:    "orderby3",
			expected: []string{"orderby1", "orderby2"},
		},
		{
			name:     "Remove orderby from empty list",
			initial:  []string{},
			field:    "orderby1",
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &Filter{Orderby: tt.initial}
			filter.RemoveOrderBy(tt.field)
			if !reflect.DeepEqual(filter.Orderby, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, filter.Orderby)
			}
		})
	}
}
func TestToQueryString(t *testing.T) {
	tests := []struct {
		name     string
		filter   *Filter
		expected string
	}{
		{
			name:     "Empty filter",
			filter:   &Filter{},
			expected: "?count=false",
		},
		{
			name: "Filter with fields",
			filter: &Filter{
				Fields: []string{"field1", "field2"},
			},
			expected: "?fields=field1%2Cfield2&count=false",
		},
		{
			name: "Filter with orderby",
			filter: &Filter{
				Orderby: []string{"field1", "field2"},
			},
			expected: "?count=false&orderby=field1%2Cfield2",
		},
		{
			name: "Filter with count true",
			filter: &Filter{
				Count: true,
			},
			expected: "?count=true",
		},
		{
			name: "Filter with count false",
			filter: &Filter{
				Count: false,
			},
			expected: "?count=false",
		},
		{
			name: "Filter with filter query",
			filter: &Filter{
				Filter: "field1 eq value1",
			},
			expected: "?filter=field1%20eq%20value1&count=false",
		},
		{
			name: "Filter with all parameters",
			filter: &Filter{
				Fields:  []string{"field1", "field2"},
				Orderby: []string{"field1", "field2"},
				Count:   true,
				Filter:  "field1 eq value1",
			},
			expected: "?filter=field1%20eq%20value1&fields=field1%2Cfield2&count=true&orderby=field1%2Cfield2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.filter.ToQueryString()
			if got != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, got)
			}
		})
	}
}
