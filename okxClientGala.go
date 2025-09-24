package galatvtr

import (
	"fmt"
	"strconv"
	"time"
)

func (c *OKXClient) ZhuanbiRedemptionAllToAccountBalance(ticker string) error {
	instId, err := ConvertTvTrickerToSingleCoinName(ticker)
	if err != nil {
		fmt.Printf("[Redemption] 转换交易对失败: %v\n", err)
		return err
	}

	var assetBalance float64
	getAssetBalanceResult, err := c.GetAssetBalance(instId)
	if err != nil {
		fmt.Printf("[Redemption] 查询资金账户余额失败: %v\n", err)
		return err
	}
	for _, data := range getAssetBalanceResult.Data {
		if data.Ccy == instId {
			if valueFloat, err := strconv.ParseFloat(data.AvailBal, 64); err == nil {
				assetBalance = valueFloat
			}
		}
	}

	getSavingsBalanceResult, err := c.GetSavingsBalance(instId)
	if err != nil {
		fmt.Printf("[Redemption] 查询稳定赚币余额失败: %v\n", err)
		return err
	}
	var savingBalance float64
	for _, data := range getSavingsBalanceResult.Data {
		if data.Ccy == instId {
			if valueFloat, err := strconv.ParseFloat(data.Amt, 64); err == nil {
				savingBalance = valueFloat
			}
		}
	}

	if savingBalance > 0 {
		request := SavingsPurchaseRedemptRequest{
			Ccy:  instId,
			Side: "redempt",
			Amt:  strconv.FormatFloat(zhuanbiFormatFloat(instId, savingBalance), 'f', 8, 64),
		}
		// 赎回到资金账户
		fmt.Printf("[Redemption] 开始从稳定赚币赎回...\n")
		_, errS := c.SavingsPurchaseRedempt(request)
		if errS == nil {
			// 等待赎回成功，一直查询资金账户直到余额变化
			fmt.Printf("[Redemption] 等待赎回到账，监控资金账户余额变化...\n")
			maxRetries := 4 // 最多等待60次，每次间隔500毫秒，总共30秒
			for i := 0; i < maxRetries; i++ {
				time.Sleep(500 * time.Millisecond) // 等待500毫秒

				// 查询当前资金账户余额
				currentAssetBalanceResult, err := c.GetAssetBalance(instId)
				if err != nil {
					fmt.Printf("[Redemption] 第%d次查询资金账户余额失败: %v\n", i+1, err)
					continue
				}

				var currentAssetBalance float64
				for _, data := range currentAssetBalanceResult.Data {
					if data.Ccy == instId {
						if valueFloat, err := strconv.ParseFloat(data.AvailBal, 64); err == nil {
							currentAssetBalance = valueFloat
						}
					}
				}

				fmt.Printf("[Redemption] 第%d次查询，当前资金账户余额: %.8f，原余额: %.8f\n", i+1, currentAssetBalance, assetBalance)

				// 如果余额有增加，说明赎回已到账
				if currentAssetBalance > assetBalance {
					fmt.Printf("[Redemption] 检测到余额增加，赎回已到账\n")
					assetBalance = currentAssetBalance
					break
				}
			}
		}
	}

	transferRequest := AssetTransferRequest{
		Ccy:  instId,
		Amt:  strconv.FormatFloat(zhuanbiFormatFloat(instId, assetBalance), 'f', 8, 64),
		From: "6",  // 资金账户
		To:   "18", // 交易账户
	}
	_, errT := c.AssetTransfer(transferRequest)
	if errT != nil {
		fmt.Printf("[Redemption] 资金划转失败: %v\n", errT)
		return errT
	}

	return nil
}

func (c *OKXClient) GalaZhuanbiRedemptionAllToAccountBalance(ticker string, assetBalance, savingBalance float64) (float64, error) {
	instId, err := ConvertTvTrickerToSingleCoinName(ticker)
	if err != nil {
		fmt.Printf("[Redemption] 转换交易对失败: %v\n", err)
		return assetBalance, err
	}

	if savingBalance > 0 {
		request := SavingsPurchaseRedemptRequest{
			Ccy:  instId,
			Side: "redempt",
			Amt:  strconv.FormatFloat(zhuanbiFormatFloat(instId, savingBalance), 'f', 8, 64),
		}
		// 赎回到资金账户
		fmt.Printf("[Redemption] 开始从稳定赚币赎回...\n")
		_, errS := c.SavingsPurchaseRedempt(request)
		if errS == nil {
			// 等待赎回成功，一直查询资金账户直到余额变化
			fmt.Printf("[Redemption] 等待赎回到账，监控资金账户余额变化...\n")
			maxRetries := 4 // 最多等待60次，每次间隔500毫秒，总共30秒
			for i := 0; i < maxRetries; i++ {
				time.Sleep(500 * time.Millisecond) // 等待500毫秒

				// 查询当前资金账户余额
				currentAssetBalance, assetBalanceErr := c.GalaGetAssetBalance(ticker)
				if assetBalanceErr != nil {
					return assetBalance, assetBalanceErr
				}

				fmt.Printf("[Redemption] 第%d次查询，当前资金账户余额: %.8f，原余额: %.8f\n", i+1, currentAssetBalance, assetBalance)

				// 如果余额有增加，说明赎回已到账
				if currentAssetBalance > assetBalance {
					fmt.Printf("[Redemption] 检测到余额增加，赎回已到账\n")
					assetBalance = currentAssetBalance
					break
				}
			}
		}
	}

	transferRequest := AssetTransferRequest{
		Ccy:  instId,
		Amt:  strconv.FormatFloat(zhuanbiFormatFloat(instId, assetBalance), 'f', 8, 64),
		From: "6",  // 资金账户
		To:   "18", // 交易账户
	}
	_, errT := c.AssetTransfer(transferRequest)
	if errT != nil {
		fmt.Printf("[Redemption] 资金划转失败: %v\n", errT)
		return assetBalance, errT
	}

	return assetBalance, nil
}

func (c *OKXClient) GalaGetTickerLast(instId string) (float64, error) {
	tickerLast, err := c.GetTickerLast(instId)
	if err != nil {
		return 0, err
	}
	// tickerLast转成float64
	tickerLastFloat, err := strconv.ParseFloat(tickerLast, 64)
	if err != nil {
		return 0, err
	}
	return tickerLastFloat, nil
}

func (c *OKXClient) GalaGetAccountBalance(instId string) (float64, error) {
	// 查询用户指定币种持仓
	ccy, err := ConvertTvTrickerToSingleCoinName(instId)
	if err != nil {
		fmt.Printf("[Redemption] 转换交易对失败: %v\n", err)
		return 0, err
	}
	var accountBalance float64
	{
		getBalanceResult, err := c.GetAccountBalance(ccy)
		if err != nil {
			return 0, err
		}
		for _, data := range getBalanceResult.Data {
			for _, detail := range data.Details {
				if detail.Ccy == ccy {
					if valueFloat, err := strconv.ParseFloat(detail.AvailBal, 64); err == nil {
						accountBalance = valueFloat
					} else {
						return 0, err
					}
				}
			}
		}
	}
	return accountBalance, nil
}

func (c *OKXClient) GalaGetAssetBalance(instId string) (float64, error) {
	ccy, err := ConvertTvTrickerToSingleCoinName(instId)
	if err != nil {
		fmt.Printf("[Redemption] 转换交易对失败: %v\n", err)
		return 0, err
	}
	var assetBalance float64
	{
		getAssetBalanceResult, err := c.GetAssetBalance(ccy)
		if err != nil {
			return 0, err
		}
		for _, data := range getAssetBalanceResult.Data {
			if data.Ccy == ccy {
				if valueFloat, err := strconv.ParseFloat(data.AvailBal, 64); err == nil {
					assetBalance = valueFloat
				} else {
					return 0, err
				}
			}
		}
	}
	return assetBalance, nil
}

func (c *OKXClient) GalaGetSavingBanlance(instId string) (float64, error) {
	ccy, err := ConvertTvTrickerToSingleCoinName(instId)
	if err != nil {
		fmt.Printf("[Redemption] 转换交易对失败: %v\n", err)
		return 0, err
	}
	var savingBalance float64
	{
		getSavingsBalanceResult, err := c.GetSavingsBalance(ccy)
		if err != nil {
			return 0, err
		}
		for _, data := range getSavingsBalanceResult.Data {
			if data.Ccy == ccy {
				if valueFloat, err := strconv.ParseFloat(data.Amt, 64); err == nil {
					savingBalance = valueFloat
				} else {
					return 0, err
				}
			}
		}
	}
	return savingBalance, nil
}
