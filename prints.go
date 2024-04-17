package cafe

import (
	"encoding/json"
	"strconv"
	"strings"
)

func (c *Cafe) JSON() string {
	m := make(map[string]interface{})
	for k, v := range c.schema {
		if v.typ == "subschema" {
			sub, ok := v.Value.(*Cafe)
			if !ok {
				continue
			}
			for k2, v2 := range sub.schema {
				m[k2] = v2.Value
			}
			continue
		}
		m[k] = v.Value
	}

	bytes, _ := json.MarshalIndent(m, "", "  ")
	return string(bytes)
}

func (c *Cafe) Env() string {
	builder := strings.Builder{}
	for k, v := range c.schema {
		if v.typ == "subschema" {
			sub, ok := v.Value.(*Cafe)
			if !ok {
				continue
			}
			for k2, v2 := range sub.schema {
				stringVal := getStringValue(v2.typ, v2.Value)
				builder.WriteString(k2)
				builder.WriteString("=")
				builder.WriteString(stringVal)
				builder.WriteString("\n")
			}
			continue
		}
		builder.WriteString(k)
		builder.WriteString("=")
		builder.WriteString(getStringValue(v.typ, v.Value))
		builder.WriteString("\n")
	}

	return builder.String()
}

func getStringValue(typ string, v interface{}) string {
	switch typ {
	case "string":
		str, ok := v.(string)
		if !ok {
			return ""
		}
		return str
	case "int":
		intVal, ok := v.(int)
		if !ok {
			return ""
		}
		return strconv.Itoa(intVal)
	case "bool":
		boolVal, ok := v.(bool)
		if !ok {
			return ""
		}
		return strconv.FormatBool(boolVal)
	default:
		return ""

	}
}
