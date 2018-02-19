package dexcomClient

func urlWithDateRange(config *Config, endpoint string, start string, end string) string {
	url := config.GetBaseUrl() + endpoint + "?startDate=" + start + "&endDate=" + end
	return url
}
