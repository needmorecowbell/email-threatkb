
/**
 * Represents a response for a resource not found (404).
 * @returns {Response} The response object.
 */
export const NotFoundResponse = () =>
	new Response(JSON.stringify({
        msg:"Oops, you seem to have went the wrong way",
        success: false}),
    {
		headers: { 'content-type': 'application/json' },
		status: 404,
	});

/**
 * Represents a generic success response (200).
 * @param {string} message - The success message.
 * @param {any} data - The data to be included in the response.
 * @returns {Response} The response object.
 */
export const GenericResponseSuccess = (message, data) =>
	new Response(JSON.stringify({
        msg: message,
        success: true,
        data: data
      }),
     {
		headers: { 'content-type': 'application/json' },
		status: 200,
	 });

/**
 * Represents a generic bad request response (400).
 * @param {string} message - The error message.
 * @returns {Response} The response object.
 */
export const GenericResponseBadRequest = (message) =>
     new Response(JSON.stringify({
        msg: message,
        success: false,
        data: null
      }),
      {
         headers: { 'content-type': 'application/json' },
         status: 400,
      });

/**
 * Represents a generic server error response (500).
 * @param {string} message - The error message.
 * @returns {Response} The response object.
 */
export const GenericResponseServerError = (message) =>
     new Response(JSON.stringify({
        msg: message,
        success: false,
        data: null
      }),
      {
         headers: { 'content-type': 'application/json' },
         status: 500,
      });