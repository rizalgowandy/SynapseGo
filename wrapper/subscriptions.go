package wrapper

/********** GLOBAL VARIABLES **********/
const _subscriptionsURL = _url + "/subscriptions"

/********** METHODS **********/

// GetSubscriptions returns all the subscriptions
func (c *ClientCredentials) GetSubscriptions(cred ClientCredentials) Users {
	return handleRequestMulti(cred, "GET", _subscriptionsURL, "subscriptions ", nil)
}
