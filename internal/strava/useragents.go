package strava

type userAgentData struct {
	UserAgent string
	SecChUa   string
}

var userAgents = []userAgentData{
	{
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.5993.117 Safari/537.36",
		SecChUa:   "\"Google Chrome\";v=\"118\", \"Chromium\";v=\"118\", \"Not=A?Brand\";v=\"8\"",
	},
	{
		UserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.5993.117 Safari/537.36",
		SecChUa:   "\"Google Chrome\";v=\"118\", \"Chromium\";v=\"118\", \"Not=A?Brand\";v=\"8\"",
	},
	{
		UserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.6 Safari/605.1.15",
		SecChUa:   "\"Not=A?Brand\";v=\"8\", \"Safari\";v=\"16.6\"",
	},
	{
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:118.0) Gecko/20100101 Firefox/118.0",
		SecChUa:   "\"Firefox\";v=\"118\", \"Not=A?Brand\";v=\"8\"",
	},
	{
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.5993.117 Safari/537.36 Edg/118.0.2088.81",
		SecChUa:   "\"Microsoft Edge\";v=\"118\", \"Chromium\";v=\"118\", \"Not=A?Brand\";v=\"8\"",
	},
	{
		UserAgent: "Mozilla/5.0 (Linux; Android 13; Pixel 7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.5993.117 Mobile Safari/537.36",
		SecChUa:   "\"Google Chrome\";v=\"118\", \"Chromium\";v=\"118\", \"Not=A?Brand\";v=\"8\"",
	},
	{
		UserAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 16_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.5 Mobile/15E148 Safari/604.1",
		SecChUa:   "\"Not=A?Brand\";v=\"8\", \"Safari\";v=\"16.5\"",
	},
	{
		UserAgent: "Mozilla/5.0 (Linux; Android 13; SAMSUNG SM-G990U1) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/20.0 Chrome/118.0.5993.117 Mobile Safari/537.36",
		SecChUa:   "\"SamsungBrowser\";v=\"20.0\", \"Chromium\";v=\"118\", \"Not=A?Brand\";v=\"8\"",
	},
	{
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.5993.117 Safari/537.36 OPR/103.0.0.0",
		SecChUa:   "\"Opera\";v=\"103\", \"Chromium\";v=\"118\", \"Not=A?Brand\";v=\"8\"",
	},
	{
		UserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.5993.117 Safari/537.36 Brave/1.61.118",
		SecChUa:   "\"Brave\";v=\"1.61\", \"Chromium\";v=\"118\", \"Not=A?Brand\";v=\"8\"",
	},
}
