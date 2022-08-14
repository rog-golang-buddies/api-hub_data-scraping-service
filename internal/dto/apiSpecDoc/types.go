package apiSpecDoc

// Type represents type of parsed API
type Type int

const (
	TypeOpenApi Type = iota
)

// SchemaType is the type of field.
type SchemaType string

const (
	Unknown    SchemaType = "UNKNOWN"
	NotDefined SchemaType = "NOT_DEFINED"
	Integer    SchemaType = "INTEGER"
	Boolean    SchemaType = "BOOLEAN"
	Number     SchemaType = "NUMBER"
	String     SchemaType = "STRING"
	Date       SchemaType = "DATE"
	Array      SchemaType = "ARRAY"  //Array type may contain single nested field with type to define full array type
	Map        SchemaType = "MAP"    //Map i.e. { "map_name": {"key1":"value1", "key2":"value2"}}
	OneOf      SchemaType = "ONE_OF" //OneOf is one of the different types (C union) (from nested fields)
	AnyOf      SchemaType = "ANY_OF" //AnyOf defines that result object can contain any set of sub-schemes
	AllOf      SchemaType = "ALL_OF" //AllOf defines that result object combines all listed objects/properties.
	Not        SchemaType = "NOT"    //Not represents type that can't be used. So it's possible to use any of types except "not"
	Object     SchemaType = "OBJECT" //Object represent object and contains set of fields inside
)

func ResolveSchemaType(strType string) SchemaType {
	switch strType {
	case "string":
		return String
	case "number":
		return Number
	case "integer":
		return Integer
	case "boolean":
		return Boolean
	case "array":
		return Array
	case "object":
		return Object
	case "":
		return NotDefined
	default:
		return Unknown
	}
}

// MethodType define type of action of the method
type MethodType string

const (
	MethodConnect MethodType = "CONNECT"
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

// ParameterType represents to what part of request parameter relates
type ParameterType string

const (
	ParameterQuery  ParameterType = "QUERY"
	ParameterHeader ParameterType = "HEADER"
	ParameterPath   ParameterType = "PATH"
	ParameterCookie ParameterType = "COOKIE"
)
