package api_endpoints

import "fmt"

func GetAllCustomDimensions() string {
	return "/customdimensions"
}

func GetCustomDimension(id ResourceID) string {
	return fmt.Sprintf("/customdimension/%v", id)
}

func CreateCustomDimension() string {
	return "/customdimension"
}

func UpdateCustomDimension(id ResourceID) string {
	return fmt.Sprintf("/customdimension/%v", id)
}

func DeleteCustomDimension(id ResourceID) string {
	return fmt.Sprintf("/customdimension/%v", id)
}

func CreatePopulator(customDimensionID ResourceID) string {
	return fmt.Sprintf("/customdimension/%v/populator", customDimensionID)
}

func UpdatePopulator(customDimensionID ResourceID, populatorID ResourceID) string {
	return fmt.Sprintf("/customdimension/%v/populator/%v", customDimensionID, populatorID)
}

func DeletePopulator(customDimensionID ResourceID, populatorID ResourceID) string {
	return fmt.Sprintf("/customdimension/%v/populator/%v", customDimensionID, populatorID)
}
