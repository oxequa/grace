package grace

type (
	printer struct {
		Err    Error
		enable bool
		fprint PrintFunc
	}
	PrintFunc func(e error)
)

func (p *printer) state() bool {
	return p.enable
}

func (p *printer) stamp(e error) {
	if p != nil && p.enable {
		p.fprint(e)
	}
}

func (p *printer) printf() PrintFunc {
	return p.fprint
}
