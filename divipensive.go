package divipensive

import (
	"fmt"
	"math"
)

type toleranceSeries int

const (
	OnePercent  toleranceSeries = 96
	TwoPercent  toleranceSeries = 48
	FivePercent toleranceSeries = 24
	TenPercent  toleranceSeries = 12
)

type decade int

const (
	DecadeOne             decade = 1
	DecadeTen             decade = 10
	DecadeHundred         decade = 100
	DecadeThousand        decade = 1000
	DecadeTenThousand     decade = 10000
	DecadeHundredThousand decade = 100000
	DecadeMillion         decade = 1000000
	DecadeTenMillion      decade = 10000000
)

var roundDigits = map[toleranceSeries]int{
	OnePercent:  2,
	TwoPercent:  2,
	FivePercent: 1,
	TenPercent:  1,
}

func IthValue(i int, N toleranceSeries, d decade) float64 {
	exp := RoundDigits(float64(i)/float64(N), 2)
	val := RoundDigits(math.Pow(10.0, exp), roundDigits[N])
	return float64(d) * val
}

func RoundDigits(n float64, digits int) float64 {
	shift := math.Pow(10, float64(digits))
	n *= shift
	n = math.Floor(n + 0.5)
	return n / shift
}

func Decades() []decade {
	return []decade{
		DecadeOne,
		DecadeTen,
		DecadeHundred,
		DecadeThousand,
		DecadeTenThousand,
		DecadeHundredThousand,
		DecadeMillion,
		DecadeTenMillion,
	}
}

func Tolerances() []toleranceSeries {
	return []toleranceSeries{
		OnePercent,
		TwoPercent,
		FivePercent,
		TenPercent,
	}
}

func ToleranceString(t toleranceSeries) string {
	switch t {
	case OnePercent:
		return "1%"
	case TwoPercent:
		return "2%"
	case FivePercent:
		return "5%"
	case TenPercent:
		return "10%"
	}
	return "invalid tolerance series"
}

func StringValue(i int, N toleranceSeries, d decade) string {
	format := fmt.Sprintf("%%.%df", roundDigits[N])
	return fmt.Sprintf(format, IthValue(i, N, d))
}
