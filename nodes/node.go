package nodes

type ValueWrapper interface {
  Wrap(interface{})
  Unwrap() interface{}
  UnwrapInto(interface{}) error
}

type NextIndexer interface {
  Next() ListNode
  NextRef() *ListNode
  AssignNext(ListNode)
  NextIs(ListNode) bool
}

type PrevIndexer interface {
  Prev() ListNode
  PrevIs(ListNode) bool
  PrevRef() *ListNode
  AssignPrev(ListNode)
}

type ListNode interface {
  ValueWrapper
  NextIndexer
  PrevIndexer
}
