package error

type Validation struct {
	Errors []ValidationMessage `json:"errors"`
}

type ValidationMessage struct {
	Field   string `json:"field"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func NewValidation() error {
	return &Validation{
		Errors: []ValidationMessage{},
	}
}

func NewValidationMessage(field, code, message string) ValidationMessage {
	return ValidationMessage{
		Field:   field,
		Code:    code,
		Message: message,
	}
}

func (e Validation) Error() string {
	return "Validation error"
}

func (e *Validation) Merge(err error) {
	if err != nil {
		e.Errors = append(e.Errors, err.(*Validation).Errors...)
	}
}

func (e *Validation) HasError() bool {
	return len(e.Errors) > 0
}

func (e *Validation) AddError(message ValidationMessage) {
	e.Errors = append(e.Errors, message)
}

func (e *Validation) ErrorOrNil() error {
	if e.HasError() {
		return e
	}

	return nil
}

func (m *ValidationMessage) Throw(err error) {
	err.(*Validation).AddError(*m)
}
