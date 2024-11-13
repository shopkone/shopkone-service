package exchang_rate

import (
	"math"
)

type GetRateIn struct {
	FromCode string
	ToCode   string
}

type GateOut struct {
	Support   bool    // 是否支持
	Rate      float64 // 汇率
	Timestamp int64   // 时间戳
}

func (s *sExchangeRate) GetRate(in GetRateIn) (out GateOut, err error) {
	if in.FromCode == "" || in.ToCode == "" {
		return out, err
	}
	rates, err := s.getRatesFromRemote()
	if err != nil { // 不处理错误
		return out, nil
	}
	fromRate := rates.Rates[in.FromCode]
	toRate := rates.Rates[in.ToCode]
	out.Rate = RoundToFixed(toRate/fromRate, 6)
	out.Support = true
	out.Timestamp = rates.TimeStamp
	return out, err
}

func RoundToFixed(value float64, decimals int) float64 {
	pow := math.Pow(10, float64(decimals))
	return math.Round(value*pow) / pow
}
