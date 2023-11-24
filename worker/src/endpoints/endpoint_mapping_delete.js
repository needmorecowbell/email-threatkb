
import { GenericResponseBadRequest, GenericResponseSuccess, GenericResponseServerError } from "../lib/responses";
import { cacheEmailMappingDelete } from "../lib/kv";

/**
 * Endpoint function for deleting an email mapping from the cache.
 * 
 * @param {Request} request - The request object.
 * @param {Object} env - The environment object.
 * @returns {Response} - The response object indicating the success or failure of the operation.
 */
export const endpointMappingDelete = async (request, env) => {
    let body = await request.json()
    if (body.forward_to === undefined || body.gateway_address === undefined) {
        return GenericResponseBadRequest("forward_to and gateway_address are required")
    }

    try {
        let result = await cacheEmailMappingDelete(body.forward_to, body.gateway_address, env)
        if (result === undefined) {
            return GenericResponseServerError("Failed to delete email mapping")
        }

    } catch (error) {
        if (error instanceof KVError) {
            return GenericResponseServerError("Error communicating with KV")
        }
        return GenericResponseServerError("Failed to delete email mapping")
    }

    return GenericResponseSuccess("Email mapping deleted")
}
