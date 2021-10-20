package storage

type genericResourceMergeFunc func(...Resource) Resource

type autoMergeResourceSet struct {
	mergeFn genericResourceMergeFunc

}

func (rs *autoMergeResourceSet) Next() bool {
	panic("implement me")
}

func (rs *autoMergeResourceSet) At() Resource {
	panic("implement me")
}

func (rs *autoMergeResourceSet) Err() error {
	panic("implement me")
}

func (rs *autoMergeResourceSet) Warnings() Warnings {
	panic("implement me")
}
