package safeguard

import (
	"net/url"
	"strings"
)

// Fields represents a list of fields to be included in the query
type Fields []string

func (f Fields) String() string {
	return strings.Join(f, ",")
}

func (f Fields) ToQueryString() string {
	return "?" + f.generateFieldQuery()
}

func (f Fields) generateFieldQuery() string {
	if len(f) == 0 {
		return ""
	}
	return "fields=" + url.PathEscape(f.String())
}

// FilterQuery represents the filter condition
type FilterQuery string

func (f FilterQuery) String() string {
	return string(f)
}

// OrderBy represents a list of fields to order the results by
type OrderBy []string

func (o OrderBy) String() string {
	return strings.Join(o, ",")
}

// Filter Example from PAM UI
// TODO: Add support for:
// (Name contains 'muster' or DomainName contains 'muster' or AccountNamespace contains 'muster' or Asset.Name contains 'muster' or PasswordProfile.EffectiveName contains 'muster' or SshKeyProfile.EffectiveName contains 'muster' or Description contains 'muster' or Tags.Name icontains 'muster' or PrivilegeGroupMembership contains 'muster')

type Filter struct {
	Fields  Fields      `json:"fields,omitempty"`
	Filter  FilterQuery `json:"filter,omitempty"`
	Orderby OrderBy     `json:"orderby,omitempty"`
	Count   bool        `json:"count,omitempty"`
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
func (f *Filter) GetFields() Fields {
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
func (f *Filter) GetOrderBy() OrderBy {
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
	return "fields=" + url.PathEscape(f.Fields.String())
}

// generateOrderByQuery generates the query string for the order by fields.
// Returns:
//   - The generated order by query string.
func (f *Filter) generateOrderByQuery() string {
	if len(f.Orderby) == 0 {
		return ""
	}
	return "orderby=" + url.PathEscape(f.Orderby.String())
}

// generateFilterQuery generates the query string for the filter conditions.
// Returns:
//   - The generated filter query string.
func (f *Filter) generateFilterQuery() string {
	if f.Filter == "" {
		return ""
	}
	return "filter=" + url.PathEscape(f.Filter.String())
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
	f.Filter = FilterQuery(string(f.Filter) + field + " " + operator + " \"" + escapedValue + "\"")
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
