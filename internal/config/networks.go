package config

import "gitlab.com/distributed_lab/kit/kv"

type Networks struct {
	m map[string]*string
}

func (c *config) Networks() Networks {
	return c.networksOnce.Do(func() interface{} {
		m := kv.MustGetStringMap(c.getter, "networks")
		ns := make(map[string]*string)

		for name, u := range m {
			str, ok := u.(string)
			if !ok {
				panic("failed to figure out networks: interface{} is not string")
			}
			ns[name] = &str
		}

		return Networks{ns}
	}).(Networks)
}

func (n Networks) Get(key string) *string {
	return n.m[key]
}
