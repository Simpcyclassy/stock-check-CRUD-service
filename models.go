package main

type Stock struct {
	amountAvailable		uint
	amountSold			int
	available			bool
	description			string
	location			string
	name				string
	price				float64
}

type StockList []Stock 