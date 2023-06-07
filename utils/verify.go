package utils

var (
	IdVerify             = Rules{"ID": []string{NotEmpty()}}
	PageInfoVerify       = Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}}
	LoginVerify          = Rules{"Username": {NotEmpty()}, "NickName": {NotEmpty()}, "Password": {NotEmpty()}}
	MenuMetaVerify       = Rules{"Title": {NotEmpty()}}
	RegisterVerify       = Rules{"Username": {NotEmpty()}, "NickName": {NotEmpty()}, "Password": {NotEmpty()}, "AuthorityID": {NotEmpty()}}
	AuthorityVerify      = Rules{"AuthorityId": {NotEmpty()}, "AuthorityName": {NotEmpty()}}
	AuthorityIdVerify    = Rules{"AuthorityId": {NotEmpty()}}
	ChangePasswordVerify = Rules{"Password": {NotEmpty()}, "NewPassword": {NotEmpty()}}
	SiteConfigVerify     = Rules{"ParentName": {NotEmpty()}, "SiteName": {NotEmpty()}, "SiteKey": {NotEmpty()}, "SiteID": {NotEmpty()}, "DirectPlayUrl": {NotEmpty()}, "CFPlayUrl": {NotEmpty()}, "CDNPlayUrl": {NotEmpty()}, "VideoCover": {NotEmpty()}, "DownloadUrl": {NotEmpty()}}
)
