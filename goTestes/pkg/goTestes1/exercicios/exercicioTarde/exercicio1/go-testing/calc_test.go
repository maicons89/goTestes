package calc_test

import (
	"testing"

	calc "github.com/maicodsantos/gotestes/pkg/goTestes1/exercicios/exercicioTarde/exercicio1/go-testing"
	"github.com/stretchr/testify/assert"
)

func TestSubtracao(t *testing.T) {
	resultadoAtual := calc.Subtracao(5, 3)
	resultadoEsperado := 0

	// if resultadoAtual != resultadoEsperado {
	// 	t.Errorf("failed: o resultado atual %v é diferente do resultado esperado %v", resultadoAtual, resultadoEsperado)
	// }

	assert.Equal(t, resultadoEsperado, resultadoAtual, "o resultado atual precisa ser igual ao resultado esperado")
}
