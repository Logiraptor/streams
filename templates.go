package streams

import (
	"github.com/clipperhouse/gen/typewriter"
)

var template = typewriter.TemplateSet{
	"base": &typewriter.Template{
		Text: `
type {{.Name}}Stream struct {
	stream chan {{.Name}}
	done   chan struct{}
}

func new{{.Name}}Stream(done chan struct{}) {{.Name}}Stream {
	if done == nil {
		done = make(chan struct{})
	}
	return {{.Name}}Stream{
		stream: make(chan {{.LocalName}}),
		done:   done,
	}
}

func {{.Name}}StreamFromSlice(a ...{{.LocalName}}) {{.Name}}Stream {
	s := new{{.Name}}Stream(nil)
	go func() {
		for _, x := range a {
			select {
			case s.stream <- x:
			case <-s.done:
				close(s.stream)
				return
			}
		}
		close(s.stream)
	}()
	return s
}

func {{.Name}}StreamFromGenerator(gen func() ({{.Name}}, bool)) {{.Name}}Stream {
	s := new{{.Name}}Stream(nil)
	go func() {
		for {
			if x, ok := gen(); ok {
				select {
				case s.stream <- x:
				case <-s.done:
					close(s.stream)
					return
				}
			} else {
				break
			}
		}
		close(s.stream)
	}()
	return s
}

func {{.Name}}StreamMerge(streams ...{{.Name}}Stream) {{.Name}}Stream {
	s := new{{.Name}}Stream(nil)
	var wg sync.WaitGroup
	for _, stream := range streams {
		wg.Add(1)
		go func(stream {{.Name}}Stream) {
			defer wg.Done()
			for x := range stream.stream {
				select {
				case s.stream <- x:
				case <-stream.done:
					return
				}
			}
		}(stream)
	}
	go func() {
		wg.Wait()
		close(s.stream)
	}()
	return s
}

func (t {{.Name}}Stream) Filter(filter func({{.LocalName}}) bool) {{.Name}}Stream {
	s := new{{.Name}}Stream(t.done)
	go func() {
		for x := range t.stream {
			if filter(x) {
				select {
				case s.stream <- x:
				case <-s.done:
					close(s.stream)
					return
				}
			}
		}
		close(s.stream)
	}()
	return s
}

func (t {{.Name}}Stream) Modify(proc func({{.LocalName}}) {{.LocalName}}) {{.Name}}Stream {
	s := new{{.Name}}Stream(t.done)
	go func() {
		for x := range t.stream {
			select {
			case s.stream <- proc(x):
			case <-s.done:
				close(s.stream)
				return
			}
		}
		close(s.stream)
	}()
	return s
}

func (t {{.Name}}Stream) Each(proc func({{.LocalName}})) {{.Name}}Stream {
	s := new{{.Name}}Stream(t.done)
	go func() {
		for x := range t.stream {
			proc(x)
			select {
			case s.stream <- x:
			case <-s.done:
				close(s.stream)
				return
			}
		}
		close(s.stream)
	}()
	return s
}

func (t {{.Name}}Stream) Drain() []{{.LocalName}} {
	var a []{{.Name}}
	for x := range t.stream {
		a = append(a, x)
	}
	return a
}

func (t {{.Name}}Stream) Iter() chan {{.LocalName}} {
	return t.stream
}

func (t {{.Name}}Stream) Next() {{.LocalName}} {
	return <-t.stream
}

func (t {{.Name}}Stream) Close() {
	close(t.done)
}

	`,
	},
}
