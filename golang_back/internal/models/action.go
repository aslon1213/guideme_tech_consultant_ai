package models

type Action struct {
	Type    string `json:"type" bson:"type"`
	Element string `json:"element" bson:"element"`
	Value   string `json:"value" bson:"value"`
}

type Document struct {
	TypeofDocument string `json:"typeofdocument" bson:"typeofdocument"`
	Name           string `json:"documentname" bson:"documentname"`
	UploadedPath   string `json:"uploadedpath" bson:"uploadedpath"`
}

func (d *Document) IsSupportedTypeofDocument() bool {

	supported_types := []string{"pdf", "txt"}
	for _, v := range supported_types {
		if v == d.TypeofDocument {
			return true
		}
	}
	return false
}
