
import { GenericResponseBadRequest, GenericResponseSuccess } from "../lib/responses";
import { cacheEmailMappingList } from "../lib/kv";

/**
 * Endpoint function for listing the email mappings from KV.
 * 
 * @param {Request} request - The request object.
 * @param {Object} env - The environment object.
 * @returns {Response} - The response object indicating the success or failure of the operation.
 */
export const endpointMappingList = async (request, env) => {

    let mappings = await cacheEmailMappingList( env)
    if (mappings === false) {
        return GenericResponseBadRequest("Failed to retrieve email mappings")
    }

    return GenericResponseSuccess("Email mappings retrieved", mappings)
}
