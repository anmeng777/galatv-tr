package galatvtr

import (
	"fmt"
	"strings"
)

func convertTradingViewTickerToGateioInstId(ticker string) (string, error) {
	// commonQuoteCurrencies 维护一个常见的计价货币列表，用于准确分割交易对。
	// 注意：将更长的代码放在前面，以避免错误匹配 (例如, 优先匹配 "USDT" 而不是 "USD")。
	var commonQuoteCurrencies = []string{"USDT", "USDC"}

	// 1. 标准化输入：转换为大写并移除交易所前缀
	processedTicker := strings.ToUpper(ticker)
	processedTicker = strings.TrimPrefix(processedTicker, "OKX:")

	// 2. 移除可能存在的.P后缀，得到交易对部分
	pair := strings.TrimSuffix(processedTicker, ".P")

	// 3. 寻找分割点并构建instId
	for _, quote := range commonQuoteCurrencies {
		if strings.HasSuffix(pair, quote) {
			base := strings.TrimSuffix(pair, quote)
			return fmt.Sprintf("%s_%s", base, quote), nil
		}
	}

	// 4. 如果无法识别交易对格式，返回错误
	return "", fmt.Errorf("无法识别的交易对格式: %s", ticker)
}
