package webpagetest

import (
	"encoding/json"
	"fmt"
	"net/url"
	"sort"
	"strconv"
)

// https://sites.google.com/a/webpagetest.org/docs/advanced-features/raw-test-results
// Pages is struct for links to various pages about test run
type Pages struct {
	Details    string `json:"details"`
	Checklist  string `json:"checklist"`
	Breakdown  string `json:"breakdown"`
	Domains    string `json:"domains"`
	ScreenShot string `json:"screenShot"`
}

// Thumbnails is struct for links to thumbnails of various images for test tun
type Thumbnails struct {
	Waterfall  string `json:"waterfall"`
	Checklist  string `json:"checklist"`
	ScreenShot string `json:"screenShot"`
}

/*
// Images is struct for links to originals of various images for test tun
type Images struct {
	Waterfall      string `json:"waterfall"`
	ConnectionView string `json:"connectionView"`
	Checklist      string `json:"checklist"`
	ScreenShot     string `json:"screenShot"`
	ScreenShotPng  string `json:"screenShotPng"`
}

// RawData is struct for links to raw data about test tun
type RawData struct {
	Headers      string `json:"headers"`
	PageData     string `json:"pageData"`
	RequestsData string `json:"requestsData"`
	Utilization  string `json:"utilization"`
}

// VideoFrame is struct for one video frame
type VideoFrame struct {
	Time  int    `json:"time"`
	Image string `json:"image"`

	VisuallyComplete int `json:"VisuallyComplete"`
}

// Domain is struct for stats about requests form particular domain
type Domain struct {
	Bytes       int    `json:"bytes"`
	Requests    int    `json:"requests"`
	CDNProvider string `json:"cdn_provider"`
	Connections int    `json:"connections"`
}

// Breakdown is struct for data for pie charts of resource distribution
type Breakdown struct {
	Color []int `json:"color"`
	Bytes int   `json:"bytes"`

	Requests int `json:"requests"`
}
*/
// Headers is struct for http headers of request and response
type Headers struct {
	Request  []string `json:"request"`
	Response []string `json:"response"`
}

type jsonRequest struct {
	IP           string `json:"ip_addr"`      // "173.194.122.199"
	Method       string `json:"method"`       // "GET"
	Host         string `json:"host"`         // "google.com"
	URL          string `json:"url"`          // "/"
	FullURL      string `json:"full_url"`     // "http://google.com/"
	ResponseCode string `json:"responseCode"` // "302",

	Protocol  string `json:"protocol"`          // "HTTP/2"
	RequestID int    `json:"request_id,string"` // "9"
	Index     int    `json:"index"`             // 0
	Number    int    `json:"number"`            // 1

	Type     int    `json:"type,string"`   // "3"
	Socket   int    `json:"socket,string"` // "22"
	Priority string `json:"priority"`      // "VeryHigh",

	// Network
	BytesOut         int `json:"bytesOut,string"`          // "397"
	BytesIn          int `json:"bytesIn,string"`           // "467"
	ServerCount      int `json:"server_count,string"`      // "11"
	ServerRTT        int `json:"server_rtt,string"`        // "26"
	ClientPort       int `json:"client_port,string"`       // "55276"
	IsSecure         int `json:"is_secure,string"`         // "0"
	CertificateBytes int `json:"certificate_bytes,string"` // "0", "3769",

	// Cache
	Expires         string `json:"expires"`           // "Tue, 14 Nov 2017 22:46:51 GMT", "-1"
	CacheControl    string `json:"cacheControl"`      // "private"
	CacheTime       int    `json:"cache_time,string"` // "0"
	ContentType     string `json:"contentType"`       // "text/html"
	ContentEncoding string `json:"contentEncoding"`   // "gzip"
	ObjectSize      int    `json:"objectSize,string"` // "256"
	CDNProvider     string `json:"cdn_provider"`      // "Google",

	// Timings
	DNSStart json.Number `json:"dns_start"` // "0"
	DNSEnd   json.Number `json:"dns_end"`   // "50"
	DNS      json.Number `json:"dns_ms"`    // "-1",

	ConnectStart json.Number `json:"connect_start"` // "50"
	ConnectEnd   json.Number `json:"connect_end"`   // "76"
	Connect      json.Number `json:"connect_ms"`    // 26,

	SSLStart json.Number `json:"ssl_start"` // "0"
	SSLEnd   json.Number `json:"ssl_end"`   // "0"
	SSL      json.Number `json:"ssl_ms"`    // "-1",

	LoadStart json.Number `json:"load_start,string"` // "76"
	LoadEnd   json.Number `json:"load_end"`          // 119
	Load      json.Number `json:"load_ms,string"`    // "43",

	TTFBStart json.Number `json:"ttfb_start"`     // "76"
	TTFBEnd   json.Number `json:"ttfb_end"`       // 119
	TTFB      json.Number `json:"ttfb_ms,string"` // "43",

	DownloadStart json.Number `json:"download_start"` // 119
	DownloadEnd   json.Number `json:"download_end"`   // 119
	Download      json.Number `json:"download_ms"`    // 0,

	AllStart json.Number `json:"all_start"` // "50"
	AllEnd   json.Number `json:"all_end"`   // 119
	All      json.Number `json:"all_ms"`    // 69,

	// Optimizations
	ScoreCache           json.Number `json:"score_cache"`            // "0"
	ScoreCDN             json.Number `json:"score_cdn"`              // "-1"
	ScoreGZip            json.Number `json:"score_gzip"`             // "-1"
	ScoreCookies         json.Number `json:"score_cookies"`          // "-1"
	ScoreKeepAlive       json.Number `json:"score_keep-alive"`       // "-1"
	ScoreMinify          json.Number `json:"score_minify"`           // "-1"
	ScoreCombine         json.Number `json:"score_combine"`          // "-1"
	ScoreCompress        json.Number `json:"score_compress"`         // "-1"
	ScoreETags           json.Number `json:"score_etags"`            // "-1"
	ScoreProgressiveJpeg json.Number `json:"score_progressive_jpeg"` // -1
	GZipTotal            json.Number `json:"gzip_total"`             // "0"
	GZipSave             json.Number `json:"gzip_save"`              // "0"
	MinifyTotal          json.Number `json:"minify_total"`           // "0"
	MinifySave           json.Number `json:"minify_save"`            // "0"
	ImageTotal           json.Number `json:"image_total"`            // "0"
	ImageSave            json.Number `json:"image_save"`             // "0"
	JpegScanCount        json.Number `json:"jpeg_scan_count"`        // "0",

	// HTTP/2
	HTTP2StreamDependency int `json:"http2_stream_dependency,string"` // "5"
	HTTP2StreamExclusive  int `json:"http2_stream_exclusive,string"`  // "1"
	HTTP2StreamID         int `json:"http2_stream_id,string"`         // "1"
	HTTP2StreamWeight     int `json:"http2_stream_weight,string"`     // "256"
	WasPushed             int `json:"was_pushed,string"`              // "0",

	// Initiator info
	Initiator         string `json:"initiator"`               // "https://www.google.cz/?gfe_rd=cr&ei=JDc5WJ2sDqSE8QfT-5SgBw&gws_rd=ssl"
	InitiatorColumn   int    `json:"initiator_column,string"` // "104"
	InitiatorDetail   string `json:"initiator_detail"`        // "{\"lineNumber\":50,\"type\":\"parser\",\"url\":\"https://www.google.cz/?gfe_rd=cr&ei=JDc5WJ2sDqSE8QfT-5SgBw&gws_rd=ssl\"}"
	InitiatorFunction string `json:"initiator_function"`      // "Xm"
	InitiatorLine     int    `json:"initiator_line,string"`   // "50"
	InitiatorType     string `json:"initiator_type"`          // "other",

	Headers Headers `json:"headers"`
}

// TestView struct tries to combine to kinds of testViews than WebPagetest returns
// With Steps in case of scripted run and without steps, when we test single url
// Because Go is strictly typed, we have to "merge" them in one data type
type TestView struct {
	Run           float64    `json:"run"`
	Tester        string     `json:"tester"`
	NumberOfSteps int        `json:"numSteps"`
	Steps         []TestStep `json:"steps"`
}

// UnmarshalJSON implements custom unmarshaling logic than mitigates "dynamic"
// nature of test result's json
func (tv *TestView) UnmarshalJSON(b []byte) error {
	var tmp struct {
		Run           float64 `json:"run"`
		Tester        string  `json:"tester"`
		NumberOfSteps int     `json:"numSteps"`
	}
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}

	tv.Run = tmp.Run
	tv.Tester = tmp.Tester
	tv.NumberOfSteps = tmp.NumberOfSteps

	// If we have "steps" array, than Unmarshal as is
	if tmp.NumberOfSteps > 1 {
		var steps struct {
			Steps []TestStep `json:"steps"`
		}
		if err := json.Unmarshal(b, &steps); err != nil {
			return err
		}
		tv.Steps = steps.Steps
		return nil
	}

	// If we have only one "step", then we emulate steps array
	var step TestStep
	if err := json.Unmarshal(b, &step); err != nil {
		return err
	}
	tv.NumberOfSteps = 1
	tv.Steps = []TestStep{step}

	return nil
}

// TestStep is struct with information of one particular test "run"
/*
type TestStep struct {
	URL    string  `json:"URL"`
	Run    int     `json:"run"`
	Date   float64 `json:"date"`   // 1479973600
	Error  string  `json:"error"`  // Timed out waiting for the browser to start.
	Result int     `json:"result"` // 99999

	Tester         string `json:"tester"`
	BrowserName    string `json:"browser_name"`    // "Google Chrome"
	BrowserVersion string `json:"browser_version"` // "54.0.2840.99",
	NumSteps       int    `json:"numSteps"`
	Step           int    `json:"step"`
	EventName      string `json:"eventName"` // "Step 1"

	PageTitle string `json:"title"`
	// Estimated RTT to Server (ms)
	ServerRTT int `json:"server_rtt"`
	// Time to First Byte (ms)
	// The First Byte time (often abbreviated as TTFB) is measured as the time from the start of
	// the initial navigation until the first byte of the base page is received by the browser (after following redirects).
	TTFB int `json:"TTFB"`
	// Time to DOM Loading - From Navigation Timing
	DOMLoading int `json:"domLoading"`
	// Browser-reported first paint time (IE-specific right now - window.performance.timing.msFirstPaint)
	FirstPaint           float64 `json:"firstPaint"`
	FirstMeaningfulPaint int     `json:"chromeUserTiming.firstMeaningfulPaint"`
	// Time from the start of the operation until the title first changed (in ms)
	TitleTime int `json:"titleTime"`
	// Time to DOM Interactive - From Navigation Timing
	DOMInteractive int `json:"domInteractive"`
	// DOM Content Loaded - From Navigation Timing
	DOMContentLoadedEventStart int `json:"domContentLoadedEventStart"`
	DOMContentLoadedEventEnd   int `json:"domContentLoadedEventEnd"` // 455,
	// Browser-reported Load Time (Navigation Timing onload)
	LoadEventStart int `json:"loadEventStart"`
	LoadEventEnd   int `json:"loadEventEnd"`
	// Load Time (onload, ms)
	// The Load Time is measured as the time from the start of the initial navigation until the beginning of the window load event (onload).
	LoadTime int `json:"loadTime"`
	DocTime  int `json:"docTime"`
	DOMTime  int `json:"domTime"`
	// Time to Start Render (ms)
	// The Start Render time is measured as the time from the start of the initial
	// navigation until the first non-white content is painted to the browser display.
	StartRender int `json:"render"`
	// Time to Visually Complete (ms)
	VisualComplete int `json:"visualComplete"`
	// Fully Loaded (ms)
	// The Fully Loaded time is measured as the time from the start of the initial navigation until
	// there was 2 seconds of no network activity after Document Complete.  This will usually
	// include any activity that is triggered by javascript after the main page loads.
	FullyLoaded int `json:"fullyLoaded"`
	// Time of the last visual change to the page (in ms, only available when video capture is enabled)
	LastVisualChange int `json:"lastVisualChange"`
	// Time until the above-the-fold stabilized (if explicitly requested)
	AboveTheFoldTime int `json:"aft"`
	SpeedIndex       int `json:"SpeedIndex"`

	// Number of DOM Elements
	// The DOM Elements metric is the count of the DOM elements on the tested page as measured at the end of the test.
	DOMElements int `json:"domElements"`

	// CPU Busy Time (ms)
	DocCPUms         float32 `json:"docCPUms"`         // 951.606
	FullyLoadedCPUms float32 `json:"fullyLoadedCPUms"` // 1294.808,

	CPUTimes    map[string]int `json:"cpuTimes"`
	CPUTimesDoc map[string]int `json:"cpuTimesDoc"`

	DocCPUpct         int     `json:"docCPUpct"`         // 39
	FullyLoadedCPUpct float64 `json:"fullyLoadedCPUpct"` // 19,

	// The number of bytes downloaded before the Document Complete time
	BytesIn         int `json:"bytesIn"`
	BytesOut        int `json:"bytesOut"`
	BytesInDoc      int `json:"bytesInDoc"`
	BytesOutDoc     int `json:"bytesOutDoc"`
	EffectiveBps    int `json:"effectiveBps"`    // 433693
	EffectiveBpsDoc int `json:"effectiveBpsDoc"` // 466135
	// Total bytes in server-supplied TLS certificates
	CertificateBytes int `json:"certificate_bytes"` // 17499,

	Connections int `json:"connections"`

	// Requests []jsonRequest `json:"requests"`

	RequestsFull int `json:"requestsFull"`
	// The number of http(s) requests before the Document Complete time
	RequestsDoc int `json:"requestsDoc"`

	Responses200   int `json:"responses_200"`
	Responses404   int `json:"responses_404"`
	ResponsesOther int `json:"responses_other"`

	OptimizationChecked  int `json:"optimization_checked"`   // 1
	ScoreCache           int `json:"score_cache"`            // 0
	ScoreCDN             int `json:"score_cdn"`              // -1
	ScoreGZip            int `json:"score_gzip"`             // -1
	ScoreCookies         int `json:"score_cookies"`          // -1
	ScoreKeepAlive       int `json:"score_keep-alive"`       // -1
	ScoreMinify          int `json:"score_minify"`           // -1
	ScoreCombine         int `json:"score_combine"`          // 100
	ScoreCompress        int `json:"score_compress"`         // -1
	ScoreETags           int `json:"score_etags"`            // -1
	ScoreProgressiveJpeg int `json:"score_progressive_jpeg"` // -1,

	GZipTotal   int `json:"gzip_total"`   // 0
	GZipSavings int `json:"gzip_savings"` // 0,

	MinifyTotal   int `json:"minify_total"`   // 0
	MinifySavings int `json:"minify_savings"` // 0,

	ImageTotal   int `json:"image_total"`   // 0
	ImageSavings int `json:"image_savings"` // 0,

	PageSpeedVersion string `json:"pageSpeedVersion"` // "1.9",

	ServerCount int `json:"server_count"` // 16,

	Cached        int `json:"cached"`         // 0,
	AdultSite     int `json:"adult_site"`     // 0,
	FixedViewport int `json:"fixed_viewport"` // 0

	BasePageCDN       string `json:"base_page_cdn"`       // "Google"
	BasePageRedirects int    `json:"base_page_redirects"` // 2
	BasePageTTFB      int    `json:"base_page_ttfb"`      // 524,

	BrowserProcessCount         int `json:"browser_process_count"`           // 8
	BrowserMainMemoryKB         int `json:"browser_main_memory_kb"`          // 69752
	BrowserWorkingSetKB         int `json:"browser_working_set_kb"`          // 136568
	BrowserOtherPrivateMemoryKB int `json:"browser_other_private_memory_kb"` // 66816,

	TimeToInteractive int `json:"TTIMeasurementEnd"` // 11846
	LastInteractive   int `json:"LastInteractive"`   // 9571

	Pages       Pages                `json:"pages"`
	Thumbnails  Thumbnails           `json:"thumbnails"`
	Images      Images               `json:"images"`
	RawData     RawData              `json:"rawData"`
	VideoFrames []VideoFrame         `json:"videoFrames"`
	Breakdown   map[string]Breakdown `json:"breakdown"`

	jsonDomains json.RawMessage   `json:"domains"`
	Domains     map[string]Domain `json:"-"` // may be empty array

	TestTiming map[string]int `json:"testTiming"`
}
*/

type TestStep struct {
	NumSteps                                   int     `json:"numSteps"`
	Run                                        int     `json:"run"`
	Tester                                     string  `json:"tester"`
	MinifyTotal                                int     `json:"minify_total"`
	Responses200                               int     `json:"responses_200"`
	TestStartOffset                            int     `json:"testStartOffset"`
	BytesOut                                   int     `json:"bytesOut"`
	GzipSavings                                int     `json:"gzip_savings"`
	RequestsFull                               int     `json:"requestsFull"`
	StartEpoch                                 int     `json:"start_epoch"`
	Connections                                int     `json:"connections"`
	BasePageCdn                                string  `json:"base_page_cdn"`
	BytesOutDoc                                int     `json:"bytesOutDoc"`
	Result                                     int     `json:"result"`
	FinalBasePageRequestID                     string  `json:"final_base_page_request_id"`
	ScoreCookies                               int     `json:"score_cookies"`
	BasePageSSLTime                            int     `json:"basePageSSLTime"`
	DocTime                                    int     `json:"docTime"`
	DomContentLoadedEventEnd                   int     `json:"domContentLoadedEventEnd"`
	ImageSavings                               int     `json:"image_savings"`
	RequestsDoc                                int     `json:"requestsDoc"`
	FirstMeaningfulPaint                       int     `json:"firstMeaningfulPaint"`
	FirstTextPaint                             int     `json:"firstTextPaint"`
	FirstPaint                                 float64 `json:"firstPaint"`
	ScoreCdn                                   int     `json:"score_cdn"`
	OptimizationChecked                        int     `json:"optimization_checked"`
	ImageTotal                                 int     `json:"image_total"`
	ScoreMinify                                int     `json:"score_minify"`
	GzipTotal                                  int     `json:"gzip_total"`
	Responses404                               int     `json:"responses_404"`
	LoadTime                                   int     `json:"loadTime"`
	URL                                        string  `json:"URL"`
	ScoreCombine                               int     `json:"score_combine"`
	FirstContentfulPaint                       int     `json:"firstContentfulPaint"`
	FirstLayout                                int     `json:"firstLayout"`
	ScoreEtags                                 int     `json:"score_etags"`
	LoadEventStart                             int     `json:"loadEventStart"`
	MinifySavings                              int     `json:"minify_savings"`
	ScoreProgressiveJpeg                       int     `json:"score_progressive_jpeg"`
	DomInteractive                             int     `json:"domInteractive"`
	ScoreGzip                                  int     `json:"score_gzip"`
	ScoreCompress                              int     `json:"score_compress"`
	DomContentLoadedEventStart                 int     `json:"domContentLoadedEventStart"`
	FinalURL                                   string  `json:"final_url"`
	BytesInDoc                                 int     `json:"bytesInDoc"`
	FirstImagePaint                            int     `json:"firstImagePaint"`
	ScoreKeepAlive                             int     `json:"score_keep-alive"`
	LoadEventEnd                               int     `json:"loadEventEnd"`
	Cached                                     int     `json:"cached"`
	ScoreCache                                 int     `json:"score_cache"`
	ResponsesOther                             int     `json:"responses_other"`
	MainFrame                                  string  `json:"main_frame"`
	FullyLoaded                                int     `json:"fullyLoaded"`
	Requests                                   int     `json:"requests"`
	FinalBasePageRequest                       int     `json:"final_base_page_request"`
	TTFB                                       int     `json:"TTFB"`
	BytesIn                                    int     `json:"bytesIn"`
	TestRunTimeMs                              int     `json:"test_run_time_ms"`
	BrowserVersion                             string  `json:"browser_version"`
	BasePageDNSServer                          string  `json:"base_page_dns_server"`
	FullyLoadedCPUms                           int     `json:"fullyLoadedCPUms"`
	PerformancePaintTimingFirstContentfulPaint float64 `json:"PerformancePaintTiming.first-contentful-paint"`
	BasePageIPPtr                              string  `json:"base_page_ip_ptr"`
	EventName                                  string  `json:"eventName"`
	Detected                                   struct {
		Widgets              string `json:"Widgets"`
		Analytics            string `json:"Analytics"`
		JavaScriptFrameworks string `json:"JavaScript Frameworks"`
		Miscellaneous        string `json:"Miscellaneous"`
		WebServers           string `json:"Web Servers"`
	} `json:"detected"`
	BasePageCname                    string  `json:"base_page_cname"`
	DocumentURL                      string  `json:"document_URL"`
	Date                             float64 `json:"date"`
	PerformancePaintTimingFirstPaint float64 `json:"PerformancePaintTiming.first-paint"`
	DomElements                      int     `json:"domElements"`
	DocumentOrigin                   string  `json:"document_origin"`
	BrowserName                      string  `json:"browser_name"`
	DetectedApps                     struct {
		JQuery          string `json:"jQuery"`
		MouseFlow       string `json:"Mouse Flow"`
		SiteCatalyst    string `json:"SiteCatalyst"`
		UnderscoreJs    string `json:"Underscore.js"`
		GoogleAnalytics string `json:"Google Analytics"`
		React           string `json:"React"`
		Nginx           string `json:"Nginx"`
		Webpack         string `json:"webpack"`
		Facebook        string `json:"Facebook"`
		RequireJS       string `json:"RequireJS"`
		LoDash          string `json:"Lo-dash"`
	} `json:"detected_apps"`
	FullyLoadedCPUpct               float64 `json:"fullyLoadedCPUpct"`
	DomComplete                     int     `json:"domComplete"`
	DocumentHostname                string  `json:"document_hostname"`
	UserTimeMarkAdobeTargetHideBody int     `json:"userTime.mark_adobe_target_hide_body"`
	UserTimes                       struct {
		MarkAdobeTargetHideBody        int `json:"mark_adobe_target_hide_body"`
		MarkAdobeTargetVisitorIDStart1 int `json:"mark_adobe_target_visitor_id_start_1"`
		MarkAdobeTargetVisitorIDStart2 int `json:"mark_adobe_target_visitor_id_start_2"`
		MarkAdobeTargetVisitorIDEnd1   int `json:"mark_adobe_target_visitor_id_end_1"`
		MarkAdobeTargetRequestStart1   int `json:"mark_adobe_target_request_start_1"`
		MarkAdobeTargetVisitorIDEnd2   int `json:"mark_adobe_target_visitor_id_end_2"`
		MarkAdobeTargetRequestStart2   int `json:"mark_adobe_target_request_start_2"`
		MarkAdobeTargetRequestEnd2     int `json:"mark_adobe_target_request_end_2"`
		MarkAdobeTargetRequestEnd1     int `json:"mark_adobe_target_request_end_1"`
		MarkAdobeTargetShowBody        int `json:"mark_adobe_target_show_body"`
	} `json:"userTimes"`
	UserTimeMarkAdobeTargetVisitorIDStart1                                   int `json:"userTime.mark_adobe_target_visitor_id_start_1"`
	UserTimeMarkAdobeTargetVisitorIDStart2                                   int `json:"userTime.mark_adobe_target_visitor_id_start_2"`
	UserTimeMarkAdobeTargetVisitorIDEnd1                                     int `json:"userTime.mark_adobe_target_visitor_id_end_1"`
	UserTimeMarkAdobeTargetRequestStart1                                     int `json:"userTime.mark_adobe_target_request_start_1"`
	UserTimeMarkAdobeTargetVisitorIDEnd2                                     int `json:"userTime.mark_adobe_target_visitor_id_end_2"`
	UserTimeMarkAdobeTargetRequestStart2                                     int `json:"userTime.mark_adobe_target_request_start_2"`
	UserTimeMarkAdobeTargetRequestEnd2                                       int `json:"userTime.mark_adobe_target_request_end_2"`
	UserTimeMarkAdobeTargetRequestEnd1                                       int `json:"userTime.mark_adobe_target_request_end_1"`
	UserTimeMarkAdobeTargetShowBody                                          int `json:"userTime.mark_adobe_target_show_body"`
	UserTimingMeasureMeasureAdobeTargetVisitorIDRequest1MboxTargetGlobalMbox int `json:"userTimingMeasure.measure_adobe_target_visitor_id_request_1_mbox_target-global-mbox"`
	UserTimingMeasures                                                       []struct {
		Name      string  `json:"name"`
		StartTime float64 `json:"startTime"`
		Duration  float64 `json:"duration"`
	} `json:"userTimingMeasures"`
	UserTimingMeasureMeasureAdobeTargetVisitorIDRequest2MboxDTPosterHome   int      `json:"userTimingMeasure.measure_adobe_target_visitor_id_request_2_mbox_DT_PosterHome"`
	UserTimingMeasureMeasureAdobeTargetRequestRequest1MboxTargetGlobalMbox int      `json:"userTimingMeasure.measure_adobe_target_request_request_1_mbox_target-global-mbox"`
	UserTimingMeasureMeasureAdobeTargetRequestRequest2MboxDTPosterHome     int      `json:"userTimingMeasure.measure_adobe_target_request_request_2_mbox_DT_PosterHome"`
	UserTime                                                               int      `json:"userTime"`
	Custom                                                                 []string `json:"custom"`
	ImagesString                                                           string   `json:"Images"`
	Colordepth                                                             int      `json:"Colordepth"`
	Resolution                                                             string   `json:"Resolution"`
	Dpi                                                                    string   `json:"Dpi"`
	SpeedIndex                                                             int      `json:"SpeedIndex"`
	VisualComplete85                                                       int      `json:"visualComplete85"`
	VisualComplete90                                                       int      `json:"visualComplete90"`
	VisualComplete95                                                       int      `json:"visualComplete95"`
	VisualComplete99                                                       int      `json:"visualComplete99"`
	VisualComplete                                                         int      `json:"visualComplete"`
	Render                                                                 int      `json:"render"`
	LastVisualChange                                                       int      `json:"lastVisualChange"`
	ChromeUserTiming                                                       []struct {
		Name string `json:"name"`
		Time int    `json:"time"`
	} `json:"chromeUserTiming"`
	ChromeUserTimingRedirectStart                 int `json:"chromeUserTiming.redirectStart"`
	ChromeUserTimingRedirectEnd                   int `json:"chromeUserTiming.redirectEnd"`
	ChromeUserTimingFetchStart                    int `json:"chromeUserTiming.fetchStart"`
	ChromeUserTimingResponseEnd                   int `json:"chromeUserTiming.responseEnd"`
	ChromeUserTimingUnloadEventStart              int `json:"chromeUserTiming.unloadEventStart"`
	ChromeUserTimingUnloadEventEnd                int `json:"chromeUserTiming.unloadEventEnd"`
	ChromeUserTimingDomLoading                    int `json:"chromeUserTiming.domLoading"`
	ChromeUserTimingFirstLayout                   int `json:"chromeUserTiming.firstLayout"`
	ChromeUserTimingFirstPaint                    int `json:"chromeUserTiming.firstPaint"`
	ChromeUserTimingFirstMeaningfulPaintCandidate int `json:"chromeUserTiming.firstMeaningfulPaintCandidate"`
	ChromeUserTimingDomInteractive                int `json:"chromeUserTiming.domInteractive"`
	ChromeUserTimingDomContentLoadedEventStart    int `json:"chromeUserTiming.domContentLoadedEventStart"`
	ChromeUserTimingDomContentLoadedEventEnd      int `json:"chromeUserTiming.domContentLoadedEventEnd"`
	ChromeUserTimingFirstContentfulPaint          int `json:"chromeUserTiming.firstContentfulPaint"`
	ChromeUserTimingFirstMeaningfulPaint          int `json:"chromeUserTiming.firstMeaningfulPaint"`
	ChromeUserTimingFirstTextPaint                int `json:"chromeUserTiming.firstTextPaint"`
	ChromeUserTimingFirstImagePaint               int `json:"chromeUserTiming.firstImagePaint"`
	ChromeUserTimingDomComplete                   int `json:"chromeUserTiming.domComplete"`
	ChromeUserTimingLoadEventStart                int `json:"chromeUserTiming.loadEventStart"`
	ChromeUserTimingLoadEventEnd                  int `json:"chromeUserTiming.loadEventEnd"`
	BlinkFeatureFirstUsed                         struct {
		AnimatedCSSFeatures []interface{} `json:"AnimatedCSSFeatures"`
		CSSFeatures         struct {
			CSSPropertyPosition                   float64 `json:"CSSPropertyPosition"`
			CSSPropertyWebkitTapHighlightColor    float64 `json:"CSSPropertyWebkitTapHighlightColor"`
			CSSPropertyMinHeight                  float64 `json:"CSSPropertyMinHeight"`
			CSSPropertyTextDecoration             float64 `json:"CSSPropertyTextDecoration"`
			CSSPropertyBorderRight                float64 `json:"CSSPropertyBorderRight"`
			CSSPropertyMaxHeight                  float64 `json:"CSSPropertyMaxHeight"`
			CSSPropertyBorderTop                  float64 `json:"CSSPropertyBorderTop"`
			CSSPropertySpeak                      float64 `json:"CSSPropertySpeak"`
			CSSPropertyTransitionDelay            float64 `json:"CSSPropertyTransitionDelay"`
			CSSPropertyPaddingLeft                float64 `json:"CSSPropertyPaddingLeft"`
			CSSPropertyPadding                    float64 `json:"CSSPropertyPadding"`
			CSSPropertyWhiteSpace                 float64 `json:"CSSPropertyWhiteSpace"`
			CSSPropertyAliasWebkitBoxSizing       float64 `json:"CSSPropertyAliasWebkitBoxSizing"`
			CSSPropertyListStyle                  float64 `json:"CSSPropertyListStyle"`
			CSSPropertyVerticalAlign              float64 `json:"CSSPropertyVerticalAlign"`
			CSSPropertyTransform                  float64 `json:"CSSPropertyTransform"`
			CSSPropertyFill                       float64 `json:"CSSPropertyFill"`
			CSSPropertyMinWidth                   float64 `json:"CSSPropertyMinWidth"`
			CSSPropertyTransition                 float64 `json:"CSSPropertyTransition"`
			CSSPropertyAliasWebkitTransformOrigin float64 `json:"CSSPropertyAliasWebkitTransformOrigin"`
			CSSPropertyBorderStyle                float64 `json:"CSSPropertyBorderStyle"`
			CSSPropertyFlexFlow                   float64 `json:"CSSPropertyFlexFlow"`
			CSSPropertyFontFamily                 float64 `json:"CSSPropertyFontFamily"`
			CSSPropertyCursor                     float64 `json:"CSSPropertyCursor"`
			CSSPropertyHeight                     float64 `json:"CSSPropertyHeight"`
			CSSPropertyFlexBasis                  float64 `json:"CSSPropertyFlexBasis"`
			CSSPropertyFontStyle                  float64 `json:"CSSPropertyFontStyle"`
			CSSPropertyBorderImageSource          float64 `json:"CSSPropertyBorderImageSource"`
			CSSPropertyFontVariant                float64 `json:"CSSPropertyFontVariant"`
			CSSPropertyBoxSizing                  float64 `json:"CSSPropertyBoxSizing"`
			CSSPropertyTextOverflow               float64 `json:"CSSPropertyTextOverflow"`
			CSSPropertyBorderColor                float64 `json:"CSSPropertyBorderColor"`
			CSSPropertyPaddingBottom              float64 `json:"CSSPropertyPaddingBottom"`
			CSSPropertyLeft                       float64 `json:"CSSPropertyLeft"`
			CSSPropertyContent                    float64 `json:"CSSPropertyContent"`
			CSSPropertyFilter                     float64 `json:"CSSPropertyFilter"`
			CSSPropertyWidth                      float64 `json:"CSSPropertyWidth"`
			CSSPropertyBorderImageSlice           float64 `json:"CSSPropertyBorderImageSlice"`
			CSSPropertyBorder                     float64 `json:"CSSPropertyBorder"`
			CSSPropertyTouchAction                float64 `json:"CSSPropertyTouchAction"`
			CSSPropertyFlex                       float64 `json:"CSSPropertyFlex"`
			CSSPropertyBackgroundColor            float64 `json:"CSSPropertyBackgroundColor"`
			CSSPropertyPaddingTop                 float64 `json:"CSSPropertyPaddingTop"`
			CSSPropertyBottom                     float64 `json:"CSSPropertyBottom"`
			CSSPropertyBorderCollapse             float64 `json:"CSSPropertyBorderCollapse"`
			CSSPropertyMargin                     float64 `json:"CSSPropertyMargin"`
			CSSPropertyTop                        float64 `json:"CSSPropertyTop"`
			CSSPropertyMarginBottom               float64 `json:"CSSPropertyMarginBottom"`
			CSSPropertyBorderSpacing              float64 `json:"CSSPropertyBorderSpacing"`
			CSSPropertyFloat                      float64 `json:"CSSPropertyFloat"`
			CSSPropertyDisplay                    float64 `json:"CSSPropertyDisplay"`
			CSSPropertyTransformOrigin            float64 `json:"CSSPropertyTransformOrigin"`
			CSSPropertyAliasWebkitUserSelect      float64 `json:"CSSPropertyAliasWebkitUserSelect"`
			CSSPropertyMarginTop                  float64 `json:"CSSPropertyMarginTop"`
			CSSPropertyBorderLeft                 float64 `json:"CSSPropertyBorderLeft"`
			CSSPropertyJustifyContent             float64 `json:"CSSPropertyJustifyContent"`
			CSSPropertyBoxShadow                  float64 `json:"CSSPropertyBoxShadow"`
			CSSPropertyMaxWidth                   float64 `json:"CSSPropertyMaxWidth"`
			CSSPropertyTextIndent                 float64 `json:"CSSPropertyTextIndent"`
			CSSPropertyObjectFit                  float64 `json:"CSSPropertyObjectFit"`
			CSSPropertyZIndex                     float64 `json:"CSSPropertyZIndex"`
			CSSPropertyBorderImageWidth           float64 `json:"CSSPropertyBorderImageWidth"`
			CSSPropertyPaddingRight               float64 `json:"CSSPropertyPaddingRight"`
			CSSPropertyFlexGrow                   float64 `json:"CSSPropertyFlexGrow"`
			CSSPropertyWebkitAppearance           float64 `json:"CSSPropertyWebkitAppearance"`
			CSSPropertyBorderRadius               float64 `json:"CSSPropertyBorderRadius"`
			CSSPropertyFlexDirection              float64 `json:"CSSPropertyFlexDirection"`
			CSSPropertyBorderImageRepeat          float64 `json:"CSSPropertyBorderImageRepeat"`
			CSSPropertyOutline                    float64 `json:"CSSPropertyOutline"`
			CSSPropertyListStyleType              float64 `json:"CSSPropertyListStyleType"`
			CSSPropertyOverflow                   float64 `json:"CSSPropertyOverflow"`
			CSSPropertyFontWeight                 float64 `json:"CSSPropertyFontWeight"`
			CSSPropertyOverflowX                  float64 `json:"CSSPropertyOverflowX"`
			CSSPropertyOverflowY                  float64 `json:"CSSPropertyOverflowY"`
			CSSPropertyTextAlign                  float64 `json:"CSSPropertyTextAlign"`
			CSSPropertyMarginRight                float64 `json:"CSSPropertyMarginRight"`
			CSSPropertyWebkitFontSmoothing        float64 `json:"CSSPropertyWebkitFontSmoothing"`
			CSSPropertyAlignItems                 float64 `json:"CSSPropertyAlignItems"`
			CSSPropertyBackground                 float64 `json:"CSSPropertyBackground"`
			CSSPropertyAliasWebkitFilter          float64 `json:"CSSPropertyAliasWebkitFilter"`
			CSSPropertyLetterSpacing              float64 `json:"CSSPropertyLetterSpacing"`
			CSSPropertyAliasWebkitTextSizeAdjust  float64 `json:"CSSPropertyAliasWebkitTextSizeAdjust"`
			CSSPropertyUserSelect                 float64 `json:"CSSPropertyUserSelect"`
			CSSPropertyLineHeight                 float64 `json:"CSSPropertyLineHeight"`
			CSSPropertyFont                       float64 `json:"CSSPropertyFont"`
			CSSPropertyOpacity                    float64 `json:"CSSPropertyOpacity"`
			CSSPropertyTextTransform              float64 `json:"CSSPropertyTextTransform"`
			CSSPropertyRight                      float64 `json:"CSSPropertyRight"`
			CSSPropertyFontSize                   float64 `json:"CSSPropertyFontSize"`
			CSSPropertyOrder                      float64 `json:"CSSPropertyOrder"`
			CSSPropertyClear                      float64 `json:"CSSPropertyClear"`
			CSSPropertyAliasWebkitTransform       float64 `json:"CSSPropertyAliasWebkitTransform"`
			CSSPropertyBorderWidth                float64 `json:"CSSPropertyBorderWidth"`
			CSSPropertyVisibility                 float64 `json:"CSSPropertyVisibility"`
			CSSPropertyClip                       float64 `json:"CSSPropertyClip"`
			CSSPropertyBorderBottom               float64 `json:"CSSPropertyBorderBottom"`
			CSSPropertyColor                      float64 `json:"CSSPropertyColor"`
			CSSPropertyMarginLeft                 float64 `json:"CSSPropertyMarginLeft"`
		} `json:"CSSFeatures"`
		Features struct {
			XMLHTTPRequestAsynchronous                                float64 `json:"XMLHttpRequestAsynchronous"`
			PlaceholderAttribute                                      float64 `json:"PlaceholderAttribute"`
			PrefixedMutationObserverConstructor                       float64 `json:"PrefixedMutationObserverConstructor"`
			XSSAuditorDisabled                                        float64 `json:"XSSAuditorDisabled"`
			XMLHTTPRequestCrossOriginWithCredentials                  float64 `json:"XMLHttpRequestCrossOriginWithCredentials"`
			V8PerformanceNavigationAttributeGetter                    float64 `json:"V8Performance_Navigation_AttributeGetter"`
			ContentSecurityPolicy                                     float64 `json:"ContentSecurityPolicy"`
			LangAttribute                                             float64 `json:"LangAttribute"`
			NavigatorVendor                                           float64 `json:"NavigatorVendor"`
			XMLDocument                                               float64 `json:"XMLDocument"`
			SameOriginApplicationScript                               float64 `json:"SameOriginApplicationScript"`
			NavigatorProductSub                                       float64 `json:"NavigatorProductSub"`
			CSSFlexibleBox                                            float64 `json:"CSSFlexibleBox"`
			SecureContextCheckPassed                                  float64 `json:"SecureContextCheckPassed"`
			V8LegacyDateParser                                        float64 `json:"V8LegacyDateParser"`
			CSSValueAppearanceNone                                    float64 `json:"CSSValueAppearanceNone"`
			CSSSelectorWebkitFileUploadButton                         float64 `json:"CSSSelectorWebkitFileUploadButton"`
			CrossOriginPropertyAccess                                 float64 `json:"CrossOriginPropertyAccess"`
			V8PerformanceTimingAttributeGetter                        float64 `json:"V8Performance_Timing_AttributeGetter"`
			CSSSelectorWebkitInnerSpinButton                          float64 `json:"CSSSelectorWebkitInnerSpinButton"`
			CSSAtRuleKeyframes                                        float64 `json:"CSSAtRuleKeyframes"`
			DocumentUnloadRegistered                                  float64 `json:"DocumentUnloadRegistered"`
			SandboxViaIFrame                                          float64 `json:"SandboxViaIFrame"`
			V8MessageChannelConstructor                               float64 `json:"V8MessageChannel_Constructor"`
			CrossOriginTextScript                                     float64 `json:"CrossOriginTextScript"`
			Picture                                                   float64 `json:"Picture"`
			HasIDClassTagAttribute                                    float64 `json:"HasIDClassTagAttribute"`
			CrossOriginApplicationScript                              float64 `json:"CrossOriginApplicationScript"`
			CSSSelectorWebkitInputPlaceholder                         float64 `json:"CSSSelectorWebkitInputPlaceholder"`
			ScrollToFragmentRequested                                 float64 `json:"ScrollToFragmentRequested"`
			SecureContextCheckForSandboxedOriginPassed                float64 `json:"SecureContextCheckForSandboxedOriginPassed"`
			V8CustomEventInitCustomEventMethod                        float64 `json:"V8CustomEvent_InitCustomEvent_Method"`
			UnprefixedUserTiming                                      float64 `json:"UnprefixedUserTiming"`
			SuppressHistoryEntryWithoutUserGesture                    float64 `json:"SuppressHistoryEntryWithoutUserGesture"`
			CrossOriginMainFrameNulledNameAccessed                    float64 `json:"CrossOriginMainFrameNulledNameAccessed"`
			V8ScreenAvailTopAttributeGetter                           float64 `json:"V8Screen_AvailTop_AttributeGetter"`
			V8ElementGetClientRectsMethod                             float64 `json:"V8Element_GetClientRects_Method"`
			UnprefixedRequestAnimationFrame                           float64 `json:"UnprefixedRequestAnimationFrame"`
			ScrollToFragmentFailWithASCII                             float64 `json:"ScrollToFragmentFailWithASCII"`
			SVGSVGElementInXMLDocument                                float64 `json:"SVGSVGElementInXMLDocument"`
			CSSSelectorWebkitSearchCancelButton                       float64 `json:"CSSSelectorWebkitSearchCancelButton"`
			FormsSubmitted                                            float64 `json:"FormsSubmitted"`
			CSPWithUnsafeEval                                         float64 `json:"CSPWithUnsafeEval"`
			PendingStylesheetAddedAfterBodyStarted                    float64 `json:"PendingStylesheetAddedAfterBodyStarted"`
			HTMLElementInnerText                                      float64 `json:"HTMLElementInnerText"`
			CookieSet                                                 float64 `json:"CookieSet"`
			FontShapingNotDefGlyphObserved                            float64 `json:"FontShapingNotDefGlyphObserved"`
			WindowPostMessage                                         float64 `json:"WindowPostMessage"`
			FormElement                                               float64 `json:"FormElement"`
			CSSValueAppearanceButton                                  float64 `json:"CSSValueAppearanceButton"`
			V8DeoptimizerDisableSpeculation                           float64 `json:"V8DeoptimizerDisableSpeculation"`
			UnprefixedPerformanceTimeline                             float64 `json:"UnprefixedPerformanceTimeline"`
			DocumentAll                                               float64 `json:"DocumentAll"`
			DocumentBeforeUnloadRegistered                            float64 `json:"DocumentBeforeUnloadRegistered"`
			CSSFilterGrayscale                                        float64 `json:"CSSFilterGrayscale"`
			V8ScreenAvailLeftAttributeGetter                          float64 `json:"V8Screen_AvailLeft_AttributeGetter"`
			HasBeforeOrAfterPseudoElement                             float64 `json:"HasBeforeOrAfterPseudoElement"`
			DuplicatedAttribute                                       float64 `json:"DuplicatedAttribute"`
			CrossOriginTextHTML                                       float64 `json:"CrossOriginTextHtml"`
			CSSSelectorPseudoFocus                                    float64 `json:"CSSSelectorPseudoFocus"`
			HTMLSlotElement                                           float64 `json:"HTMLSlotElement"`
			MultipleOriginsInTimingAllowOrigin                        float64 `json:"MultipleOriginsInTimingAllowOrigin"`
			CertificateTransparencyNonCompliantSubresourceInMainFrame float64 `json:"CertificateTransparencyNonCompliantSubresourceInMainFrame"`
			CertificateTransparencyNonCompliantResourceInSubframe     float64 `json:"CertificateTransparencyNonCompliantResourceInSubframe"`
			V8StrictMode                                              float64 `json:"V8StrictMode"`
			CSSAtRuleWebkitKeyframes                                  float64 `json:"CSSAtRuleWebkitKeyframes"`
			AddEventListenerThirdArgumentIsObject                     float64 `json:"AddEventListenerThirdArgumentIsObject"`
			CryptoGetRandomValues                                     float64 `json:"CryptoGetRandomValues"`
			CSSAtRuleFontFace                                         float64 `json:"CSSAtRuleFontFace"`
			V8ElementGetBoundingClientRectMethod                      float64 `json:"V8Element_GetBoundingClientRect_Method"`
			V8EventInitEventMethod                                    float64 `json:"V8Event_InitEvent_Method"`
			PrefixedPageVisibility                                    float64 `json:"PrefixedPageVisibility"`
			SVGSVGElementInDocument                                   float64 `json:"SVGSVGElementInDocument"`
			SecureContextCheckFailed                                  float64 `json:"SecureContextCheckFailed"`
			SVGSVGElement                                             float64 `json:"SVGSVGElement"`
			CSSGradient                                               float64 `json:"CSSGradient"`
			LangAttributeOnHTML                                       float64 `json:"LangAttributeOnHTML"`
			CSSAtRuleMedia                                            float64 `json:"CSSAtRuleMedia"`
			InputTypeText                                             float64 `json:"InputTypeText"`
			CSSSelectorWebkitUnknownPseudo                            float64 `json:"CSSSelectorWebkitUnknownPseudo"`
			StarInTimingAllowOrigin                                   float64 `json:"StarInTimingAllowOrigin"`
			CleanScriptElementWithNonce                               float64 `json:"CleanScriptElementWithNonce"`
			SecureContextCheckForSandboxedOriginFailed                float64 `json:"SecureContextCheckForSandboxedOriginFailed"`
			V8SloppyMode                                              int     `json:"V8SloppyMode"`
			CookieGet                                                 float64 `json:"CookieGet"`
		} `json:"Features"`
	} `json:"blinkFeatureFirstUsed"`
	Step            int `json:"step"`
	EffectiveBps    int `json:"effectiveBps"`
	EffectiveBpsDoc int `json:"effectiveBpsDoc"`
	DomTime         int `json:"domTime"`
	Aft             int `json:"aft"`
	TitleTime       int `json:"titleTime"`
	DomLoading      int `json:"domLoading"`
	ServerRtt       int `json:"server_rtt"`
	SmallImageCount int `json:"smallImageCount"`
	BigImageCount   int `json:"bigImageCount"`
	MaybeCaptcha    int `json:"maybeCaptcha"`
	Pages           struct {
		Details    string `json:"details"`
		Checklist  string `json:"checklist"`
		Breakdown  string `json:"breakdown"`
		Domains    string `json:"domains"`
		ScreenShot string `json:"screenShot"`
	} `json:"pages"`
	Thumbnails struct {
		Waterfall  string `json:"waterfall"`
		Checklist  string `json:"checklist"`
		ScreenShot string `json:"screenShot"`
	} `json:"thumbnails"`
	Images struct {
		Waterfall      string `json:"waterfall"`
		ConnectionView string `json:"connectionView"`
		Checklist      string `json:"checklist"`
		ScreenShot     string `json:"screenShot"`
	} `json:"images"`
	RawData struct {
		Headers      string `json:"headers"`
		PageData     string `json:"pageData"`
		RequestsData string `json:"requestsData"`
		Utilization  string `json:"utilization"`
		Trace        string `json:"trace"`
	} `json:"rawData"`
	Domains struct {
		DertourDe struct {
			Bytes       int `json:"bytes"`
			Requests    int `json:"requests"`
			Connections int `json:"connections"`
		} `json:"dertour.de"`
		MetricsDertourDe struct {
			Bytes       int `json:"bytes"`
			Requests    int `json:"requests"`
			Connections int `json:"connections"`
		} `json:"metrics.dertour.de"`
		WwwDertourDe struct {
			Bytes       int `json:"bytes"`
			Requests    int `json:"requests"`
			Connections int `json:"connections"`
		} `json:"www.dertour.de"`
		WwwDwin1Com struct {
			Bytes       int    `json:"bytes"`
			Requests    int    `json:"requests"`
			CdnProvider string `json:"cdn_provider"`
			Connections int    `json:"connections"`
		} `json:"www.dwin1.com"`
		T2SymcbCom struct {
			Bytes       int `json:"bytes"`
			Requests    int `json:"requests"`
			Connections int `json:"connections"`
		} `json:"t2.symcb.com"`
		FontsGstaticCom struct {
			Bytes       int    `json:"bytes"`
			Requests    int    `json:"requests"`
			CdnProvider string `json:"cdn_provider"`
			Connections int    `json:"connections"`
		} `json:"fonts.gstatic.com"`
		SBtstaticCom struct {
			Bytes       int    `json:"bytes"`
			Requests    int    `json:"requests"`
			CdnProvider string `json:"cdn_provider"`
			Connections int    `json:"connections"`
		} `json:"s.btstatic.com"`
		TiSymcdCom struct {
			Bytes       int `json:"bytes"`
			Requests    int `json:"requests"`
			Connections int `json:"connections"`
		} `json:"ti.symcd.com"`
		WwwGoogleCom struct {
			Bytes       int    `json:"bytes"`
			Requests    int    `json:"requests"`
			CdnProvider string `json:"cdn_provider"`
			Connections int    `json:"connections"`
		} `json:"www.google.com"`
		SThebrighttagCom struct {
			Bytes       int `json:"bytes"`
			Requests    int `json:"requests"`
			Connections int `json:"connections"`
		} `json:"s.thebrighttag.com"`
		DBtttagCom struct {
			Bytes       int    `json:"bytes"`
			Requests    int    `json:"requests"`
			CdnProvider string `json:"cdn_provider"`
			Connections int    `json:"connections"`
		} `json:"d.btttag.com"`
		DertourBtttagCom struct {
			Bytes       int    `json:"bytes"`
			Requests    int    `json:"requests"`
			CdnProvider string `json:"cdn_provider"`
			Connections int    `json:"connections"`
		} `json:"dertour.btttag.com"`
		BatBingCom struct {
			Bytes       int `json:"bytes"`
			Requests    int `json:"requests"`
			Connections int `json:"connections"`
		} `json:"bat.bing.com"`
		WwwFacebookCom struct {
			Bytes       int    `json:"bytes"`
			Requests    int    `json:"requests"`
			CdnProvider string `json:"cdn_provider"`
			Connections int    `json:"connections"`
		} `json:"www.facebook.com"`
		AssetsAdobedtmCom struct {
			Bytes       int    `json:"bytes"`
			Requests    int    `json:"requests"`
			CdnProvider string `json:"cdn_provider"`
			Connections int    `json:"connections"`
		} `json:"assets.adobedtm.com"`
		WwwGoogleAnalyticsCom struct {
			Bytes       int    `json:"bytes"`
			Requests    int    `json:"requests"`
			CdnProvider string `json:"cdn_provider"`
			Connections int    `json:"connections"`
		} `json:"www.google-analytics.com"`
		WwwGoogleadservicesCom struct {
			Bytes       int    `json:"bytes"`
			Requests    int    `json:"requests"`
			CdnProvider string `json:"cdn_provider"`
			Connections int    `json:"connections"`
		} `json:"www.googleadservices.com"`
		AjaxGoogleapisCom struct {
			Bytes       int    `json:"bytes"`
			Requests    int    `json:"requests"`
			CdnProvider string `json:"cdn_provider"`
			Connections int    `json:"connections"`
		} `json:"ajax.googleapis.com"`
		O2MouseflowCom struct {
			Bytes       int `json:"bytes"`
			Requests    int `json:"requests"`
			Connections int `json:"connections"`
		} `json:"o2.mouseflow.com"`
		CdnMouseflowCom struct {
			Bytes       int    `json:"bytes"`
			Requests    int    `json:"requests"`
			CdnProvider string `json:"cdn_provider"`
			Connections int    `json:"connections"`
		} `json:"cdn.mouseflow.com"`
		DertouristikonlinegmTtOmtrdcNet struct {
			Bytes       int `json:"bytes"`
			Requests    int `json:"requests"`
			Connections int `json:"connections"`
		} `json:"dertouristikonlinegm.tt.omtrdc.net"`
		CmEveresttechNet struct {
			Bytes       int `json:"bytes"`
			Requests    int `json:"requests"`
			Connections int `json:"connections"`
		} `json:"cm.everesttech.net"`
		GoogleadsGDoubleclickNet struct {
			Bytes       int    `json:"bytes"`
			Requests    int    `json:"requests"`
			CdnProvider string `json:"cdn_provider"`
			Connections int    `json:"connections"`
		} `json:"googleads.g.doubleclick.net"`
		StatsGDoubleclickNet struct {
			Bytes       int    `json:"bytes"`
			Requests    int    `json:"requests"`
			CdnProvider string `json:"cdn_provider"`
			Connections int    `json:"connections"`
		} `json:"stats.g.doubleclick.net"`
		ConnectFacebookNet struct {
			Bytes       int    `json:"bytes"`
			Requests    int    `json:"requests"`
			CdnProvider string `json:"cdn_provider"`
			Connections int    `json:"connections"`
		} `json:"connect.facebook.net"`
		TrackAdformNet struct {
			Bytes       int `json:"bytes"`
			Requests    int `json:"requests"`
			Connections int `json:"connections"`
		} `json:"track.adform.net"`
		DertouristikDemdexNet struct {
			Bytes       int `json:"bytes"`
			Requests    int `json:"requests"`
			Connections int `json:"connections"`
		} `json:"dertouristik.demdex.net"`
		DpmDemdexNet struct {
			Bytes       int `json:"bytes"`
			Requests    int `json:"requests"`
			Connections int `json:"connections"`
		} `json:"dpm.demdex.net"`
	} `json:"domains"`
	Breakdown struct {
		HTML struct {
			Color             []int `json:"color"`
			Bytes             int   `json:"bytes"`
			BytesUncompressed int   `json:"bytesUncompressed"`
			Requests          int   `json:"requests"`
		} `json:"html"`
		Js struct {
			Color             []int `json:"color"`
			Bytes             int   `json:"bytes"`
			BytesUncompressed int   `json:"bytesUncompressed"`
			Requests          int   `json:"requests"`
		} `json:"js"`
		CSS struct {
			Color             []int `json:"color"`
			Bytes             int   `json:"bytes"`
			BytesUncompressed int   `json:"bytesUncompressed"`
			Requests          int   `json:"requests"`
		} `json:"css"`
		Image struct {
			Color             []int `json:"color"`
			Bytes             int   `json:"bytes"`
			BytesUncompressed int   `json:"bytesUncompressed"`
			Requests          int   `json:"requests"`
		} `json:"image"`
		Flash struct {
			Color             []int `json:"color"`
			Bytes             int   `json:"bytes"`
			BytesUncompressed int   `json:"bytesUncompressed"`
			Requests          int   `json:"requests"`
		} `json:"flash"`
		Font struct {
			Color             []int `json:"color"`
			Bytes             int   `json:"bytes"`
			BytesUncompressed int   `json:"bytesUncompressed"`
			Requests          int   `json:"requests"`
		} `json:"font"`
		Video struct {
			Color             []int `json:"color"`
			Bytes             int   `json:"bytes"`
			BytesUncompressed int   `json:"bytesUncompressed"`
			Requests          int   `json:"requests"`
		} `json:"video"`
		Other struct {
			Color             []int `json:"color"`
			Bytes             int   `json:"bytes"`
			BytesUncompressed int   `json:"bytesUncompressed"`
			Requests          int   `json:"requests"`
		} `json:"other"`
	} `json:"breakdown"`
}

type TestRun struct {
	FirstView  TestView `json:"firstView"`
	RepeatView TestView `json:"repeatView"`
}

type ResultData struct {
	Connectivity

	ID       string `json:"id"`
	URL      string `json:"url"`
	Summary  string `json:"summary"`
	TestUrl  string `json:"testUrl"`
	Location string `json:"location"`
	Label    string `json:"label"`
	From     string `json:"from"`

	Mobile           int    `json:"mobile"`
	Completed        int    `json:"completed"`
	Tester           string `json:"tester"`
	TesterDNS        string `json:"testerDNS"`
	FirstViewOnly    bool   `json:"fvonly"`
	SuccessfulFVRuns int    `json:"successfulFVRuns"`
	SuccessfulRVRuns int    `json:"successfulRVRuns"`

	Runs map[string]TestRun `json:"runs"`
}

// GetMedianRun will calculate and return median run by given metric and step
// Step is 0-based
func (rd *ResultData) GetMedianRun(step int, metric string) (*TestRun, error) {
	var testRun TestRun
	var firstViewValues []float64
	var repeatViewValues []float64
	firstViewValueMap := make(map[float64]string, 0)
	repeatViewValueMap := make(map[float64]string, 0)

	for idx, run := range rd.Runs {
		if step >= len(run.FirstView.Steps) || step >= len(run.RepeatView.Steps) {
			continue
		}
		var firstViewValue float64
		var repeatViewValue float64
		switch metric {
		case "speedindex":
			firstViewValue = float64(run.FirstView.Steps[step].SpeedIndex)
			repeatViewValue = float64(run.RepeatView.Steps[step].SpeedIndex)
		case "loadtime":
			firstViewValue = float64(run.FirstView.Steps[step].LoadTime)
			repeatViewValue = float64(run.RepeatView.Steps[step].LoadTime)
		case "fullyloaded":
			firstViewValue = float64(run.FirstView.Steps[step].FullyLoaded)
			repeatViewValue = float64(run.RepeatView.Steps[step].FullyLoaded)
		default:
			return nil, fmt.Errorf("unsupported metric: %s", metric)
		}
		firstViewValueMap[firstViewValue] = idx
		repeatViewValueMap[repeatViewValue] = idx
		firstViewValues = append(firstViewValues, firstViewValue)
		repeatViewValues = append(repeatViewValues, repeatViewValue)
	}
	sort.Float64s(firstViewValues)
	sort.Float64s(repeatViewValues)

	var firstRunIdx, repeatRunIdx string
	if len(firstViewValues) == 1 {
		firstRunIdx = firstViewValueMap[firstViewValues[0]]
	} else {
		idx := len(firstViewValues) / 2
		if len(firstViewValues)%2 == 0 {
			idx++
		}
		firstRunIdx = firstViewValueMap[firstViewValues[idx]]
	}
	testRun.FirstView = rd.Runs[firstRunIdx].FirstView

	if len(repeatViewValues) == 1 {
		repeatRunIdx = repeatViewValueMap[repeatViewValues[0]]
	} else {
		idx := len(repeatViewValues) / 2
		if len(repeatViewValues)%2 == 0 {
			idx++
		}
		repeatRunIdx = repeatViewValueMap[repeatViewValues[idx]]
	}
	testRun.RepeatView = rd.Runs[repeatRunIdx].RepeatView

	return &testRun, nil
}

func (w *WebPageTest) GetTestResult(testID string) (*ResultData, error) {
	query := url.Values{}
	query.Add("test", testID)
	query.Add("requests", "0")
	query.Add("average", "0")
	query.Add("standard", "0")

	body, err := w.query("/jsonResult.php", query)
	if err != nil {
		return nil, err
	}
	var resultData *ResultData
	resultData, err = parseResultResponse(body)
	if err != nil {
		return nil, err
	}

	return resultData, nil
}

func parseResultResponse(rawResponse []byte) (*ResultData, error) {
	var err error
	var responose struct {
		StatusCode int             `json:"statusCode"`
		StatusText string          `json:"statusText"`
		Data       json.RawMessage `json:"data"`
	}

	if err = json.Unmarshal(rawResponse, &responose); err != nil {
		return nil, err
	}
	if responose.StatusCode != 200 {
		return nil, fmt.Errorf("Unexpected status %d: %v",
			responose.StatusCode, responose.StatusText)
	}

	var resultData ResultData
	if err = json.Unmarshal(responose.Data, &resultData); err != nil {
		return nil, err
	}

	// Unset value is "0"
	if resultData.Connectivity.RawPacketLossRate == nil || string(*resultData.Connectivity.RawPacketLossRate) == "\"0\"" {
		resultData.Connectivity.PacketLossRate = 0
	} else {
		resultData.Connectivity.PacketLossRate, _ = strconv.Atoi(string(*resultData.Connectivity.RawPacketLossRate))
	}

	return &resultData, nil
}
