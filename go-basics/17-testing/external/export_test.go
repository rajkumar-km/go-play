package search

// As part of white box testing, we need to test the isSorted() functionality defined in the
// production package "search".
// - Assume we supposed to write the tests in an external package "search_test" to avoid some
//   import cycle.
// - In such case, we can export a internal member like this in a "*_test.go" so that it is
//   exported only to the test environment. Typically the file is named "export_test.go"
var IsSorted = isSorted