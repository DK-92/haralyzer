package structs

/*
  Describes the HAR model as per W3C specification
  https://w3c.github.io/web-performance/specs/HAR/Overview.html
*/

type MainLog struct {
	Log Log `json:"log"`
}

// Log This object represents the root of the exported data. This object MUST be present and its name MUST be "log". The object contains the following name/value pairs:
type Log struct {
	Version string  `json:"version"`
	Creator Creator `json:"creator"`
	Browser Browser `json:"browser,omitempty"`
	Pages   []Page  `json:"pages,omitempty"`
	Entries []Entry `json:"entries,omitempty"`
}

// Creator This object contains information about the log creator application and contains the following name/value pairs:
type Creator struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Comment string `json:"comment,omitempty"`
}

// Browser This object contains information about the browser that created the log and contains the following name/value pairs:
type Browser struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Comment string `json:"comment,omitempty"`
}

// Page This object represents list of exported pages.
type Page struct {
	StartedDateTime string     `json:"startedDateTime"`
	ID              string     `json:"id"`
	Title           string     `json:"title"`
	PageTimings     PageTiming `json:"pageTimings"`
	Comment         string     `json:"comment,omitempty"`
}

// PageTiming This object describes timings for various events (states) fired during the page load. All times are specified in milliseconds. If a time info is not available appropriate field is set to -1.
type PageTiming struct {
	OnContentLoad int    `json:"onContentLoad,omitempty"`
	OnLoad        int    `json:"onLoad,omitempty"`
	Comment       string `json:"comment,omitempty"`
}

// Entry This object represents an array with all exported HTTP requests. Sorting entries by startedDateTime (starting from the oldest) is preferred way how to export data since it can make importing faster. However the reader application should always make sure the array is sorted (if required for the import).
type Entry struct {
	Pageref         string   `json:"pageref,omitempty"`
	StartedDateTime string   `json:"startedDateTime"`
	Time            int      `json:"time"`
	Request         Request  `json:"request"`
	Response        Response `json:"response"`
	Cache           Cache    `json:"cache"`
	Timings         Timing   `json:"timings"`
	ServerIPAddress string   `json:"serverIPAddress,omitempty"`
	Connection      string   `json:"connection,omitempty"`
	Comment         string   `json:"comment,omitempty"`
}

// Request This object contains detailed info about performed request.
type Request struct {
	Method      string        `json:"method"`
	URL         string        `json:"url"`
	HTTPVersion string        `json:"httpVersion"`
	Cookies     []Cookie      `json:"cookies"`
	Headers     []Header      `json:"headers"`
	QueryString []QueryString `json:"queryString"`
	PostData    PostData      `json:"postData,omitempty"`
	HeaderSize  int           `json:"headersSize"`
	BodySize    int           `json:"bodySize"`
	Comment     string        `json:"comment,omitempty"`
}

// Response This object contains detailed info about the response.
type Response struct {
	Status      int      `json:"status"`
	StatusText  string   `json:"statusText"`
	HTTPVersion string   `json:"httpVersion"`
	Cookies     []Cookie `json:"cookies"`
	Headers     []Header `json:"headers"`
	Content     Content  `json:"content"`
	HeaderSize  int      `json:"headersSize"`
	BodySize    int      `json:"bodySize"`
	Comment     string   `json:"comment,omitempty"`
}

// Cookie This object contains list of all cookies (used in <request> and <response> objects).
type Cookie struct {
	Name     string `json:"name"`
	Value    string `json:"value"`
	Path     string `json:"path,omitempty"`
	Domain   string `json:"domain,omitempty"`
	Expires  string `json:"expires,omitempty"`
	HTTPOnly bool   `json:"httpOnly"`
	Secure   bool   `json:"secure"`
	Comment  string `json:"comment,omitempty"`
}

// Header This object contains list of all headers (used in <request> and <response> objects).
type Header struct {
	Name    string `json:"name"`
	Value   string `json:"value"`
	Comment string `json:"comment"`
}

// QueryString This object contains list of all parameters & values parsed from a query string, if any (embedded in <request> object).
type QueryString struct {
	Name    string `json:"name"`
	Value   string `json:"value"`
	Comment string `json:"comment"`
}

// PostData This object describes posted data, if any (embedded in <request> object).
type PostData struct {
	MimeType string  `json:"mimeType"`
	Params   []Param `json:"params"`
	Text     string  `json:"text"`
	Comment  string  `json:"comment,omitempty"`
}

// Param List of posted parameters, if any (embedded in <postData> object).
type Param struct {
	Name        string `json:"name"`
	Value       string `json:"value,omitempty"`
	FileName    string `json:"fileName,omitempty"`
	ContentType string `json:"contentType,omitempty"`
	Comment     string `json:"comment,omitempty"`
}

// Content This object describes details about response content (embedded in <response> object).
type Content struct {
	Size        int    `json:"size"`
	Compression int    `json:"compression,omitempty"`
	MimeType    string `json:"mimeType"`
	Text        string `json:"text,omitempty"`
	Encoding    string `json:"encoding"`
	Comment     string `json:"comment,omitempty"`
}

// Cache This objects contains info about a request coming from browser cache.
type Cache struct {
	BeforeRequest BeforeAfterRequest `json:"beforeRequest"`
	AfterRequest  BeforeAfterRequest `json:"afterRequest"`
	Comment       string             `json:"comment,omitempty"`
}

// BeforeAfterRequest Both beforeRequest and afterRequest object share the following structure.
type BeforeAfterRequest struct {
	Expires    string `json:"expires,omitempty"`
	LastAccess string `json:"lastAccess"`
	ETag       string `json:"eTag"`
	HitCount   int    `json:"hitCount"`
	Comment    string `json:"comment,omitempty"`
}

// Timing This object describes various phases within request-response round trip. All times are specified in milliseconds.
type Timing struct {
	Blocked int    `json:"blocked,omitempty"`
	DNS     int    `json:"dns,omitempty"`
	Connect int    `json:"connect,omitempty"`
	Send    int    `json:"send"`
	Wait    int    `json:"wait"`
	Receive int    `json:"receive"`
	SSL     int    `json:"ssl"`
	Comment string `json:"comment,omitempty"`
}
