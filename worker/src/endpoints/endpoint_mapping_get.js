
import { GenericResponseBadRequest, GenericResponseSuccess, GenericResponseServerError } from "../lib/responses";
import { mappingExists, CacheEmailMapping, cacheEmailMappingGet, generateNewGatewayAddress } from "../lib/kv";

/**
 * Endpoint function for adding a new email mapping to the cache.
 * 
 * @param {Request} request - The request object.
 * @param {Object} env - The environment object.
 * @returns {Response} - The response object indicating the success or failure of the operation.
 */
export const endpointMappingGet = async (request, env) => {
    let body = await request.json();
    console.debug(`Received request to /mapping/add to add ${body.gateway_address} -> ${body.forward_to}`)
    if (body.forward_to === undefined || body.gateway_address === undefined) {
        console.debug("bad request: forward_to  and gateway_address are required")
        return GenericResponseBadRequest("forward_to and gateway_address are required")
    }
    
    var mapping
   
    try {
        mapping = await cacheEmailMappingGet(body.gateway_address, body.forward_to, env)
        if (mapping === null) {
            return GenericResponseBadRequest("Mapping does not exist")
        }
    } catch (error) {
        return GenericResponseServerError("Failed to retrieve email mapping: "+error)

    }
    console.debug(`Retrieved mapping for ${body.gateway_address} -> ${body.forward_to}`)
    return GenericResponseSuccess("Mapping retrieved", mapping)
}
