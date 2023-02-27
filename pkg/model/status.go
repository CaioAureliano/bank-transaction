package model

type Status int

const (
	REQUESTED Status = iota + 1
	PROCESSING
	SUCCESS
	FAILED
)
