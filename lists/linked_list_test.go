package lists

import (
  "fmt"
  "testing"

  // Remote
  "github.com/stretchr/testify/assert"

  // Local
  "github.com/brent-soles/dstructs/nodes"

)

func Test_ValueUnwrapping(t *testing.T) {
  root := NewGenericNode("testing123")

  var v string
  err := root.UnwrapInto(&v)

  assert.Equal(t, nil, err, fmt.Sprintf("Unexpected error: %v", err))
  assert.Equal(t, "testing123", v, "UnwrapInto failed.")

  v = "123testing"

  var unchanged string
  err = root.UnwrapInto(&unchanged)
  assert.Equal(t, nil, err, fmt.Sprintf("Unexpected error: %v", err))
  assert.Equal(t, "testing123", unchanged, "UnwrapInto failed.")
}

func Test_WorkingWithNodes(t *testing.T) {
  root := NewGenericNode("test1")
  next1 := NewGenericNode("test2")
  next2 := NewGenericNode("test3")

  root.AssignNext(next1)
  next1.AssignNext(next2)

  assert.Equal(t, "test1", root.Unwrap(), "")
  assert.Equal(t, "test2", root.Next().Unwrap(), "")
  assert.Equal(t, "test3", root.Next().Next().Unwrap(), "")
  assert.Nil(t, root.Next().Next().Next(), "")
}

func Test_ListConstruction(t *testing.T) {
  l := NewLinkedList()
  l.Append(NewGenericNode(10))
  l.Append(NewGenericNode(20))

  var c []int
  l.CollectInto(&c)

  assert.Equal(t, []int{10, 20}, c, "")
}

func Test_ListConstructionWithList(t *testing.T) {
  l2 := NewLinkedList(1, 2, 3, 4, 5)

  var c2 []int
  l2.CollectInto(&c2)
  assert.Equal(t, []int{1, 2, 3, 4, 5}, c2, "")
}

func Test_Delete(t *testing.T) {
  l := NewLinkedList(1, 2, 3, 4, 5)

  err := l.DeleteAt(0)
  assert.Nil(t, err)

  var c2 []int
  l.CollectInto(&c2)
  assert.Equal(t, []int{2, 3, 4, 5}, c2, "")

  c2 = []int{}

  err = l.DeleteAt(3)
  l.CollectInto(&c2)
  assert.Nil(t, err)
  assert.Equal(t, []int{2, 3, 4}, c2, "")

  err = l.DeleteAt(5)
  assert.Error(t, err)

  err = l.DeleteAt(3)
  assert.Error(t, err)

  err = l.DeleteAt(1)
  var c3 []int
  l.CollectInto(&c3)
  assert.Nil(t, err)
  assert.Equal(t, []int{2, 4}, c3, "")
  assert.Equal(t, uint(2), l.Size(), "")
}

func Test_At(t *testing.T) {
  vals := []interface{}{10, 20, 30, 40}
  l := NewLinkedList(vals...)

  v0, _ := l.At(0)
  assert.Equal(t, 10, v0.Unwrap(), "")

  v2, _ := l.At(2)
  assert.Equal(t, 30, v2.Unwrap(), "")

  _, err := l.At(4)
  assert.Error(t, err)
}

func Test_Prepend(t *testing.T) {
  l := NewLinkedList()

  l.Prepend(NewGenericNode("world"))
  l.Prepend(NewGenericNode("Hello"))

  var c []string
  l.CollectInto(&c)
  assert.Equal(t, []string{"Hello", "world"}, c, "")
}

func Test_ForEach(t *testing.T) {
  vals := []interface{}{10, 20, 30, 40}
  l := NewLinkedList(vals...)

  l.ForEach(func(n nodes.ListNode, i int) {
    assert.Equal(t, vals[i], n.Unwrap(), "")
  })
}

func Test_Filter(t *testing.T) {
  vals := []interface{}{10, 20, 30, 40}
  l := NewLinkedList(vals...)

  fl := l.Filter(func(n nodes.ListNode, i int) bool {
    return n.Unwrap() == 20 || n.Unwrap() == 40
  })

  var c []int
  fl.CollectInto(&c)
  assert.Equal(t, []int{20, 40}, c, "")
}

func Test_InsertAt(t *testing.T) {
  vals := []interface{}{10, 20, 30, 40}
  l := NewLinkedList(vals...)

  l.InsertAt(1, NewGenericNode(15))

  var c []int
  l.CollectInto(&c)
  assert.Equal(t, []int{10, 15, 20, 30, 40}, c, "")

  c = []int{}
  l.InsertAt(100, NewGenericNode(100))
  l.CollectInto(&c)
  assert.Equal(t, []int{10, 15, 20, 30, 40, 100}, c, "")

  c = []int{}
  l.InsertAt(-1, NewGenericNode(-1))
  l.CollectInto(&c)
  assert.Equal(t, []int{-1, 10, 15, 20, 30, 40, 100}, c, "")
}

