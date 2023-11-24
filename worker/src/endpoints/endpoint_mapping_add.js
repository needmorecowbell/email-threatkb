
import { GenericResponseBadRequest, GenericResponseSuccess, GenericResponseServerError } from "../lib/responses";
import { mappingExists, CacheEmailMapping, cacheEmailMappingAdd, generateNewGatewayAddress } from "../lib/kv";

/**
 * Endpoint function for adding a new email mapping to the cache.
 * 
 * @param {Request} request - The request object.
 * @param {Object} env - The environment object.
 * @returns {Response} - The response object indicating the success or failure of the operation.
 */
export const endpointMappingAdd = async (request, env) => {
    let body = await request.json();

    if(body.forward_to === undefined){
        return GenericResponseBadRequest("forward_to is required")
    }
    if(body.gateway_address === undefined){
        body.gateway_address = await generateNewGatewayAddress(env)
    }

    var result
    try {
        result = await mappingExists(body.gateway_address, body.forward_to, env);

    } catch (error) {
        if (error instanceof KVError) {
            return GenericResponseServerError("Error communicating with KV")
        }
        return GenericResponseServerError("Failed to determine if email mapping exists")
    }
    
    if(result === undefined) {
        return GenericResponseBadRequest("Couldn't determine if mapping exists")
    }else if(result === true){
        return GenericResponseBadRequest("Email mapping already exists")
    }
    let new_mapping = CacheEmailMapping(body.forward_to, body.gateway_address)
    result = await cacheEmailMappingAdd(new_mapping, "", env)
    if(result === undefined){
        return GenericResponseServerError("Failed to add email mapping to KV")
    }

    return GenericResponseSuccess("Email mapping added")
}
