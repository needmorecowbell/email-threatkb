/**
 * @module kv
 * @description This module provides functions for interacting with the KV store.
 */
import { KVError } from "./errors"

/**
 * @constant {string} PrefixMapping- The prefix for mapping-related keys in the KV store.
 */
export const PrefixMapping = "mapping"

/**
 * Creates a new email mapping object for caching.
 * If gateway_address is not provided, a new one will be generated.
 * @param {string} forward_to - The email address to forward to.
 * @param {string} [gateway_address=""] - The gateway address to use.
 * @returns {Object} The email mapping object.
 */
export const CacheEmailMapping = (forward_to, gateway_address) => {
    return {
        forward_to: forward_to,
        date_created: new Date().toISOString(),
        gateway_address: gateway_address,
    }
}

/**
 * Checks if a mapping exists in the KV store.
 * @param {string} gateway_address - The gateway address.
 * @param {string} forward_to - The forward address.
 * @param {Object} env - The environment object.
 * @returns {Promise<boolean>} A promise that resolves to true if the mapping exists, otherwise false.
 */
export async function mappingExists(gateway_address, forward_to, env) {
    console.debug(`Checking if mapping exists for ${gateway_address} -> ${forward_to}`)
    try {
        let result = await cacheEmailMappingGet(gateway_address, forward_to, env)
        if (result === undefined) {
            throw KVError("Failed to retrieve mapping")
        }
        return result !== null
    } catch (error) {
        return error
    }
}

/**
 * Generates a new gateway address.
 * @returns {Promise<string>} The new gateway address.
 */
export async function generateNewGatewayAddress(env) {
    console.debug("Generating new gateway address")
    let gen_new_name = true
    let new_name = ""
    let new_address = ""
    let mappings = await cacheEmailMappingList(env)

    while(gen_new_name){
        new_name = `${await generateRandomWord()}-${await generateRandomWord()}-${await generateRandomWord()}`
        gen_new_name = false
        new_address = `${new_name}@${env.GATEWAY_DOMAIN}`

        for (let i = 0; i < mappings.length; i++) {
            if (mappings[i].gateway_address === new_address) {
                gen_new_name = true
                console.debug(`Generated name ${new_name} already exists, generating new name`)
                break
            }
        }
    }
    return new_address
};

/**
 * Generates a random word using an API.
 * @returns {Promise<string>} A promise that resolves to a random word.
 */
export async function generateRandomWord() {
    console.debug("Retrieving random word from api")
    let response = await fetch("https://random-word-api.herokuapp.com/word?number=1")
    let json = await response.json()
    return json[0]
}

/**
 * Retrieves and caches the email mapping list from the environment's KV store.
 * @param {Object} env - The environment object.
 * @returns {Promise<Array>} - A promise that resolves to an array of email mappings.
 */
export async function cacheEmailMappingList(env) {
    console.debug("Retrieving all email mappings")
    let results = await env.KV.list({ prefix: PrefixMapping });
    if (results === undefined) {
        console.error("Failed to retrieve email mappings")
        throw KVError("Failed to retrieve email mappings")
    }
    console.debug(`Retrieved ${results.keys.length} email mapping keys`)
    let mappings = []
    for (let i = 0; i < results.keys.length; i++) {
        console.debug(`Retrieving mapping for ${results.keys[i].name}`)
        let mapping = await env.KV.get(results.keys[i].name)
        if (mapping === undefined) {
            console.error("Failed to retrieve email mapping from key: " + results.keys[i].name)
            throw KVError("Failed to retrieve email mapping from key: " + results.keys[i].name)
        }
        mappings.push(mapping)
    }
    return mappings
}

/**
 * Retrieves an email mapping from the cache.
 * @param {string} gateway_address - The gateway address.
 * @param {string} forward_to - The forward address.
 * @param {Object} env - The environment object.
 * @returns {Promise<Object>} A promise that resolves to an email mapping, or null.
 */
export async function cacheEmailMappingGet(gateway_address, forward_to, env) {
    console.debug(`Retrieving email mapping for ${gateway_address} -> ${forward_to}`)
    try {
        let result = await env.KV.get(buildKey(gateway_address, forward_to))
        return result
    } catch (e) {
        console.error(e)
        throw KVError("Failed to retrieve email mapping:" + e)
    }
}

/**
 * Deletes the email mapping from the cache when passed a CacheEmailMapping object.
 * @param {Object} cache_email_mapping - The email mapping to delete from the cache.
 * @param {Object} env - The environment object.
 * @returns {Promise<boolean>} A promise that resolves to true if the email mapping is successfully deleted, otherwise false.
 */
export async function cacheEmailMappingDeleteByMapping(cache_email_mapping, env) {
    console.debug(`Deleting email mapping for ${cache_email_mapping.gateway_address} -> ${cache_email_mapping.forward_to}`)
    let result = await env.KV.delete(buildKey(cache_email_mapping))
    return result
}

/**
 * Deletes the email mapping from the cache when passed a forward-to and gateway address.
 * @param {string} forward_to - The email address to forward to.
 * @param {string} gateway_address - The gateway address.
 * @param {Object} env - The environment object.
 * @returns {Promise<boolean>} A promise that resolves to true if the email mapping is successfully deleted, otherwise false.
 */
export async function cacheEmailMappingDelete(forward_to, gateway_address, env) {
    console.debug(`Deleting email mapping for ${gateway_address} -> ${forward_to}`)
    try {
        let result = await env.KV.delete(buildKey(gateway_address, forward_to))
        return result
    } catch (error) {
        console.error(error)
        throw KVError("Failed to delete email mapping: " + error)
    }
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
    console.debug(`Adding email mapping for ${gateway_address} -> ${forward_to}`)
    if (gateway_address === "") {
        console.debug("No gateway address provided, generating new one...")
        gateway_address = generateNewGatewayAddress()
        console.debug(`Generated gateway address: ${gateway_address}`)
    }

    let mapping = CacheEmailMapping(forward_to, gateway_address, env)
    console.debug(`Adding mapping to KV as: KEY:${buildKey(mapping)} VALUE:${JSON.stringify(mapping)}`)
    try {
        let result = env.KV.put(buildKeyByMapping(mapping), JSON.stringify(mapping))
        return result
    } catch (error) {
        console.error(error)
        throw KVError("Failed to add email mapping to KV")
    }

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
