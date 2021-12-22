package tckimlik

var Config = configuration{
	INVALIDLENGTH: "Invalid length",
	INVALIDIDENTITY: "Invalid TC Identity Number",
	INVALIDIDENTITYVALIDATION: "Invalid Identity Validation",
	INVALIDSUM: "Invalid Sum Of Identity",
}

type configuration struct {
	INVALIDLENGTH string
	INVALIDIDENTITY string
	INVALIDIDENTITYVALIDATION string
	INVALIDSUM string
}
