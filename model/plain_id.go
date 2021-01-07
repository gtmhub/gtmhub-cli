package model

type IDResponse struct {
	ID string `json:"id"`
}

type FullIDResponse struct {
	Items []IDResponse `json:"items"`
}

type IDResponses []IDResponse

func (ids IDResponses) ToQueryIDs() string {
	if len(ids) == 0 {
		return ""
	}

	idStr := ""

	for _, id := range ids {
		if len(idStr) == 0 {
			idStr = "\"" + id.ID + "\""
		} else {
			idStr = idStr + ", \"" + id.ID + "\""
		}
	}

	return "[" + idStr + "]"
}
