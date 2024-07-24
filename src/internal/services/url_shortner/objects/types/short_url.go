package types

const (
	ShortUrlUsageHistoryStatusFailed    = "failed"
	ShortUrlUsageHistoryStatusSuccess   = "success"
	ShortUrlUsageHistoryStatusForbidden = "forbidden"
)

const (
	ShortUrlTypeProxy    = "proxy"
	ShortUrlTypeRedirect = "redirect"
)

type (
	// ShortUrlType - тип короткого url.
	ShortUrlType string

	// ShortUrlUsageHistoryStatus - статус короткого url в истории использований.
	ShortUrlUsageHistoryStatus string
)
