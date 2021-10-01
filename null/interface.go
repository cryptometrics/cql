package null

type Interface struct {
	Value interface{}
	Valid bool
}

func NewInterfaceFloat(f *float64) (i Interface) {
	if f == nil {
		return
	}
	i.Value = &f
	i.Valid = true
	return
}

func NewInterfaceString(str *string) (i Interface) {
	if str == nil {
		return
	}
	i.Value = &str
	i.Valid = true
	return
}
