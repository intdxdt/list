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

		g.It("should test list util funcs", func() {
			var lst = NewList[string](1)
			g.Assert(lst.Pop() == nil).IsTrue()
			g.Assert(lst.PopLeft() == nil).IsTrue()
			lst.AppendLeft("first")
			lst.AppendLeft("second")
			g.Assert(lst.First().(string)).Equal("second")
			g.Assert(lst.PopLeft().(string)).Equal("second")
			g.Assert(lst.PopLeft().(string)).Equal("first")
			g.Assert(lst.First() == nil).IsTrue()
			lst.AppendLeft("first")
			lst.AppendLeft("second")
			lst.Pop()
			g.Assert(lst.First().(string)).Equal("second")
			lst.Pop()
			g.Assert(lst.Last() == nil).IsTrue()
			g.Assert(lst.At(1) == nil).IsTrue()

			lst.Append("foo")
			lst.AppendLeft("bar")
			lst.Append("baz")
			var each_item = []string{"bar", "foo", "baz"}
			lst.Each(func(o string, i int) bool {
				g.Assert(o).Equal(each_item[i])
				return true
			})
			//head and tail should be the same
			g.Assert(lst.Len()).Equal(3)
			g.Assert(lst.First().(string)).Equal("bar")
			g.Assert(lst.Last().(string)).Equal("baz")
			g.Assert(lst.Pop().(string)).Equal("baz")
			g.Assert(lst.PopLeft().(string)).Equal("bar")
			g.Assert(lst.First().(string)).Equal("foo")
			g.Assert(lst.Last().(string)).Equal("foo")
			g.Assert(lst.PopLeft().(string)).Equal("foo")

		})
	})
}
