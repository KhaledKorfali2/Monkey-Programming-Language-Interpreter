// object/object.go 

package object

type ObjectType string

type Object interface {
	Type()	ObjectType
	Inspect()	string
}
