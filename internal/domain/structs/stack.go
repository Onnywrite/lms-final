package structs

type Stack[T any] struct {
	data []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		data: make([]T, 0, 1024),
	}
}

func (st *Stack[T]) Push(val T) {
	st.data = append(st.data, val)
}

func (st *Stack[T]) TryPop(res *T) (can bool) {
	if !st.TryTop(res) {
		return false
	}
	newSize := len(st.data) - 1
	st.data = st.data[:newSize]
	return true
}

func (st *Stack[T]) Pop() (res T) {
	if !st.TryPop(&res) {
		panic("Stack is empty")
	}
	return
}

func (st *Stack[T]) TryTop(res *T) (can bool) {
	if len(st.data) < 1 {
		return false
	}
	*res = st.data[len(st.data)-1]
	return true
}

func (st *Stack[T]) Top() (res T) {
	if !st.TryTop(&res) {
		panic("Stack is empty")
	}
	return
}
