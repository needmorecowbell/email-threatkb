/**
 * @module kv
 * @description This module provides functions for interacting with the KV store.
 */


/**
 * @constant {string} PrefixMapping- The prefix for mapping-related keys in the KV store.
 */
export const PrefixMapping= "mapping"

/**
 * Creates a new email mapping object for caching.
 * If gateway_address is not provided, a new one will be generated.
 * @param {string} forward_to - The email address to forward to.
 * @param {string} [gateway_address=""] - The gateway address to use.
 * @returns {Object} The email mapping object.
 */
export const CacheEmailMapping = (forward_to, gateway_address = "") => {
    let address = gateway_address !== "" ? gateway_address : generateNewGatewayAddress(env);
    return {
        forward_to: forward_to,
        date_created: new Date().toISOString(),
        gateway_address: address,
    }
}

/**
 * Checks if a mapping exists in the KV store.
 * @param {string} gateway_address - The gateway address.
 * @param {string} forward_to - The forward address.
 * @param {Object} env - The environment object.
 * @returns {Promise<boolean>} A promise that resolves to true if the mapping exists, false otherwise.
 */
export async function mappingExists(gateway_address, forward_to, env) {
    console.log("Checking if mapping exists")
    console.log(env.KV)
    let result = await env.KV.get(buildKey(gateway_address, forward_to))
    console.log("Result: ",result)
    return result !== null
}

/**
 * Generates a new gateway address.
 * @returns {Promise<string>} The new gateway address.
 */
export async function generateNewGatewayAddress(env) {
    // generate a string shortname for the email address that doesn't already exist in the cache, example could look like "selfish-zebra-1234"
    let new_name = `${generateRandomWord()}-${generateRandomWord()}-${generateRandomWord()}`
    let name_exists = true
    let mappings = await cacheEmailMappingList(env)

   // instead of a while loop, check using the mappings variable and a for loop
   for (let i = 0; i < mappings.length; i++) {
       if (mappings[i].gateway_address === new_name) {
           name_exists = true
           break
       }
    }

    return new_name
};

/**
 * Generates a random word using an API.
 * @returns {Promise<string>} A promise that resolves to a random word.
 */
export async function generateRandomWord() {
    let response = await fetch("https://random-word-api.herokuapp.com/word?number=1")
    let json = await response.json()
    return json[0]
}
/**
 * Retrieves a list of email mappings from the cache.
 * @param {Object} env - The environment object.
 * @returns {Promise<Array>} A promise that resolves to an array of email mappings.
 */
export async function cacheEmailMappingList(env) {
    let results = await env.KV.list({ prefix: PrefixMapping });
    return results
}

/**
 * Retrieves an email mapping from the cache.
 * @param {string} gateway_address - The gateway address.
 * @param {string} forward_to - The forward address.
 * @param {Object} env - The environment object.
 * @returns {Promise<Object>} A promise that resolves to an email mapping, or null.
 */
export async function cacheEmailMappingGet(gateway_address, forward_to, env) {
    let result = await env.KV.get(buildKey(gateway_address, forward_to))
    return result
}

/**
 * Deletes the email mapping from the cache when passed a CacheEmailMapping object.
 * @param {Object} cache_email_mapping - The email mapping to delete from the cache.
 * @param {Object} env - The environment object.
 * @returns {Promise<boolean>} A promise that resolves to true if the email mapping is successfully deleted, otherwise false.
 */
export async function cacheEmailMappingDeleteByMapping(cache_email_mapping, env) {
    let result = await env.KV.delete(buildKey(cache_email_mapping))
    return result
}

/**
 * Deletes the email mapping from the cache when passed a forward-to and gateway address.
 * @param {Object} cache_email_mapping - The email mapping to delete from the cache.
 * @param {Object} env - The environment object.
 * @returns {Promise<boolean>} A promise that resolves to true if the email mapping is successfully deleted, otherwise false.
 */
export async function cacheEmailMappingDelete(forward_to, gateway_address, env) {
    let result = await env.KV.delete(buildKey(gateway_address, forward_to))
    return result
}


/**
 * Adds a new email mapping to the cache.
 * If gateway_address is not provided, a new one will be generated.
 * @param {string} forward_to - The email address to forward to.
 * @param {string} [gateway_address=""] - The gateway address to use.
 * @param {Object} env - The environment object.
 * @returns {Promise<boolean>} A promise that resolves to true if the mapping is successfully added to the cache, otherwise false.
 */
export async function cacheEmailMappingAdd(forward_to, gateway_address, env) {
    if (gateway_address === "") {
        gateway_address = generateNewGatewayAddress()
    }

    let mapping = CacheEmailMapping(forward_to, gateway_address, env)
    let result = env.KV.put(buildKeyByMapping(mapping), JSON.stringify(mapping))
    return result
}

/**
 * Builds a key for the KV store based on a cache_email_mapping object.
 * @param {Object} cache_email_mapping - The cache_email_mapping object.
 * @returns {string} The generated key.
 */
async function buildKeyByMapping(cache_email_mapping) {
    return `${PrefixMapping}_${cache_email_mapping.gateway_address}_${cache_email_mapping.forward_to}`
}

/**
 * Builds a key for the KV store based on the provided gateway address and forward address.
 * @param {string} gateway_address - The gateway address.
 * @param {string} forward_to - The forward address.
 * @returns {string} The built key.
 */
async function buildKey(gateway_address, forward_to) {
    return `${PrefixMapping}_${gateway_address}_${forward_to}`
}
