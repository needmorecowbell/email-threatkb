
import { GenericResponseServerError, GenericResponseSuccess } from "../lib/responses";
import { cacheEmailMappingList } from "../lib/kv";
import { KVError } from "../lib/errors";

/**
 * Endpoint function for listing the email mappings from KV.
 * 
 * @param {Request} request - The request object.
 * @param {Object} env - The environment object.
 * @returns {Response} - The response object indicating the success or failure of the operation.
 */
export const endpointMappingList = async (request, env) => {
    let mappings = []
    try {
        mappings = await cacheEmailMappingList(env)
    } catch (err) {
        if (err instanceof KVError) {
            return GenericResponseServerError("Error communicating with KV")
        }
        return GenericResponseServerError("Failed to retrieve email mappings")
    }


    return GenericResponseSuccess("Email mappings retrieved", mappings)
}
