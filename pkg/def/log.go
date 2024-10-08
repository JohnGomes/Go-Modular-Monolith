package def

// Log field names.
const (
	LogServer   = "server"   // "OpenAPI", "gRPC", "Prometheus metrics", etc.
	LogRemoteIP = "remoteIP" // IP address.
	LogAddr     = "addr"     // host:port.
	LogHost     = "host"     // DNS hostname or IPv4/IPv6 address.
	LogPort     = "port"     // TCP/UDP port number.
	LogFunc     = "func"     // RPC/event handler method name, REST resource path.
	LogUserName = "userName"
	LogGRPCCode = "grpcCode"
)
