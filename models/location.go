package models

//Record struct that will house the DynamoDB records.
type LocRecord struct {
	UserID       string 
	ProductIdentifier	 string  
	Location	string `json:"location"`
}

/*type productRecord struct {
	UserID       string
	ProductIdentifier string
	ProductName string
	Quantity       int
}*/
