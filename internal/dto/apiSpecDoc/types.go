package apiSpecDoc

//Type represents type of parsed API
type Type int

const (
	TypeOpenApi Type = iota
)

//SchemaType is the type of field.
type SchemaType string

const (
	Integer SchemaType = "INTEGER"
	Number  SchemaType = "NUMBER"
	String  SchemaType = "STRING"
	Date    SchemaType = "DATE"
	Array   SchemaType = "ARRAY"  //Array type may contain single nested field with type to define full array type
	Map     SchemaType = "MAP"    //Map i.e. { "map_name": {"key1":"value1", "key2":"value2"}}
	OneOf   SchemaType = "ONE_OF" //OneOf is one of the different types (from nested fields)
	Object  SchemaType = "OBJECT" //Object represent object and contains set of fields inside
)

//MethodType define type of action of the method
type MethodType string

const (
	MethodGet     MethodType = "GET"
	MethodPut     MethodType = "PUT"
	MethodPost    MethodType = "POST"
	MethodDelete  MethodType = "DELETE"
	MethodOptions MethodType = "OPTIONS"
	MethodHead    MethodType = "HEAD"
	MethodPatch   MethodType = "PATCH"
	MethodTrace   MethodType = "TRACE"
	//Unary/bidirectional stream etc. for the rpc
)

//ParameterType represents to what part of request parameter relates
type ParameterType int

const (
	ParameterQuery ParameterType = iota
	ParameterHeader
	ParameterPath
	ParameterCookie
)
