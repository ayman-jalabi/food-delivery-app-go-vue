package helpers

import (
	"main/models"
)

//func ItemContains(items []models.ItemJson, item models.ItemJson) bool {
//	for _, item := range items {
//		if item == items  {
//			return true
//		}
//	}
//	return false
//}

// StringContains is a simple function to check if a string slice that's being looped over contains an item or not
func StringContains(slice []string, targetString string) bool {
	for _, item := range slice {
		if item == targetString {
			return true
		}
	}
	return false
}

// OpeningAndClosingHoursContains is a simple function to check if a WorkingHoursJson slice that's being looped over contains an item or not
func OpeningAndClosingHoursContains(workingHours []models.WorkingHoursJson, opening string, closing string) bool {
	for _, hours := range workingHours {
		if hours.Opening == opening && hours.Closing == closing {
			return true
		}
	}
	return false
}
