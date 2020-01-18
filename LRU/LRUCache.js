'use strict';

class LRUCache {
    constructor(config) {
        this.maxEntries = parseInt((config || {}).size) || 1; //check if number and parse
        if (this.maxEntries < 1) this.maxEntries = 1; //make sure the max size is more than 0
        this.cache = new Map(); //initialize the storage

        this.invalidateKeys();
    }

    get(key) {
        if (this.validateKey(key)) {
            let item = this.cache.get(key);
            if (item) // refresh key
            {
                this.del(key);
                this.cache.set(key, item);
            }
            return (item || {}).data || null;
        } else
            return null;
    }

    put(key, val, ttlInSeconds = -1) { //add kv to cache. Optional TTL setting
        try {
            let isValidated = this.validateKey(key) && this.validate(val);

            if (this.cache.has(key)) // refresh key
                this.cache.delete(key);
            else if (this.cache.size == this.maxEntries) // evict oldest
                this.cache.delete(this.cache.keys().next().value); //the first item is the one which was inserted first
            this.cache.set(key, {
                data: val, ttl: this._setTTL(ttlInSeconds)
            });

            return [null, key];
        } catch (error) {
            return [error, key];
        }
    }

    _setTTL(ttlInSeconds) { //get TTL in epoch time
        var currentDate = new Date();
        currentDate.setSeconds(currentDate.getSeconds() + (ttlInSeconds > 0 ? ttlInSeconds : 9999999999));
        return currentDate.getTime();
    }

    reset() { //clear the cache
        this.cache.clear();
    }

    del(key) { //delete key
        if (this.cache.has(key)) this.cache.delete(key);
    }

    getLimit() { //get max size count
        return this.maxEntries;
    }

    count() { //get current cache size
        return this.cache.size;
    }

    validateKey(key) { //make sure the key is only either a string or a number
        if (this.validate(key) && (typeof key === "string" || typeof key === "number")) return true;
        else throw "Invalid Key Type. Only String and Integer allowed";
    }

    validate(value) { //Check if value is not null or undefined
        if (value == null) throw "Key/Value Cannot be Null or Undefined";
        else return true;
    }

    invalidateKeys() {
        this.invalidateCacheEvery(1).then(() => {
            let currentTime = new Date().getTime();
            for (const [key, value] of this.cache.entries()) {
                if (value.ttl <= currentTime) {
                    this.del(key);
                }
            }
        }).finally(() => {
            this.invalidateKeys(); //keep checking
        });
    }

    invalidateCacheEvery(inSeconds) {
        return new Promise(resolve => setTimeout(resolve, inSeconds * 1000));
    }
}

if (module) module.exports = LRUCache;