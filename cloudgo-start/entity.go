package server

type Book struct {
	ISDN   string "'json':'isdn'"
	Title  string "'json':'title'"
	Author string "'json':'author'"
	Pages  int    "'json':''pages"
}
type JsonResponse struct {
	Meta interface{} "json:'status'"
	Data interface{} "json:'data'"
}
type JsonErrorResponse struct {
	Error *ApiError "json:'error'"
}
type ApiError struct {
	Status int16  "json:'status'"
	Title  string "json:'title'"
}
