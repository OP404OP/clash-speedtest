package unlock

import (
	"io"
	"net/http"
	"strings"
)

// TestLineTV 测试 LINE TV 解锁情况
func TestLineTV(client *http.Client) *StreamResult {
	result := &StreamResult{
		Platform: "LINE TV",
	}

	req, err := http.NewRequest("GET", "https://www.linetv.tw/", nil)
	if err != nil {
		result.Status = "Failed"
		result.Info = "Create Request Error"
		return result
	}

	req.Header.Set("User-Agent", UA_Browser)
	req.Header.Set("Accept-Language", "zh-TW")

	resp, err := client.Do(req)
	if err != nil {
		result.Status = "Failed"
		result.Info = "Network Connection Error"
		return result
	}
	defer resp.Body.Close()

	// 检查重定向URL
	finalURL := resp.Request.URL.String()
	if strings.Contains(finalURL, "not-available") {
		result.Status = "Failed"
		result.Info = "Region Not Available"
		return result
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		result.Status = "Failed"
		result.Info = "Read Response Error"
		return result
	}

	htmlContent := string(body)

	if strings.Contains(htmlContent, "LINE TV") && !strings.Contains(htmlContent, "not available") {
		result.Status = "Success"
		result.Region = "TW"
	} else {
		result.Status = "Failed"
		result.Info = "Region Not Available"
	}

	return result
}