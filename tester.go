// This demonstrates godoc.org issue #657.
//
// 	https://github.com/golang/gddo/issues/657
//
// Formatting of the Do function looks like this
//
// 	func Do( a int, b string, )
//
// where it should rather be
//
// 	func Do(a int, b string)
//
// This bug shows when a function is formatted with its parameters on extra
// lines like this:
//
// 	func Do(
// 		a int,
// 		b string,
// 	) {}
package tester

func Do(
	a int,
	b string,
) {
}
