package provider

var registry = map[string]Provider{}

func Register(p Provider)              { registry[p.Name()] = p }
func Get(name string) (Provider, bool) { p, ok := registry[name]; return p, ok }
func All() []Provider {
	out := make([]Provider, 0, len(registry))
	for _, p := range registry {
		out = append(out, p)
	}
	return out
}
