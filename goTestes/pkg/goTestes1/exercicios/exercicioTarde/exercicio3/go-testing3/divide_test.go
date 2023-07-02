package divide

import (
	"testing"
)

func TestDivisao(t *testing.T) {
	// Testando o caso em que o denominador é diferente de 0
	numerator := 10
	denominator := 2
	expectedResult := 5
	expectedError := error(nil)

	result, err := Divisao(numerator, denominator)

	if result != expectedResult {
		t.Errorf("Resultado esperado: %d, Resultado obtido: %d", expectedResult, result)
	}

	if err != expectedError {
		t.Errorf("Erro esperado: %v, Erro obtido: %v", expectedError, err)
	}
}

/* func TestDivisao(t *testing.T) {
Testando o caso em que o denominador é igual a 0
numerator = 10
denominator = 0
expectedResult = 0
expectedError = errors.New("O denominador não pode ser 0")

result, err = Divisao(numerator, denominator)

if result != expectedResult {
	t.Errorf("Resultado esperado: %d, Resultado obtido: %d", expectedResult, result)
}

if err.Error() != expectedError.Error() {
	t.Errorf("Erro esperado: %v, Erro obtido: %v", expectedError, err)
}
} */
