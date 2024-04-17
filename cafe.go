package cafe

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Schema map[string]*Type

type Cafe struct {
	schema Schema
}

type Type struct {
	key        string
	typ        string
	isRequired bool
	Value      interface{}
	defaultVal interface{}
}

func (t *Type) Require() *Type {
	t.isRequired = true
	return t
}

func (t *Type) Default(v interface{}) *Type {
	t.defaultVal = v
	return t
}

func (t *Type) Key(k string) *Type {
	t.key = k
	return t
}

func (t *Type) Before() error {
	if t.key == "" {
		return Err_KEY_IS_REQUIRED
	}
	return nil
}

func (t *Type) Validate() error {
	if t.isRequired && t.Value == nil {
		return buildRequiredKeyMissing(t.key)
	}
	return nil
}

func String(k string) *Type {
	return &Type{typ: "string", isRequired: false, key: k}
}

func Int(k string) *Type {
	return &Type{typ: "int", isRequired: false, key: k}
}

func Bool(k string) *Type {
	return &Type{typ: "bool", isRequired: false, key: k}
}

func SubSchema(k string, s Schema) *Type {
	return &Type{typ: "subschema", isRequired: false, key: k, Value: &Cafe{
		schema: s,
	}}
}

// func Object(k string) *Type {
// 	return &Type{typ: "object", isRequired: false, key: k}
// }

func (s *Cafe) Initialize() error {
	var err error
	for _, v := range (*s).schema {
		err = v.Before()
		if err != nil {
			return err
		}
		val := os.Getenv(v.key)
		if val == "" && v.defaultVal != nil {
			val = fmt.Sprintf("%v", v.defaultVal)
		}
		switch v.typ {
		case "string":
			if v.isRequired && val == "" {
				return buildRequiredKeyMissing(v.key)
			}
			v.Value = val
		case "int":
			v.Value, err = strconv.Atoi(val)
			if err != nil {
				println(err.Error())
			}
		case "bool":
			v.Value, err = strconv.ParseBool(val)
			if err != nil {
				println(err.Error())
			}
		case "subschema":
			sub, ok := v.Value.(*Cafe)
			if !ok {
				return fmt.Errorf("subschema is not a Cafe")
			}
			err = sub.Initialize()
			if err != nil {
				return err
			}
		case "object":
			// TODO
		default:
			return fmt.Errorf("unknown type: %s", v.typ)
		}
		err = v.Validate()
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Cafe) GetString(k string) (string, error) {
	spilt := strings.Split(k, ".")
	if len(spilt) == 1 {
		v, err := s.fetch(k, "string")
		if err != nil {
			return "", err
		}
		return v.(string), nil
	}
	return "", nil
}

func (s *Cafe) GetInt(k string) (int, error) {
	spilt := strings.Split(k, ".")
	if len(spilt) == 1 {
		v, err := s.fetch(k, "int")
		if err != nil {
			return 0, err
		}
		return v.(int), nil
	}
	return 0, nil
}

func (s *Cafe) GetBool(k string) (bool, error) {
	spilt := strings.Split(k, ".")
	if len(spilt) == 1 {
		v, err := s.fetch(k, "bool")
		if err != nil {
			return false, err
		}
		return v.(bool), nil
	}
	return false, nil
}

func (s *Cafe) GetSubSchema(k string) (*Cafe, error) {
	spilt := strings.Split(k, ".")
	if len(spilt) == 1 {
		v, err := s.fetch(k, "subschema")
		if err != nil {
			return nil, err
		}
		return v.(*Cafe), nil
	}
	return nil, nil
}

func (s *Cafe) fetch(k string, typ string) (interface{}, error) {
	v, ok := (*s).schema[k]
	if !ok {
		return "", fmt.Errorf(k + str_UNREGISTERED)
	}
	if v.typ != typ {
		return "", fmt.Errorf(fmt.Sprintf(str_NON_MATCHED_FETCH, k, v.typ, typ))
	}
	return v.Value, nil
}

func New(s Schema) (*Cafe, error) {
	c := NewCafeSchema(s)
	err := c.Initialize()
	if err != nil {
		return nil, err
	}
	return c, nil
}

func NewCafeSchema(s Schema) *Cafe {
	return &Cafe{schema: s}
}

// func (s *Cafe) GetInt(k string) (int, error) {
