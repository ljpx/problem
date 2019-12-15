package problem

// Details is the main item of the package.  When serialized, it produces a
// valid RFC-7807 problem details structure.
type Details struct {
	Type      string      `json:"type"`
	Title     string      `json:"title"`
	Detail    string      `json:"detail,omitempty"`
	Specifics interface{} `json:"specifics,omitempty"`
	Error     string      `json:"error,omitempty"`
}

// AttachError attaches an error to the calling Details object.
func (d *Details) AttachError(err error) {
	if err != nil {
		d.Error = err.Error()
	} else {
		d.Error = ""
	}
}
