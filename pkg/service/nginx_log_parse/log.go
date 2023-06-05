package nginx_log_parse

import (
	"encoding/json"
)

type NginxLog struct {
	Msec                   string `json:"msec"`
	Connection             string `json:"connection"`
	ConnectionRequests     string `json:"connection_requests"`
	Pid                    string `json:"pid"`
	RequestID              string `json:"request_id"`
	RequestBody            string `json:"request_body"`
	RequestLength          string `json:"request_length"`
	RemoteAddr             string `json:"remote_addr"`
	RemoteUser             string `json:"remote_user"`
	RemotePort             string `json:"remote_port"`
	TimeLocal              string `json:"time_local"`
	TimeISO8601            string `json:"time_iso8601"`
	Request                string `json:"request"`
	RequestURI             string `json:"request_uri"`
	Args                   string `json:"args"`
	Status                 string `json:"status"`
	BodyBytesSent          string `json:"body_bytes_sent"`
	BytesSent              string `json:"bytes_sent"`
	HTTPReferer            string `json:"http_referer"`
	HTTPUserAgent          string `json:"http_user_agent"`
	HTTPXForwardedFor      string `json:"http_x_forwarded_for"`
	HTTPHost               string `json:"http_host"`
	ServerName             string `json:"server_name"`
	RequestTime            string `json:"request_time"`
	Upstream               string `json:"upstream"`
	UpstreamConnectTime    string `json:"upstream_connect_time"`
	UpstreamHeaderTime     string `json:"upstream_header_time"`
	UpstreamResponseTime   string `json:"upstream_response_time"`
	UpstreamResponseLength string `json:"upstream_response_length"`
	UpstreamCacheStatus    string `json:"upstream_cache_status"`
	SSLProtocol            string `json:"ssl_protocol"`
	SSLCipher              string `json:"ssl_cipher"`
	Scheme                 string `json:"scheme"`
	RequestMethod          string `json:"request_method"`
	ServerProtocol         string `json:"server_protocol"`
	Pipe                   string `json:"pipe"`
	GzipRatio              string `json:"gzip_ratio"`
	HTTPCfRay              string `json:"http_cf_ray"`
}

func newLogLine(line string) (*NginxLog, error) {
	log := &NginxLog{}
	err := json.Unmarshal([]byte(line), log)
	if err != nil {
		return nil, err
	}
	return log, nil
}
