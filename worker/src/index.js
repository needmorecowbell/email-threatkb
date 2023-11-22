import { handleRequest } from "./lib/routes";
import { handleEmail } from "./lib/email";

export default {
  async fetch(request, env, ctx) {
    return handleRequest(request, env)
  },

  async email(message, env, ctx) {
    handleEmail(message, env)
  }
};




