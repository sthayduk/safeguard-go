package safeguard

import (
	"reflect"
	"strings"
	"testing"
)

func TestAddField(t *testing.T) {
	tests := []struct {
		name     string
		initial  Fields
		field    string
		expected Fields
	}{
		{
			name:     "Add field to empty list",
			initial:  Fields{},
			field:    "newField",
			expected: Fields{"newField"},
		},
		{
			name:     "Add field to non-empty list",
			initial:  Fields{"field1", "field2"},
			field:    "newField",
			expected: Fields{"field1", "field2", "newField"},
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
		initial  Fields
		field    string
		expected Fields
	}{
		{
			name:     "Remove field from list with one element",
			initial:  Fields{"field1"},
			field:    "field1",
			expected: Fields{},
		},
		{
			name:     "Remove field from list with multiple elements",
			initial:  Fields{"field1", "field2", "field3"},
			field:    "field2",
			expected: Fields{"field1", "field3"},
		},
		{
			name:     "Remove field that does not exist",
			initial:  Fields{"field1", "field2"},
			field:    "field3",
			expected: Fields{"field1", "field2"},
		},
		{
			name:     "Remove field from empty list",
			initial:  Fields{},
			field:    "field1",
			expected: Fields{},
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
		initial  Fields
		expected Fields
	}{
		{
			name:     "Get fields from empty list",
			initial:  Fields{},
			expected: Fields{},
		},
		{
			name:     "Get fields from non-empty list",
			initial:  Fields{"field1", "field2"},
			expected: Fields{"field1", "field2"},
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
		initial  []FilterQuery
		field    string
		operator FilterOperator
		value    string
		expected string
	}{
		{
			name:     "Add filter to empty filter",
			initial:  []FilterQuery{},
			field:    "field1",
			operator: OpEqual,
			value:    "value1",
			expected: "field1 eq 'value1'",
		},
		{
			name:     "Add filter to non-empty filter",
			initial:  []FilterQuery{FilterQuery("field1 eq 'value1'")},
			field:    "field2",
			operator: OpNotEqual,
			value:    "value2",
			expected: "field1 eq 'value1' and field2 ne 'value2'",
		},
		{
			name:     "Add filter with special characters",
			initial:  []FilterQuery{},
			field:    "field1",
			operator: OpContains,
			value:    `value'with\special*chars`,
			expected: `field1 contains 'value\'with\\special\*chars'`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &Filter{Filter: tt.initial}
			filter.AddFilter(tt.field, tt.operator, tt.value)

			// Convert []FilterQuery to a string for easier comparison
			result := FilterQueries(filter.Filter).String()

			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestAddOrderBy(t *testing.T) {
	tests := []struct {
		name     string
		initial  OrderBy
		field    string
		expected OrderBy
	}{
		{
			name:     "Add orderby to empty list",
			initial:  OrderBy{},
			field:    "newOrderBy",
			expected: OrderBy{"newOrderBy"},
		},
		{
			name:     "Add orderby to non-empty list",
			initial:  OrderBy{"orderby1", "orderby2"},
			field:    "newOrderBy",
			expected: OrderBy{"orderby1", "orderby2", "newOrderBy"},
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
		initial  OrderBy
		expected OrderBy
	}{
		{
			name:     "Get orderby from empty list",
			initial:  OrderBy{},
			expected: OrderBy{},
		},
		{
			name:     "Get orderby from non-empty list",
			initial:  OrderBy{"orderby1", "orderby2"},
			expected: OrderBy{"orderby1", "orderby2"},
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
		initial  OrderBy
		field    string
		expected OrderBy
	}{
		{
			name:     "Remove orderby from list with one element",
			initial:  OrderBy{"orderby1"},
			field:    "orderby1",
			expected: OrderBy{},
		},
		{
			name:     "Remove orderby from list with multiple elements",
			initial:  OrderBy{"orderby1", "orderby2", "orderby3"},
			field:    "orderby2",
			expected: OrderBy{"orderby1", "orderby3"},
		},
		{
			name:     "Remove orderby that does not exist",
			initial:  OrderBy{"orderby1", "orderby2"},
			field:    "orderby3",
			expected: OrderBy{"orderby1", "orderby2"},
		},
		{
			name:     "Remove orderby from empty list",
			initial:  OrderBy{},
			field:    "orderby1",
			expected: OrderBy{},
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
				Fields: Fields{"field1", "field2"},
			},
			expected: "?fields=field1%2Cfield2&count=false",
		},
		{
			name: "Filter with orderby",
			filter: &Filter{
				Orderby: OrderBy{"field1", "field2"},
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
				Filter: []FilterQuery{FilterQuery(`field1 eq 'value1'`)},
			},
			expected: `?filter=%28field1%20eq%20%27value1%27%29&count=false`,
		},
		{
			name: "Filter with all parameters",
			filter: &Filter{
				Fields:  Fields{"field1", "field2"},
				Orderby: OrderBy{"field1", "field2"},
				Count:   true,
				Filter:  []FilterQuery{FilterQuery(`field1 eq 'value1'`)},
			},
			expected: `?filter=%28field1%20eq%20%27value1%27%29&fields=field1%2Cfield2&count=true&orderby=field1%2Cfield2`,
		},
		{
			name: "Filter with multiple filter queries",
			filter: &Filter{
				Filter: []FilterQuery{
					FilterQuery(`field1 eq 'value1'`),
					FilterQuery(`field2 contains 'value2'`),
				},
			},
			expected: `?filter=%28field1%20eq%20%27value1%27%20and%20field2%20contains%20%27value2%27%29&count=false`,
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
			name:     "Fields with multiple values",
			input:    Fields{"field1", "field2", "field3"},
			expected: "field1,field2,field3",
		},
		{
			name:     "Empty Fields",
			input:    Fields{},
			expected: "",
		},
		{
			name:     "FilterQuery with value",
			input:    FilterQuery("name eq 'John'"),
			expected: "name eq 'John'",
		},
		{
			name:     "Empty FilterQuery",
			input:    FilterQuery(""),
			expected: "",
		},
		{
			name:     "OrderBy with multiple values",
			input:    OrderBy{"name", "age", "date"},
			expected: "name,age,date",
		},
		{
			name:     "Empty OrderBy",
			input:    OrderBy{},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got string
			switch v := tt.input.(type) {
			case Fields:
				got = v.String()
			case FilterQuery:
				got = v.String()
			case OrderBy:
				got = v.String()
			}
			if got != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, got)
			}
		})
	}
}

func TestFilterOperator(t *testing.T) {
	tests := []struct {
		name     string
		operator FilterOperator
		expected string
	}{
		{
			name:     "Equal operator",
			operator: OpEqual,
			expected: "eq",
		},
		{
			name:     "Contains operator",
			operator: OpContains,
			expected: "contains",
		},
		{
			name:     "And operator",
			operator: OpAnd,
			expected: "and",
		},
		{
			name:     "Or operator",
			operator: OpOr,
			expected: "or",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.operator.String()
			if got != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, got)
			}
		})
	}
}

func TestFilterQueriesMethods(t *testing.T) {
	tests := []struct {
		name            string
		queries         FilterQueries
		operator        FilterOperator
		expectedString  string
		expectedGrouped string
	}{
		{
			name:            "Single query",
			queries:         FilterQueries{FilterQuery("name eq 'John'")},
			operator:        OpAnd,
			expectedString:  "name eq 'John'",
			expectedGrouped: "name eq 'John'",
		},
		{
			name:            "Multiple queries with AND",
			queries:         FilterQueries{FilterQuery("name eq 'John'"), FilterQuery("age gt '30'")},
			operator:        OpAnd,
			expectedString:  "name eq 'John' and age gt '30'",
			expectedGrouped: "(name eq 'John' and age gt '30')",
		},
		{
			name:            "Multiple queries with OR",
			queries:         FilterQueries{FilterQuery("name eq 'John'"), FilterQuery("name eq 'Jane'")},
			operator:        OpOr,
			expectedString:  "name eq 'John' or name eq 'Jane'",
			expectedGrouped: "(name eq 'John' or name eq 'Jane')",
		},
		{
			name:            "Empty queries",
			queries:         FilterQueries{},
			operator:        OpAnd,
			expectedString:  "",
			expectedGrouped: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test StringWithOperator method
			gotString := tt.queries.StringWithOperator(tt.operator)
			if gotString != tt.expectedString {
				t.Errorf("StringWithOperator: expected %q, got %q", tt.expectedString, gotString)
			}

			// Test GroupedWithOperator method
			gotGrouped := tt.queries.GroupedWithOperator(tt.operator)
			if gotGrouped != tt.expectedGrouped {
				t.Errorf("GroupedWithOperator: expected %q, got %q", tt.expectedGrouped, gotGrouped)
			}

			// Test String method (uses OpAnd by default)
			if tt.operator == OpAnd {
				gotDefault := tt.queries.String()
				if gotDefault != tt.expectedString {
					t.Errorf("String: expected %q, got %q", tt.expectedString, gotDefault)
				}
			}
		})
	}
}

func TestAddComplexSearchFilter(t *testing.T) {
	tests := []struct {
		name         string
		searchValue  string
		fields       map[string]FilterOperator
		contentCheck bool     // Whether to check content components rather than exact match
		expected     string   // Used when contentCheck is false
		contains     []string // Components that should be present when contentCheck is true
	}{
		{
			name:        "Search across two fields",
			searchValue: "test",
			fields: map[string]FilterOperator{
				"Name":        OpContains,
				"Description": OpContains,
			},
			contentCheck: true,
			contains: []string{
				"Name contains 'test'",
				"Description contains 'test'",
				" or ",
				// Check for opening and closing parentheses
				"(",
				")",
			},
		},
		{
			name:        "Search with special characters",
			searchValue: "O'Reilly",
			fields: map[string]FilterOperator{
				"Author": OpEqual,
				"Title":  OpContains,
			},
			contentCheck: true,
			contains: []string{
				"Author eq 'O\\'Reilly'",
				"Title contains 'O\\'Reilly'",
				" or ",
				"(",
				")",
			},
		},
		{
			name:         "Search with empty fields",
			searchValue:  "test",
			fields:       map[string]FilterOperator{},
			contentCheck: false,
			expected:     "",
		},
		{
			name:        "Search with different operators",
			searchValue: "test",
			fields: map[string]FilterOperator{
				"Name": OpStartsWith,
				"Tags": OpIContains,
			},
			contentCheck: true,
			contains: []string{
				"Name sw 'test'",
				"Tags icontains 'test'",
				" or ",
				"(",
				")",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &Filter{}
			filter.AddComplexSearchFilter(tt.searchValue, tt.fields)

			var result string
			if len(filter.Filter) > 0 {
				result = filter.Filter[0].String()
			}

			if tt.contentCheck {
				// Check that result begins with "(" and ends with ")"
				if len(result) > 0 && (!strings.HasPrefix(result, "(") || !strings.HasSuffix(result, ")")) {
					t.Errorf("expected result to be wrapped in parentheses, got: %q", result)
				}

				// Check that all expected components are present
				for _, component := range tt.contains {
					if !strings.Contains(result, component) {
						t.Errorf("expected result to contain %q, but it didn't. Got: %q", component, result)
					}
				}

				// Check that we have the right number of OR operators
				if len(tt.fields) > 1 {
					orCount := strings.Count(result, " or ")
					expectedOrCount := len(tt.fields) - 1
					if orCount != expectedOrCount {
						t.Errorf("expected %d 'or' operators, got %d. Result: %q",
							expectedOrCount, orCount, result)
					}
				}
			} else {
				if result != tt.expected {
					t.Errorf("expected %q, got %q", tt.expected, result)
				}
			}
		})
	}
}

func TestAddSearchFilter(t *testing.T) {
	filter := &Filter{}
	filter.AddSearchFilter("muster")

	if len(filter.Filter) != 1 {
		t.Fatalf("expected 1 filter, got %d", len(filter.Filter))
	}

	result := filter.Filter[0].String()

	// Check that all expected fields are present
	expectedFields := []string{
		"Name contains 'muster'",
		"DomainName contains 'muster'",
		"AccountNamespace contains 'muster'",
		"Asset.Name contains 'muster'",
		"Description contains 'muster'",
		"Tags.Name icontains 'muster'",
	}

	for _, expected := range expectedFields {
		if !strings.Contains(result, expected) {
			t.Errorf("expected result to contain %q, but it didn't", expected)
		}
	}

	// Check that it's a proper grouped query with parentheses and OR operators
	if !strings.HasPrefix(result, "(") || !strings.HasSuffix(result, ")") {
		t.Errorf("expected result to be wrapped in parentheses")
	}

	orCount := strings.Count(result, " or ")
	if orCount < 8 { // At least 8 "or"s for 9 fields
		t.Errorf("expected at least 8 OR operators, got %d", orCount)
	}
}

func TestEscapeSpecialChars(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "String with no special chars",
			input:    "normal string",
			expected: "normal string",
		},
		{
			name:     "String with single quote",
			input:    "O'Reilly",
			expected: "O\\'Reilly",
		},
		{
			name:     "String with backslash",
			input:    "C:\\Files\\Data",
			expected: "C:\\\\Files\\\\Data",
		},
		{
			name:     "String with asterisk",
			input:    "wild*card",
			expected: "wild\\*card",
		},
		{
			name:     "String with mixed special chars",
			input:    "The 'wild*\\' string",
			expected: "The \\'wild\\*\\\\\\' string",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := escapeSpecialChars(tt.input)
			if got != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, got)
			}
		})
	}
}

func TestGenerateFilterQueryParentheses(t *testing.T) {
	tests := []struct {
		name         string
		filter       *Filter
		expected     string
		contentCheck bool     // Set to true to check for content rather than exact match
		contains     []string // Strings that should be present when doing content check
	}{
		{
			name: "Complex filter with already wrapped condition",
			filter: &Filter{
				Filter: []FilterQuery{FilterQuery("(Name contains 'test' or Description contains 'test')")},
			},
			expected: `?filter=%28Name%20contains%20%27test%27%20or%20Description%20contains%20%27test%27%29&count=false`,
		},
		{
			name: "Filter with unwrapped condition",
			filter: &Filter{
				Filter: []FilterQuery{FilterQuery("Name contains 'test'")},
			},
			expected: `?filter=%28Name%20contains%20%27test%27%29&count=false`,
		},
		{
			name: "Filter with multiple conditions",
			filter: &Filter{
				Filter: []FilterQuery{
					FilterQuery("Name contains 'test'"),
					FilterQuery("Active eq 'true'"),
				},
			},
			expected: `?filter=%28Name%20contains%20%27test%27%20and%20Active%20eq%20%27true%27%29&count=false`,
		},
		{
			name: "Complex search with AddComplexSearchFilter",
			filter: func() *Filter {
				f := &Filter{}
				fields := map[string]FilterOperator{
					"Name":        OpContains,
					"Description": OpContains,
				}
				f.AddComplexSearchFilter("test", fields)
				return f
			}(),
			// Map iteration order is not guaranteed, so we check for content instead of exact match
			contentCheck: true,
			contains: []string{
				"filter=",
				"Name%20contains%20%27test%27",
				"Description%20contains%20%27test%27",
				"%28", // Opening parenthesis
				"%29", // Closing parenthesis
				"or",
				"count=false",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.filter.ToQueryString()

			if tt.contentCheck {
				// Check that all expected parts are present
				for _, part := range tt.contains {
					if !strings.Contains(got, part) {
						t.Errorf("expected result to contain %q, but it didn't.\nGot: %q", part, got)
					}
				}
			} else {
				if got != tt.expected {
					t.Errorf("expected %q, got %q", tt.expected, got)
				}
			}
		})
	}
}

func TestIsAlreadyWrapped(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Simple wrapped expression",
			input:    "(Name contains 'test')",
			expected: true,
		},
		{
			name:     "Unwrapped expression",
			input:    "Name contains 'test'",
			expected: false,
		},
		{
			name:     "Complex wrapped expression",
			input:    "(Name contains 'test' or Description contains 'test')",
			expected: true,
		},
		{
			name:     "Nested parentheses with proper outer wrapping",
			input:    "((Name contains 'test') or Description contains 'test')",
			expected: true, // Changed from false to true - outer parentheses do form a single enclosing group
		},
		{
			name:     "Unbalanced opening",
			input:    "(Name contains 'test'",
			expected: false,
		},
		{
			name:     "Unbalanced closing",
			input:    "Name contains 'test')",
			expected: false,
		},
		{
			name:     "With whitespace",
			input:    "  (Name contains 'test')  ",
			expected: true,
		},
		{
			name:     "Multiple expressions each wrapped but not overall",
			input:    "(Name eq 'John') and (Age gt '30')",
			expected: false,
		},
		{
			name:     "Complex nested expression properly wrapped",
			input:    "((Name eq 'John') and (Age gt '30'))",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isAlreadyWrapped(tt.input)
			if got != tt.expected {
				t.Errorf("isAlreadyWrapped(%q): expected %v, got %v", tt.input, tt.expected, got)
			}
		})
	}
}
