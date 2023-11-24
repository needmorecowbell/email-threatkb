
import { GenericResponseSuccess } from "../lib/responses"
/**
 * Endpoint function for the index route.
 * @param {Request} request - The request object.
 * @returns {Response} The response object.
 */
const endpointIndex = async (request,env) => {
	console.debug("Received request to /, returning welcome banner")
	return GenericResponseSuccess("Welcome to the eml funnel worker")
}

export default endpointIndex