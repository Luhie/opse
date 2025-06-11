package software

import "time"

// 공통 만료일 계산 유틸
func CalculateExpireDate(grace uint32) string {
	now := time.Now()
	expire := now.Add(time.Duration(grace) * 24 * time.Hour)
	return expire.Format("2006-01-02")
}
