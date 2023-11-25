/*
Package example demonstrates writing example functions as part of Go test files.

Go test files can contain functions starting with "Example" keyword and it servers three purposes.

 1. Example functions serves as Go documentation: An example can show a working application by
    assembling different package APIs. This can be more easy to understand than the text
    description about the package. godoc scans the Example functions and shows as example under
    the corresponding library function.
 2. Example functions are executable tests and its standard output can be verified against the one
    mentioned in the final comment after "// Output: ".
 3. Hands on experimentation is easy with example functions. Go playgound automatically shows as
    editable code for experimentation.
*/
package example

func BubbleSort(v []int) {
	for i:=1; i < len(v); i++ {
		for j:=0; j < len(v)-i; j++ {
			if v[j] > v[j+1] {
				v[j],v[j+1] = v[j+1],v[j]
			}			
		}
	}
}