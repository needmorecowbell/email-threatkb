
import { GenericResponseBadRequest, GenericResponseSuccess } from "../lib/responses";
import { mappingExists, CacheEmailMapping, cacheEmailMappingAdd } from "../lib/kv";

/**
 * Endpoint function for adding a new email mapping to the cache.
 * 
 * @param {Request} request - The request object.
 * @param {Object} env - The environment object.
 * @returns {Response} - The response object indicating the success or failure of the operation.
 */
export const endpointMappingAdd = async (request, env) => {
    let body = await request.json()
    if(await mappingExists(body.gateway_address, body.forward_to, env)) {
        console.log("Mapping already exists")
        return GenericResponseBadRequest("Email mapping already exists")
    }

    let new_mapping = CacheEmailMapping(body.forward_to, body.gateway_address)
    let result = await cacheEmailMappingAdd(new_mapping,"", env)
    console.log(result)
    if (result === false) {
        return GenericResponseBadRequest("Failed to add email mapping")
    }

    return GenericResponseSuccess("Email mapping added")
}
