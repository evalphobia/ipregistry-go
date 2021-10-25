package ipregistry

type ErrData struct {
	StatusCode int    `json:"status_code"`
	Code       string `json:"code"`
	Message    string `json:"message"`
	Resolution string `json:"resolution"`
}

func (r *ErrData) SetStatusCode(code int) {
	r.StatusCode = code
}

func (r ErrData) HasError() bool {
	return r.StatusCode >= 400
}

// IPListResponse is a response data of Batch IP Lookup API.
type IPListResponse struct {
	ErrData

	Results []IPResponse `json:"results"`
}

// IPResponse is a response data of Single IP Lookup API.
type IPResponse struct {
	ErrData

	IP         string     `json:"ip"`
	Type       string     `json:"type"`
	Hostname   string     `json:"hostname"`
	Carrier    Carrier    `json:"carrier"`
	Connection Connection `json:"connection"`
	Currency   Currency   `json:"currency"`
	Location   Location   `json:"location"`
	Security   Security   `json:"security"`
	TimeZone   TimeZone   `json:"time_zone"`
}

type Carrier struct {
	Name string `json:"name"`
	MCC  string `json:"mcc"`
	MNC  string `json:"mnc"`
}

type Connection struct {
	ASN          int64  `json:"asn"`
	Domain       string `json:"domain"`
	Organization string `json:"organization"`
	Route        string `json:"route"`
	Type         string `json:"type"`
}

type Currency struct {
	Code         string `json:"code"`
	Name         string `json:"name"`
	NameNative   string `json:"name_native"`
	Plural       string `json:"plural"`
	PluralNative string `json:"plural_native"`
	Symbol       string `json:"symbol"`
	SymbolNative string `json:"symbol_native"`
	Format       Format `json:"format"`
}

type Format struct {
	Negative CurrencyFormat `json:"negative"`
	Positive CurrencyFormat `json:"positive"`
}

type CurrencyFormat struct {
	Prefix string `json:"prefix"`
	Suffix string `json:"suffix"`
}

type Location struct {
	Continent Continent `json:"continent"`
	Country   Country   `json:"country"`
	Region    Region    `json:"region"`
	City      string    `json:"city"`
	Postal    string    `json:"postal"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Language  Language  `json:"language"`
	InEU      bool      `json:"in_eu"`
}

type Continent struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type Country struct {
	Area              int64      `json:"area"`
	Borders           []string   `json:"borders"`
	CallingCode       string     `json:"calling_code"`
	Capital           string     `json:"capital"`
	Code              string     `json:"code"`
	Name              string     `json:"name"`
	Population        int64      `json:"population"`
	PopulationDensity float64    `json:"population_density"`
	Flag              Flag       `json:"flag"`
	Languages         []Language `json:"languages"`
	TLD               string     `json:"tld"`
}

type Flag struct {
	Emoji        string `json:"emoji"`
	EmojiUnicode string `json:"emoji_unicode"`
	EmojiTwo     string `json:"emojitwo"`
	Noto         string `json:"noto"`
	Twemoji      string `json:"twemoji"`
	Wikimedia    string `json:"wikimedia"`
}

type Language struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Native string `json:"native"`
}

type Region struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type Security struct {
	IsBogon         bool `json:"is_bogon"`
	IsCloudProvider bool `json:"is_cloud_provider"`
	IsTor           bool `json:"is_tor"`
	IsTorExit       bool `json:"is_tor_exit"`
	IsProxy         bool `json:"is_proxy"`
	IsAnonymous     bool `json:"is_anonymous"`
	IsAbuser        bool `json:"is_abuser"`
	IsAttacker      bool `json:"is_attacker"`
	IsThreat        bool `json:"is_threat"`
}

type TimeZone struct {
	ID               string `json:"id"`
	Abbreviation     string `json:"abbreviation"`
	CurrentTime      string `json:"current_time"`
	Name             string `json:"name"`
	Offset           int64  `json:"offset"`
	InDaylightSaving bool   `json:"in_daylight_saving"`
}

// UserAgentResponse is a response data of User-Agent Parsing API.
type UserAgentResponse struct {
	ErrData

	Results []UserAgent `json:"results"`
}

type UserAgent struct {
	Header       string `json:"header"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Version      string `json:"version"`
	VersionMajor string `json:"version_major"`
	Device       Device `json:"device"`
	Engine       Engine `json:"engine"`
	OS           OS     `json:"os"`
}

type Device struct {
	Brand string `json:"brand"`
	Name  string `json:"name"`
	Type  string `json:"type"`
}

type Engine struct {
	Name         string `json:"name"`
	Type         string `json:"type"`
	Version      string `json:"version"`
	VersionMajor string `json:"version_major"`
}

type OS struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Version string `json:"version"`
}
