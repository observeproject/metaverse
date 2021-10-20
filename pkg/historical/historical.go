package historical

import (
	"github.com/prometheus/common/model"
)

type RecordString struct {
	Timestamp int64
	Value     string
}

type RecordNumber struct {
	Timestamp int64
	Value     uint64
}

type RecordStrings []RecordString

func (ss RecordStrings) Len() int {
	return len(ss)
}

func (ss RecordStrings) Less(i, j int) bool {
	return ss[i].Timestamp < ss[j].Timestamp
}

func (ss RecordStrings) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}

type RecordNumbers []RecordNumber

func (ns RecordNumbers) Len() int {
	return len(ns)
}

func (ns RecordNumbers) Less(i, j int) bool {
	return ns[i].Timestamp < ns[j].Timestamp
}

func (ns RecordNumbers) Swap(i, j int) {
	ns[i], ns[j] = ns[j], ns[i]
}

type TimePair struct {
	Begin model.Time
	End   model.Time
}

type TimePairs []TimePair

func (ts TimePairs) Len() int {
	return len(ts)
}

func (ts TimePairs) Less(i, j int) bool {
	return ts[i].Begin < ts[j].Begin || ts[i].End < ts[j].End
}

func (ts TimePairs) Swap(i, j int) {
	ts[i], ts[j] = ts[j], ts[i]
}
