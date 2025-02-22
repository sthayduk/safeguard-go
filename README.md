# Safeguard Go

Safeguard Go is a Go library designed to provide filtering, ordering, and querying capabilities for your Go applications.

## Installation

To install the library, use the following command:

```sh
go get github.com/sthayduk/safeguard-go
```

## Usage

### Creating a Filter

To create a filter, initialize a new `Filter` struct:

```go
import "github.com/sthayduk/safeguard-go/client"

filter := &client.Filter{}
```

### Adding Fields

To add fields to the filter:

```go
filter.AddField("fieldName")
```

### Removing Fields

To remove fields from the filter:

```go
filter.RemoveField("fieldName")
```

### Getting Fields

To get the list of fields:

```go
fields := filter.GetFields()
```

### Adding Filters

To add a filter condition:

```go
filter.AddFilter("fieldName", "operator", "value")
```

### Adding Order By

To add an order by condition:

```go
filter.AddOrderBy("fieldName")
```

### Getting Order By

To get the list of order by conditions:

```go
orderBy := filter.GetOrderBy()
```

### Removing Order By

To remove an order by condition:

```go
filter.RemoveOrderBy("fieldName")
```

### Converting to Query String

To convert the filter to a query string:

```go
queryString := filter.ToQueryString()
```

## Running Tests

To run the tests, use the following command:

```sh
go test ./...
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License.
