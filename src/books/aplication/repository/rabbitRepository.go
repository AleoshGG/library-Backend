package repository

type IRabbit interface {
	NotifyOfLend(id_reader int, return_date string)
	NotifyOfReturn(id_reader int)
}