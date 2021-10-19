package diag

type Diag struct {
	warnings []error
	errors   []error
}

func (d *Diag) SaveError(err error) {
	d.errors = append(d.errors, err)
}

func (d *Diag) SaveWarning(warning error) {
	d.warnings = append(d.warnings, warning)
}

func (d *Diag) HasErrors() bool {
	return len(d.errors) != 0
}

func (d *Diag) GetErrors() []error {
	return d.errors
}
