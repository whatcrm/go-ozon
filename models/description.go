package models

type TreeCategoryResponse struct {
	Result []struct {
		DescriptionCategoryId int    `json:"description_category_id"`
		CategoryName          string `json:"category_name"`
		Disabled              bool   `json:"disabled"`
		Children              []struct {
			DescriptionCategoryId int    `json:"description_category_id"`
			CategoryName          string `json:"category_name"`
			Disabled              bool   `json:"disabled"`
			Children              []struct {
				TypeName string        `json:"type_name"`
				TypeId   int           `json:"type_id"`
				Disabled bool          `json:"disabled"`
				Children []interface{} `json:"children"`
			} `json:"children"`
		} `json:"children"`
	} `json:"result"`
}

type AttributeCategoryRequest struct {
	DescriptionCategoryId int    `json:"description_category_id"`
	Language              string `json:"language"`
	TypeId                int    `json:"type_id"`
}

type AttributeCategoryResponse struct {
	Result []struct {
		Id                  int    `json:"id"`
		AttributeComplexId  int    `json:"attribute_complex_id"`
		Name                string `json:"name"`
		Description         string `json:"description"`
		Type                string `json:"type"`
		IsCollection        bool   `json:"is_collection"`
		IsRequired          bool   `json:"is_required"`
		IsAspect            bool   `json:"is_aspect"`
		MaxValueCount       int    `json:"max_value_count"`
		GroupName           string `json:"group_name"`
		GroupId             int    `json:"group_id"`
		DictionaryId        int    `json:"dictionary_id"`
		CategoryDependent   bool   `json:"category_dependent"`
		ComplexIsCollection bool   `json:"complex_is_collection"`
	} `json:"result"`
}

type AttributeValueRequest struct {
	AttributeId           int    `json:"attribute_id"`
	DescriptionCategoryId int    `json:"description_category_id"`
	Language              string `json:"language,omitempty"`
	LastValueId           int    `json:"last_value_id,omitempty"`
	Limit                 int    `json:"limit"`
	TypeId                int    `json:"type_id"`
	Value                 string `json:"value,omitempty"`
}

type AttributeValueResponse struct {
	Result []struct {
		Id      int    `json:"id"`
		Value   string `json:"value"`
		Info    string `json:"info"`
		Picture string `json:"picture"`
	} `json:"result"`
	HasNext bool `json:"has_next,omitempty"`
}
