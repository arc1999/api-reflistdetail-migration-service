package service

import (
	"api-reflistdetail-migration-service/dao"
	"api-reflistdetail-migration-service/transformer"
	"log"
	"os"
	"strconv"
)

var d dao.RefListDetailDao

type RefListService struct {
}

func (is RefListService) Migrate() {
	totaldoc, err := d.GetCount()
	if err != nil {
		log.Fatal(err)
	}
	perpage := os.Getenv("N_PER_PAGE")
	nperpage, err := strconv.ParseInt(perpage, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	var i int64
	i=0
	var noofpages = totaldoc/nperpage
	log.Println(totaldoc)
	log.Println(noofpages)
	log.Println(nperpage)
	for i < noofpages {
		mlabtest, err := d.Paginate(i*nperpage, nperpage)
		if err != nil {
			log.Fatal(err)
		}
		labtests := transformer.Transform(mlabtest)

		err = d.BulkInsert(labtests, nperpage)
		if err != nil {
			log.Fatal(err)
		}
		i++
	}
	mlabtest, err := d.Paginate(i*nperpage, totaldoc - (nperpage*(i)))
	if err != nil {
		log.Fatal(err)
	}
	labtests := transformer.Transform(mlabtest)
	err = d.BulkInsert(labtests, nperpage)
	if err != nil {
		log.Fatal(err)
	}
}
