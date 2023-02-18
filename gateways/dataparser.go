package gateways

type DataParser interface {
	ParseData(jsonData []byte) (error, map[string]string)
}
