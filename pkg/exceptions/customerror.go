package exceptions

const (
	ERRDOMAIN       Status = 1
	ERRBUSSINESS    Status = 2
	ERRSYSTEM       Status = 3
	ERRNOTFOUND     Status = 4
	ERRREPOSITORY   Status = 5
	ERRUNKNOWN      Status = 6
	ERRAUTHORIZED   Status = 7
	ERRFORBIDDEN    Status = 8
	ERRPRECONDITION Status = 9
)

type Status int

type CustomError struct {
	Status Status
	Errors error
}
