// Package dsn implements dsn parse with struct bind
/*
The DSN format is similar to a URI, and the DSN structure is as follows.

Network:[//[username[:password]@]address[:port][,address[:port]]][/path][?query][#fragment]

The main difference from URI is that scheme is replaced by network, host is replaced with address and supports multiple addresses.
Network has the same meaning as network in net package, tcp, udp, unix, etc., address supports multiple use ',' split, if
Network is a local sock protocol such as unix, using Path, there is only one

The dsn package mainly provides Parse, Bind and validate functions.

Parse parses the dsn string into a DSN struct, and the DSN struct is almost identical to the url.URL

Bind provides the ability to bind DSN data to a struct, specifying the bound fields via tag dsn:"key,[default]", currently supports two types of data binding

Built-in variable key:
Network string tcp, udp, unix, etc., refer to the network in the net package
Username string
Password string
Address string or []string address can be bound to string or []string, if string then take address first

Query: The data on the query can be retrieved by query.name

Arrays can be obtained by passing multiple

Array=1&array=2&array3 -> []int `tag:"query.array"`

Struct support nesting

Foo.sub.name=hello&foo.tm=hello

Struct Foo {
Tm string `dsn:"query.tm"`
Sub struct {
Name string `dsn:"query.name"`
} `dsn:"query.sub"`
}

Default: By dsn:"key,[default]" The default value does not support arrays for the time being.

Ignore Bind: Ignore Bind by dsn:"-"

Custom Bind: can implement encoding.TextUnmarshaler custom Bind implementation at the same time

Validate: Reference https://github.com/go-playground/validator

Use reference: example_test.go

DSN naming convention:

In the absence of historical legacy, try to use the names of Address, Network, Username, Password, etc. instead of the previous names such as Proto and Addr.

Query naming reference, starting with a hump lowercase:

Timeout universal timeout
dialTimeout connection establishment timeout
readTimeout read operation timed out
writeTimeout write operation timed out
readsTimeout batch read timeout
writesTimeout batch write timeout
*/
package dsn
