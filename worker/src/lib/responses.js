// 404 not found
export const NotFoundResponse = () =>
	new Response(JSON.stringify({
        msg:"Oops, you seem to have went the wrong way",
        success: false}),
    {
		headers: { 'content-type': 'application/json' },
		status: 404,
	});

// 200 Success
export const GenericResponseSuccess = (message,data) =>
	new Response(JSON.stringify({
        msg:message,
        success: true,
        data: data
      }),
     {
		headers: { 'content-type': 'application/json' },
		status: 200,
	 });
    


// 400 Bad Reques
export const GenericResponseBadRequest = (message) =>
     new Response(JSON.stringify({
        msg:message,
        success: false,
        data: null
      }),
      {
         headers: { 'content-type': 'application/json' },
         status: 400,
      });
 

// 500 Error code
export const GenericResponseServerError = (message) =>
     new Response(JSON.stringify({
        msg:message,
        success: false,
        data: null
      }),
      {
         headers: { 'content-type': 'application/json' },
         status: 500,
      });