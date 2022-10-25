package constants

// var MongoCollections map[string]string = map[string]string{
// 	"BLOGS": "blogs",
// }

var MongoCollections = struct{ BLOGS string }{
	BLOGS: "blogs",
}
