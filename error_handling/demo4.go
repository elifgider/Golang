package error_handling

import (
	"fmt"
)

type borderException struct {
	parameter int
	massage   string
}

func (b *borderException) Error() string {
	return fmt.Sprintf("%d %s ", b.parameter, b.massage)
}

func TahminEt2(tahmin int) (string, error) {
	if tahmin < 1 || tahmin > 100 {
		return "", &borderException{tahmin, "sınırlar arasında bir sayi giriniz."}
	}
	return "bildiniz", nil
}
