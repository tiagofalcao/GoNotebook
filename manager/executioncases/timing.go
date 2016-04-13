package executioncases

import (
  "time"
)

type timing struct {
  s      time.Time
  i      time.Time
  e      time.Time
}

func (t *timing) start() {
  t.s = time.Now()
}

func (t *timing) read() {
  t.i = time.Now()
}

func (t *timing) end() {
  t.e = time.Now()
}

func (t timing) reading() time.Duration  {
  return t.i.Sub(t.s)
}

func (t timing) computing() time.Duration  {
  return t.e.Sub(t.i)
}

func (t timing) complete() time.Duration  {
  return t.e.Sub(t.s)
}
