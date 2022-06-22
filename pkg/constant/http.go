package constant

import "time"

// UserAgent used by HTTP client
const UserAgent = "Mozilla/5.0 (compatible; GoodReads/1.0; +https://github.com/adi221/good-reads)"

// DefaultTimeout for HTTP requests
const DefaultTimeout = time.Duration(5 * time.Second)
