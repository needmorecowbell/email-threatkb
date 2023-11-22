import Router from "./router"
import endpointIndex from "../endpoints"
import { NotFoundResponse } from "./responses"

export async function handleRequest(request, env) {
    const router = new Router()
  
    router.get('/', () => endpointIndex(request))
    router.all(() => { return NotFoundResponse() })
  
    const response = await router.route(request)
    return response
  }
  
