package models

type Action struct {
	Type        string `json:"type" bson:"type"`
	Description string `json:"description" bson:"description"`
	Deeplink    Deeplink
}

type Deeplink struct {
	Url    string            `json:"url" bson:"url"`
	Params map[string]string `json:"params" bson:"params"`
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
