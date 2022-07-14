package util

import (
	"reflect"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	MAP_STRUCT_TAG    = "mapstruct"
	STRUCT_TO_MAP_TAG = "json"
	TIME_LAYOUT_TAG   = "formattime"
	DATE_TIME_LAYOUT1 = "2006-01-02"          // yyy-mm-dd
	DATE_TIME_LAYOUT2 = "2006-01-02 15:04:05" // yy-mm-dd hh:mm:ss
)

func MapStruct(obj interface{}, m map[string]string) error {
	for name, value := range m {
		err := mapField(obj, name, value)
		if err != nil {
			logrus.WithError(err).Infof("Map field %s fail", name)
		}
	}

	return nil
}

func mapField(obj interface{}, tag string, value string) error {
	structValue := reflect.ValueOf(obj).Elem()

	for i := 0; i < structValue.NumField(); i++ {
		field := structValue.Type().Field(i)

		if field.Tag.Get(MAP_STRUCT_TAG) == tag {
			name := field.Name
			structFieldValue := structValue.FieldByName(name)

			if !structFieldValue.IsValid() {
				errors.Errorf("No such field: %s", name)
			}
			if !structFieldValue.CanSet() {
				errors.Errorf("Cannot set %s field value", name)
			}

			structFieldType := structFieldValue.Type()
			val := reflect.ValueOf(value)

			if structFieldType == val.Type() {
				structFieldValue.Set(val)
				return nil
			} else {
				switch structFieldValue.Interface().(type) {
				case float64:
					if valNew, err := strconv.ParseFloat(value, 64); err == nil {
						structFieldValue.Set(reflect.ValueOf(valNew))
						return nil
					}

				case float32:
					if valNew, err := strconv.ParseFloat(value, 32); err == nil {
						structFieldValue.Set(reflect.ValueOf(valNew))
						return nil
					}

				case int64:
					if valNew, err := strconv.ParseInt(value, 10, 64); err == nil {
						structFieldValue.Set(reflect.ValueOf(valNew))
						return nil
					}

				case int:
					if valNew, err := strconv.Atoi(value); err == nil {
						structFieldValue.Set(reflect.ValueOf(valNew))
						return nil
					}

				case time.Time:
					layout := DATE_TIME_LAYOUT2
					if val, ok := field.Tag.Lookup(TIME_LAYOUT_TAG); ok {
						switch val {
						case "DATE_TIME_LAYOUT1":
							layout = DATE_TIME_LAYOUT1
						}
					}

					if valNew, err := time.Parse(layout, value); err == nil {
						structFieldValue.Set(reflect.ValueOf(valNew))
						return nil
					}

				case bool:
					if valNew, err := strconv.ParseBool(value); err == nil {
						structFieldValue.Set(reflect.ValueOf(valNew))
						return nil
					}

				case string:
					structFieldValue.Set(reflect.ValueOf(value))
					return nil
				}
			}

			return errors.Errorf("Cannot parse type of %s field value", name)
		}
	}

	return errors.Errorf("Tag: %s does not exist", tag)
}

func StructToMap(obj interface{}) map[string]interface{} {
	mapResult := map[string]interface{}{}
	structValue := reflect.ValueOf(obj).Elem()

	for i := 0; i < structValue.NumField(); i++ {
		field := structValue.Type().Field(i)
		name := field.Name
		structFieldValue := structValue.FieldByName(name)

		if tagName := field.Tag.Get(STRUCT_TO_MAP_TAG); tagName != "" {
			mapResult[tagName] = structFieldValue.Interface()
		} else {
			mapResult[name] = structFieldValue.Interface()
		}
	}

	return mapResult
}
