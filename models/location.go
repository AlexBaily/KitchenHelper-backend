package models

//Record struct that will house the DynamoDB records.
type LocRecord struct {
	UserID            string `json:UserID`
	ProductIdentifier string `json:productIdentifier`
	Location          string `json:"location"`
}

type ProductRecord struct {
	ProductIdentifier string `json:productIdentifier`
	ProductName       string `json:productName`
	Quantity          int    `json:quantity`
}
