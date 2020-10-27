package main
import "api-reflistdetail-migration-service/service"

var s service.RefListService

func main() {
	s.Migrate()
}
