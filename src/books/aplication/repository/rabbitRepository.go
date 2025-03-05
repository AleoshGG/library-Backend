package repository

type IRabbit interface {
	NotifyOfLend()
	NotifyOfReturn()
}