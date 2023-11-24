
/**
 * Handles the email message by sending it to the processor and handling the metadata.
 * @param {object} message - The email message object.
 * @param {object} env - The environment object.
 * @returns {Promise<void>} - A promise that resolves when the email handling is complete.
 */
export async function handleEmail(message, env) {
    const init = {
        method: "POST",
        body: message.raw,
        headers: {
            "content-type": "application/json;charset=UTF-8",
        }
    };

    const response = await fetch(buildURL(env.PROCESSOR_SCHEMA, env.PROCESSOR_HOST, ""), init);
    const metadata = await gatherResponse(response);
    handleMetadata(metadata, message);
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
 * @returns {Promise<void>} - A promise that resolves when the metadata handling is complete.
 */
export async function handleMetadata(metadata, message) {
    if (metadata.status === "malicious") {
        await message.forward(VAULT_EMAIL);
    } else {
        await message.forward(PASSTHROUGH_EMAIL);
    }
}

/**
 * buildURL constructs a URL based on the provided schema, host, and path.
 * @param {string} schema - The URL schema (e.g., "http" or "https").
 * @param {string} host - The URL host.
 * @param {string} path - The URL path.
 * @returns {Promise<string>} - A promise that resolves with the constructed URL as a string.
 */
export async function buildURL(schema, host, path) {
    return `${schema}://${host}/${path}`;
}