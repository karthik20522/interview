##  About
To implement a "Least-Recently Used" (LRU) Cache. A correctly constructed LRU Cache will have its maximum capacity set at the time of construction and when adding new keys that cause the capacity to be exceeded, the "least recently used" item will be identified and discarded.

### Prerequisites
- git
- NPM

### Installation

```
$ git clone -b chisel https://github.com/karthik20522/interview.git
$ npm install
```

### Commands

-   Run test cases
```
$ npm test
```

## Overview
For this implementation I used (ES6) Map object as Map keeps track of the insertion order and maintains O(1) CRUD operations. When iterating over the Map, the first item is the one which was inserted first. In result, `values.keys().next()` returns the least-recently used key. In order to “remember” the most-recently used items, we re-insert them by first — `.delete()` followed by `.put()` — when reading items from the cache.

### Initialization

```
let lruCache = new LRUCache({ 
	size: 3 
});
```

Where `size` is the maximum number of items before evicting the least recently used items. If *no size* value is provided, the default size of 1 is set.

### Functions
- put (key, value, ttl) 
	- Set the value of the key and mark the key as most recently used. `ttl` is an optional parameter where one can set the expiry time in seconds which then will be automatically evicted. 
	```
		returns [err, key];
	```
- get (key)
	- Query the value of the key and mark the key as most recently used.
- del (key)
	- Remove the value from the cache.
- reset ()
	- Clear the cache.
- count ()
	- Get current cache count


### Usage
```
let lruCache = new LRUCache({ 
	size: 3 
});

//insert
let [err, key] = lruCache.put("key1", 'value1');
if(err) { /*Handle Error*/ }

//read
let value = lruCache.get("key1");

//delete by Key
lruCache.del("key1");

//reset
lruCache.reset();
```
