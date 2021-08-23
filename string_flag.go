package corviper

import (
	"flag"
)

// StringFlag ...
type StringFlag struct {
	flagValue *string
	viperName string
}

// NewStringFlag ...
func NewStringFlag(name, value, usage, viperName string) StringFlag {
	return StringFlag{
		flagValue: flag.String(name, value, usage),
		viperName: viperName,
	}
}

// HasChanged Checks if the value has changed
func (f StringFlag) HasChanged() bool {
	return f.flagValue != nil && len(*f.flagValue) > 0
}

// Name returns the name of the value in Viper
func (f StringFlag) Name() string {
	return f.viperName
}

// ValueString returns the value of the flag as a string
func (f StringFlag) ValueString() string {
	if f.flagValue == nil {
		return ""
	}
	return *f.flagValue
}

// ValueType returns the type of the flag as a string
func (f StringFlag) ValueType() string {
	return "string"
}
