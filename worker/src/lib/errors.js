/**
 * Represents an error specific to the KV (Key-Value) store.
 * @class
 * @extends Error
 * @name KVError
 * @param {string} message - The error message.
 */
export const KVError = class extends Error {
    constructor(message) {
        super(message);
        this.name = "KVError";
    }
}