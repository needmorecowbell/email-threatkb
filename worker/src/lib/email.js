import {cacheEmailMappingList} from "./kv"
/**
 * Handles the email message by sending it to the processor and handling the metadata.
 * @param {object} message - The email message object.
 * @param {object} env - The environment object.
 * @returns {Promise<void>} - A promise that resolves when the email handling is complete.
 */
export async function handleEmail(message, env) {
    console.debug(`Entered handle email for message: ${message.from}->${message.to}`)
    const init = {
        method: "POST",
        body: message.raw,
        headers: {
            "content-type": "application/json;charset=UTF-8",
        }
    };
    console.debug("Requesting metadata from processor")
    const response = await fetch(buildURL(env.PROCESSOR_SCHEMA, env.PROCESSOR_HOST, ""), init);

    console.debug("Gathering Response from processor")
    const metadata = await gatherResponse(response);

    await handleMetadata(metadata, message,env);
}

/**
 * gatherResponse awaits and returns a response body as a string.
 * Use await gatherResponse(..) in an async function to get the response body
 * @param {Response} response - The response object.
 * @returns {Promise<string>} - A promise that resolves with the response body as a string.
 */
export async function gatherResponse(response) {
    const { headers } = response;
    const contentType = headers.get("content-type") || "";

    if (contentType.includes("application/json")) {
        return JSON.stringify(await response.json());
    }
    return response.text();
}

/**
 * handleMetadata takes the metadata response and analyzes it to determine if the message will ultimately be forwarded.
 * @param {object} metadata - The metadata object.
 * @param {object} message - The email message object.
 * @param {object} env - The environment object.
 * @returns {Promise<void>} - A promise that resolves when the metadata handling is complete.
 */
export async function handleMetadata(metadata, message,env) {
    console.debug(`Handling metadata: ${metadata}`)

    if (metadata.status === "malicious") {
        console.debug("Message is malicious, forwarding to vault: ", env.VAULT_EMAIL)
        message.forward(env.VAULT_EMAIL);
    } else {
        console.debug("Message is not malicious, forwarding to gateway: ", env.GATEWAY_EMAIL)
        await handleEmailForwarding(message,env)
    }
}

/**
 * Handles email forwarding based on the provided message and environment.
 * 
 * @param {object} message - The email message object.
 * @param {object} env - The environment object.
 * @returns {void}
 */
export  async function handleEmailForwarding(message, env){
    let mappings = await cacheEmailMappingList(env)
    mappings.forEach(mapping => {
        if (mapping.gateway_email === message.to){
            console.log(`Forwarding message from ${message.from} to ${mapping.forward_email}`)
            message.forward(mapping.forward_email)
            return
        }
    });

    console.debug(`No mapping found for ${message.to}, dropping message...`)
}

/**
 * buildURL constructs a URL based on the provided schema, host, and path.
 * @param {string} schema - The URL schema (e.g., "http" or "https").
 * @param {string} host - The URL host.
 * @param {string} path - The URL path.
 * @returns {Promise<string>} - A promise that resolves with the constructed URL as a string.
 */
export function buildURL(schema, host, path) {
    return `${schema}://${host}/${path}`;
}