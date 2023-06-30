package service

func Go(fn func()) {
	go fn()
}
