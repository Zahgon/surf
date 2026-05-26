package chrome

import (
	"github.com/enetx/g"
	"github.com/enetx/http"
	"github.com/enetx/surf/header"
	"github.com/enetx/surf/profiles"
)

// --- Header order maps -------------------------------------------------------

var headerOrderDesktop = g.Map[string, g.Slice[string]]{
	http.MethodGet: {
		":method",
		":authority",
		":scheme",
		":path",
		header.SEC_CH_UA,
		header.SEC_CH_UA_MOBILE,
		header.SEC_CH_UA_PLATFORM,
		header.AUTHORIZATION,
		header.UPGRADE_INSECURE_REQUESTS,
		header.USER_AGENT,
		header.ACCEPT,
		header.SEC_FETCH_SITE,
		header.SEC_FETCH_MODE,
		header.SEC_FETCH_USER,
		header.SEC_FETCH_DEST,
		header.REFERER,
		header.ACCEPT_ENCODING,
		header.ACCEPT_LANGUAGE,
		header.COOKIE,
		header.PRIORITY,
	},

	http.MethodGet + "http3": {
		":method",
		":authority",
		":scheme",
		":path",
		header.SEC_CH_UA,
		header.SEC_CH_UA_MOBILE,
		header.SEC_CH_UA_PLATFORM,
		header.AUTHORIZATION,
		header.UPGRADE_INSECURE_REQUESTS,
		header.USER_AGENT,
		header.ACCEPT,
		header.SEC_FETCH_SITE,
		header.SEC_FETCH_MODE,
		header.SEC_FETCH_USER,
		header.SEC_FETCH_DEST,
		header.REFERER,
		header.ACCEPT_ENCODING,
		header.ACCEPT_LANGUAGE,
		header.COOKIE,
		header.PRIORITY,
	},

	http.MethodPost: {
		":method",
		":authority",
		":scheme",
		":path",
		header.CONTENT_LENGTH,
		header.PRAGMA,
		header.CACHE_CONTROL,
		header.SEC_CH_UA_PLATFORM,
		header.AUTHORIZATION,
		header.USER_AGENT,
		header.SEC_CH_UA,
		header.CONTENT_TYPE,
		header.SEC_CH_UA_MOBILE,
		header.ACCEPT,
		header.ORIGIN,
		header.SEC_FETCH_SITE,
		header.SEC_FETCH_MODE,
		header.SEC_FETCH_DEST,
		header.REFERER,
		header.ACCEPT_ENCODING,
		header.ACCEPT_LANGUAGE,
		header.COOKIE,
		header.PRIORITY,
	},

	http.MethodPost + "http3": {
		":method",
		":authority",
		":scheme",
		":path",
		header.CONTENT_LENGTH,
		header.PRAGMA,
		header.CACHE_CONTROL,
		header.SEC_CH_UA_PLATFORM,
		header.AUTHORIZATION,
		header.USER_AGENT,
		header.SEC_CH_UA,
		header.CONTENT_TYPE,
		header.SEC_CH_UA_MOBILE,
		header.ACCEPT,
		header.ORIGIN,
		header.SEC_FETCH_SITE,
		header.SEC_FETCH_MODE,
		header.SEC_FETCH_DEST,
		header.REFERER,
		header.ACCEPT_ENCODING,
		header.ACCEPT_LANGUAGE,
		header.COOKIE,
		header.PRIORITY,
	},
}

// headerOrderMobile is a placeholder mobile variant. On the day real Chrome Android header
// ordering is observed, this map is the single point to substitute it without touching desktop.
// The literal is a physical copy of headerOrderDesktop so the two maps can diverge independently.
var headerOrderMobile = g.Map[string, g.Slice[string]]{
	http.MethodGet: {
		":method",
		":authority",
		":scheme",
		":path",
		header.SEC_CH_UA,
		header.SEC_CH_UA_MOBILE,
		header.SEC_CH_UA_PLATFORM,
		header.AUTHORIZATION,
		header.UPGRADE_INSECURE_REQUESTS,
		header.USER_AGENT,
		header.ACCEPT,
		header.SEC_FETCH_SITE,
		header.SEC_FETCH_MODE,
		header.SEC_FETCH_USER,
		header.SEC_FETCH_DEST,
		header.REFERER,
		header.ACCEPT_ENCODING,
		header.ACCEPT_LANGUAGE,
		header.COOKIE,
		header.PRIORITY,
	},

	http.MethodGet + "http3": {
		":method",
		":authority",
		":scheme",
		":path",
		header.SEC_CH_UA,
		header.SEC_CH_UA_MOBILE,
		header.SEC_CH_UA_PLATFORM,
		header.AUTHORIZATION,
		header.UPGRADE_INSECURE_REQUESTS,
		header.USER_AGENT,
		header.ACCEPT,
		header.SEC_FETCH_SITE,
		header.SEC_FETCH_MODE,
		header.SEC_FETCH_USER,
		header.SEC_FETCH_DEST,
		header.REFERER,
		header.ACCEPT_ENCODING,
		header.ACCEPT_LANGUAGE,
		header.COOKIE,
		header.PRIORITY,
	},

	http.MethodPost: {
		":method",
		":authority",
		":scheme",
		":path",
		header.CONTENT_LENGTH,
		header.PRAGMA,
		header.CACHE_CONTROL,
		header.SEC_CH_UA_PLATFORM,
		header.AUTHORIZATION,
		header.USER_AGENT,
		header.SEC_CH_UA,
		header.CONTENT_TYPE,
		header.SEC_CH_UA_MOBILE,
		header.ACCEPT,
		header.ORIGIN,
		header.SEC_FETCH_SITE,
		header.SEC_FETCH_MODE,
		header.SEC_FETCH_DEST,
		header.REFERER,
		header.ACCEPT_ENCODING,
		header.ACCEPT_LANGUAGE,
		header.COOKIE,
		header.PRIORITY,
	},

	http.MethodPost + "http3": {
		":method",
		":authority",
		":scheme",
		":path",
		header.CONTENT_LENGTH,
		header.PRAGMA,
		header.CACHE_CONTROL,
		header.SEC_CH_UA_PLATFORM,
		header.AUTHORIZATION,
		header.USER_AGENT,
		header.SEC_CH_UA,
		header.CONTENT_TYPE,
		header.SEC_CH_UA_MOBILE,
		header.ACCEPT,
		header.ORIGIN,
		header.SEC_FETCH_SITE,
		header.SEC_FETCH_MODE,
		header.SEC_FETCH_DEST,
		header.REFERER,
		header.ACCEPT_ENCODING,
		header.ACCEPT_LANGUAGE,
		header.COOKIE,
		header.PRIORITY,
	},
}

var headerCache = profiles.NewHeaderCache(headerOrderDesktop, headerOrderMobile)

// --- Static header set (Variant.BuildHeaders) --------------------------------

// buildHeadersDesktop constructs the desktop Chrome 145 request header set.
func buildHeadersDesktop(os profiles.OSKey) *g.MapOrd[g.String, g.String] {
	_ = "STUB: not implemented"
	return nil
}

// buildHeadersMobile constructs the placeholder mobile Chrome 145 request header set.
// On the day real Chrome Android header set diverges from desktop (different Accept-Encoding,
// shorter sec-ch-ua, different ordering / inserts), replace this body — it is the single point
// of substitution for the entire mobile header set.
func buildHeadersMobile(os profiles.OSKey) *g.MapOrd[g.String, g.String] {
	_ = "STUB: not implemented"
	return nil
}

// --- Per-request header pipeline (Variant.Headers) ---------------------------

// DesktopApplier applies the desktop Chrome request-header pipeline. Wired into chrome.Desktop.
var DesktopApplier = profiles.NewApplier(insertDesktopHeaders, insertDesktopHeaders, headerCache, false)

// MobileApplier applies the mobile Chrome request-header pipeline. Wired into chrome.Mobile.
var MobileApplier = profiles.NewApplier(insertMobileHeaders, insertMobileHeaders, headerCache, true)

func insertDesktopHeaders[T ~string](headers *g.MapOrd[T, T], method string) {
	_ = "STUB: not implemented"
	return
}

// insertMobileHeaders is a placeholder mobile variant. On the day the real Chrome Android header
// inserts diverge from desktop, this function is the single point to substitute them.
func insertMobileHeaders[T ~string](headers *g.MapOrd[T, T], method string) {
	_ = "STUB: not implemented"
	return
}
