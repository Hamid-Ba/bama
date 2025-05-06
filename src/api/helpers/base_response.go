package helpers

import "github.com/Hamid-Ba/bama/api/validators"

type BaseResponse struct {
	Result           any                           `json:"result"`
	Success          bool                          `json:"success"`
	ResultCode       int                           `json:"result_code"`
	ValidationErrors *[]validators.ValidationError `json:"validation_errors"`
	Error            any                           `json: "error"`
}

func GenerateBaseResponse(result any, success bool, result_code int) *BaseResponse {
	return &BaseResponse{
		Result:     result,
		Success:    success,
		ResultCode: result_code,
	}
}

func GenerateBaseResponseWithError(result any, success bool, result_code int, err error) *BaseResponse {
	return &BaseResponse{
		Result:     result,
		Success:    success,
		ResultCode: result_code,
		Error:      err.Error(),
	}
}

func GenerateBaseResponseWithAnyError(result any, success bool, result_code int, err any) *BaseResponse {
	return &BaseResponse{
		Result:     result,
		Success:    success,
		ResultCode: result_code,
		Error:      err,
	}
}

func GenerateBaseResponseWithValidationError(result any, success bool, result_code int, err error) *BaseResponse {
	return &BaseResponse{
		Result:           result,
		Success:          success,
		ResultCode:       result_code,
		ValidationErrors: validators.GetValidationErrors(err),
	}
}
