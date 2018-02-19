package dexcomClient

const GRANT_TYPE = "grant_type=authorization_code&"

func UrlWithDateRange(config *Config, endpoint string, start string, end string) string {
	url := config.GetBaseUrl() + endpoint + "?startDate=" + start + "&endDate=" + end
	return url
}
