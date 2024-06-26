package handlers

/*
Test for GetAllAPIServicesHandler flow

1. It will test all the incoming params are valid or not should return 400 error with valid error code in response
	i. It will check per_page is present. If not return error
	ii. It will check valid value of the per_page is integer and >1. If not return error
	iii. It will check if page and filter[name][contains] is present or not. If not return error
	iv. It will check valid value of the per_page is integer and >1. If not return error
2. It will check the error scenario for the FindAllAPIServices service api call. If error, handle in response
3. It will check the success scenario for the FindAllAPIServices service api call and return 200
	with reGetAllAPIServicesResponse in response.
*/
