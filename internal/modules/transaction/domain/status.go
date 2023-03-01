package domain

type Status int

const (
	REQUESTED Status = iota + 1
	PROCESSING
	SUCCESS
	FAILED
)

func (s Status) String() string {
	values := map[Status]string{
		REQUESTED:  "REQUEST",
		PROCESSING: "PROCESSING",
		SUCCESS:    "SUCCESS",
		FAILED:     "FAILED",
	}

	return values[s]
}
