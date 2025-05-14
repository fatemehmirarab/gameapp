package richerror

type RichError struct {
	Operation    Op
	WrappedError error
	Message      string
	Kind         Kind
	Meta         map[string]interface{}
}

type Kind int

type Op string

func New(op Op) RichError {
	return RichError{Operation: op}
}

func (r RichError) WithError(err error) RichError {
	r.WrappedError = err
	return r
}

func (r RichError) WithMessage(message string) RichError {
	r.Message = message
	return r
}

func (r RichError) WithKind(kind Kind) RichError {
	r.Kind = kind
	return r
}

func (r RichError) WithMeta(meta map[string]interface{}) RichError {
	r.Meta = meta
	return r
}

/*func New(args ...interface{}) RichError {
	r := RichError{}
	for _, arg := range args {
		switch arg.(type) {
		case Op:
			r.Operation = arg.(Op)
		case string:
			r.Message = arg.(string)
		case error:
			r.WrappedError = arg.(error)
		case Kind:
			r.Kind = arg.(Kind)
		case map[string]interface{}:
			r.Meta = arg.(map[string]interface{})

		}
	}
	return r
}

*/

/*func New(operation string , wrappedError error , message string ,kind Kind , meta  map[string] interface{}) RichError{
	return RichError{
		WrappedError: wrappedError,
		Message : message,
		Kind: kind,
		Operation :operation ,
		Meta: meta,
	}
}
*/

func (r RichError) Error() string {
	return r.Message
}
