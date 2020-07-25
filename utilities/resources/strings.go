package resources

func prepareMessage(message string) map[string]interface{} {
	result := make(map[string]interface{})
	result["error"] = message
	return result
}

var NotFound = prepareMessage("Not Found")
var InvalidCredintials = prepareMessage("Invalid username or password")
var InvalidRequest = prepareMessage("Invalid request")
var UsernameLengthCannotBeLessThan3Chars = prepareMessage("Username cannot be less than 3 characters")
var PasswordLengthCannotBeLessThan6Chars = prepareMessage("Password cannot be less than 6 characters")
var UnableToAccessTheDatabase = prepareMessage("Unable to access the database")
var UsernameAlreadyExists = prepareMessage("The specified username is already exists. Please select another username.")
var CannotCreateToken = prepareMessage("Cannot create token due to a server error")
var CannotParseToken = prepareMessage("Cannot parse the specified token")
var InvalidPage = prepareMessage("Invalid page")
var UnableToUpdateTheRecord = prepareMessage("Unable to update the database record")
