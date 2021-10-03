package services

import (
	"awesomeProject1/models"
	"awesomeProject1/repos/inmemoryrepos"
	"github.com/stretchr/testify/suite"
	"sync"
)

type ProductServiceImplTest struct{
	datastore *models.InmemoryDatastore
	repos  *inmemoryrepos.ProductRepoImpl
	mutex *sync.Mutex
	suite.Suite
}

func (pst *ProductServiceImplTest) SetUpTest() {

}

//func (pst *ProductServiceImplTest) BeforeTest(suitename ,testname string) {
//	pst.datastore:=InitialiseInmemoryDatastore()
//
//}
