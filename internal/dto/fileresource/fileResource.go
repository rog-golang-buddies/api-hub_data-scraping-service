package fileresource

// FileResource representation of file resource
type FileResource struct {
	//File name if exists, else empty
	Name string

	//Original link to file
	Link string

	//File content
	Content []byte

	//Type of the API specification file (json/yaml ...)
	Type AsdFileType
}
