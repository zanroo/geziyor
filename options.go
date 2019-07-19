package geziyor

import (
	"github.com/zanroo/geziyor/cache"
	"github.com/zanroo/geziyor/client"
	"github.com/zanroo/geziyor/export"
	"github.com/zanroo/geziyor/metrics"
	"github.com/zanroo/geziyor/middleware"
	"time"
)

// Options is custom options type for Geziyor
type Options struct {
	// AllowedDomains is domains that are allowed to make requests
	// If empty, any domain is allowed
	AllowedDomains []string

	// Cache storage backends.
	// - Memory
	// - Disk
	// - LevelDB
	Cache cache.Cache

	// Policies for caching.
	// - Dummy policy (default)
	// - RFC2616 policy
	CachePolicy cache.Policy

	// Response charset detection for decoding to UTF-8
	CharsetDetectDisabled bool

	// Concurrent requests limit
	ConcurrentRequests int

	// Concurrent requests per domain limit. Uses request.URL.Host
	// Subdomains are different than top domain
	ConcurrentRequestsPerDomain int

	// If set true, cookies won't send.
	CookiesDisabled bool

	// For extracting data
	Exporters []export.Exporter

	// Disable logging by setting this true
	LogDisabled bool

	// Max body reading size in bytes. Default: 1GB
	MaxBodySize int64

	// Maximum redirection time. Default: 10
	MaxRedirect int

	// Scraper metrics exporting type. See metrics.Type
	MetricsType metrics.Type

	// ParseFunc is callback of StartURLs response.
	ParseFunc func(g *Geziyor, r *client.Response)

	// If true, HTML parsing is disabled to improve performance.
	ParseHTMLDisabled bool

	// Request delays
	RequestDelay time.Duration

	// RequestDelayRandomize uses random interval between 0.5 * RequestDelay and 1.5 * RequestDelay
	RequestDelayRandomize bool

	// Called before requests made to manipulate requests
	RequestMiddlewares []middleware.RequestProcessor

	// Called after response received
	ResponseMiddlewares []middleware.ResponseProcessor

	// Which HTTP response codes to retry.
	// Other errors (DNS lookup issues, connections lost, etc) are always retried.
	// Default: []int{500, 502, 503, 504, 522, 524, 408}
	RetryHTTPCodes []int

	// Maximum number of times to retry, in addition to the first download.
	// Set -1 to disable retrying
	// Default: 2
	RetryTimes int

	// If true, disable robots.txt checks
	RobotsTxtDisabled bool

	// StartRequestsFunc called on scraper start
	StartRequestsFunc func(g *Geziyor)

	// First requests will made to this url array. (Concurrently)
	StartURLs []string

	// Timeout is global request timeout
	Timeout time.Duration

	// Revisiting same URLs is disabled by default
	URLRevisitEnabled bool

	// User Agent.
	// Default: "Geziyor 1.0"
	UserAgent string
}
