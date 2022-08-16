package fileresource

type AsdFileType int

const (
	Undefined AsdFileType = iota
	OpenApi
	YamlOpenApi
	JsonOpenAPI
)
