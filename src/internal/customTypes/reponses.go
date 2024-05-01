package customTypes

type Message string
type Files []File

type FilesResponse struct {
	Results struct {
		Message Message `json:"message"`
		Data    Files   `json:"data"`
	} `json:"results"`
}

//type DirsResponse struct {
//	Results struct {
//		Message Message `json:"message"`
//		Data    Files   `json:"data"`
//	} `json:"results"`
//}
