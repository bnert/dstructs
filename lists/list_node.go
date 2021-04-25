package lists

import (
  "errors"
  "reflect"

  // Local
  "github.com/brent-soles/dstructs/nodes"
)

type GenericNode struct {
  nodes.ListNode
  value interface{}
  next  nodes.ListNode
  prev  nodes.ListNode
}

func NewGenericNode(v interface{}) *GenericNode {
  return &GenericNode{ value: v, next: nil, prev: nil }
}

func validateInterfacePtr(v reflect.Value) (reflect.Value, error) {
  if v.Kind() != reflect.Ptr || !v.IsValid() {
    return reflect.ValueOf(nil), errors.New("cannot unwrap: must unwrap into pointer type")
  }

  v = v.Elem()
  if !v.CanSet() {
    return reflect.ValueOf(nil), errors.New("cannot unwrap: unable to set")
  }
  return v, nil
}

func (g *GenericNode) Wrap(v interface{}) {
  g.value = v
}

func (g *GenericNode) Unwrap() interface{} {
  return g.value
}

func (g *GenericNode) UnwrapInto(v interface{}) error {
  iv, err := validateInterfacePtr(reflect.ValueOf(v))
  if err != nil {
    return err
  }

  rv := reflect.ValueOf(g.value)
  iv.Set(rv)
  return nil
}

func (g *GenericNode) Next() nodes.ListNode {
  return g.next
}

func (g *GenericNode) NextRef() *nodes.ListNode {
  return &g.next
}

func (g *GenericNode) AssignNext(n nodes.ListNode) {
  g.next = n
}

func (g *GenericNode) NextIs(n nodes.ListNode) bool {
  return g.next == n
}

func (g *GenericNode) Prev() nodes.ListNode {
  return g.prev
}

func (g *GenericNode) PrevRef() *nodes.ListNode {
  return &g.prev
}


func (g *GenericNode) AssignPrev(n nodes.ListNode) {
  g.prev = n
}

func (g *GenericNode) PrevIs(n nodes.ListNode) bool {
  return g.prev == n
}
