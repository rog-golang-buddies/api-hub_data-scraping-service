package model

type AsdFileType int

const (
	Undefined AsdFileType = iota
	YamlOpenApi
	JsonOpenAPI
)
