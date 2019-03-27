package enums

type JsonResultCode int

// 视频排序方式
type VideoSortType int

//Category 排序方式
type CategorySortType int

const (
	JRCodeSucc JsonResultCode = iota
	JRCodeFailed
	JRCode302 = 302 //跳转至地址
	JRCode401 = 401 //未授权访问
)

const (
	Deleted = iota - 1
	Disabled
	Enabled
)

// 视频排序方式
const (
	VSMostView VideoSortType = iota
	VSNewVideo
	TopRateVideo
	MostViewsVideo
)


// Category 排序方式
const (
	DefaultSort CategorySortType = iota
	Alphabetically
	TopRateCategory
	MostViewCategory
	MostVidesCategory
)