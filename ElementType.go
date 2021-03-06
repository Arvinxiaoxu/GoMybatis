package GoMybatis

type ElementType = string

const (
	Element_ResultMap ElementType = "resultMap"
	Element_Insert    ElementType = "insert"
	Element_Delete    ElementType = "delete"
	Element_Update    ElementType = `update`
	Element_Select    ElementType = "select"

	Element_bind ElementType = "bind"

	Element_String    ElementType = "string"
	Element_If        ElementType = `if`
	Element_Trim      ElementType = "trim"
	Element_Foreach   ElementType = "foreach"
	Element_Set       ElementType = "set"
	Element_choose    ElementType = "choose"
	Element_when      ElementType = "when"
	Element_otherwise ElementType = "otherwise"
)
