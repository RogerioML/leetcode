package lists

import "testing"

func TestMergeTwoLists(t *testing.T) {
	// Cria duas listas encadeadas de exemplo
	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}}
	list2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}

	// Chama a função para mesclar as duas listas
	mergedList := MergeTwoLists(list1, list2)

	// Define o resultado esperado como uma lista encadeada de valores em ordem crescente
	expected := &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4, Next: nil}}}}}}

	// Compara a lista mesclada com o resultado esperado
	for mergedList != nil {
		if mergedList.Val != expected.Val {
			t.Errorf("Valor esperado %d, mas obtido %d", expected.Val, mergedList.Val)
		}
		mergedList = mergedList.Next
		expected = expected.Next
	}
}
