// Generated by gen_types. Don't modify manually!

package core

import (
	"fmt"
	"io"
)

func EnsureObjectIsComparable(obj Object, pattern string) Comparable {
	switch c := obj.(type) {
	case Comparable:
		return c
	default:
		if pattern == "" {
			pattern = "%s"
		}
		msg := fmt.Sprintf("Expected %s, got %s", "Comparable", obj.GetType().ToString(false))
		panic(RT.NewError(fmt.Sprintf(pattern, msg)))
	}
}

func EnsureArgIsComparable(args []Object, index int) Comparable {
	switch c := args[index].(type) {
	case Comparable:
		return c
	default:
		panic(RT.NewArgTypeError(index, c, "Comparable"))
	}
}

func EnsureObjectIsVector(obj Object, pattern string) *Vector {
	switch c := obj.(type) {
	case *Vector:
		return c
	default:
		if pattern == "" {
			pattern = "%s"
		}
		msg := fmt.Sprintf("Expected %s, got %s", "Vector", obj.GetType().ToString(false))
		panic(RT.NewError(fmt.Sprintf(pattern, msg)))
	}
}

func EnsureArgIsVector(args []Object, index int) *Vector {
	switch c := args[index].(type) {
	case *Vector:
		return c
	default:
		panic(RT.NewArgTypeError(index, c, "Vector"))
	}
}

func EnsureObjectIsChar(obj Object, pattern string) Char {
	switch c := obj.(type) {
	case Char:
		return c
	default:
		if pattern == "" {
			pattern = "%s"
		}
		msg := fmt.Sprintf("Expected %s, got %s", "Char", obj.GetType().ToString(false))
		panic(RT.NewError(fmt.Sprintf(pattern, msg)))
	}
}

func EnsureArgIsChar(args []Object, index int) Char {
	switch c := args[index].(type) {
	case Char:
		return c
	default:
		panic(RT.NewArgTypeError(index, c, "Char"))
	}
}

func EnsureObjectIsString(obj Object, pattern string) String {
	switch c := obj.(type) {
	case String:
		return c
	default:
		if pattern == "" {
			pattern = "%s"
		}
		msg := fmt.Sprintf("Expected %s, got %s", "String", obj.GetType().ToString(false))
		panic(RT.NewError(fmt.Sprintf(pattern, msg)))
	}
}

func EnsureArgIsString(args []Object, index int) String {
	switch c := args[index].(type) {
	case String:
		return c
	default:
		panic(RT.NewArgTypeError(index, c, "String"))
	}
}

func EnsureObjectIsSymbol(obj Object, pattern string) Symbol {
	switch c := obj.(type) {
	case Symbol:
		return c
	default:
		if pattern == "" {
			pattern = "%s"
		}
		msg := fmt.Sprintf("Expected %s, got %s", "Symbol", obj.GetType().ToString(false))
		panic(RT.NewError(fmt.Sprintf(pattern, msg)))
	}
}

func EnsureArgIsSymbol(args []Object, index int) Symbol {
	switch c := args[index].(type) {
	case Symbol:
		return c
	default:
		panic(RT.NewArgTypeError(index, c, "Symbol"))
	}
}

func EnsureObjectIsKeyword(obj Object, pattern string) Keyword {
	switch c := obj.(type) {
	case Keyword:
		return c
	default:
		if pattern == "" {
			pattern = "%s"
		}
		msg := fmt.Sprintf("Expected %s, got %s", "Keyword", obj.GetType().ToString(false))
		panic(RT.NewError(fmt.Sprintf(pattern, msg)))
	}
}

func EnsureArgIsKeyword(args []Object, index int) Keyword {
	switch c := args[index].(type) {
	case Keyword:
		return c
	default:
		panic(RT.NewArgTypeError(index, c, "Keyword"))
	}
}

func EnsureObjectIsRegex(obj Object, pattern string) *Regex {
	switch c := obj.(type) {
	case *Regex:
		return c
	default:
		if pattern == "" {
			pattern = "%s"
		}
		msg := fmt.Sprintf("Expected %s, got %s", "Regex", obj.GetType().ToString(false))
		panic(RT.NewError(fmt.Sprintf(pattern, msg)))
	}
}

func EnsureArgIsRegex(args []Object, index int) *Regex {
	switch c := args[index].(type) {
	case *Regex:
		return c
	default:
		panic(RT.NewArgTypeError(index, c, "Regex"))
	}
}

func EnsureObjectIsBoolean(obj Object, pattern string) Boolean {
	switch c := obj.(type) {
	case Boolean:
		return c
	default:
		if pattern == "" {
			pattern = "%s"
		}
		msg := fmt.Sprintf("Expected %s, got %s", "Boolean", obj.GetType().ToString(false))
		panic(RT.NewError(fmt.Sprintf(pattern, msg)))
	}
}

func EnsureArgIsBoolean(args []Object, index int) Boolean {
	switch c := args[index].(type) {
	case Boolean:
		return c
	default:
		panic(RT.NewArgTypeError(index, c, "Boolean"))
	}
}

func EnsureObjectIsTime(obj Object, pattern string) Time {
	switch c := obj.(type) {
	case Time:
		return c
	default:
		if pattern == "" {
			pattern = "%s"
		}
		msg := fmt.Sprintf("Expected %s, got %s", "Time", obj.GetType().ToString(false))
		panic(RT.NewError(fmt.Sprintf(pattern, msg)))
	}
}

func EnsureArgIsTime(args []Object, index int) Time {
	switch c := args[index].(type) {
	case Time:
		return c
	default:
		panic(RT.NewArgTypeError(index, c, "Time"))
	}
}

func EnsureObjectIsNumber(obj Object, pattern string) Number {
	switch c := obj.(type) {
	case Number:
		return c
	default:
		if pattern == "" {
			pattern = "%s"
		}
		msg := fmt.Sprintf("Expected %s, got %s", "Number", obj.GetType().ToString(false))
		panic(RT.NewError(fmt.Sprintf(pattern, msg)))
	}
}

func EnsureArgIsNumber(args []Object, index int) Number {
	switch c := args[index].(type) {
	case Number:
		return c
	default:
		panic(RT.NewArgTypeError(index, c, "Number"))
	}
}

func EnsureObjectIsSeqable(obj Object, pattern string) Seqable {
	switch c := obj.(type) {
	case Seqable:
		return c
	default:
		if pattern == "" {
			pattern = "%s"
		}
		msg := fmt.Sprintf("Expected %s, got %s", "Seqable", obj.GetType().ToString(false))
		panic(RT.NewError(fmt.Sprintf(pattern, msg)))
	}
}

func EnsureArgIsSeqable(args []Object, index int) Seqable {
	switch c := args[index].(type) {
	case Seqable:
		return c
	default:
		panic(RT.NewArgTypeError(index, c, "Seqable"))
	}
}

func EnsureObjectIsCallable(obj Object, pattern string) Callable {
	switch c := obj.(type) {
	case Callable:
		return c
	default:
		if pattern == "" {
			pattern = "%s"
		}
		msg := fmt.Sprintf("Expected %s, got %s", "Callable", obj.GetType().ToString(false))
		panic(RT.NewError(fmt.Sprintf(pattern, msg)))
	}
}

func EnsureArgIsCallable(args []Object, index int) Callable {
	switch c := args[index].(type) {
	case Callable:
		return c
	default:
		panic(RT.NewArgTypeError(index, c, "Callable"))
	}
}

func EnsureObjectIsType(obj Object, pattern string) *Type {
	switch c := obj.(type) {
	case *Type:
		return c
	default:
		if pattern == "" {
			pattern = "%s"
		}
		msg := fmt.Sprintf("Expected %s, got %s", "Type", obj.GetType().ToString(false))
		panic(RT.NewError(fmt.Sprintf(pattern, msg)))
	}
}

func EnsureArgIsType(args []Object, index int) *Type {
	switch c := args[index].(type) {
	case *Type:
		return c
	default:
		panic(RT.NewArgTypeError(index, c, "Type"))
	}
}

func EnsureObjectIsMeta(obj Object, pattern string) Meta {
	switch c := obj.(type) {
	case Meta:
		return c
	default:
		if pattern == "" {
			pattern = "%s"
		}
		msg := fmt.Sprintf("Expected %s, got %s", "Meta", obj.GetType().ToString(false))
		panic(RT.NewError(fmt.Sprintf(pattern, msg)))
	}
}

func EnsureArgIsMeta(args []Object, index int) Meta {
	switch c := args[index].(type) {
	case Meta:
		return c
	default:
		panic(RT.NewArgTypeError(index, c, "Meta"))
	}
}

func EnsureObjectIsInt(obj Object, pattern string) Int {
	switch c := obj.(type) {
	case Int:
		return c
	default:
		if pattern == "" {
			pattern = "%s"
		}
		msg := fmt.Sprintf("Expected %s, got %s", "Int", obj.GetType().ToString(false))
		panic(RT.NewError(fmt.Sprintf(pattern, msg)))
	}
}

func EnsureArgIsInt(args []Object, index int) Int {
	switch c := args[index].(type) {
	case Int:
		return c
	default:
		panic(RT.NewArgTypeError(index, c, "Int"))
	}
}

func EnsureObjectIsDouble(obj Object, pattern string) Double {
	switch c := obj.(type) {
	case Double:
		return c
	default:
		if pattern == "" {
			pattern = "%s"
		}
		msg := fmt.Sprintf("Expected %s, got %s", "Double", obj.GetType().ToString(false))
		panic(RT.NewError(fmt.Sprintf(pattern, msg)))
	}
}

func EnsureArgIsDouble(args []Object, index int) Double {
	switch c := args[index].(type) {
	case Double:
		return c
	default:
		panic(RT.NewArgTypeError(index, c, "Double"))
	}
}

func EnsureObjectIsStack(obj Object, pattern string) Stack {
	switch c := obj.(type) {
	case Stack:
		return c
	default:
		if pattern == "" {
			pattern = "%s"
		}
		msg := fmt.Sprintf("Expected %s, got %s", "Stack", obj.GetType().ToString(false))
		panic(RT.NewError(fmt.Sprintf(pattern, msg)))
	}
}

func EnsureArgIsStack(args []Object, index int) Stack {
	switch c := args[index].(type) {
	case Stack:
		return c
	default:
		panic(RT.NewArgTypeError(index, c, "Stack"))
	}
}

func EnsureObjectIsMap(obj Object, pattern string) Map {
	switch c := obj.(type) {
	case Map:
		return c
	default:
		if pattern == "" {
			pattern = "%s"
		}
		msg := fmt.Sprintf("Expected %s, got %s", "Map", obj.GetType().ToString(false))
		panic(RT.NewError(fmt.Sprintf(pattern, msg)))
	}
}

func EnsureArgIsMap(args []Object, index int) Map {
	switch c := args[index].(type) {
	case Map:
		return c
	default:
		panic(RT.NewArgTypeError(index, c, "Map"))
	}
}

func EnsureObjectIsSet(obj Object, pattern string) Set {
	switch c := obj.(type) {
	case Set:
		return c
	default:
		if pattern == "" {
			pattern = "%s"
		}
		msg := fmt.Sprintf("Expected %s, got %s", "Set", obj.GetType().ToString(false))
		panic(RT.NewError(fmt.Sprintf(pattern, msg)))
	}
}

func EnsureArgIsSet(args []Object, index int) Set {
	switch c := args[index].(type) {
	case Set:
		return c
	default:
		panic(RT.NewArgTypeError(index, c, "Set"))
	}
}

func EnsureObjectIsAssociative(obj Object, pattern string) Associative {
	switch c := obj.(type) {
	case Associative:
		return c
	default:
		if pattern == "" {
			pattern = "%s"
		}
		msg := fmt.Sprintf("Expected %s, got %s", "Associative", obj.GetType().ToString(false))
		panic(RT.NewError(fmt.Sprintf(pattern, msg)))
	}
}

func EnsureArgIsAssociative(args []Object, index int) Associative {
	switch c := args[index].(type) {
	case Associative:
		return c
	default:
		panic(RT.NewArgTypeError(index, c, "Associative"))
	}
}

func EnsureObjectIsReversible(obj Object, pattern string) Reversible {
	switch c := obj.(type) {
	case Reversible:
		return c
	default:
		if pattern == "" {
			pattern = "%s"
		}
		msg := fmt.Sprintf("Expected %s, got %s", "Reversible", obj.GetType().ToString(false))
		panic(RT.NewError(fmt.Sprintf(pattern, msg)))
	}
}

func EnsureArgIsReversible(args []Object, index int) Reversible {
	switch c := args[index].(type) {
	case Reversible:
		return c
	default:
		panic(RT.NewArgTypeError(index, c, "Reversible"))
	}
}

func EnsureObjectIsNamed(obj Object, pattern string) Named {
	switch c := obj.(type) {
	case Named:
		return c
	default:
		if pattern == "" {
			pattern = "%s"
		}
		msg := fmt.Sprintf("Expected %s, got %s", "Named", obj.GetType().ToString(false))
		panic(RT.NewError(fmt.Sprintf(pattern, msg)))
	}
}

func EnsureArgIsNamed(args []Object, index int) Named {
	switch c := args[index].(type) {
	case Named:
		return c
	default:
		panic(RT.NewArgTypeError(index, c, "Named"))
	}
}

func EnsureObjectIsComparator(obj Object, pattern string) Comparator {
	switch c := obj.(type) {
	case Comparator:
		return c
	default:
		if pattern == "" {
			pattern = "%s"
		}
		msg := fmt.Sprintf("Expected %s, got %s", "Comparator", obj.GetType().ToString(false))
		panic(RT.NewError(fmt.Sprintf(pattern, msg)))
	}
}

func EnsureArgIsComparator(args []Object, index int) Comparator {
	switch c := args[index].(type) {
	case Comparator:
		return c
	default:
		panic(RT.NewArgTypeError(index, c, "Comparator"))
	}
}

func EnsureObjectIsRatio(obj Object, pattern string) *Ratio {
	switch c := obj.(type) {
	case *Ratio:
		return c
	default:
		if pattern == "" {
			pattern = "%s"
		}
		msg := fmt.Sprintf("Expected %s, got %s", "Ratio", obj.GetType().ToString(false))
		panic(RT.NewError(fmt.Sprintf(pattern, msg)))
	}
}

func EnsureArgIsRatio(args []Object, index int) *Ratio {
	switch c := args[index].(type) {
	case *Ratio:
		return c
	default:
		panic(RT.NewArgTypeError(index, c, "Ratio"))
	}
}

func EnsureObjectIsNamespace(obj Object, pattern string) *Namespace {
	switch c := obj.(type) {
	case *Namespace:
		return c
	default:
		if pattern == "" {
			pattern = "%s"
		}
		msg := fmt.Sprintf("Expected %s, got %s", "Namespace", obj.GetType().ToString(false))
		panic(RT.NewError(fmt.Sprintf(pattern, msg)))
	}
}

func EnsureArgIsNamespace(args []Object, index int) *Namespace {
	switch c := args[index].(type) {
	case *Namespace:
		return c
	default:
		panic(RT.NewArgTypeError(index, c, "Namespace"))
	}
}

func EnsureObjectIsVar(obj Object, pattern string) *Var {
	switch c := obj.(type) {
	case *Var:
		return c
	default:
		if pattern == "" {
			pattern = "%s"
		}
		msg := fmt.Sprintf("Expected %s, got %s", "Var", obj.GetType().ToString(false))
		panic(RT.NewError(fmt.Sprintf(pattern, msg)))
	}
}

func EnsureArgIsVar(args []Object, index int) *Var {
	switch c := args[index].(type) {
	case *Var:
		return c
	default:
		panic(RT.NewArgTypeError(index, c, "Var"))
	}
}

func EnsureObjectIsError(obj Object, pattern string) Error {
	switch c := obj.(type) {
	case Error:
		return c
	default:
		if pattern == "" {
			pattern = "%s"
		}
		msg := fmt.Sprintf("Expected %s, got %s", "Error", obj.GetType().ToString(false))
		panic(RT.NewError(fmt.Sprintf(pattern, msg)))
	}
}

func EnsureArgIsError(args []Object, index int) Error {
	switch c := args[index].(type) {
	case Error:
		return c
	default:
		panic(RT.NewArgTypeError(index, c, "Error"))
	}
}

func EnsureObjectIsFn(obj Object, pattern string) *Fn {
	switch c := obj.(type) {
	case *Fn:
		return c
	default:
		if pattern == "" {
			pattern = "%s"
		}
		msg := fmt.Sprintf("Expected %s, got %s", "Fn", obj.GetType().ToString(false))
		panic(RT.NewError(fmt.Sprintf(pattern, msg)))
	}
}

func EnsureArgIsFn(args []Object, index int) *Fn {
	switch c := args[index].(type) {
	case *Fn:
		return c
	default:
		panic(RT.NewArgTypeError(index, c, "Fn"))
	}
}

func EnsureObjectIsDeref(obj Object, pattern string) Deref {
	switch c := obj.(type) {
	case Deref:
		return c
	default:
		if pattern == "" {
			pattern = "%s"
		}
		msg := fmt.Sprintf("Expected %s, got %s", "Deref", obj.GetType().ToString(false))
		panic(RT.NewError(fmt.Sprintf(pattern, msg)))
	}
}

func EnsureArgIsDeref(args []Object, index int) Deref {
	switch c := args[index].(type) {
	case Deref:
		return c
	default:
		panic(RT.NewArgTypeError(index, c, "Deref"))
	}
}

func EnsureObjectIsAtom(obj Object, pattern string) *Atom {
	switch c := obj.(type) {
	case *Atom:
		return c
	default:
		if pattern == "" {
			pattern = "%s"
		}
		msg := fmt.Sprintf("Expected %s, got %s", "Atom", obj.GetType().ToString(false))
		panic(RT.NewError(fmt.Sprintf(pattern, msg)))
	}
}

func EnsureArgIsAtom(args []Object, index int) *Atom {
	switch c := args[index].(type) {
	case *Atom:
		return c
	default:
		panic(RT.NewArgTypeError(index, c, "Atom"))
	}
}

func EnsureObjectIsRef(obj Object, pattern string) Ref {
	switch c := obj.(type) {
	case Ref:
		return c
	default:
		if pattern == "" {
			pattern = "%s"
		}
		msg := fmt.Sprintf("Expected %s, got %s", "Ref", obj.GetType().ToString(false))
		panic(RT.NewError(fmt.Sprintf(pattern, msg)))
	}
}

func EnsureArgIsRef(args []Object, index int) Ref {
	switch c := args[index].(type) {
	case Ref:
		return c
	default:
		panic(RT.NewArgTypeError(index, c, "Ref"))
	}
}

func EnsureObjectIsKVReduce(obj Object, pattern string) KVReduce {
	switch c := obj.(type) {
	case KVReduce:
		return c
	default:
		if pattern == "" {
			pattern = "%s"
		}
		msg := fmt.Sprintf("Expected %s, got %s", "KVReduce", obj.GetType().ToString(false))
		panic(RT.NewError(fmt.Sprintf(pattern, msg)))
	}
}

func EnsureArgIsKVReduce(args []Object, index int) KVReduce {
	switch c := args[index].(type) {
	case KVReduce:
		return c
	default:
		panic(RT.NewArgTypeError(index, c, "KVReduce"))
	}
}

func EnsureObjectIsPending(obj Object, pattern string) Pending {
	switch c := obj.(type) {
	case Pending:
		return c
	default:
		if pattern == "" {
			pattern = "%s"
		}
		msg := fmt.Sprintf("Expected %s, got %s", "Pending", obj.GetType().ToString(false))
		panic(RT.NewError(fmt.Sprintf(pattern, msg)))
	}
}

func EnsureArgIsPending(args []Object, index int) Pending {
	switch c := args[index].(type) {
	case Pending:
		return c
	default:
		panic(RT.NewArgTypeError(index, c, "Pending"))
	}
}

func EnsureObjectIsFile(obj Object, pattern string) *File {
	switch c := obj.(type) {
	case *File:
		return c
	default:
		if pattern == "" {
			pattern = "%s"
		}
		msg := fmt.Sprintf("Expected %s, got %s", "File", obj.GetType().ToString(false))
		panic(RT.NewError(fmt.Sprintf(pattern, msg)))
	}
}

func EnsureArgIsFile(args []Object, index int) *File {
	switch c := args[index].(type) {
	case *File:
		return c
	default:
		panic(RT.NewArgTypeError(index, c, "File"))
	}
}

func EnsureObjectIsio_Reader(obj Object, pattern string) io.Reader {
	switch c := obj.(type) {
	case io.Reader:
		return c
	default:
		if pattern == "" {
			pattern = "%s"
		}
		msg := fmt.Sprintf("Expected %s, got %s", "io.Reader", obj.GetType().ToString(false))
		panic(RT.NewError(fmt.Sprintf(pattern, msg)))
	}
}

func EnsureArgIsio_Reader(args []Object, index int) io.Reader {
	switch c := args[index].(type) {
	case io.Reader:
		return c
	default:
		panic(RT.NewArgTypeError(index, c, "io.Reader"))
	}
}

func EnsureObjectIsio_Writer(obj Object, pattern string) io.Writer {
	switch c := obj.(type) {
	case io.Writer:
		return c
	default:
		if pattern == "" {
			pattern = "%s"
		}
		msg := fmt.Sprintf("Expected %s, got %s", "io.Writer", obj.GetType().ToString(false))
		panic(RT.NewError(fmt.Sprintf(pattern, msg)))
	}
}

func EnsureArgIsio_Writer(args []Object, index int) io.Writer {
	switch c := args[index].(type) {
	case io.Writer:
		return c
	default:
		panic(RT.NewArgTypeError(index, c, "io.Writer"))
	}
}

func EnsureObjectIsStringReader(obj Object, pattern string) StringReader {
	switch c := obj.(type) {
	case StringReader:
		return c
	default:
		if pattern == "" {
			pattern = "%s"
		}
		msg := fmt.Sprintf("Expected %s, got %s", "StringReader", obj.GetType().ToString(false))
		panic(RT.NewError(fmt.Sprintf(pattern, msg)))
	}
}

func EnsureArgIsStringReader(args []Object, index int) StringReader {
	switch c := args[index].(type) {
	case StringReader:
		return c
	default:
		panic(RT.NewArgTypeError(index, c, "StringReader"))
	}
}

func EnsureObjectIsio_RuneReader(obj Object, pattern string) io.RuneReader {
	switch c := obj.(type) {
	case io.RuneReader:
		return c
	default:
		if pattern == "" {
			pattern = "%s"
		}
		msg := fmt.Sprintf("Expected %s, got %s", "io.RuneReader", obj.GetType().ToString(false))
		panic(RT.NewError(fmt.Sprintf(pattern, msg)))
	}
}

func EnsureArgIsio_RuneReader(args []Object, index int) io.RuneReader {
	switch c := args[index].(type) {
	case io.RuneReader:
		return c
	default:
		panic(RT.NewArgTypeError(index, c, "io.RuneReader"))
	}
}

func EnsureObjectIsChannel(obj Object, pattern string) *Channel {
	switch c := obj.(type) {
	case *Channel:
		return c
	default:
		if pattern == "" {
			pattern = "%s"
		}
		msg := fmt.Sprintf("Expected %s, got %s", "Channel", obj.GetType().ToString(false))
		panic(RT.NewError(fmt.Sprintf(pattern, msg)))
	}
}

func EnsureArgIsChannel(args []Object, index int) *Channel {
	switch c := args[index].(type) {
	case *Channel:
		return c
	default:
		panic(RT.NewArgTypeError(index, c, "Channel"))
	}
}
