
import Router from "./router"
import endpointIndex from "../endpoints"
import { endpointMappingAdd } from "../endpoints/endpoint_mapping_add"
import { endpointMappingList } from "../endpoints/endpoint_mapping_list"
import { endpointMappingDelete } from "../endpoints/endpoint_mapping_delete"
import { NotFoundResponse } from "./responses"

/**
 * Handles incoming requests and routes them to the appropriate endpoint.
 * @param {Request} request - The incoming request object.
 * @param {Object} env - The environment object.
 * @returns {Promise<Response>} - The response object.
 */
export async function handleRequest(request, env) {
  const router = new Router()

  router.get('/', () => endpointIndex(request,env));
  router.post('/mapping/add', () => endpointMappingAdd(request, env));
  router.post('/mapping/delete', () => endpointMappingDelete(request, env));
  router.get('/mapping/list', () => endpointMappingList(request, env));

  router.all(() => { return NotFoundResponse() })

  const response = await router.route(request)
  return response
}

