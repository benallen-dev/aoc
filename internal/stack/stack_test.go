package stack

import (
	"testing"
	"reflect"
)

// TestPush tests push
func TestPush(t *testing.T) {
	testStack := Stack[int]{}
	expected := New([]int{1,3})

	testStack.Push(1)
	testStack.Push(3)

	if !reflect.DeepEqual(testStack, expected) {
		t.Errorf("Expected %v, but got %v", expected, testStack)
	}
}

func TestPop(t *testing.T) {
	testStack := New([]int{1,1,2,3,5})
	expected := []int{5,3,2,1,1}

	for _, e := range expected {
		p, err := testStack.Pop()
		if p != e || err != nil {
			t.Errorf("Expected %v, but got %v", e, p)
		}
	}
}

func TestPeek(t *testing.T) {
	testStack := New([]int{-1,420,69,1337,69420})

	p, err := testStack.Peek()

	if p != 69420 || err != nil {
		t.Errorf("Expected 69420, but got %v", p)
	}

	_, _ = testStack.Pop()

	p, err = testStack.Peek()

	if p != 1337 || err != nil {
		t.Errorf("Expected 1337, but got %v", p)
	}
}
		
