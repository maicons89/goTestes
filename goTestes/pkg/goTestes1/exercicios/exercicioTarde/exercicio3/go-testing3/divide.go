package divide

import "errors"

func Divisao(num, den int) (int, error) {
	if den == 0 {
		return 0, errors.New("O denominador não pode ser 0")
	}
	return num / den, nil
}
