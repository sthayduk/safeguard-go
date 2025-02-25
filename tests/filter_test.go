package tests

import (
	"reflect"
	"testing"

	"github.com/sthayduk/safeguard-go/client"
)

func TestAddField(t *testing.T) {
	tests := []struct {
		name     string
		initial  client.Fields
		field    string
		expected client.Fields
	}{
		{
			name:     "Add field to empty list",
			initial:  client.Fields{},
			field:    "newField",
			expected: client.Fields{"newField"},
		},
		{
			name:     "Add field to non-empty list",
			initial:  client.Fields{"field1", "field2"},
			field:    "newField",
			expected: client.Fields{"field1", "field2", "newField"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &client.Filter{Fields: tt.initial}
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
		initial  client.Fields
		field    string
		expected client.Fields
	}{
		{
			name:     "Remove field from list with one element",
			initial:  client.Fields{"field1"},
			field:    "field1",
			expected: client.Fields{},
		},
		{
			name:     "Remove field from list with multiple elements",
			initial:  client.Fields{"field1", "field2", "field3"},
			field:    "field2",
			expected: client.Fields{"field1", "field3"},
		},
		{
			name:     "Remove field that does not exist",
			initial:  client.Fields{"field1", "field2"},
			field:    "field3",
			expected: client.Fields{"field1", "field2"},
		},
		{
			name:     "Remove field from empty list",
			initial:  client.Fields{},
			field:    "field1",
			expected: client.Fields{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &client.Filter{Fields: tt.initial}
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
		initial  client.Fields
		expected client.Fields
	}{
		{
			name:     "Get fields from empty list",
			initial:  client.Fields{},
			expected: client.Fields{},
		},
		{
			name:     "Get fields from non-empty list",
			initial:  client.Fields{"field1", "field2"},
			expected: client.Fields{"field1", "field2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &client.Filter{Fields: tt.initial}
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
		initial  client.FilterQuery
		field    string
		operator string
		value    string
		expected client.FilterQuery
	}{
		{
			name:     "Add filter to empty filter",
			initial:  "",
			field:    "field1",
			operator: "eq",
			value:    "value1",
			expected: `field1 eq "value1"`,
		},
		{
			name:     "Add filter to non-empty filter",
			initial:  `field1 eq "value1"`,
			field:    "field2",
			operator: "ne",
			value:    "value2",
			expected: `field1 eq "value1"field2 ne "value2"`,
		},
		{
			name:     "Add filter with special characters",
			initial:  "",
			field:    "field1",
			operator: "contains",
			value:    `value"with\special*chars`,
			expected: `field1 contains "value\"with\\special\*chars"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &client.Filter{Filter: tt.initial}
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
		initial  client.OrderBy
		field    string
		expected client.OrderBy
	}{
		{
			name:     "Add orderby to empty list",
			initial:  client.OrderBy{},
			field:    "newOrderBy",
			expected: client.OrderBy{"newOrderBy"},
		},
		{
			name:     "Add orderby to non-empty list",
			initial:  client.OrderBy{"orderby1", "orderby2"},
			field:    "newOrderBy",
			expected: client.OrderBy{"orderby1", "orderby2", "newOrderBy"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &client.Filter{Orderby: tt.initial}
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
		initial  client.OrderBy
		expected client.OrderBy
	}{
		{
			name:     "Get orderby from empty list",
			initial:  client.OrderBy{},
			expected: client.OrderBy{},
		},
		{
			name:     "Get orderby from non-empty list",
			initial:  client.OrderBy{"orderby1", "orderby2"},
			expected: client.OrderBy{"orderby1", "orderby2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &client.Filter{Orderby: tt.initial}
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
		initial  client.OrderBy
		field    string
		expected client.OrderBy
	}{
		{
			name:     "Remove orderby from list with one element",
			initial:  client.OrderBy{"orderby1"},
			field:    "orderby1",
			expected: client.OrderBy{},
		},
		{
			name:     "Remove orderby from list with multiple elements",
			initial:  client.OrderBy{"orderby1", "orderby2", "orderby3"},
			field:    "orderby2",
			expected: client.OrderBy{"orderby1", "orderby3"},
		},
		{
			name:     "Remove orderby that does not exist",
			initial:  client.OrderBy{"orderby1", "orderby2"},
			field:    "orderby3",
			expected: client.OrderBy{"orderby1", "orderby2"},
		},
		{
			name:     "Remove orderby from empty list",
			initial:  client.OrderBy{},
			field:    "orderby1",
			expected: client.OrderBy{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &client.Filter{Orderby: tt.initial}
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
		filter   *client.Filter
		expected string
	}{
		{
			name:     "Empty filter",
			filter:   &client.Filter{},
			expected: "?count=false",
		},
		{
			name: "client.Filter with fields",
			filter: &client.Filter{
				Fields: client.Fields{"field1", "field2"},
			},
			expected: "?fields=field1%2Cfield2&count=false",
		},
		{
			name: "client.Filter with orderby",
			filter: &client.Filter{
				Orderby: client.OrderBy{"field1", "field2"},
			},
			expected: "?count=false&orderby=field1%2Cfield2",
		},
		{
			name: "client.Filter with count true",
			filter: &client.Filter{
				Count: true,
			},
			expected: "?count=true",
		},
		{
			name: "client.Filter with count false",
			filter: &client.Filter{
				Count: false,
			},
			expected: "?count=false",
		},
		{
			name: "client.Filter with filter query",
			filter: &client.Filter{
				Filter: client.FilterQuery(`field1 eq "value1"`),
			},
			expected: `?filter=field1%20eq%20%22value1%22&count=false`,
		},
		{
			name: "client.Filter with all parameters",
			filter: &client.Filter{
				Fields:  client.Fields{"field1", "field2"},
				Orderby: client.OrderBy{"field1", "field2"},
				Count:   true,
				Filter:  client.FilterQuery(`field1 eq "value1"`),
			},
			expected: `?filter=field1%20eq%20%22value1%22&fields=field1%2Cfield2&count=true&orderby=field1%2Cfield2`,
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

// Add new tests for String() methods
func TestTypeStrings(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected string
	}{
		{
			name:     "client.Fields with multiple values",
			input:    client.Fields{"field1", "field2", "field3"},
			expected: "field1,field2,field3",
		},
		{
			name:     "Empty client.Fields",
			input:    client.Fields{},
			expected: "",
		},
		{
			name:     "client.FilterQuery with value",
			input:    client.FilterQuery("name eq 'John'"),
			expected: "name eq 'John'",
		},
		{
			name:     "Empty client.FilterQuery",
			input:    client.FilterQuery(""),
			expected: "",
		},
		{
			name:     "OrderBy with multiple values",
			input:    client.OrderBy{"name", "age", "date"},
			expected: "name,age,date",
		},
		{
			name:     "Empty OrderBy",
			input:    client.OrderBy{},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got string
			switch v := tt.input.(type) {
			case client.Fields:
				got = v.String()
			case client.FilterQuery:
				got = v.String()
			case client.OrderBy:
				got = v.String()
			}
			if got != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, got)
			}
		})
	}
}
