export const PrefixMapping = 'mapping';

const WORKER_URL = import.meta.env.WORKER_URL;
/**
 * Retrieves a mapping from the server.
 * @param {string} gateway_address - The gateway address.
 * @param {string} forward_to - The forward address.
 * @returns {any} - The mapping data or null if an error occurred.
 */
export async function mappingGet(gateway_address, forward_to){
    const init = {method:"POST", body:JSON.stringify({gateway_address:gateway_address,forward_to:forward_to})};
    const resp = await fetch(`${WORKER_URL}/mapping/get`, init);
    var mapping 
    try {
        mapping = await resp.json();
    } catch (error) {
        console.error(error);
        return null
    }

    if(mapping.success === false){
        console.debug(mapping.msg)
        return null;
    }
    return mapping.data;
}


/**
 * Fetches the list of mappings from the server.
 * @returns {Promise<Array>} A promise that resolves to an array of mappings.
 */
export async function mappingList(){
    const resp = await fetch(`${WORKER_URL}/mapping/list`);
    let mappings = await resp.json();
    return mappings.data;
}

/**
 * Builds a key for the KV store based on the provided gateway address and forward address.
 * @param {string} gateway_address - The gateway address.
 * @param {string} forward_to - The forward address.
 * @returns {string} The built key.
 */
export function buildKey(gateway_address, forward_to) {
    return `${PrefixMapping}_${gateway_address}_${forward_to}`
}
