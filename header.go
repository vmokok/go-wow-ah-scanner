package ahsc

import (
	"net/http"
	"time"
)

const (
	HeaderAccept                  = "Accept"
	HeaderBattleNetNamespace      = "Battlenet-Namespace"
	HeaderIfModifiedSince         = "If-Modified-Since"
	HeaderLastModified            = "Last-Modified"
	HeaderCacheControl            = "Cache-Control"
	HeaderXTraceTraceid           = "X-Trace-Traceid"
	HeaderBattlenetSchema         = "Battlenet-Schema"
	HeaderXFrameOptions           = "X-Frame-Options"
	HeaderXContentTypeOptions     = "X-Content-Type-Options"
	HeaderDate                    = "Date"
	HeaderBattlenetSchemaRevision = "Battlenet-Schema-Revision"
	HeaderVary                    = "Vary"
	HeaderXTraceSpanid            = "X-Trace-Spanid"
	HeaderXTraceParentspanid      = "X-Trace-Parentspanid"
	HeaderContentType             = "Content-Type"
	HeaderConnection              = "Connection"
	HeaderBlizzardTokenExpires    = "Blizzard-Token-Expires"
	HeaderServer                  = "Server"
)

type Header struct {
	LastModified            time.Time
	CacheControl            string
	XTraceTraceid           string
	BattlenetSchema         string
	XFrameOptions           string
	XContentTypeOptions     string
	Date                    time.Time
	BattlenetSchemaRevision string
	Vary                    string
	XTraceSpanid            string
	XTraceParentspanid      string
	ContentType             string
	Connection              string
	BattlenetNamespace      string
	BlizzardTokenExpires    string
	Server                  string
}

func mapHeader(responseHeader http.Header) *Header {
	header := &Header{
		CacheControl:            responseHeader.Get(HeaderCacheControl),
		XTraceTraceid:           responseHeader.Get(HeaderXTraceTraceid),
		BattlenetSchema:         responseHeader.Get(HeaderBattlenetSchema),
		XFrameOptions:           responseHeader.Get(HeaderXFrameOptions),
		XContentTypeOptions:     responseHeader.Get(HeaderXContentTypeOptions),
		BattlenetSchemaRevision: responseHeader.Get(HeaderBattlenetSchemaRevision),
		Vary:                    responseHeader.Get(HeaderVary),
		XTraceSpanid:            responseHeader.Get(HeaderXTraceSpanid),
		XTraceParentspanid:      responseHeader.Get(HeaderXTraceParentspanid),
		ContentType:             responseHeader.Get(HeaderContentType),
		Connection:              responseHeader.Get(HeaderConnection),
		BattlenetNamespace:      responseHeader.Get(HeaderBattleNetNamespace),
		BlizzardTokenExpires:    responseHeader.Get(HeaderBlizzardTokenExpires),
		Server:                  responseHeader.Get(HeaderServer),
	}
	if responseHeader.Get(HeaderDate) != "" {
		var err error
		header.Date, err = time.Parse(http.TimeFormat, responseHeader.Get(HeaderDate))
		if err != nil {
			header.Date = time.Time{}
		}
	}
	if responseHeader.Get(HeaderLastModified) != "" {
		var err error
		header.LastModified, err = time.Parse(http.TimeFormat, responseHeader.Get(HeaderLastModified))
		if err != nil {
			header.LastModified = time.Time{}
		}
	}

	return header
}
