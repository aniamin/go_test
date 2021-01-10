package models

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm/clause"
)


func SetupModels() *gorm.DB {

	dsn := "host=postgres user=postgres password=admin123 dbname=postgresdb port=5432 sslmode=disable TimeZone=Asia/Dhaka"
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	
	db.AutoMigrate(
		&BeanType{}, 
		&Carrier{}, 
		&CarrierBeanType{}, 
		&Delivery{}, 
		&Driver{}, 
		&Supplier{}, 
		&SupplierBeanType{},
	)

	// Initialize value
	var bean_types = []BeanType {
		{ID: 1, Name: "Arabica"}, 
		{ID: 2, Name: "Robusta"}, 
		{ID: 3, Name: "Liberica"},
	}
	
	db.Clauses(clause.OnConflict{DoNothing: true}).Create( &bean_types )

	var carriers = []Carrier {
		{ID: 1, Name: "FedEx"}, 
		{ID: 2, Name: "UPS"}, 
		{ID: 3, Name: "USPS"}, 
		{ID: 4, Name: "DHL"}, 
		{ID: 5, Name: "EMS"},
	}

	db.Clauses(clause.OnConflict{DoNothing: true}).Create( &carriers )

	var carrier_bean_types = []CarrierBeanType {
		{ID: 1, Carrier_id: 1, Bean_type_id: 1},
		{ID: 2, Carrier_id: 1, Bean_type_id: 2},
		{ID: 3, Carrier_id: 2, Bean_type_id: 2},
		{ID: 4, Carrier_id: 4, Bean_type_id: 3},
		{ID: 5, Carrier_id: 5, Bean_type_id: 1},
		{ID: 6, Carrier_id: 5, Bean_type_id: 2},
	}

	db.Clauses(clause.OnConflict{DoNothing: true}).Create( &carrier_bean_types )

	var deliveries = []Delivery {
		
		{ID: 1, SupplierId:	1, DriverId: 1},
		{ID: 2, SupplierId:	8, DriverId: 2},
		{ID: 3, SupplierId:	8, DriverId: 3},
		{ID: 4, SupplierId:	4, DriverId: 3},
	}

	db.Clauses(clause.OnConflict{DoNothing: true}).Create( &deliveries )

	var drivers = []Driver {
		{ID: 1, Name: "Driver 1", CarrierId: 1},
		{ID: 2, Name: "Driver 2", CarrierId: 2},
		{ID: 3, Name: "Driver 3", CarrierId: 4},
	}
	
	db.Clauses(clause.OnConflict{DoNothing: true}).Create( &drivers )


	var suppliers = []Supplier {
		{ID: 1, Name: "Vietnam" },
		{ID: 2, Name: "Brazil" },
		{ID: 3, Name: "Colombia" },
		{ID: 4, Name: "Indonesia" },
		{ID: 5, Name: "Honduras" },
		{ID: 6, Name: "India" },
		{ID: 7, Name: "Ethioia" },
		{ID: 8, Name: "Uganda" },
	}
	
	db.Clauses(clause.OnConflict{DoNothing: true}).Create( &suppliers )
	

	var supplier_bean_types = []SupplierBeanType {
		{ID: 1 , SupplierId: 1, BeanTypeId: 2},	
		{ID: 2 , SupplierId: 4, BeanTypeId: 2},	
		{ID: 3 , SupplierId: 4, BeanTypeId: 1},	
		{ID: 4 , SupplierId: 6, BeanTypeId: 2},	
		{ID: 5 , SupplierId: 6, BeanTypeId: 1},	
		{ID: 6 , SupplierId: 2, BeanTypeId: 1},	
		{ID: 7 , SupplierId: 3, BeanTypeId: 1},	
		{ID: 8 , SupplierId: 7, BeanTypeId: 1},	
		{ID: 9 , SupplierId: 8, BeanTypeId: 1},	
		{ID: 10, SupplierId: 8, BeanTypeId:	2},
		{ID: 11, SupplierId: 5, BeanTypeId:	3},
		{ID: 12, SupplierId: 2, BeanTypeId:	3},
	}
	
	db.Clauses(clause.OnConflict{DoNothing: true}).Create( &supplier_bean_types )
	


	return db
}