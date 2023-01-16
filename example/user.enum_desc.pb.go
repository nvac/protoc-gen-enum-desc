package example

var (
	Status_name_desc = map[int32]string{
		0: "",
		1: "ACTIVE",
		2: "DELETED",
	}

	Status_value_desc = map[string]string{
		"_":       "",
		"Active":  "ACTIVE",
		"Deleted": "DELETED",
	}
)

func NewStatusFromNumber(number int32) *Status {
	e := Status(number)
	return &e
}

func NewStatusFromValue(value string) *Status {
	return NewStatusFromNumber(Status_value[value])
}

func (e Status) GetDesc() string {
	return Status_name_desc[int32(e)]
}

func (Status) GetDescFromNumber(number int32) string {
	return Status_name_desc[number]
}

func (Status) GetDescFromName(name string) string {
	return Status_value_desc[name]
}
