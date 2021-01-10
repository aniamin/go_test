package main

import (
	"log"
	"net"

	"go_test/models"
	"go_test/pb"
	context "golang.org/x/net/context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
)


type server struct{
	DB *gorm.DB
}


func get_invalid_deliveries(db *gorm.DB) []*pb.Delivery {

	var results []*pb.Delivery
	
	var sql_cmd = `WITH all_compination AS (
		SELECT suppliers.id AS supplier_id,
			   drivers.id   AS driver_id
		FROM drivers
				 CROSS JOIN suppliers),
		 valid_delivery AS (
			 SELECT DISTINCT suppliers.id AS supplier_id,
							 drivers.id   AS driver_id
			 FROM bean_types
					  LEFT JOIN carrier_bean_types ON bean_types.id = carrier_bean_types.bean_type_id
					  LEFT JOIN carriers ON carriers.id = carrier_bean_types.carrier_id
					  LEFT JOIN drivers ON drivers.carrier_id = carriers.id
					  LEFT JOIN supplier_bean_types ON supplier_bean_types.bean_type_id = carrier_bean_types.bean_type_id
					  LEFT JOIN suppliers ON suppliers.id = supplier_bean_types.supplier_id AND drivers.id IS NOT NULL
		 )
	SELECT all_compination.*
	FROM all_compination
			 LEFT JOIN valid_delivery ON valid_delivery.supplier_id = all_compination.supplier_id AND
										 valid_delivery.driver_id = all_compination.driver_id
	WHERE valid_delivery.supplier_id IS NULL
	  AND valid_delivery.driver_id IS NULL`
	db.Raw(sql_cmd).Scan(&results)

	return results
}


func main() {

	// init db
	db := models.SetupModels()

	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterInvalidDeliveriesServiceServer(s, &server{DB: db})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}


func (s *server) FindInvalidDeliveries(ctx context.Context, r *pb.InvalidDeliveriesRequest) (*pb.InvalidDeliveriesResponse, error) {
	
	var deliveries []*pb.Delivery
	deliveries = get_invalid_deliveries(s.DB)

	return &pb.InvalidDeliveriesResponse{
		Deliveries: deliveries,
	}, nil
}
