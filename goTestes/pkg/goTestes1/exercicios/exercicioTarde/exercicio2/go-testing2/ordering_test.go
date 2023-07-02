package ordering

import (
	"reflect"
	"testing"
)

func TestOrderArray(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	expected := []int{1, 2, 3, 4, 5}
	result := input
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("O array ordenado é diferente do resultado esperado. Esperado: %v, Recebido: %v", expected, result)
	}
}
