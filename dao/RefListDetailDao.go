package dao

import (
	"api-reflistdetail-migration-service/db"
	"api-reflistdetail-migration-service/model"
	"context"
	log "github.com/sirupsen/logrus"
	gormbulk "github.com/t-tiger/gorm-bulk-insert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

type RefListDetailDao struct {
}

func (d RefListDetailDao) Paginate(pagenumber int64, nperpage int64) ([]model.ReferenceListDetailMongo, error) {
	log.Println(pagenumber,nperpage)
	findOptions := options.Find()
	findOptions.SetLimit(nperpage)
	findOptions.SetSort(bson.M{})
	findOptions.SetSkip(pagenumber)

	db := db.GetMongoDB()
	cur, err := db.Collection(os.Getenv("DATA_MONGODB_COLLECTION")).Find(context.TODO(), bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())
	var jobs []model.ReferenceListDetailMongo
	for cur.Next(context.TODO()) {
		var job model.ReferenceListDetailMongo
		err := cur.Decode(&job)
		if err != nil {
			return nil, err
		}
		jobs = append(jobs, job)
	}
	log.Println(len(jobs))
	return jobs, nil
}
func (d RefListDetailDao) GetCount() (int64, error) {
	db := db.GetMongoDB()
	return db.Collection(os.Getenv("DATA_MONGODB_COLLECTION")).CountDocuments(context.TODO(),bson.D{})
}
func (d RefListDetailDao) BulkInsert(Entity []model.ReferenceListDetails, nperpage int64) error {
	sqldb := db.GetMysqlDB()
	b := make([]interface{}, len(Entity))
	for i := range Entity {
		b[i] = Entity[i]
	}
	err := gormbulk.BulkInsert(sqldb, b, int(nperpage))
	if err != nil {
		log.Printf("error in saving labtest")
		return err
	}
	return nil

}
