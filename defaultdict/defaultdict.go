package defaultdict

type DefaultDict struct {
	mapping       map[interface{}]interface{}
	defaultGetter func() interface{}
}

func New(defaultGetter func() interface{}) *DefaultDict {
	return &DefaultDict{
		mapping:       map[interface{}]interface{}{},
		defaultGetter: defaultGetter,
	}
}

func (d *DefaultDict) Get(key interface{}) interface{} {
	if value, ok := d.mapping[key]; ok {
		return value
	}
	return d.defaultGetter()
}

func (d *DefaultDict) Update(key interface{}, fn func(value interface{}) interface{}) {
	value, ok := d.mapping[key]
	if !ok {
		value = d.defaultGetter()
	}
	d.mapping[key] = fn(value)
}

func (d *DefaultDict) Set(key, value interface{}) {
	d.mapping[key] = value
}

func IntDefault() interface{}   { return int(0) }
func Int32Default() interface{} { return int32(0) }
func Int64Default() interface{} { return int64(0) }

func Float64Default() interface{} { return float64(0) }
func Float32Default() interface{} { return float32(0) }
