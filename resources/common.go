package resources

type ServiceError struct {
	ErrorCode uint
	ErrorMsg  string
}

type ServiceResultData struct {
	Data interface{}
}

type ServiceResult struct {
	Code int
	ServiceResultData
	IsError bool
	Error   ServiceError
}
