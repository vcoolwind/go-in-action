package errs

type bizerror struct {
	errs string
}
func (this bizerror) Error() string{
	return this.errs
}

func NewBizError(msg string) error{
	return  bizerror{errs:msg}
}