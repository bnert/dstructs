package lists

import (
  "errors"
  "reflect"

  // Local
  "github.com/brent-soles/dstructs/nodes"
)

type ListCallback func(nodes.ListNode, int) (nodes.ListNode, error)
type ListVoidFunc func(nodes.ListNode, int)
type ListFilterFunc func(nodes.ListNode, int) bool

type List interface {
  Size() uint
  Prepend(nodes.ListNode)
  Append(nodes.ListNode)
  At(int) (nodes.ListNode, error)
  InsertAt(int, nodes.ListNode)
  Delete(nodes.ListNode) error
  DeleteAt(int) error

  Find(ListFilterFunc) (nodes.ListNode, error)
  FindRef(ListFilterFunc) (*nodes.ListNode, error)

  ForEach(ListVoidFunc)
  Filter(ListFilterFunc) List
  Map(ListCallback) List

  Collect() []interface{}
  CollectInto(interface{}) error
}

var (
  ListFindByValue = func(v interface{}) ListFilterFunc {
    return func(n nodes.ListNode, i int) bool { return v == n.Unwrap() }
  }

  ListAtIndex = func(i int) ListFilterFunc {
    return func(n nodes.ListNode, idx int) bool { return i == idx }
  }
)

type LinkedList struct {
  List
  head nodes.ListNode
  size uint
}

func NewLinkedList(i ...interface{}) *LinkedList {
  l := &LinkedList{ head: nil, size: 0 }
  for _, v := range i {
    l.Append(NewGenericNode(v))
  }
  return l
}

func (l *LinkedList) Size() uint {
  return l.size
}

func (l *LinkedList) Head() nodes.ListNode {
  return l.head
}

func (l *LinkedList) Prepend(n nodes.ListNode) {
  n.AssignNext(l.head)
  l.head = n
  l.size++
}

func (l *LinkedList) Append(n nodes.ListNode) {
  head := &l.head
  for ; *head != nil; head = (*head).NextRef() {}
  (*head) = n
  l.size++
}

func (l *LinkedList) At(i int) (nodes.ListNode, error) {
  if i >= int(l.size) || i < 0 {
    return nil, errors.New("exceeded bounds")
  }

  return l.Find(ListAtIndex(i))
}

func (l *LinkedList) InsertAt(i int, n nodes.ListNode) {
  if i >= int(l.size) {
    l.Append(n)
    return
  }

  if i < 0 {
    l.Prepend(n)
    return
  }

  node, _ := l.FindRef(ListAtIndex(i))
  n.AssignNext(*node)
  (*node) = n
}

func (l *LinkedList) DeleteAt(i int) error {
  if i >= int(l.size) || i < 0 {
    return errors.New("exceeded bounds")
  }

  node, err := l.FindRef(ListAtIndex(i))
  if err != nil {
    return err
  }

  (*node) = (*node).Next()
  l.size--

  return nil
}

func (l *LinkedList) Find(condition ListFilterFunc) (nodes.ListNode, error) {
  i := 0
  for head := l.head; head != nil; head = head.Next() {
    if condition(head, i) {
      return head, nil
    }
    i++
  }

  return nil, errors.New("unable to find node")
}

func (l *LinkedList) FindRef(condition ListFilterFunc) (*nodes.ListNode, error) {
  i := 0
  head := &l.head
  for ; *head != nil; head = (*head).NextRef() {
    if condition(*head, i) {
      return head, nil
    }
    i++
  }
  return nil, errors.New("unable to find node")
}

func (l *LinkedList) ForEach(fn ListVoidFunc) {
  for i, head := 0, l.head; head != nil; head = head.Next() {
    fn(head, i)
    i++
  }
}

func (l *LinkedList) Filter(fn ListFilterFunc) List {
  list := NewLinkedList()
  for i, head := 0, l.head; head != nil; head = head.Next() {
    if fn(head, i) {
      list.Append(NewGenericNode(head.Unwrap()))
    }
  }
  return list
}

func (l *LinkedList) Collect() []interface{} {
  c := make([]interface{}, 0)
  for head := l.head; head != nil; head = head.Next() {
    c = append(c, head.Unwrap())
  }
  return c
}

func (l *LinkedList) CollectInto(v interface{}) error {
  c := l.Collect()
  rv, err := validateInterfacePtr(reflect.ValueOf(v))
  if err != nil {
    return err
  }

  for _, cv := range c {
    rv.Set(reflect.Append(rv, reflect.ValueOf(cv)))
  }
  return nil
}

