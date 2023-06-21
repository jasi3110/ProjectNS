package masterRepo

import (
	"OnlineShop/models"
	"OnlineShop/utls"
	"log"
)

type AdImagesInterface interface {
	AdImagesCreate(obj *models.AdImages) (bool, string)
	AdImageGetAll() ([]string, bool, string)

	
}
type AdImagesStruct struct {
}

func (image *AdImagesStruct) AdImagesCreate(obj *models.AdImages) (bool, string) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		log.Panic("DB Disconnceted in AdImages Create")
	}

	url:=models.Imageurl(obj.Imageurl)

	err := Db.QueryRow(`INSERT INTO "adimages" (imageurl,createdon)values($1,$2)RETURNING id`,
url,utls.GetCurrentDate()).Scan(&obj.Id)
	if err != nil {
		log.Panic("Error in AdImages Create QueryRow :", err)
		return false, "Something Went Wrong"
	}
	
	defer func() {
		Db.Close()
	}()
	return true, "Successfully Created"
}

func (image *AdImagesStruct) AdImageGetAll() ([]string, bool, string) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		log.Panic("DB Disconnected in  AdImages GetAll ")
	}
    
	var images [] string
	url:=""
	basicUrl := "https://drive.google.com/uc?export=view&id="
	
	query, err := Db.Query(`SELECT imageurl FROM "adimages" WHERE isdeleted=0`)
	if err != nil {
		log.Panic("Error in AdImages GetAll QueryRow :", err)
		return images, false, "Something Went Wrong"
	}
	for query.Next(){
	err = query.Scan(&url)

	images=append(images,basicUrl + url)
	
	}
	if query.Err(); err != nil {
		log.Panic("Error in AdImages GetAll QueryRow Scan :", err)
		return images, false, "Something Went Wrong"
	}
	
	defer func() {
		Db.Close()
		query.Close()
	}()
	return images, true, "Successfully Completed"
}
 