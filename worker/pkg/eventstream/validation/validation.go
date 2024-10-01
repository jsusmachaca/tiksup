package validation

import "go.mongodb.org/mongo-driver/bson/primitive"

func IsValidObjectID(s string) string {
	objID, err := primitive.ObjectIDFromHex(s)
	if err != nil {
		randomId := "c5ecece22fb00fa4da29cf01"
		return randomId
	}
	return objID.Hex()
}
