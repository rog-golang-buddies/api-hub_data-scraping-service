package apiSpecDoc

// ApiSpecDoc represents full API Specification document
// with all required data to view and execute requests
// Propose to not store something like components at the initial stage - leave it as update to simplify logic
// So each ApiMethod need to keep all schema (some duplication, but we could update it faster with extending of the model)
type ApiSpecDoc struct {
	Title string

	Description string

	//Type of API definition
	//It's possible to be OpenApi, gRPC and others
	Type Type

	//Groups are just grouped lists of ApiMethod's , supposed to be displayed as folders
	//It relates to:
	//Tags - Open API specs
	//Services - RPCs
	Groups []*Group

	//Methods at the root level without groups
	Methods []*ApiMethod
}

// Group represents some grouping rule
// Tags for the open API
// Services for the gRPC
type Group struct {
	//Name of the group
	Name string

	//Description of the group
	Description string

	//Methods is a set of request methods related to the group
	Methods []*ApiMethod
}

// ApiMethod represents particular API method to call
type ApiMethod struct {
	//Path to send request
	//For example /open/api/info/
	//It can contain path parameters,
	//They can be expressed with some placeholders like /api/v1/item/{id}/someth/{someth}
	//This placeholder definitions need to be searched as Parameters with Parameter.In equals ParameterType.Path
	Path string

	//Name of API method
	Name string

	//Description of API method if exists
	Description string

	//Type represents type of request method
	Type MethodType

	//RequestBody is a description of ApiMethod request body, nil if no request body
	RequestBody *RequestBody

	//Servers represent available paths for requests with description.
	//In the case of swagger we need to calculate it taking the deepest nested servers definition
	Servers []Server

	//Parameters contains all possible request parameters
	//Including query, path, header and cookie parameters for the Open API
	//If Parameter.In is header - then it represents headers, if cookies - cookies,
	//so it's header and cookies full representation as well (not just a parametrisation)
	Parameters []Parameter

	//ExternalDoc represents link to external method documentation (if exists)
	//
	ExternalDoc *ExternalDoc
}

//Schema represents type or structure of request/response/metadata
//Also type, properties and description of the specific Schema field
//This is not the same as swagger schema!
//
//General purpose of the Schema is to provide description for the UI how to
//generate sample method body for example. So it purpose to provide information to view to form sample request.
//
//Schema focuses on description of json body structure (or parameter).
//Constraints supposed to be additional field in the schema.
/*
For example object:
{
	"int_val": 5,
	"arr": [5, 6, 7, 8],
	"obj": {
		"field1": "test",
		"field2": "test2"
	}
}
can be expressed as:
	dto.Schema{
		Type: dto.Object,
		Fields: []dto.Schema{
			{
				Key:  "int_val",
				Type: dto.Integer,
			},
			{
				Key: "arr",
				Type: dto.Array,
				Fields: []dto.Schema{
					{
						Type: dto.Integer,
					},
				},
			},
			{
				Key: "obj",
				Type: dto.Object,
				Fields: []dto.Schema{
					{
						Key: "field1",
						Type: dto.String,
					},
					{
						Key: "field2",
						Type: dto.String,
					},
				},
			},
		},
	}
*/
type Schema struct {
	//Key is a field name
	Key string

	//Type is a component field type
	//To provide some validation or specific input on UI
	Type SchemaType

	//Description of the field if exist
	Description string

	//Nested describe nested object/properties
	//If it's object and contain nested fields
	Fields []*Schema
}

// MediaTypeObject represents schema for the different media types
// i.e. "application/json" and etc.
type MediaTypeObject struct {
	Schema *Schema
}

// RequestBody is a representation of request body
type RequestBody struct {
	Description string

	//Content represents request object for the different media types
	Content map[string]*MediaTypeObject

	//Required define is request body required
	Required bool
}

// Server represents server description
// To use in the address line of view
// https://swagger.io/specification/#server-object
type Server struct {
	//Url to access server (no templates, real url)
	Url string

	//Description of the particular server (relates to Open API spec)
	Description string
}

// Parameter is abstraction of additional data to request
// It is headers for the REST API and metadata for the gRPC
type Parameter struct {
	Name string

	//In represents what part of request parameter relates
	//https://swagger.io/specification/#parameter-object
	In ParameterType

	Description string

	Schema Schema

	//Required defines is parameter required
	Required bool
}

// ExternalDoc may be available for the Open API
// And contain link to description of request method
// Maybe useful also for debugging - to find failed API description rapidly
type ExternalDoc struct {
	Description string

	//Url to external documentation about resource
	Url string
}
