package handlers

/*
Test for GetAPIServiceDetailsHandler flow

1. It will test all the incoming params service_id is present and valid or not. It should return 400 error with
	valid error code in response in case it's invalid
2. It will check the error scenario for the FindAPIServiceDetails service api call. If error, handle in response
3. It will check the success scenario for the FindAPIServiceDetails service api call and return 200
	with reGetAllAPIServicesResponse in response.
*/
