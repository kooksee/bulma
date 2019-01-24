package bulma

type _sizeMap struct {
	IsSmall  bool
	IsMedium bool
	IsLarge  bool
	IsNormal bool
}

type _size struct {
	IsSmall  string
	IsMedium string
	IsLarge  string
	IsNormal string
}

var Size = _size{
	IsSmall:  "is-small",
	IsMedium: "is-medium",
	IsLarge:  "is-large",
	IsNormal: "is-normal",
}
