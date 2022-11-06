package model

type ApplicationError struct {
	Message string
}

func (e *ApplicationError) Error() string {
	if e.Message == "" {
		return "Application error"
	}
	return e.Message
}

type ConversionError struct {
	Message string
}

func (e *ConversionError) Error() string {
	if e.Message == "" {
		return "Conversion error"
	}
	return e.Message
}

type JsonError struct {
	Message string `json:"message"`
}

func (e *JsonError) Error() string {
	if e.Message == "" {
		return "Json error"
	}
	return e.Message
}

type ProfessorNotFoundError struct {
	Message string
}

func (e *ProfessorNotFoundError) Error() string {
	if e.Message == "" {
		return "Professor not found"
	}
	return e.Message
}

type AlunoNotFoundError struct {
	Message string
}

func (e *AlunoNotFoundError) Error() string {
	if e.Message == "" {
		return "Aluno not found"
	}
	return e.Message
}

type ValidationError struct {
	Message string
	Errors  map[string][]string
}

func (e *ValidationError) Error() string {
	if e.Message == "" {
		return "Validation error"
	}
	return e.Message
}

func (e *ValidationError) HasErrors() bool {
	return len(e.Errors) > 0
}

func (e *ValidationError) AddError(field string, message string) {
	if e.Errors == nil {
		e.Errors = make(map[string][]string)
	}
	e.Errors[field] = append(e.Errors[field], message)
}

func (e *ValidationError) AddErrorIf(condition bool, field string, message string) {
	if condition {
		e.AddError(field, message)
	}
}

type BadCredentialsError struct {
	Message string
}

func (e *BadCredentialsError) Error() string {
	if e.Message == "" {
		return "Bad credentials"
	}
	return e.Message
}

type JwtTokenError struct {
	Message string
}

func (e *JwtTokenError) Error() string {
	if e.Message == "" {
		return "Jwt token error"
	}
	return e.Message
}
