package util

import (
	"time"

	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

type NSTime time.Time

const nsTimeLayout = "2006-01-02T15:04:05.000000000"

func (ct *NSTime) UnmarshalEasyJSON(in *jlexer.Lexer) {
	t, err := time.Parse(nsTimeLayout, in.UnsafeString())
	if err != nil {
		in.AddError(err)
		return
	}
	*ct = NSTime(t)
}

func (ct NSTime) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(time.Time(ct).Format(nsTimeLayout))
}

func (ct NSTime) Time() time.Time {
	return time.Time(ct)
}
