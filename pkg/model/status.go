package model

type Status int

const (
	REQUESTED Status = iota
	PROCESSING
	SUCCESS
	FAILED
)
