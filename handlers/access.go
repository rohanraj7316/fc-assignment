package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type updateAccessRequest struct {
	FileName         string `json:"fn"`
	FileAccessStatus bool   `json:"access"`
}

// AccessBody in-memory cache for checking
// the access status of the file.
var AccessBody map[string]bool = make(map[string]bool)

// UpdateAccess update file permitions
func UpdateAccess(w http.ResponseWriter, r *http.Request) error {
	var updateAccessObject updateAccessRequest

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	json.Unmarshal(reqBody, &updateAccessObject)

	if _, ok := AccessBody[updateAccessObject.FileName]; ok {
		AccessBody[updateAccessObject.FileName] = updateAccessObject.FileAccessStatus
	} else {
		AccessBody[updateAccessObject.FileName] = updateAccessObject.FileAccessStatus
	}

	/** DEP - add the access in file */
	// TODO - couldn't able to update file access into the ACCESS.txt.
	// fd, err := utils.GetAccessFileDiscriptor()
	// if err != nil {
	// 	return err
	// }
	// defer fd.Close()

	// _, err = fmt.Fprintln(fd, updateAccessObject)
	// if err != nil {
	// 	return err
	// }

	return nil
}
