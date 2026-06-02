package linkedlist_test

import (
	"testing"

	"github.com/murilosolino/data-structures-and-algorithms/internal/linkedlist"
)

func TestAddToFront(t *testing.T) {
	value := 1
	linkedlist := linkedlist.NewLinkedList()
	linkedlist.AddToFront(value)
	head := linkedlist.GetHead().Value
	tail := linkedlist.GetTail().Value
	if head != 1 {
		t.Errorf("head: esperado %v, recebido %v", value, head)
	}
	if tail != 1 {
		t.Errorf("tail: esperado %v, recebido %v", value, head)
	}

	value = 2
	linkedlist.AddToFront(value)
	head = linkedlist.GetHead().Value
	tail = linkedlist.GetTail().Value
	if head != 2 {
		t.Errorf("head: esperado %v, recebido %v", value, head)
	}
	if tail != 1 {
		t.Errorf("tail: esperado %v, recebido %v", value, head)
	}

	value = 3
	linkedlist.AddToFront(value)
	head = linkedlist.GetHead().Value
	tail = linkedlist.GetTail().Value
	if head != 3 {
		t.Errorf("head: esperado %v, recebido %v", value, head)
	}
	if tail != 1 {
		t.Errorf("tail: esperado %v, recebido %v", value, head)
	}
}

func TestAddToEnd(t *testing.T) {
	value := 1
	linkedlist := linkedlist.NewLinkedList()
	linkedlist.AddToEnd(value)
	head := linkedlist.GetHead().Value
	tail := linkedlist.GetTail().Value
	if head != 1 {
		t.Errorf("head: esperado %v, recebido %v", value, head)
	}
	if tail != 1 {
		t.Errorf("tail: esperado %v, recebido %v", value, head)
	}

	value = 2
	linkedlist.AddToEnd(value)
	head = linkedlist.GetHead().Value
	tail = linkedlist.GetTail().Value
	if head != 1 {
		t.Errorf("head: esperado %v, recebido %v", value, head)
	}
	if tail != 2 {
		t.Errorf("tail: esperado %v, recebido %v", value, head)
	}

	value = 3
	linkedlist.AddToEnd(value)
	head = linkedlist.GetHead().Value
	tail = linkedlist.GetTail().Value
	if head != 1 {
		t.Errorf("head: esperado %v, recebido %v", value, head)
	}
	if tail != 3 {
		t.Errorf("tail: esperado %v, recebido %v", value, head)
	}
}

func TestRemoveToFront(t *testing.T) {
	linkedlist := linkedlist.NewLinkedList()
	linkedlist.AddToEnd(1)
	linkedlist.AddToEnd(2)
	linkedlist.AddToEnd(3)
	linkedlist.RemoveToFront()
	if linkedlist.GetHead().Value != 2 {
		t.Errorf("esperado: %v, recebido %v", 2, linkedlist.GetHead().Value)
	}
}

func TestSearchValueInLinkedList(t *testing.T) {
	linkedlist := linkedlist.NewLinkedList()
	linkedlist.AddToEnd(1)
	linkedlist.AddToEnd(2)
	linkedlist.AddToEnd(3)
	linkedlist.AddToEnd(4)
	linkedlist.AddToEnd(5)
	index := linkedlist.SearchValueInLinkedList(4)
	if index != 3 {
		t.Errorf("esperado: %v, recebido %v", 3, index)
	}
}
func TestAddInTheMiddle(t *testing.T) {
	linkedlist := linkedlist.NewLinkedList()
	linkedlist.AddToEnd(1)
	linkedlist.AddToEnd(2)
	linkedlist.AddToEnd(3)
	linkedlist.AddToEnd(4)
	linkedlist.AddToEnd(5)
	linkedlist.AddInTheMiddle(100)

	index := linkedlist.SearchValueInLinkedList(100)
	LastIndex := linkedlist.SearchValueInLinkedList(5)

	if index >= LastIndex {
		linkedlist.PrintLinkedList()
		t.Errorf("Elemento não foi inserido no meio da lista. Index do último nó: %v, Index do elemento inserido: %v", LastIndex, index)
	}
}

func TestGetFromTheMiddle(t *testing.T) {
	linkedlist := linkedlist.NewLinkedList()
	linkedlist.AddToEnd(1)
	linkedlist.AddToEnd(2)
	linkedlist.AddToEnd(3)
	linkedlist.AddToEnd(4)

	middleNode := linkedlist.GetFromTheMiddle()

	if middleNode.Value != 3 {
		t.Errorf("Não foi possível capturar o nó no meio da lista. esperado: %v, recebido: %v", 3, middleNode.Value)
	}
}
