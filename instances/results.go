package instances

type result struct {
	instance uint64
	value    interface{}
}

type ResultManager struct {
	Results []interface{}
	count   uint64
	comm    chan result
}

func ResultMan(instances uint64) *ResultManager {
	return &ResultManager{
		Results: make([]interface{}, instances),
		comm:    make(chan result, instances),
	}
}

func (r *ResultManager) Set(instance uint64, value interface{}) {
	r.comm <- result{instance, value}
}

func (r *ResultManager) Join() {
	l := uint64(len(r.Results))
	for r.count < l {
		i := <-r.comm
		r.Results[i.instance] = i.value
		r.count++
	}
}
