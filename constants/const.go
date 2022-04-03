package constants

const RadiusToSearchDriver = 9750000

const ErrorInvalidCoordinates = "longitude and latitude should be in the right range " +
	"(-180<=longitude<=180 and -90<=latitude<=90)"
const ErrorUnprocessableCoordinates = "longitude and latitude should be number and not empty"
const ErrorURLNotFound = "requested url does not exist"
const ErrorDriverApiDoesNotRespond = "driver api does not respond"
const ErrorTokenCreation = "token could not be created"
const ErrorMalformedMissingToken = "Missing or malformed JWT"
const ErrorAuthHeaderRequired = "authorization header is required"
const ErrorWrongAuthHeader = "wrong authorization header"
const ErrorWrongSigningMethod = "unexpected signing method"
const ErrorInvalidToken = "invalid token"
