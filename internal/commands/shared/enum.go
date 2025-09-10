package shared

type Verb string

const (
	List   Verb = "list"
	Get    Verb = "get"
	Create Verb = "create"
	Update Verb = "update"
	Delete Verb = "delete"
	Import Verb = "import"
)

func (v Verb) String() string {
	return string(v)
}

type Resource string

const (
	Gateway Resource = "gateways"
	Device  Resource = "devices"
)

func (r Resource) String() string {
	return string(r)
}
