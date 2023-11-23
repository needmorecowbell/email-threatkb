/**
 * Entry point of the Cloudflare worker.
 * @module index
 */

import { handleRequest } from "./lib/routes";
import { handleEmail } from "./lib/email";

export default {
  /**
   * Fetch event handler for the Cloudflare worker.
   * @function fetch
   * @param {Request} request - The incoming request object.
   * @param {Object} env - The environment variables.
   * @param {Object} ctx - The Cloudflare worker context.
   * @returns {Promise<Response>} The response to be sent back.
   */
  async fetch(request, env, ctx) {
    return handleRequest(request, env)
  },

  /**
   * Email event handler for the Cloudflare worker.
   * @function email
   * @param {Object} message - The email message object.
   * @param {Object} env - The environment variables.
   * @param {Object} ctx - The Cloudflare worker context.
   */
  async email(message, env, ctx) {
    handleEmail(message, env)
  }
};
