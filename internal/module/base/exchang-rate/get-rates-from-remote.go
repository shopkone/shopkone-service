package exchang_rate

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"shopkone-service/internal/module/base/cache/sCache"
	"shopkone-service/utility/code"
)

type getRateFromRemoteResult struct {
	TimeStamp int64
	Rates     map[string]float64
}

const url = "https://openexchangerates.org/api/latest.json?app_id=f8c0abd665da43bca2f531fa8e90f918"

func (s *sExchangeRate) getRatesFromRemote() (result getRateFromRemoteResult, err error) {
	cacheKey := sCache.EXCHANGE_RATE_PREFIX_KEY + "expand_rate"

	// 从缓存中查找，如果有的话，则直接返回
	if err = sCache.NewCache().Get(cacheKey, &result); err != nil {
		return result, err
	}
	if result.TimeStamp > 0 {
		return result, nil
	}

	// 发起 GET 请求
	resp, err := http.Get(url)
	if err != nil {
		return result, code.MarketExchangeRateError
	}
	defer resp.Body.Close()

	// 检查 HTTP 状态码
	if resp.StatusCode != http.StatusOK {
		return result, code.MarketExchangeRateError
	}

	// 读取响应 Body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, code.MarketExchangeRateError
	}

	// 解析
	if err = json.Unmarshal(body, &result); err != nil {
		return result, code.MarketExchangeRateError
	}

	// 1小时有效期
	if err = sCache.NewCache().Set(cacheKey, result, 60); err != nil {
		return result, err
	}

	return result, err
}
