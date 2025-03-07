package v3

type Contact struct {
	Email string `json:"email,omitempty"`
}

type License struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}

type Info struct {
	Title          string   `json:"title,omitempty"`
	Description    string   `json:"description,omitempty"`
	TermsOfService string   `json:"termsOfService,omitempty"`
	Contact        *Contact `json:"contact,omitempty"`
	License        *License `json:"license,omitempty"`
	Version        string   `json:"version,omitempty"`
}

type Server struct {
	Url string `json:"url,omitempty"`
}

type ExternalDocs struct {
	Description string `json:"description,omitempty"`
	Url         string `json:"url,omitempty"`
}

type Tag struct {
	Name         string        `json:"name,omitempty"`
	Description  string        `json:"description,omitempty"`
	ExternalDocs *ExternalDocs `json:"externalDocs,omitempty"`
}

type In string

const (
	InQuery In = "query"
	// TODO
	InPath In = "path"
	// TODO
	InHeader In = "header"
	// TODO
	InCookie In = "cookie"
)

type Example struct {
	// TODO
}

type Encoding struct {
	// TODO
}

type MediaType struct {
	Schema   *Schema             `json:"schema,omitempty"`
	Example  interface{}         `json:"example,omitempty"`
	Examples map[string]Example  `json:"examples,omitempty"`
	Encoding map[string]Encoding `json:"encoding,omitempty"`
}

type Content struct {
	TextPlain *MediaType `json:"text/plain,omitempty"`
	Json      *MediaType `json:"application/json,omitempty"`
	FormUrl   *MediaType `json:"application/x-www-form-urlencoded,omitempty"`
	Stream    *MediaType `json:"application/octet-stream,omitempty"`
	FormData  *MediaType `json:"multipart/form-data,omitempty"`
	Default   *MediaType `json:"*/*,omitempty"`
}

type Parameter struct {
	Name            string      `json:"name,omitempty"`
	In              In          `json:"in,omitempty"`
	Description     string      `json:"description,omitempty"`
	Required        bool        `json:"required,omitempty"`
	Deprecated      bool        `json:"deprecated,omitempty"`
	Example         interface{} `json:"example,omitempty"`
	Schema          *Schema     `json:"schema,omitempty"`
	Style           string      `json:"style,omitempty"`
	Explode         bool        `json:"explode,omitempty"`
	AllowReserved   bool        `json:"allowReserved,omitempty"`
	Content         *Content    `json:"content,omitempty"`
	AllowEmptyValue bool        `json:"allowEmptyValue,omitempty"`
}

type RequestBody struct {
	Description string   `json:"description,omitempty"`
	Content     *Content `json:"content,omitempty"`
	Required    bool     `json:"required,omitempty"`
	Ref         string   `json:"$ref,omitempty"`
}

type Header struct {
	Ref         string      `json:"$ref,omitempty"`
	Description string      `json:"description,omitempty"`
	Required    bool        `json:"required,omitempty"`
	Deprecated  bool        `json:"deprecated,omitempty"`
	Example     interface{} `json:"example,omitempty"`
	Schema      *Schema     `json:"schema,omitempty"`
}

type Link struct {
	// TODO
}

type Response struct {
	Description string   `json:"description,omitempty"`
	Content     *Content `json:"content,omitempty"`
	// TODO
	Headers map[string]Header `json:"headers,omitempty"`
	Links   map[string]Link   `json:"links,omitempty"`
	Ref     string            `json:"$ref,omitempty"`
}

type Responses struct {
	Resp200 *Response `json:"200,omitempty"`
	Resp400 *Response `json:"400,omitempty"`
	Resp401 *Response `json:"401,omitempty"`
	Resp403 *Response `json:"403,omitempty"`
	Resp404 *Response `json:"404,omitempty"`
	Resp405 *Response `json:"405,omitempty"`
	Default *Response `json:"default,omitempty"`
}

type Callback struct {
	// TODO
}

type Security struct {
	// TODO
}

type Operation struct {
	Tags         []string            `json:"tags,omitempty"`
	Summary      string              `json:"summary,omitempty"`
	Description  string              `json:"description,omitempty"`
	OperationId  string              `json:"operationId,omitempty"`
	Parameters   []Parameter         `json:"parameters,omitempty"`
	RequestBody  *RequestBody        `json:"requestBody,omitempty"`
	Responses    *Responses          `json:"responses,omitempty"`
	Deprecated   bool                `json:"deprecated,omitempty"`
	ExternalDocs *ExternalDocs       `json:"externalDocs,omitempty"`
	Callbacks    map[string]Callback `json:"callbacks,omitempty"`
	Security     []Security          `json:"security,omitempty"`
	Servers      []Server            `json:"servers,omitempty"`
}

type Path struct {
	Get    *Operation `json:"get,omitempty"`
	Post   *Operation `json:"post,omitempty"`
	Put    *Operation `json:"put,omitempty"`
	Delete *Operation `json:"delete,omitempty"`
	// TODO
	Parameters []Parameter `json:"parameters,omitempty"`
}

type SecurityScheme struct {
	// TODO
}

type Discriminator struct {
	PropertyName string            `json:"propertyName,omitempty"`
	Mapping      map[string]string `json:"mapping,omitempty"`
}

type Schema struct {
	Ref              string             `json:"$ref,omitempty"`
	Title            string             `json:"title,omitempty"`
	Type             Type               `json:"type,omitempty"`
	Properties       map[string]*Schema `json:"properties,omitempty"`
	Format           Format             `json:"format,omitempty"`
	Items            *Schema            `json:"items,omitempty"`
	Description      string             `json:"description,omitempty"`
	Default          interface{}        `json:"default,omitempty"`
	Example          interface{}        `json:"example,omitempty"`
	Deprecated       bool               `json:"deprecated,omitempty"`
	Discriminator    *Discriminator     `json:"discriminator,omitempty"`
	Nullable         bool               `json:"nullable,omitempty"`
	Maximum          interface{}        `json:"maximum,omitempty"`
	Minimum          interface{}        `json:"minimum,omitempty"`
	ExclusiveMaximum interface{}        `json:"exclusiveMaximum,omitempty"`
	ExclusiveMinimum interface{}        `json:"exclusiveMinimum,omitempty"`
	MaxLength        int                `json:"maxLength,omitempty"`
	MinLength        int                `json:"minLength,omitempty"`
	Required         []string           `json:"required,omitempty"`
	Enum             []interface{}      `json:"enum,omitempty"`
	AllOf            []*Schema          `json:"allOf,omitempty"`
	OneOf            []*Schema          `json:"oneOf,omitempty"`
	AnyOf            []*Schema          `json:"anyOf,omitempty"`
	Not              []*Schema          `json:"not,omitempty"`
	// AdditionalProperties *Schema or bool
	AdditionalProperties interface{} `json:"additionalProperties,omitempty"`
	Pattern              interface{} `json:"pattern,omitempty"`
}

type Components struct {
	Schemas       map[string]Schema      `json:"schemas,omitempty"`
	RequestBodies map[string]RequestBody `json:"requestBodies,omitempty"`
	Responses     map[string]Response    `json:"responses,omitempty"`
	// TODO
	Parameters map[string]Parameter `json:"parameters,omitempty"`
	// TODO
	Examples map[string]Example `json:"examples,omitempty"`
	// TODO
	Headers map[string]Header `json:"headers,omitempty"`
	// TODO
	SecuritySchemes map[string]SecurityScheme `json:"securitySchemes,omitempty"`
	// TODO
	Links map[string]Link `json:"links,omitempty"`
	// TODO
	Callbacks map[string]Callback `json:"callbacks,omitempty"`
}

type Api struct {
	Openapi      string          `json:"openapi,omitempty"`
	Info         *Info           `json:"info,omitempty"`
	Servers      []Server        `json:"servers,omitempty"`
	Tags         []Tag           `json:"tags,omitempty"`
	Paths        map[string]Path `json:"paths,omitempty"`
	Components   *Components     `json:"components,omitempty"`
	ExternalDocs *ExternalDocs   `json:"externalDocs,omitempty"`
}

type Type string

const (
	IntegerT Type = "integer"
	StringT  Type = "string"
	BooleanT Type = "boolean"
	NumberT  Type = "number"
	ObjectT  Type = "object"
	ArrayT   Type = "array"
)

type Format string

const (
	Int32F    Format = "int32"
	Int64F    Format = "int64"
	FloatF    Format = "float"
	DoubleF   Format = "double"
	DateTimeF Format = "date-time"
	BinaryF   Format = "binary"
)

var (
	Any = &Schema{
		Type: ObjectT,
	}
	Int = &Schema{
		Type:   IntegerT,
		Format: Int32F,
	}
	Int64 = &Schema{
		Type:   IntegerT,
		Format: Int64F,
	}
	String = &Schema{
		Type: StringT,
	}
	Time = &Schema{
		Type:   StringT,
		Format: DateTimeF,
	}
	Bool = &Schema{
		Type: BooleanT,
	}
	Float32 = &Schema{
		Type:   NumberT,
		Format: FloatF,
	}
	Float64 = &Schema{
		Type:   NumberT,
		Format: DoubleF,
	}
	File = &Schema{
		Type:   StringT,
		Format: BinaryF,
	}
	FileArray = &Schema{
		Type:  ArrayT,
		Items: File,
	}
)
