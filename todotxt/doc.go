// todotxt package implements a Todo type compatible with todo.txt format.
// It aims to implement [encoding.TextMarshaler] and
// [encoding.TextUnmarshaler], in which the prior should be based in an
// implementation of [fmt.Stringer]
//
// Reference: https://github.com/todotxt/todo.txt
package todotxt
