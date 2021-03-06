package constants

const (
	HeaderContentTypeJson = "application/json"

	KeyPassword   = "1234567890"

	UriKeyCreate     = "/keys"
	UriAccountInfo   = "/auth/accounts/%v"           // format is /auth/accounts/{address}
	UriTransfer      = "/bank/accounts/%s/transfers" // format is /bank/accounts/{address}/transfers
	UriTxBroadcast   = "/tx/broadcast"

	// http status code
	StatusCodeOk       = 200
	StatusCodeConflict = 409

	//
	MockDefaultGas     = "20000"
	MockDefaultFee     = "5iris"
	Denom              = "iris"
)
