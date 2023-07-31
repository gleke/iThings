package relationDB

import (
	"time"
)

// 仓库

// StockLocation Inventory Locations
type StockLocation struct {
	ID                       int32     `gorm:"column:id;primary_key" json:"id"`
	LocationID               int32     `gorm:"column:location_id" json:"location_id"`                               // Parent Location
	Posx                     int32     `gorm:"column:posx" json:"posx"`                                             // Corridor (X)
	Posy                     int32     `gorm:"column:posy" json:"posy"`                                             // Shelves (Y)
	Posz                     int32     `gorm:"column:posz" json:"posz"`                                             // Height (Z)
	CompanyID                int32     `gorm:"column:company_id" json:"company_id"`                                 // Company
	RemovalStrategyID        int32     `gorm:"column:removal_strategy_id" json:"removal_strategy_id"`               // Removal Strategy
	CyclicInventoryFrequency int32     `gorm:"column:cyclic_inventory_frequency" json:"cyclic_inventory_frequency"` // Inventory Frequency (Days)
	WarehouseID              int32     `gorm:"column:warehouse_id" json:"warehouse_id"`                             // Warehouse
	StorageCategoryID        int32     `gorm:"column:storage_category_id" json:"storage_category_id"`               // Storage Category
	CreateUId                int32     `gorm:"column:create_uid" json:"create_uid"`                                 // Created by
	WriteUId                 int32     `gorm:"column:write_uid" json:"write_uid"`                                   // Last Updated by
	Name                     string    `gorm:"column:name" json:"name"`                                             // Location Name
	CompleteName             string    `gorm:"column:complete_name" json:"complete_name"`                           // Full Location Name
	Usage                    string    `gorm:"column:usage" json:"usage"`                                           // Location Type
	ParentPath               string    `gorm:"column:parent_path" json:"parent_path"`                               // Parent Path
	Barcode                  string    `gorm:"column:barcode" json:"barcode"`                                       // Barcode
	LastInventoryDate        time.Time `gorm:"column:last_inventory_date" json:"last_inventory_date"`               // Last Effective Inventory
	NextInventoryDate        time.Time `gorm:"column:next_inventory_date" json:"next_inventory_date"`               // Next Expected Inventory
	Comment                  string    `gorm:"column:comment" json:"comment"`                                       // Additional Information
	Active                   bool      `gorm:"column:active" json:"active"`                                         // Active
	ScrapLocation            bool      `gorm:"column:scrap_location" json:"scrap_location"`                         // Is a Scrap Location?
	ReturnLocation           bool      `gorm:"column:return_location" json:"return_location"`                       // Is a Return Location?
	ReplenishLocation        bool      `gorm:"column:replenish_location" json:"replenish_location"`                 // Replenish Location
	CreateDate               time.Time `gorm:"column:create_date" json:"create_date"`                               // Created on
	WriteDate                time.Time `gorm:"column:write_date" json:"write_date"`                                 // Last Updated on
	ValuationInAccountID     int32     `gorm:"column:valuation_in_account_id" json:"valuation_in_account_id"`       // Stock Valuation Account (Incoming)
	ValuationOutAccountID    int32     `gorm:"column:valuation_out_account_id" json:"valuation_out_account_id"`     // Stock Valuation Account (Outgoing)
}

// TableName the name of table in database
func (t *StockLocation) TableName() string {
	return "stock_location"
}

// StockMove Stock Move
type StockMove struct {
	ID                    int32       `gorm:"column:id;primary_key" json:"id"`
	Sequence              int32       `gorm:"column:sequence" json:"sequence"`                                 // Sequence
	CompanyID             int32       `gorm:"column:company_id" json:"company_id"`                             // Company
	ProductID             int32       `gorm:"column:product_id" json:"product_id"`                             // Product
	ProductUom            int32       `gorm:"column:product_uom" json:"product_uom"`                           // UoM
	LocationID            int32       `gorm:"column:location_id" json:"location_id"`                           // Source Location
	LocationDestID        int32       `gorm:"column:location_dest_id" json:"location_dest_id"`                 // Destination Location
	PartnerID             int32       `gorm:"column:partner_id" json:"partner_id"`                             // Destination Address
	PickingID             int32       `gorm:"column:picking_id" json:"picking_id"`                             // Transfer
	GroupID               int32       `gorm:"column:group_id" json:"group_id"`                                 // Procurement Group
	RuleID                int32       `gorm:"column:rule_id" json:"rule_id"`                                   // Stock Rule
	PickingTypeID         int32       `gorm:"column:picking_type_id" json:"picking_type_id"`                   // Operation Type
	OriginReturnedMoveID  int32       `gorm:"column:origin_returned_move_id" json:"origin_returned_move_id"`   // Origin return move
	RestrictPartnerID     int32       `gorm:"column:restrict_partner_id" json:"restrict_partner_id"`           // Owner
	WarehouseID           int32       `gorm:"column:warehouse_id" json:"warehouse_id"`                         // Warehouse
	PackageLevelID        int32       `gorm:"column:package_level_id" json:"package_level_id"`                 // Package Level
	NextSerialCount       int32       `gorm:"column:next_serial_count" json:"next_serial_count"`               // Number of SN
	OrderpointID          int32       `gorm:"column:orderpoint_id" json:"orderpoint_id"`                       // Original Reordering Rule
	ProductPackagingID    int32       `gorm:"column:product_packaging_id" json:"product_packaging_id"`         // Packaging
	CreateUId             int32       `gorm:"column:create_uid" json:"create_uid"`                             // Created by
	WriteUId              int32       `gorm:"column:write_uid" json:"write_uid"`                               // Last Updated by
	Name                  string      `gorm:"column:name" json:"name"`                                         // Description
	Priority              string      `gorm:"column:priority" json:"priority"`                                 // Priority
	State                 string      `gorm:"column:state" json:"state"`                                       // Status
	Origin                string      `gorm:"column:origin" json:"origin"`                                     // Source Document
	ProcureMethod         string      `gorm:"column:procure_method" json:"procure_method"`                     // Supply Method
	Reference             string      `gorm:"column:reference" json:"reference"`                               // Reference
	NextSerial            string      `gorm:"column:next_serial" json:"next_serial"`                           // First SN
	ReservationDate       time.Time   `gorm:"column:reservation_date" json:"reservation_date"`                 // Date to Reserve
	DescriptionPicking    string      `gorm:"column:description_picking" json:"description_picking"`           // Description of Picking
	ProductQty            interface{} `gorm:"column:product_qty" json:"product_qty"`                           // Real Quantity
	ProductUomQty         interface{} `gorm:"column:product_uom_qty" json:"product_uom_qty"`                   // Demand
	QuantityDone          interface{} `gorm:"column:quantity_done" json:"quantity_done"`                       // Quantity Done
	Scrapped              bool        `gorm:"column:scrapped" json:"scrapped"`                                 // Scrapped
	PropagateCancel       bool        `gorm:"column:propagate_cancel" json:"propagate_cancel"`                 // Propagate cancel and split
	IsInventory           bool        `gorm:"column:is_inventory" json:"is_inventory"`                         // Inventory
	Additional            bool        `gorm:"column:additional" json:"additional"`                             // Whether the move was added after the picking's confirmation
	Date                  time.Time   `gorm:"column:date" json:"date"`                                         // Date Scheduled
	DateDeadline          time.Time   `gorm:"column:date_deadline" json:"date_deadline"`                       // Deadline
	DelayAlertDate        time.Time   `gorm:"column:delay_alert_date" json:"delay_alert_date"`                 // Delay Alert Date
	CreateDate            time.Time   `gorm:"column:create_date" json:"create_date"`                           // Created on
	WriteDate             time.Time   `gorm:"column:write_date" json:"write_date"`                             // Last Updated on
	PriceUnit             interface{} `gorm:"column:price_unit" json:"price_unit"`                             // Unit Price
	AnalyticAccountLineID int32       `gorm:"column:analytic_account_line_id" json:"analytic_account_line_id"` // Analytic Account Line
	ToRefund              bool        `gorm:"column:to_refund" json:"to_refund"`                               // Update quantities on SO/PO
	SaleLineID            int32       `gorm:"column:sale_line_id" json:"sale_line_id"`                         // Sale Line
	PurchaseLineID        int32       `gorm:"column:purchase_line_id" json:"purchase_line_id"`                 // Purchase Order Line
	CreatedPurchaseLineID int32       `gorm:"column:created_purchase_line_id" json:"created_purchase_line_id"` // Created Purchase Order Line
}

// TableName the name of table in database
func (t *StockMove) TableName() string {
	return "stock_move"
}

// StockMoveLine Product Moves (Stock Move Line)
type StockMoveLine struct {
	ID                  int32       `gorm:"column:id;primary_key" json:"id"`
	PickingID           int32       `gorm:"column:picking_id" json:"picking_id"`                       // Transfer
	MoveID              int32       `gorm:"column:move_id" json:"move_id"`                             // Stock Operation
	CompanyID           int32       `gorm:"column:company_id" json:"company_id"`                       // Company
	ProductID           int32       `gorm:"column:product_id" json:"product_id"`                       // Product
	ProductUomID        int32       `gorm:"column:product_uom_id" json:"product_uom_id"`               // Unit of Measure
	PackageID           int32       `gorm:"column:package_id" json:"package_id"`                       // Source Package
	PackageLevelID      int32       `gorm:"column:package_level_id" json:"package_level_id"`           // Package Level
	LotID               int32       `gorm:"column:lot_id" json:"lot_id"`                               // Lot/Serial Number
	ResultPackageID     int32       `gorm:"column:result_package_id" json:"result_package_id"`         // Destination Package
	OwnerID             int32       `gorm:"column:owner_id" json:"owner_id"`                           // From Owner
	LocationID          int32       `gorm:"column:location_id" json:"location_id"`                     // From
	LocationDestID      int32       `gorm:"column:location_dest_id" json:"location_dest_id"`           // To
	CreateUId           int32       `gorm:"column:create_uid" json:"create_uid"`                       // Created by
	WriteUId            int32       `gorm:"column:write_uid" json:"write_uid"`                         // Last Updated by
	ProductCategoryName string      `gorm:"column:product_category_name" json:"product_category_name"` // Product Category
	LotName             string      `gorm:"column:lot_name" json:"lot_name"`                           // Lot/Serial Number Name
	State               string      `gorm:"column:state" json:"state"`                                 // Status
	Reference           string      `gorm:"column:reference" json:"reference"`                         // Reference
	DescriptionPicking  string      `gorm:"column:description_picking" json:"description_picking"`     // Description picking
	ReservedQty         interface{} `gorm:"column:reserved_qty" json:"reserved_qty"`                   // Real Reserved Quantity
	ReservedUomQty      interface{} `gorm:"column:reserved_uom_qty" json:"reserved_uom_qty"`           // Reserved
	QtyDone             interface{} `gorm:"column:qty_done" json:"qty_done"`                           // Done
	Date                time.Time   `gorm:"column:date" json:"date"`                                   // Date
	CreateDate          time.Time   `gorm:"column:create_date" json:"create_date"`                     // Created on
	WriteDate           time.Time   `gorm:"column:write_date" json:"write_date"`                       // Last Updated on
}

// TableName the name of table in database
func (t *StockMoveLine) TableName() string {
	return "stock_move_line"
}
