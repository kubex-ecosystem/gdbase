package models

import (
	"fmt"
	"reflect"

	"github.com/goccy/go-json"
	ci "github.com/kubex-ecosystem/gdbase/internal/interfaces"
)

type Model interface {
	ci.IReference // Aqui tem nome e ID (uuid)
	Validate() error
}

var ModelList = make([]interface{}, 0) /*{
	//&UserImpl{},
	//&Product{},
	//&CustomerImpl{},
	//&Order{},
})*/
var ModelRegistryMap = map[reflect.Type]any{
	//strings.ToLower("User"):     reflect.TypeOf(UserImpl{}),
	//strings.ToLower("Product"):  reflect.TypeOf(Product{}),
	//strings.ToLower("Customer"): reflect.TypeOf(CustomerImpl{}),
	//strings.ToLower("Order"):    reflect.TypeOf(Order{}),
	//strings.ToLower("Ping"):     reflect.TypeOf(PingImpl{}),
}

type ModelRegistryImpl[T Model] struct {
	Dt *T     `json:"data"`
	St []byte `json:"status"`
}
type ModelRegistryInterface interface {
	GetType() reflect.Type
	FromModel(model interface{}) ModelRegistryInterface
	FromSerialized(data []byte) (ModelRegistryInterface, error)
	ToModel() interface{}
}

func (m *ModelRegistryImpl[T]) GetType() reflect.Type { return reflect.TypeFor[T]() }

func (m *ModelRegistryImpl[T]) FromModel(model interface{}) ModelRegistryInterface {
	if model == nil {
		return nil
	}
	// Ficou assim para evitar o loop de importação, ta lindão!
	md, ok := model.(*T)
	if ok {
		vl := reflect.ValueOf(md)
		if !vl.IsValid() || vl.IsNil() {
			return nil
		}
		if (*md).Validate() != nil {
			return nil
		}

		// Agora são ponteiros pro mesmo valor?
		// Se não for um ponteiro, cria um novo ponteiro
		m.Dt = md
		m.St, _ = json.Marshal(m.Dt)
	}
	if m.Dt == nil {
		return nil
	}
	// Verifica se o tipo do modelo está registrado
	if _, ok := ModelRegistryMap[reflect.TypeOf(m.Dt)]; !ok {
		return nil
	}
	return m
}
func (m *ModelRegistryImpl[T]) FromSerialized(data []byte) (ModelRegistryInterface, error) {
	var mdr ModelRegistryImpl[T]
	if err := json.Unmarshal(data, &mdr); err != nil {
		return nil, err
	}
	// Retorna o tipo que está implícito na estrutura pelo generic T
	// Assim não é preciso armazenar o tipo do modelo
	// Verifica se o tipo do modelo está registrado
	if _, ok := ModelRegistryMap[mdr.GetType()]; !ok {
		return nil, fmt.Errorf("model %s not found", mdr.GetType())
	}
	return &mdr, nil
}
func (m *ModelRegistryImpl[T]) ToModel() interface{} {
	if model, ok := ModelRegistryMap[m.GetType()]; ok {
		// Verifica se está nulo o objeto e se existe de fato no map.
		// Se não existir, retorna nil
		if model != nil {
			return model
		}
	}
	return nil
}

func RegisterModel(modelType reflect.Type) error {
	// Ferrou porque não tem mais como guardar o nome.. rsrs
	if _, exists := ModelRegistryMap[modelType]; exists {
		return fmt.Errorf("model %s já registrado", modelType.String())
	}
	// O map armazena valores pelo tipo do modelo, então como estamos só
	// registrando o tipo, não precisamos guardar valor. O nome está implícito
	// na interface Model. Só implementar lá. rsrs
	ModelRegistryMap[modelType] = nil
	return nil
}
func NewModelRegistry[T Model]() ModelRegistryInterface {
	return &ModelRegistryImpl[T]{}
}
func NewModelRegistryFromModel[T Model](model interface{}) ModelRegistryInterface {
	mr := ModelRegistryImpl[T]{}
	return mr.FromModel(model)
}
func NewModelRegistryFromSerialized[T Model](data []byte) (ModelRegistryInterface, error) {
	mr := ModelRegistryImpl[T]{}
	return mr.FromSerialized(data)
}

// func ParseConditionClause(conditions ...any) (string, []any, error) {
// 	if len(conditions) == 0 {
// 		return "", nil, fmt.Errorf("no conditions provided")
// 	}

// 	var whereClause strings.Builder
// 	var clauseSliceObj = make([]any, 0) // É slice de interfaces, então a merda que vier de valor é só jogar e
// 	// dane-se o Gorm.. Ele se vira pra lá. A gente evita o mínimo/médio de dentro da responsabilidade dele....

// 	// Cada posição da condição pode ser um slice, um map ou qualquer outra coisa...
// 	for _, condition := range conditions {
// 		value := reflect.ValueOf(condition)

// 		// Então primeiro detecto se é slice, se for, trato como slice
// 		switch reflect.TypeOf(condition).Kind() {
// 		case reflect.Slice, reflect.SliceOf(reflect.TypeFor[map[string]string]()).Kind():
// 			if value.Len() == 0 {
// 				// If the slice is empty, assign the deserialized object to the slice
// 				gl.Log("debug", "Query conditions are empty")
// 				continue
// 			} else if value.Len() == 1 {
// 				// If the slice has only one element, assign it directly
// 				clauseSliceObj = append(clauseSliceObj, value.Index(0).Interface())
// 				continue
// 			} else if value.Len() > 1 {
// 				if value.Type().Elem().Kind() == reflect.Map {
// 					for i := 0; i < value.Len(); i++ {
// 						clauseSliceObj = append(clauseSliceObj, value.Index(i).Interface())
// 					}
// 				}
// 				continue
// 			} else {
// 				for i := 0; i < value.Len(); i++ {
// 					clauseSliceObj = append(clauseSliceObj, value.Index(i).Interface())
// 				}
// 			}
// 		// Segundo verifico se é map, se for, trato como map
// 		case reflect.Map:
// 			if value.Len() == 0 {
// 				// If the map is empty, assign the deserialized object to the map
// 				clauseSliceObj = reflect.ValueOf(conditions).Interface().([]any)
// 			} else {
// 				for _, key := range value.MapKeys() {
// 					clauseSliceObj = append(clauseSliceObj, fmt.Sprintf("%s = ?", key.String()))
// 					clauseSliceObj = append(clauseSliceObj, value.MapIndex(key).Interface())
// 				}
// 			}
// 		default:
// 			// If the type is neither a slice nor a map, assign the first object to m.object
// 			if len(conditions) == 0 {
// 				gl.Log("debug", "Query conditions are empty")
// 				continue
// 			}
// 			if len(conditions) > 1 {
// 				gl.Log("debug", "Multiple query conditions found")
// 				continue
// 			}
// 			clauseSliceObj = append(clauseSliceObj, condition)
// 		}
// 	}

// 	var whereClause strings.Builder
// 	for i, condition := range clauseSliceObj {
// 		switch keyType := condition.(type) {
// 		case int:
// 			whereClause.WriteString(fmt.Sprintf("column%d = ?", i))
// 		case string:
// 			whereClause.WriteString(fmt.Sprintf("%s = ?", keyType))
// 		default:
// 			return "", nil, fmt.Errorf("unsupported type for where clause: %T", keyType)
// 		}
// 		if i < len(clauseSliceObj)-1 {
// 			whereClause.WriteString(" AND ")
// 		}
// 	}
// 	return whereClause.String(), clauseSliceObj, nil
// }
