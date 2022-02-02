package bandit

import (
	"errors"
	"math"
)

// формула расчета рейтинга по алгоритму UCB1. Xj + Sqrt( 2ln(N) /Nj )
// где  Xj - среднее количество кликов за просмотр баннера (кол-во кликов поделить на кол-во просмотров)
// N - общее кол-во показов баннеров в слоте
// Nj - кол-во просмотров баннера, для которого расчитываем рейтинг
func calculateRating(clicksCount int, viewsCount int, viewsTotal int) (float64, error) {
	if clicksCount > viewsCount {
		return 0, errors.New("Неверное количество кликов. Кликов не может быть больше просмотров. ")
	}
	if viewsCount > viewsTotal {
		return 0, errors.New("Неверное количество просмотров баннера. " +
			"У одного баннера не может быть просмотров больше чем у всех вместевзятых. ")
	}

	averageProfit := float64(clicksCount) / float64(viewsCount)
	numenator := calculateNumenator(viewsTotal)
	rating := averageProfit + math.Sqrt(numenator/float64(viewsCount))

	return math.Round(rating*100) / 100, nil
}

// берем ранее расчитанное значение, либо считаем
func calculateNumenator(viewsTotal int) float64 {
	return 2 * Ln(float64(viewsTotal))
}

func Ln(x float64) float64 {
	return math.Log(x) / math.Log(math.E)
}
