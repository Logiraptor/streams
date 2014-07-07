package streams

import (
	"github.com/clipperhouse/gen/typewriter"

	"io"
)

func init() {
	err := typewriter.Register(NewStreamWriter())
	if err != nil {
		panic(err)
	}
}

type StreamWriter struct {
}

func NewStreamWriter() *StreamWriter {
	return &StreamWriter{}
}

func (s *StreamWriter) Name() string {
	return "stream"
}

// Validate is called for every Type, prior to further action, to answer two questions:
// a) that your TypeWriter will write for this Type; return false if your TypeWriter intends not to write a file for this Type at all.
// b) that your TypeWriter considers the declaration (i.e., Tags) valid; return err if not.
func (s *StreamWriter) Validate(t typewriter.Type) (bool, error) {
	return true, nil
}

// WriteHeader writer to the top of the generated code, before the package declaration; intended for licenses or general documentation.
func (s *StreamWriter) WriteHeader(w io.Writer, t typewriter.Type) {
	m := `// Stream implementation inspired by http://blog.golang.org/pipelines	
`

	w.Write([]byte(m))
}

// Imports is a slice of imports required for the type; each will be written into the imports declaration.
func (s *StreamWriter) Imports(t typewriter.Type) []typewriter.ImportSpec {
	return []typewriter.ImportSpec{
		{Path: "sync"},
	}
}

// WriteBody writes to the body of the generated code, following package declaration, headers and imports. This is the meat.
func (s *StreamWriter) WriteBody(w io.Writer, t typewriter.Type) {
	tmpl, _ := template.Get("base")
	err := tmpl.Execute(w, &t)
	if err != nil {
		panic(err)
	}
}
