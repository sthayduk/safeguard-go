// Package safeguard provides functionality for building and manipulating API filters.
// It allows for creating complex query strings with various filtering operations,
// field selections, and ordering capabilities.
package safeguard

import (
	"net/url"
	"strings"
)

// FilterOperator represents the operator used in filter conditions.
// These operators define how values in filter expressions should be compared.
type FilterOperator string

// String returns the string representation of the filter operator.
func (o FilterOperator) String() string {
	return string(o)
}

// Standard filter operators supported by the API.
const (
	OpEqual              FilterOperator = "eq"        // Equal
	OpNotEqual           FilterOperator = "ne"        // Not equal
	OpGreaterThan        FilterOperator = "gt"        // Greater than
	OpGreaterThanOrEqual FilterOperator = "ge"        // Greater than or equal
	OpLessThan           FilterOperator = "lt"        // Less than
	OpLessThanOrEqual    FilterOperator = "le"        // Less than or equal
	OpAnd                FilterOperator = "and"       // Logical AND
	OpOr                 FilterOperator = "or"        // Logical OR
	OpNot                FilterOperator = "not"       // Logical NOT
	OpContains           FilterOperator = "contains"  // Contains substring (case sensitive)
	OpIEqual             FilterOperator = "ieq"       // Case-insensitive equals
	OpIContains          FilterOperator = "icontains" // Case-insensitive contains
	OpStartsWith         FilterOperator = "sw"        // Starts with (case sensitive)
	OpIStartsWith        FilterOperator = "isw"       // Case-insensitive starts with
	OpEndsWith           FilterOperator = "ew"        // Ends with (case sensitive)
	OpIEndsWith          FilterOperator = "iew"       // Case-insensitive ends with
	OpIn                 FilterOperator = "in"        // Value is in a set
)

// Fields represents a list of fields to be included in the query results.
// This allows API consumers to specify which fields they want returned.
type Fields []string

// String returns a comma-separated list of field names.
func (f Fields) String() string {
	return strings.Join(f, ",")
}

// ToQueryString converts the Fields to a URL query string parameter.
// Returns a string in the format "?fields=field1,field2,field3".
func (f Fields) ToQueryString() string {
	return "?" + f.generateFieldQuery()
}

// generateFieldQuery builds the fields portion of the query string.
// Returns an empty string if no fields are specified.
func (f Fields) generateFieldQuery() string {
	if len(f) == 0 {
		return ""
	}
	return "fields=" + url.PathEscape(f.String())
}

// FilterQuery represents a single filter condition or a group of conditions.
// Example: "name eq 'value'" or "(field1 eq 'value1' and field2 eq 'value2')".
type FilterQuery string

// String returns the string representation of the filter query.
func (f FilterQuery) String() string {
	return string(f)
}

// FilterQueries represents a collection of FilterQuery objects.
// Multiple queries are typically combined with a logical operator like AND or OR.
type FilterQueries []FilterQuery

// String returns the string representation of filter queries joined by 'and'.
// This is a convenience method that uses StringWithOperator with the AND operator.
func (fq FilterQueries) String() string {
	return fq.StringWithOperator(OpAnd)
}

// StringWithOperator returns the string representation of filter queries joined by the specified operator.
// Parameters:
//   - operator: The operator (typically AND or OR) to join the filter queries.
//
// Returns:
//   - A string with all filter queries joined by the specified operator.
func (fq FilterQueries) StringWithOperator(operator FilterOperator) string {
	if len(fq) == 0 {
		return ""
	}
	if len(fq) == 1 {
		return fq[0].String()
	}

	queries := make([]string, len(fq))
	for i, q := range fq {
		queries[i] = q.String()
	}
	return strings.Join(queries, " "+operator.String()+" ")
}

// GroupedWithOperator returns the filter queries as a grouped expression with parentheses
// using the specified operator (and, or) between conditions.
// Parameters:
//   - operator: The operator to use for joining the conditions.
//
// Returns:
//   - A string with parentheses around the joined conditions if there are multiple conditions.
func (fq FilterQueries) GroupedWithOperator(operator FilterOperator) string {
	if len(fq) <= 1 {
		return fq.StringWithOperator(operator)
	}
	return "(" + fq.StringWithOperator(operator) + ")"
}

// OrderBy represents a list of fields to order the results by.
// The order of fields in the slice determines their precedence in sorting.
type OrderBy []string

// String returns a comma-separated list of fields to order by.
func (o OrderBy) String() string {
	return strings.Join(o, ",")
}

// Filter represents a complete set of query parameters for filtering API results.
// It combines field selection, filtering conditions, ordering, and count options.
type Filter struct {
	Fields  Fields        `json:"fields,omitempty"`  // Fields to include in the response
	Filter  []FilterQuery `json:"filter,omitempty"`  // Filter conditions to apply
	Orderby OrderBy       `json:"orderby,omitempty"` // Fields to order the results by
	Count   bool          `json:"count,omitempty"`   // Whether to include a count of total results
}

// AddField adds a field to the list of fields to be included in the query.
// Parameters:
//   - field: The field name to be added.
func (f *Filter) AddField(field string) {
	f.Fields = append(f.Fields, field)
}

// RemoveField removes a field from the list of fields to be included in the query.
// If the field doesn't exist in the list, no change is made.
// Parameters:
//   - field: The field name to be removed.
func (f *Filter) RemoveField(field string) {
	for i, v := range f.Fields {
		if v == field {
			f.Fields = append(f.Fields[:i], f.Fields[i+1:]...)
			break
		}
	}
}

// GetFields returns the list of fields to be included in the query response.
// Returns:
//   - A copy of the Fields slice.
func (f *Filter) GetFields() Fields {
	return f.Fields
}

// AddOrderBy adds a field to the list of fields to order the results by.
// Parameters:
//   - field: The field name to add to the order by list.
func (f *Filter) AddOrderBy(field string) {
	f.Orderby = append(f.Orderby, field)
}

// RemoveOrderBy removes a field from the list of fields to order the results by.
// If the field doesn't exist in the list, no change is made.
// Parameters:
//   - field: The field name to be removed from the order by list.
func (f *Filter) RemoveOrderBy(field string) {
	for i, v := range f.Orderby {
		if v == field {
			f.Orderby = append(f.Orderby[:i], f.Orderby[i+1:]...)
			break
		}
	}
}

// GetOrderBy returns the list of fields used to order the results.
// Returns:
//   - A copy of the OrderBy slice.
func (f *Filter) GetOrderBy() OrderBy {
	return f.Orderby
}

// ToQueryString generates a complete URL query string based on all filter parameters.
// The query string includes fields, filter conditions, ordering, and count options.
// Returns:
//   - The fully formatted query string starting with "?".
func (f *Filter) ToQueryString() string {
	// Pre-allocate slice with capacity of 4 (max number of possible parameters)
	queryParams := make([]string, 0, 4)

	// Directly append the results of generate functions if they're not empty
	if filter := f.generateFilterQuery(); filter != "" {
		queryParams = append(queryParams, filter)
	}
	if fields := f.generateFieldQuery(); fields != "" {
		queryParams = append(queryParams, fields)
	}

	// Count parameter is always included based on the test cases
	queryParams = append(queryParams, f.generateCountQuery())

	if orderBy := f.generateOrderByQuery(); orderBy != "" {
		queryParams = append(queryParams, orderBy)
	}

	return "?" + strings.Join(queryParams, "&")
}

// generateFieldQuery builds the fields portion of the query string.
// Returns an empty string if no fields are specified.
// Returns:
//   - The fields query parameter string or empty string if no fields are specified.
func (f *Filter) generateFieldQuery() string {
	if len(f.Fields) == 0 {
		return ""
	}
	return "fields=" + url.PathEscape(f.Fields.String())
}

// generateOrderByQuery builds the orderby portion of the query string.
// Returns an empty string if no order by fields are specified.
// Returns:
//   - The orderby query parameter string or empty string if no orderby fields are specified.
func (f *Filter) generateOrderByQuery() string {
	if len(f.Orderby) == 0 {
		return ""
	}
	return "orderby=" + url.PathEscape(f.Orderby.String())
}

// generateFilterQuery builds the filter portion of the query string.
// Returns an empty string if no filter conditions are specified.
// Filter expressions with multiple conditions are wrapped in parentheses.
// Returns:
//   - The filter query parameter string or empty string if no filter conditions are specified.
func (f *Filter) generateFilterQuery() string {
	// return Empty String if no filters are set
	if len(f.Filter) == 0 {
		return ""
	}

	filterExpr := FilterQueries(f.Filter).String()

	// If we have multiple conditions, wrap in parentheses
	// For single conditions, only add parentheses if not already wrapped
	if len(f.Filter) > 1 {
		filterExpr = "(" + filterExpr + ")"
	} else if len(f.Filter) == 1 && !isAlreadyWrapped(filterExpr) {
		filterExpr = "(" + filterExpr + ")"
	}

	return "filter=" + url.PathEscape(filterExpr)
}

// isAlreadyWrapped checks if a string is already wrapped in properly balanced parentheses
// that enclose the entire expression as a single group.
// Parameters:
//   - s: The string to check.
//
// Returns:
//   - true if the string is properly wrapped in balanced parentheses, false otherwise.
func isAlreadyWrapped(s string) bool {
	s = strings.TrimSpace(s)
	if !strings.HasPrefix(s, "(") || !strings.HasSuffix(s, ")") {
		return false
	}

	// Check if the first opening parenthesis matches the last closing one
	count := 0
	for i, r := range s {
		if r == '(' {
			count++
		} else if r == ')' {
			count--
		}

		// If we reach zero before the end, the outer parentheses
		// don't form a single group around the whole expression
		if count == 0 && i < len(s)-1 {
			return false
		}

		// Unbalanced - more closing than opening
		if count < 0 {
			return false
		}
	}

	// If count is 0, all parentheses are balanced
	return count == 0
}

// generateCountQuery builds the count portion of the query string.
// Returns:
//   - The count query parameter string (either "count=true" or "count=false").
func (f *Filter) generateCountQuery() string {
	if !f.Count {
		return "count=false"
	}
	return "count=true"
}

// AddFilter adds a new filter condition to the Filter.
// The condition is created using the specified field, operator, and value.
// Special characters in the value are escaped automatically.
// Parameters:
//   - field: The field name to filter on.
//   - operator: The operator to use for the filter condition.
//   - value: The value to compare against (will be escaped for special characters).
func (f *Filter) AddFilter(field string, operator FilterOperator, value string) {
	escapedValue := escapeSpecialChars(value)
	f.Filter = append(f.Filter, FilterQuery(field+" "+operator.String()+" '"+escapedValue+"'"))
}

// AddComplexSearchFilter adds a search filter across multiple fields with OR conditions.
// This creates a grouped filter like: (field1 op 'value' or field2 op 'value' or field3 op 'value').
// Parameters:
//   - value: The value to search for across all specified fields.
//   - fields: A map where keys are field names and values are the operators to use.
func (f *Filter) AddComplexSearchFilter(value string, fields map[string]FilterOperator) {
	if len(fields) == 0 {
		return
	}

	escapedValue := escapeSpecialChars(value)
	conditions := make(FilterQueries, 0, len(fields))

	for field, operator := range fields {
		conditions = append(conditions, FilterQuery(field+" "+operator.String()+" '"+escapedValue+"'"))
	}

	// Convert to a single filter query with OR operators and parentheses
	complexFilter := FilterQuery(conditions.GroupedWithOperator(OpOr))
	f.Filter = append(f.Filter, complexFilter)
}

// escapeSpecialChars escapes special characters that need to be encoded in filter values.
// Characters escaped: backslash (\), single quote ('), and asterisk (*).
// Parameters:
//   - value: The string value to escape.
//
// Returns:
//   - The string with special characters escaped.
func escapeSpecialChars(value string) string {
	escapedChars := map[rune]string{
		'\\': "\\\\",
		'\'': "\\'",
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

// AddSearchFilter adds a predefined search across common searchable fields.
// This implements a standard search pattern matching the PAM UI functionality.
// Parameters:
//   - searchTerm: The term to search for across multiple predefined fields.
func (f *Filter) AddSearchFilter(searchTerm string) {
	searchFields := map[string]FilterOperator{
		"Name":                          OpContains,
		"DomainName":                    OpContains,
		"AccountNamespace":              OpContains,
		"Asset.Name":                    OpContains,
		"PasswordProfile.EffectiveName": OpContains,
		"SshKeyProfile.EffectiveName":   OpContains,
		"Description":                   OpContains,
		"Tags.Name":                     OpIContains,
		"PrivilegeGroupMembership":      OpContains,
	}

	f.AddComplexSearchFilter(searchTerm, searchFields)
}
