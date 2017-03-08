package siridb

// CprotoReqQuery for sending queries
const CprotoReqQuery = 0

// CprotoReqInsert for sending inserts
const CprotoReqInsert = 1

// CprotoReqAuth for authentication
const CprotoReqAuth = 2

// CprotoReqPing for ping on the connection
const CprotoReqPing = 3

// CprotoReqInfo for requesting database info
const CprotoReqInfo = 4

// CprotoReqLoadDB for loading a new database
const CprotoReqLoadDB = 5

// CprotoReqRegisterServer for registering a new server
const CprotoReqRegisterServer = 6

// CprotoReqFileServers for requesting a server.dat file
const CprotoReqFileServers = 7

// CprotoReqFileUsers for requesting a users.dat file
const CprotoReqFileUsers = 8

// CprotoReqFileGroups for requesting a groups.dat file
const CprotoReqFileGroups = 9

// CprotoResQuery on query response
const CprotoResQuery = 0

// CprotoResInsert on insert response
const CprotoResInsert = 1

// CprotoResAuthSuccess on authentication success
const CprotoResAuthSuccess = 2

// CprotoResAck on ack
const CprotoResAck = 3

// CprotoResInfo on database info response
const CprotoResInfo = 4

// CprotoResFile on request file response
const CprotoResFile = 5

// CprotoErrMsg general error
const CprotoErrMsg = 64

// CprotoErrQuery on query error
const CprotoErrQuery = 65

// CprotoErrInsert on insert error
const CprotoErrInsert = 66

// CprotoErrServer on server error
const CprotoErrServer = 67

// CprotoErrPool on server error
const CprotoErrPool = 68

// CprotoErrUserAccess on server error
const CprotoErrUserAccess = 69

// CprotoErr on server error
const CprotoErr = 70

// CprotoErrNotAuthenticated on server error
const CprotoErrNotAuthenticated = 71

// CprotoErrAuthCredentials on server error
const CprotoErrAuthCredentials = 72

// CprotoErrAuthUnknownDb on server error
const CprotoErrAuthUnknownDb = 73

// CprotoErrLoadingDb on server error
const CprotoErrLoadingDb = 74

// CprotoErrFile on server error
const CprotoErrFile = 75
