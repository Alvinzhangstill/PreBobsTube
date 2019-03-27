package initial

import "github.com/MobileCPX/PreBobsTube/models/web"

func InitDatabase() {
	web.InitTagsJSONToTable()
	web.InitCategoryJOSNToTable()
	web.InitPornStarsJOSNToTable()
	web.InitVideoJOSNToTable()
}
