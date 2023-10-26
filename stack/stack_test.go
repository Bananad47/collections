package stack

import "testing"

func TestStack_IsEmpty(t *testing.T) {
	st := NewStack[int]()
	if !st.IsEmpty() {
		t.Errorf("stack is empty")
	}
}

func TestStack(t *testing.T) {
	tests := []struct {
		el   int
		exc  int
		name string
	}{
		{1, 1, "first item"},
		{2, 2, "second item"},
		{3, 3, "third item"},
		{4, 4, "fourth item"},
	}
	st := NewStack[int]()
	for _, el := range tests {
		t.Run(el.name, func(t *testing.T) {
			st.Push(el.el)
			if st.Top() != el.exc {
				t.Errorf("excepted: %d, got: %d", el.el, el.exc)
			}
		})
	}
	st.Pop()
	if st.Top() != tests[len(tests)-2].el {
		t.Errorf("excepted: %d, got %d", st.Top(), tests[len(tests)-2].el)
	}

}
