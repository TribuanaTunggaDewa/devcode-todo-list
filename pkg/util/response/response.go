package response

type Meta struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type responseHelper struct {
	Error   errorHelper
	Success successHelper
}

var Constant responseHelper = responseHelper{
	Error:   errorConstant,
	Success: successConstant,
}
