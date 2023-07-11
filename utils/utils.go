package utils

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
)

type PrefixedFormatter struct {
	Prefix string
	logrus.Formatter
	TextColor color.Attribute
}

func (f *PrefixedFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	prefix := color.New(f.TextColor).SprintFunc()(f.Prefix)
	entry.Message = fmt.Sprintf("%s %s", prefix, entry.Message)
	return f.Formatter.Format(entry)
}
