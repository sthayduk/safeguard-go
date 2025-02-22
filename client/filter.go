package client

import (
	"net/url"
	"strings"
)

type Filter struct {
	Fields  []string `json:"fields,omitempty"`
	Filter  string   `json:"filter,omitempty"`
	Orderby []string `json:"orderby,omitempty"`
	Count   bool     `json:"count,omitempty"`
}

// AddField adds a field to the list of fields to be included in the query.
// Parameters:
//   - field: The field to be added.
func (f *Filter) AddField(field string) {
	f.Fields = append(f.Fields, field)
}

// RemoveField removes a field from the list of fields to be included in the query.
// Parameters:
//   - field: The field to be removed.
func (f *Filter) RemoveField(field string) {
	for i, v := range f.Fields {
		if v == field {
			f.Fields = append(f.Fields[:i], f.Fields[i+1:]...)
			break
		}
	}
}

// GetFields returns the list of fields to be included in the query.
func (f *Filter) GetFields() []string {
	return f.Fields
}

// AddOrderBy adds a field to the list of fields to order the results by.
// Parameters:
//   - field: The field to be added to the order by list.
func (f *Filter) AddOrderBy(field string) {
	f.Orderby = append(f.Orderby, field)
}

// RemoveOrderBy removes a field from the list of fields to order the results by.
// Parameters:
//   - field: The field to be removed from the order by list.
func (f *Filter) RemoveOrderBy(field string) {
	for i, v := range f.Orderby {
		if v == field {
			f.Orderby = append(f.Orderby[:i], f.Orderby[i+1:]...)
			break
		}
	}
}

// GetOrderBy returns the list of fields to order the results by.
func (f *Filter) GetOrderBy() []string {
	return f.Orderby
}

// ToQueryString generates the query string based on the filter's fields, filter conditions, order by fields, and count flag.
// Returns:
//   - The generated query string.
func (f *Filter) ToQueryString() string {
	var queryParams []string
	if filterQuery := f.generateFilterQuery(); filterQuery != "" {
		queryParams = append(queryParams, filterQuery)
	}
	if fieldQuery := f.generateFieldQuery(); fieldQuery != "" {
		queryParams = append(queryParams, fieldQuery)
	}
	if countQuery := f.generateCountQuery(); countQuery != "" {
		queryParams = append(queryParams, countQuery)
	}
	if orderByQuery := f.generateOrderByQuery(); orderByQuery != "" {
		queryParams = append(queryParams, orderByQuery)
	}

	return "?" + strings.Join(queryParams, "&")
}

// generateFieldQuery generates the query string for the fields to be included in the query.
// Returns:
//   - The generated field query string.
func (f *Filter) generateFieldQuery() string {
	if len(f.Fields) == 0 {
		return ""
	}

	var builder strings.Builder
	builder.Grow(len(f.Fields) * 10) // Preallocate capacity based on the number of order by fields
	for i, field := range f.Fields {
		builder.WriteString(field)
		if i < len(f.Fields)-1 {
			builder.WriteString(",")
		}
	}

	query := "fields=" + url.PathEscape(builder.String())
	return query
}

// generateOrderByQuery generates the query string for the order by fields.
// Returns:
//   - The generated order by query string.
func (f *Filter) generateOrderByQuery() string {
	if len(f.Orderby) == 0 {
		return ""
	}

	var builder strings.Builder
	builder.Grow(len(f.Orderby) * 10) // Preallocate capacity based on the number of order by fields
	for i, orderby := range f.Orderby {
		builder.WriteString(orderby)
		if i < len(f.Orderby)-1 {
			builder.WriteString(",")
		}
	}

	query := "orderby=" + url.PathEscape(builder.String())
	return query
}

// generateFilterQuery generates the query string for the filter conditions.
// Returns:
//   - The generated filter query string.
func (f *Filter) generateFilterQuery() string {
	if f.Filter == "" {
		return ""
	}

	query := "filter=" + url.PathEscape(f.Filter)
	return query
}

// generateCountQuery generates the query string for the count flag.
// Returns:
//   - The generated count query string.
func (f *Filter) generateCountQuery() string {
	if !f.Count {
		return "count=false"
	}

	return "count=true"
}

// AddFilter adds a new filter condition to the Filter.
// The condition is specified by the field, operator, and value parameters.
// The value is escaped to handle special characters.
// Parameters:
//   - field: The field to filter on.
//   - operator: The operator to use for the filter condition (e.g., 'eq', 'ne', 'gt', 'ge', 'lt', 'le', 'and', 'or', 'not', 'contains', 'ieq', 'icontains', 'sw', 'isw', 'ew', 'iew', 'in').
//   - value: The value to compare the field against.
func (f *Filter) AddFilter(field, operator, value string) {
	escapedValue := escapeSpecialChars(value)
	f.Filter += field + " " + operator + " " + escapedValue
}

// escapeSpecialChars escapes special characters in the filter value.
// Parameters:
//   - value: The value to escape.
//
// Returns:
//   - The escaped value.
func escapeSpecialChars(value string) string {
	escapedChars := map[rune]string{
		'\\': "\\\\",
		'"':  "\\\"",
		'*':  "\\*",
	}

	escaped := ""
	for _, char := range value {
		if esc, ok := escapedChars[char]; ok {
			escaped += esc
		} else {
			escaped += string(char)
		}
	}
	return escaped
}
