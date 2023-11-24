
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
    console.debug(`Received request to /mapping/add to add ${body.gateway_address} -> ${body.forward_to}`)
    if (body.forward_to === undefined) {
        console.debug("bad request: forward_to is required")
        return GenericResponseBadRequest("forward_to is required")
    }
    if (body.gateway_address === undefined) {
        console.debug("No gateway address provided, generating new one...")
        body.gateway_address = await generateNewGatewayAddress(env)
    }

    var result
    try {
        result = await mappingExists(body.gateway_address, body.forward_to, env);
    } catch (error) {
        if (error instanceof KVError) {
            console.debug("Error communicating with KV")
            return GenericResponseServerError("Error communicating with KV")
        }
        console.debug("Failed to determine if email mapping exists")
        return GenericResponseServerError("Failed to determine if email mapping exists")
    }

    if (result === undefined) {
        console.debug("Couldn't determine if mapping exists")
        return GenericResponseBadRequest("Couldn't determine if mapping exists")
    } else if (result === true) {
        console.debug(`Mapping already exists for ${body.gateway_address} -> ${body.forward_to}`)
        return GenericResponseBadRequest("Email mapping already exists")
    }

    try {
        await cacheEmailMappingAdd(body.forward_to, body.gateway_address, env)

    } catch (error) {
        console.debug(`Mapping already exists for ${body.gateway_address} -> ${body.forward_to}`)
        return GenericResponseServerError("Failed to add email mapping")

    }
    console.debug(`Added mapping for ${body.gateway_address} -> ${body.forward_to}`)
    return GenericResponseSuccess("Email mapping added")
}
