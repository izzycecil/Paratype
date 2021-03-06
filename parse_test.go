package main

import (
	"Paratype/paraparse"
	"Paratype"
	"testing"
	//"fmt"
	//"Paratype/context"
)


// Used to run a test and check the error. The data is then printed.
func RunTest(code string, file bool, t *testing.T) {
	flist, err := paraparse.Setup(code, file)
	if err != nil {
		t.Error(err)
		return
	}
	main.RunParatype(4, "testout", true, flist)
}

func TestFile(t *testing.T) {
	RunTest("x.para", true, t)
}

func TestString(t *testing.T) {
	RunTest("type int\ntype float\nfunc f() int\n=g(int,float)\nfunc g(A,B) A throws errorType\n=A\n", false, t)
}

// PASS
// typeclass Num inherits Zin
// typeclass Zin
// type z implements Zin
// func foo(d, A) iNT
// 		=x
func Test1(t *testing.T) {
	RunTest("typeclass Num inherits Zin\ntypeclass Zin\ntype z implements Zin\nfunc foo(d, A) iNT\n=x\n", false, t)
}

// PASS
// typeclass Num
// typeclass Zun
// type y implements Zun, Num
// type z implements Num
// func foo constrain A <Num, Zun> (d, A, y) iNT throws bigError, gError
// 		=x
func Test2(t *testing.T) {
	//RunTest("typeclass Num\ntypeclass Zun\ntype y implements Zun, Num\ntype z implements Num\nfunc foo constrain A <Num, Zun> (d, A, y) iNT throws bigError, gError\n=x\n", false, t)
	RunTest("type int\ntype float\ntype boat\nfunc f(A) boat\n=g(m(A), h(A), int)\nfunc g(B, C, B) D\n=D\nfunc m(int) A\n=A\nfunc h(A) float\n=float\n", false, t)

}

// PASS
// FOO: 4 Parameters
// GOO: 3 Parameters
// typeclass Num
// typeclass Zun
// type y implements Zun, Num
// type z implements Num
// func foo constrain A <Num, Zun> (d, A, y) iNt throws bigError, gError
// 		=x
// func goo(x, y) float throws veryBigError, someError, moreError
// 		=x
func TestMultFuncs(t *testing.T) {
	RunTest("typeclass Num\ntypeclass Zun\ntype y implements Zun, Num\ntype z implements Num\nfunc foo constrain A <Num, Zun> (d, A, y) iNt throws bigError, gError\n=x\nfunc goo(x, y) float throws veryBigError, someError, moreError\n=x\n", false, t)
}

// PASS
// typeclass Num
// type y implements Num
// func foo constrain A <Num> (y) int
// 		= y
func TestSimpleConstraint(t *testing.T) {
	RunTest("typeclass Num\ntype y implements Num\nfunc foo constrain A <Num> (y) int\n=y\n", false, t)
}

// PASS
// typeclass Num
// func foo(d, A) int
// 		=bar(A)
// func bar(int) B
// 		=int
func TestComposedFuncs(t *testing.T) {
	RunTest("typeclass Num\ntype int\ntype d\nfunc foo(d, A) int\n=bar(A)\nfunc bar(int) B\n=int\n", false, t)
}

// PASS
// func foo(d, A) int
// 		=bar(baz(ban(A)))
// func bar(int) B
// 		=int
// func baz(int) B
//		=int
// func ban(int) B
//		=baq()
// func baq() B
//		=B
func TestParents(t *testing.T) {
	RunTest("type d\ntype int\nfunc foo(d, A) int\n=bar(baz(ban(A)))\nfunc bar(int) B\n=int\nfunc baz(int) B\n=int\nfunc ban(int) B\n=baq()\nfunc baq() B\n=B\n", false, t)
}

func TestReallySimple(t *testing.T) {
	RunTest("func foo(A) A\n=A\n", false, t)
}

func TestFavorite(t *testing.T) {
	RunTest("typeclass Zat inherits Num\ntypeclass Num\ntype float implements Num\ntype int implements Num\nfunc f() float\n=g(int)\nfunc q() int\n=g(float)\nfunc g constrain A<Zat> (A) B\n=h(A)\nfunc h(A) B\n=B\n", false, t)

}

func TestComposition(t *testing.T) {
	RunTest("typeclass Num\ntype int implements Num\ntype float implements Num\nfunc f(B, B) int\n=bar(baz(B), ban(B))\nfunc bar(A, A) int\n=int\nfunc baz(float) int\n=int\nfunc ban(int) int\n=int\n", false, t)
}

// PASS
// Tests correct creation of the children map.
// func foo(A) int
// 		=bar(baz(A), float)
// func bar(int, float) int
// 		=int
// func baz(int) A
// 		=int
func TestChildren(t *testing.T) {
	RunTest("type int\ntype float\nfunc foo(A) int\n=bar(baz(A), float)\nfunc bar(int, float) int\n=int\nfunc baz(int) int\n=int\n", false, t)
}

// FAIL
// This should fail. Check if a non-existent function fails correctly.
// func foo(A) int
// 		=bar(bal(A), float)
// func bar(int, float) int
// 		=int
// func baz(int) int
// 		=int
func TestNonExistentChild(t *testing.T) {
	RunTest("func foo(A) int\n=bar(bal(A), float)\nfunc bar(int, float) int\n=int\nfunc baz(int) int\n=int\n", false, t)
}

// PASS
// Test the situation where a function has one child function.
// func foo(A) int
// 		=bar(A)
// func bar(A) int
// 		=int
func TestOnlyChild(t *testing.T) {
	RunTest("func foo(A) int\n=bar(A)\nfunc bar(A) int\n=int\n", false, t)
}

// Should fail because Zin does not exist.
// typeclass Num inherits Zin
// typeclass Yin
// type z implements Yin
// func foo(d, A) iNT
// 		=d
func TestFail1(t *testing.T) {
	RunTest("typeclass Num inherits Zin\ntypeclass Yin\ntype z implements Yin\nfunc foo(d, A) iNT\n=d\n", false, t)
}

