package transformer

import (
	"api-reflistdetail-migration-service/model"
	"time"
)

func Transform(mreflists []model.ReferenceListDetailMongo) []model.ReferenceListDetails{
	var reflists []model.ReferenceListDetails
	for _, mreflist := range mreflists {
		var reflist model.ReferenceListDetails
		reflist.DateUpdated = time.Now()
		reflist.ID = mreflist.ID
		reflist.DateCreated = mreflist.DateCreated
		reflist.Description=mreflist.Description
		reflist.Code=mreflist.Code
		reflist.Type=mreflist.Type
		reflist.Name=mreflist.Name
		reflists = append(reflists, reflist)
	}
	return reflists
}
