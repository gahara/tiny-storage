package constants

const FILE_ROUTE = "files"
const DIR_ROUTE = "dirs"

const (
	DirDoesNotExist        = "Dir does not exist"
	DirAlreadyExists       = "Dir already exists"
	DirCreated             = "Dir was created with name"
	SomethingWentWrong     = "Something went wrong"
	FileNameWasNotProvided = "File name was not provided"
	CouldNotParseResponse  = "Could not parse response"
)

const (
	StatusTextOk      = "OK"
	StatusTextDeleted = "DELETED"
	StatusTextError   = "ERROR"
	StatusTextCreated = "CREATED"
)

const (
	NotFoundMessage           = "Not found"
	SomethingWentWrongMessage = "Something went wrong"
	BadRequest                = "Bad Request"
)

const (
	testDB = "testgorm.db"
	GormDB = "gorm.db"
)
