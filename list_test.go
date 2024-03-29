package list

import (
	"fmt"
	"github.com/franela/goblin"
	"testing"
)

var array = []float64{0, 1, 0, 2, 3, 4, 3, 3.3, 9, 29, 3.1, 0.1,
	1.1, 1.81, 0.91, 0.81, 0.71, 0.88, 0.82, 0.81,
}

func TestSSet(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("List", func() {

		g.It("should test list of items", func() {

			var lst = NewList[string]()
			lst.Append("foo")
			//head and tail should be the same
			g.Assert(lst.First()).Eql(lst.Last())
			lst.Append("bar")
			lst.Append("baz")
			lst.Append("boo")
			lst.Append("asd")
			lst.Append("dsa")
			lst.Append("elf")
			lst.Append("fro")
			lst.Append("gap")
			lst.Append("hoo")
			lst.Append("ike")
			lst.Append("jut")
			lst.Append("kni")
			lst.Append("lam")
			lst.Append("mut")
			lst.Append("nop")
			lst.Append("orc")
			//print
			fmt.Println(lst)
			g.Assert(lst.Len()).Equal(17)
			g.Assert(lst.Len()).Equal(17)
			g.Assert(lst.Pop()).Equal("orc")
			g.Assert(lst.PopLeft()).Equal("foo")

			g.Assert(lst.Len()).Equal(15)
			g.Assert(lst.At(0)).Equal("bar")

			g.Assert(lst.At(1)).Equal("baz")
			g.Assert(lst.At(-0)).Equal("bar")
			g.Assert(lst.At(-1)).Equal("nop")
			g.Assert(lst.At(5)).Equal("elf")

			g.Assert(lst.Len()).Equal(15)
			lst.Clear()
			g.Assert(lst.Len()).Equal(0)

		})

		g.It("should test panics1", func() {
			var panics = 0
			var panicFn = func() {
				if err := recover(); err != nil {
					panics += 1
				}
			}
			var lst = NewList[string](1)
			func() {
				defer panicFn()
				lst.Pop() //panic
			}()
			func() {
				defer panicFn()
				lst.PopLeft() //panic
			}()
			g.Assert(panics == 2).IsTrue()
		})
		g.It("should test panics2", func() {
			var panics = 0
			var panicFn = func() {
				if err := recover(); err != nil {
					panics += 1
				}
			}
			var lst = NewList[string](1)
			lst.AppendLeft("first")
			lst.AppendLeft("second")

			lst.PopLeft()
			lst.PopLeft()
			func() {
				defer panicFn()
				lst.First() //panic
			}()
			func() {
				defer panicFn()
				lst.Last() //panic
			}()
			func() {
				defer panicFn()
				lst.At(1) //panic
			}()
			g.Assert(panics == 3).IsTrue()
		})

		g.It("should test list util funcs", func() {
			var lst = NewList[string](1)
			lst.AppendLeft("first")
			lst.AppendLeft("second")
			g.Assert(lst.First()).Equal("second")
			g.Assert(lst.PopLeft()).Equal("second")
			g.Assert(lst.PopLeft()).Equal("first")
			lst.AppendLeft("first")
			lst.AppendLeft("second")
			lst.Pop()
			g.Assert(lst.First()).Equal("second")
			lst.Pop()

			lst.Append("foo")
			lst.AppendLeft("bar")
			lst.Append("baz")
			var eachItem = []string{"bar", "foo", "baz"}
			lst.Each(func(o string, i int) bool {
				g.Assert(o).Equal(eachItem[i])
				return true
			})
			//head and tail should be the same
			g.Assert(lst.Len()).Equal(3)
			g.Assert(lst.First()).Equal("bar")
			g.Assert(lst.Last()).Equal("baz")
			g.Assert(lst.Pop()).Equal("baz")
			g.Assert(lst.PopLeft()).Equal("bar")
			g.Assert(lst.First()).Equal("foo")
			g.Assert(lst.Last()).Equal("foo")
			g.Assert(lst.PopLeft()).Equal("foo")

		})
	})
}
