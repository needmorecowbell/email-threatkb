
import { GenericResponseBadRequest, GenericResponseSuccess } from "../lib/responses";
import { cacheEmailMappingDelete } from "../lib/kv";
import { CacheEmailMapping } from "../lib/kv";

/**
 * Endpoint function for deleting an email mapping from the cache.
 * 
 * @param {Request} request - The request object.
 * @param {Object} env - The environment object.
 * @returns {Response} - The response object indicating the success or failure of the operation.
 */
export const endpointMappingDelete = async (request, env) => {
    let body = await request.json()
    let result = await cacheEmailMappingDelete(body.forward_to, body.gateway_address, env)
    if (result === undefined) {
        return GenericResponseBadRequest("Failed to delete email mapping")
    }

    return GenericResponseSuccess("Email mapping deleted")
}
