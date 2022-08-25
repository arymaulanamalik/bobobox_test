package utils

import "github.com/leekchan/accounting"

func GenerateFormatPriceIDR(price float64) string {

	ac := accounting.Accounting{Symbol: "IDR ", Precision: 2, Thousand: ".", Decimal: ","}

	return ac.FormatMoney(price)

}
