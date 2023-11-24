
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
    console.debug(`Received request to /mapping/delete to delete ${body.gateway_address} -> ${body.forward_to}`)
    if (body.forward_to === undefined || body.gateway_address === undefined) {
        console.debug("bad request: forward_to and gateway_address are required")
        return GenericResponseBadRequest("forward_to and gateway_address are required")
    }

    try {
        await cacheEmailMappingDelete(body.forward_to, body.gateway_address, env)
    } catch (error) {
        if (error instanceof KVError) {
            console.debug("Error communicating with KV")
            return GenericResponseServerError("Error communicating with KV")
        }
        console.debug("Failed to delete email mapping"+error)
        return GenericResponseServerError("Failed to delete email mapping:"+error)
    }
    console.debug(`Deleted mapping for ${body.gateway_address} -> ${body.forward_to}`)
    return GenericResponseSuccess("Email mapping deleted")
}
