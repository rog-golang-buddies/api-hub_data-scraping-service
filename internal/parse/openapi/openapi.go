package openapi

import (
	"context"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/dto/apiSpecDoc"
)

func parseOpenAPI(ctx context.Context, content []byte) (*openapi3.T, error) {
	loader := openapi3.Loader{Context: ctx, IsExternalRefsAllowed: false}
	doc, err := loader.LoadFromData(content)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func openapiToApiSpec(openapi *openapi3.T) *apiSpecDoc.ApiSpecDoc {
	asd := apiSpecDoc.ApiSpecDoc{
		Type:    apiSpecDoc.TypeOpenApi,
		Methods: make([]*apiSpecDoc.ApiMethod, 0),
	}

	groups := tagToGroup(openapi.Tags)
	groupMap := make(map[string]*apiSpecDoc.Group)
	for _, group := range groups {
		groupMap[group.Name] = group
	}

	asd.Groups = groups

	populateMethods(&asd, openapi.Paths)
	return &asd
}

// tagToGroup creates group with empty methods
func tagToGroup(tags []*openapi3.Tag) []*apiSpecDoc.Group {
	groups := make([]*apiSpecDoc.Group, 0, len(tags))
	if tags == nil {
		return groups
	}
	for _, tag := range tags {
		group := new(apiSpecDoc.Group)
		group.Name = tag.Name
		group.Description = tag.Description
		group.Methods = make([]*apiSpecDoc.ApiMethod, 0)
		groups = append(groups, group)
	}
	return groups
}

func populateMethods(asd *apiSpecDoc.ApiSpecDoc, paths openapi3.Paths) {
	groupMap := make(map[string]*apiSpecDoc.Group)
	for _, group := range asd.Groups {
		groupMap[group.Name] = group
	}
	for url, path := range paths {
		for httpMethod, operation := range path.Operations() {
			method := new(apiSpecDoc.ApiMethod)
			method.Path = url
			method.Description = operation.Description
			method.Type = apiSpecDoc.MethodType(httpMethod)
			if operation.RequestBody != nil {
				method.RequestBody = convertBody(operation.RequestBody.Value)
			}
			//TODO too much nested loops/if's extract to methods/refactor
			if operation.Tags != nil && len(operation.Tags) > 0 {
				addedToAnyGroup := false
				for _, tag := range operation.Tags {
					if group, ok := groupMap[tag]; ok {
						group.Methods = append(group.Methods, method)
						addedToAnyGroup = true
					} else {
						//TODO log record here. WARN or ERROR ?
					}
				}
				if !addedToAnyGroup {
					asd.Methods = append(asd.Methods, method)
				}
			} else {
				asd.Methods = append(asd.Methods, method)
			}
		}
	}
}

func convertBody(body *openapi3.RequestBody) *apiSpecDoc.RequestBody {
	specBody := new(apiSpecDoc.RequestBody)
	specBody.Description = body.Description
	specContent := make(map[string]*apiSpecDoc.MediaTypeObject)
	for cType, content := range body.Content {
		if content.Schema == nil || content.Schema.Value == nil {
			continue
		}
		specContent[cType] = &apiSpecDoc.MediaTypeObject{Schema: convertSchema("", content.Schema.Value)}
	}
	specBody.Content = specContent
	return specBody
}

func convertSchema(key string, schema *openapi3.Schema) *apiSpecDoc.Schema {
	resSchema := new(apiSpecDoc.Schema)
	if schema == nil {
		return resSchema
	}
	resSchema.Key = key
	resSchema.Description = schema.Description
	resSchema.Type = apiSpecDoc.ResolveSchemaType(schema.Type)
	resSchema.Fields = make([]*apiSpecDoc.Schema, 0)
	switch resSchema.Type {
	case apiSpecDoc.Object:
		//If the type is an Object it can be an Object or Map. The map represents additional properties - can be only one of Object/Map
		if schema.Properties != nil {
			for pKey, prop := range schema.Properties {
				resSchema.Fields = append(resSchema.Fields, convertSchema(pKey, prop.Value))
			}
		} else if schema.AdditionalProperties != nil {
			resSchema.Type = apiSpecDoc.Map
			resSchema.Fields = append(resSchema.Fields, convertSchema("", schema.AdditionalProperties.Value))
		}
	case apiSpecDoc.Array:
		if schema.Items != nil {
			resSchema.Fields = append(resSchema.Fields, convertSchema("", schema.Items.Value))
		}
	case apiSpecDoc.NotDefined:
		//If type is not defined it means that here one of the "combine" types is used. So need to check them all
		switch true {
		case schema.OneOf != nil && len(schema.OneOf) > 0:
			resSchema.Type = apiSpecDoc.OneOf
			for _, sch := range schema.OneOf {
				resSchema.Fields = append(resSchema.Fields, convertSchema("", sch.Value))
			}
		case schema.AnyOf != nil && len(schema.AnyOf) > 0:
			resSchema.Type = apiSpecDoc.AnyOf
			for _, sch := range schema.AnyOf {
				resSchema.Fields = append(resSchema.Fields, convertSchema("", sch.Value))
			}
		case schema.AllOf != nil && len(schema.AllOf) > 0:
			resSchema.Type = apiSpecDoc.AllOf
			for _, sch := range schema.AllOf {
				resSchema.Fields = append(resSchema.Fields, convertSchema("", sch.Value))
			}
		case schema.Not != nil:
			resSchema.Type = apiSpecDoc.Not
			resSchema.Fields = append(resSchema.Fields, convertSchema("", schema.Not.Value))
		}
	}

	return resSchema
}
