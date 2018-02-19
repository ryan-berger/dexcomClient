package dexcomClient

func UrlWithDateRange(config *Config, endpoint string, start string, end string) string {
	url := config.GetBaseUrl() + endpoint + "?startDate=" + start + "&endDate=" + end
	return url
}
