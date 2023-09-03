package data

type Category int

const (
	Inbox Category = iota
	Project
	Area
	Resource
	Archive
)

//type Lifecycle struct {
//}

type Asset struct {
	Id        int
	Content   string
	Category  Category
	CreatedAt int
	UpdatedAt int
}

func GetCategoryAsString(category Category) string {
	switch category {
	case Inbox:
		return "inbox"
	case Project:
		return "project"
	case Area:
		return "area"
	case Resource:
		return "resource"
	case Archive:
		return "archive"
	default:
		return "uncategorized"
	}
}
