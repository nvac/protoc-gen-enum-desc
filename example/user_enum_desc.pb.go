package example

var (
	User_Mime_name_desc = map[int32]string{
		0: "UNSPECIFIED",
		1: "VIDEO",
		2: "IMAGE",
	}

	User_Mime_value_desc = map[string]string{
		"Mime_Unspecified": "UNSPECIFIED",
		"Mime_Video":       "VIDEO",
		"Mime_Image":       "IMAGE",
	}
)

func NewUser_MimeFromNumber(number int32) *User_Mime {
	e := User_Mime(number)
	return &e
}

func NewUser_MimeFromValue(value string) *User_Mime {
	return NewUser_MimeFromNumber(User_Mime_value[value])
}

func (e User_Mime) GetDesc() string {
	return User_Mime_name_desc[int32(e)]
}

func (User_Mime) GetDescFromNumber(number int32) string {
	return User_Mime_name_desc[number]
}

func (User_Mime) GetDescFromName(name string) string {
	return User_Mime_value_desc[name]
}

var (
	User_NestedUser_NestedMime_name_desc = map[int32]string{
		0: "UNSPECIFIED",
		1: "VIDEO",
		2: "IMAGE",
	}

	User_NestedUser_NestedMime_value_desc = map[string]string{
		"NestedMime_Unspecified": "UNSPECIFIED",
		"NestedMime_Video":       "VIDEO",
		"NestedMime_Image":       "IMAGE",
	}
)

func NewUser_NestedUser_NestedMimeFromNumber(number int32) *User_NestedUser_NestedMime {
	e := User_NestedUser_NestedMime(number)
	return &e
}

func NewUser_NestedUser_NestedMimeFromValue(value string) *User_NestedUser_NestedMime {
	return NewUser_NestedUser_NestedMimeFromNumber(User_NestedUser_NestedMime_value[value])
}

func (e User_NestedUser_NestedMime) GetDesc() string {
	return User_NestedUser_NestedMime_name_desc[int32(e)]
}

func (User_NestedUser_NestedMime) GetDescFromNumber(number int32) string {
	return User_NestedUser_NestedMime_name_desc[number]
}

func (User_NestedUser_NestedMime) GetDescFromName(name string) string {
	return User_NestedUser_NestedMime_value_desc[name]
}

var (
	Status_name_desc = map[int32]string{
		0: "UNSPECIFIED",
		1: "ACTIVE",
		2: "DELETED",
	}

	Status_value_desc = map[string]string{
		"Status_Unspecified": "UNSPECIFIED",
		"Status_Active":      "ACTIVE",
		"Status_Deleted":     "DELETED",
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

var (
	DayOfWeek_name_desc = map[int32]string{
		0: "星期天",
		1: "星期一",
		2: "星期二",
		3: "星期三",
		4: "星期四",
		5: "星期五",
		6: "星期六",
	}

	DayOfWeek_value_desc = map[string]string{
		"DayOfWeek_Sunday":    "星期天",
		"DayOfWeek_Monday":    "星期一",
		"DayOfWeek_Tuesday":   "星期二",
		"DayOfWeek_Wednesday": "星期三",
		"DayOfWeek_Thursday":  "星期四",
		"DayOfWeek_Friday":    "星期五",
		"DayOfWeek_Saturday":  "星期六",
	}
)

func NewDayOfWeekFromNumber(number int32) *DayOfWeek {
	e := DayOfWeek(number)
	return &e
}

func NewDayOfWeekFromValue(value string) *DayOfWeek {
	return NewDayOfWeekFromNumber(DayOfWeek_value[value])
}

func (e DayOfWeek) GetDesc() string {
	return DayOfWeek_name_desc[int32(e)]
}

func (DayOfWeek) GetDescFromNumber(number int32) string {
	return DayOfWeek_name_desc[number]
}

func (DayOfWeek) GetDescFromName(name string) string {
	return DayOfWeek_value_desc[name]
}
