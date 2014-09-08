package api

// api error codes
const (
	API_INVALID_ACTION          = 1000
	API_INVALID_TARGET          = 1001
	API_INVALID_DATA            = 1002
	API_INVALID_VERSION         = 1003
	API_LOCATION_DOES_NOT_EXIST = 1004
	API_INVALID_REQUEST         = 1005
)

type Api_Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}