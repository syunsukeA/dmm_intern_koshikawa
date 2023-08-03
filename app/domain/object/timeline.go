package object

import (
	_ "fmt"
)

type Timeline struct {
	Timeline []*TimelineElement
}

type TimelineElement struct {
	Account *Account  `json:"account,omitempty"`
	Status  *Status	  `json:"status,omitempty"`
}
