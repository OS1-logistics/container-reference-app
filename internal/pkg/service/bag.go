package service

import "fmt"

type BagServiceInterface interface {
	GetBags()
}

type BagService struct {
}

func NewBagService() BagService {
	return BagService{}
}

func (s BagService) GetBags() {
	fmt.Println("invoked GetBags")
}
