package helper

import (
	"errors"
	"fmt"
	"strconv"
)

func getKeyFromValue(originalMap map[string]int64, valueToSearch int64) (string, error) {
	for key, val := range originalMap {
		if val == valueToSearch {
			return key, nil
		}
	}
	return "", errors.New("Key not found")
}

// Decimal string to int conversions

func DecimalStrToInt64(str string) (int64, error) {
	n, err := strconv.ParseInt(str, 10, 64)
	if err == nil {
		return n, nil
	}
	return -1, nil
}
func DecimalStrToInt32(str string) (int32, error) {
	n, err := strconv.ParseInt(str, 10, 32)
	if err == nil {
		return int32(n), nil
	}
	return -1, nil
}
func DecimalStrToInt64WithoutError(str string) int64 {
	n, err := strconv.ParseInt(str, 10, 64)
	if err == nil {
		return n
	}
	return -1
}
func DecimalStrToInt32WithoutError(str string) int32 {
	n, err := strconv.ParseInt(str, 10, 32)
	if err == nil {
		return int32(n)
	}
	return -1
}

func BulkDecimalStrToInt64WithoutError(strs []string) []int64 {
	vals := make([]int64, 0)
	for _, str := range strs {
		vals = append(vals, DecimalStrToInt64WithoutError(str))
	}
	return vals
}

func DecimalInt64ToStr(num int64) string {
	return strconv.FormatInt(num, 10)
}

func BulkDecimalInt64ToStr(nums []int64) []string {
	strs := make([]string, 0)
	for _, num := range nums {
		strs = append(strs, DecimalInt64ToStr(num))
	}
	return strs
}

// Hex string to int conversions

func HexStrToInt64(str string) (int64, error) {
	return strconv.ParseInt(str, 16, 64)
}

func HexStrToInt64WoErr(str string) int64 {
	val, err := strconv.ParseInt(str, 16, 64)
	if err == nil {
		return val
	}
	return -1
}

func BulkHexStrToInt64WoErr(strs []string) []int64 {
	intVals := make([]int64, 0)
	for _, str := range strs {
		val, err := strconv.ParseInt(str, 16, 64)
		if err == nil {
			intVals = append(intVals, val)
		} else {
			intVals = append(intVals, -1)
		}
	}
	return intVals
}

func HexInt64ToStr(num int64) string {
	return strconv.FormatInt(num, 16)
}

func BulkHexInt64ToStr(nums []int64) []string {
	strs := make([]string, 0)
	for _, num := range nums {
		strs = append(strs, HexInt64ToStr(num))
	}
	return strs
}

// Int string conversions

func Int64ToHexStr(i int64) string {
	return fmt.Sprintf("%x", i)
}
func Int64ToDecimalStr(i int64) string {
	return strconv.FormatInt(i, 10)
}

// OrgId conversions

func OrgIdToInt64(orgId string) (int64, error) {
	if val, ok := OrgsMap[orgId]; ok {
		return val, nil
	}
	val2, err := DecimalStrToInt64(orgId)
	if err == nil {
		return val2, nil
	}
	return -1, err
}
func OrgIdToInt64WithoutError(orgId string) int64 {
	if val, ok := OrgsMap[orgId]; ok {
		return val
	}
	val2, err := DecimalStrToInt64(orgId)
	if err == nil {
		return val2
	}
	return -1
}
func OrgToString(orgId int64) (string, error) {
	if val, err := getKeyFromValue(OrgsMap, orgId); err == nil {
		return val, nil
	}
	return Int64ToDecimalStr(orgId), nil
}
func OrgToStringWithoutError(orgId int64) string {
	if val, err := getKeyFromValue(OrgsMap, orgId); err == nil {
		return val
	}
	return Int64ToDecimalStr(orgId)
}

// CompanyId conversions

func CompanyIdToInt64(companyId string) (int64, error) {
	if val, ok := CompanyIdsMap[companyId]; ok {
		return val, nil
	}
	val2, err := DecimalStrToInt64(companyId)
	if err == nil {
		return val2, nil
	}
	return -1, err
}
func CompanyIdToInt64WithoutError(companyId string) int64 {
	if val, ok := CompanyIdsMap[companyId]; ok {
		return val
	}
	val2, err := DecimalStrToInt64(companyId)
	if err == nil {
		return val2
	}
	return -1
}
func CompanyIdToString(companyId int64) (string, error) {
	if val, err := getKeyFromValue(CompanyIdsMap, companyId); err == nil {
		return val, nil
	}
	return Int64ToDecimalStr(companyId), nil
}
func CompanyIdToStringWithoutError(companyId int64) string {
	if val, err := getKeyFromValue(CompanyIdsMap, companyId); err == nil {
		return val
	}
	return Int64ToDecimalStr(companyId)
}

// SeriesId conversion

func SeriesIdToInt64(seriesId string) (int64, error) {
	val, err := DecimalStrToInt64(seriesId)
	if err == nil {
		return val, nil
	}
	return -1, err
}

func SeriesIdToInt64WithoutError(seriesId string) int64 {
	val, err := DecimalStrToInt64(seriesId)
	if err == nil {
		return val
	}
	return -1
}

func SeriesIdToStrWithoutError(seriesId int64) string {
	return DecimalInt64ToStr(seriesId)
}

func BulkSeriesIdToInt64WithoutError(seriesIds []string) []int64 {
	intVals := make([]int64, 0)
	for _, seriesId := range seriesIds {
		intVals = append(intVals, SeriesIdToInt64WithoutError(seriesId))
	}
	return intVals
}

// TODO: Add mapper for certificationProgramId
// TODO: Add mapper for moduleId
// TODO: Add mapper for userID

// Int <-> Bool
func BoolToInt(boolean bool) int {
	if boolean {
		return 1
	} else {
		return 0
	}
}
func IntToBool(number int) bool {
	if number == 0 {
		return false
	} else {
		return true
	}
}

/**
This function will convert decimal string to hex string.
*/
func DecimalStrToHexStrWithoutError(str string) string {
	res, err := strconv.ParseInt(str, 10, 64)
	if err == nil {
		return fmt.Sprintf("%x", res)
	}
	return ""
}
