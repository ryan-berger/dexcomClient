package dexcomClient

func urlWithDateRange(config *Config, endpoint string, start string, end string) string {
	url := config.getBaseUrl() + endpoint + "?startDate=" + start + "&endDate=" + end
	return url
}
