/**
 * handleEmail is the top-level function for any emails that are forwarded to the function. It will ultimately choose to drop or forward based off of what it thinks of the metadata.
 * @param {*} message 
 * @param {*} env 
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
 * @param {Response} response
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
 * handleMetadata takes the metadata response and analyzes it to determine if the message will ultimately be forwarded
 * @param {object} metadata
 * @param message
 */
export async function handleMetadata(metadata, message) {
    if (metadata.status === "malicious") {
        await message.forward(VAULT_EMAIL)
    } else {
        await message.forward(PASSTHROUGH_EMAIL)
    }
}

export async function buildURL(schema,host,path){
    return `${schema}://${host}/${path}`
}