const LRUCache = require('./LRUCache')
jest.setTimeout(30000);


/*********************************************/
/*  CONFIG TESTS  */
/*********************************************/
test('Set and Get Cache Max Entries Size', () => {
    let lruCache = new LRUCache({ size: 99 });
    expect(lruCache.getLimit()).toBe(99);
})

test('Get default Cache Max Entries Size of 1 when no Size is provided', () => {
    let lruCache = new LRUCache();
    expect(lruCache.getLimit()).toBe(1);
})

test('Validate Config Size Value - Is Number and convert to Int', () => {
    let lruCache = new LRUCache({ size: 1 });
    expect(lruCache.getLimit()).toBe(1);

    lruCache = new LRUCache({ size: "2" });
    expect(lruCache.getLimit()).toBe(2);
})

test('Validate Config Size Value - Should be Greater than Zero else set Default Value of 1', () => {
    let lruCache = new LRUCache({ size: -1 });
    expect(lruCache.getLimit()).toBe(1);
})

/*********************************************/
/*  GET AND SET VALUE TESTS */
/*********************************************/
test('Set and Get Cache', () => {
    let lruCache = new LRUCache({ size: 3 });
    lruCache.put("key1", "value1");
    lruCache.put("key2", "value2");
    lruCache.put("key3", "value3");

    expect(lruCache.get("key1")).toBe("value1");
    expect(lruCache.get("key2")).toBe("value2");
    expect(lruCache.get("key3")).toBe("value3");
})

test('Validate Values - Ignore Null and Undefined', () => {
    let lruCache = new LRUCache({ size: 3 });
    lruCache.put("key0", "hello world");

    let [nullErr, key1] = lruCache.put("key1", null);
    let [undefinedErr, key2]= lruCache.put("key2", undefined);    
    let [nullKeyErr, key3] = lruCache.put(null, "Bad Key");
    let [badKeyErr, key4] = lruCache.put({"bad": "key"}, "Bad Key");

    expect(nullErr).toBe("Key/Value Cannot be Null or Undefined");
    expect(undefinedErr).toBe("Key/Value Cannot be Null or Undefined");
    expect(nullKeyErr).toBe("Key/Value Cannot be Null or Undefined");
    expect(badKeyErr).toBe("Invalid Key Type. Only String and Integer allowed");

    expect(lruCache.get("key0")).toBe("hello world");
    expect(lruCache.get("key1")).toBe(null);
    expect(lruCache.get("key2")).toBe(null);
    expect(lruCache.get("key3")).toBe(null);
    expect(lruCache.get("key4")).toBe(null);
});

test('Validate Key as String and Int Only', () => {
    let lruCache = new LRUCache({ size: 4 });
    lruCache.put("key1", "hello string");
    lruCache.put(1, "hello int");
    lruCache.put(true, "bad boolean key");
    lruCache.put({ "bad": "object" }, "bad object key");

    expect(lruCache.count()).toBe(2);

    expect(lruCache.get("key1")).toBe("hello string");
    expect(lruCache.get(1)).toBe("hello int");
});

test('Get NULL when invalid key is used', () => {
    let lruCache = new LRUCache();
    expect(lruCache.get("invalidKey")).toBe(null);
});

/*********************************************/
/*  EVICTION TESTS  */
/*********************************************/
test('Evict recently used key', () => {
    let lruCache = new LRUCache({ size: 3 });
    lruCache.put("key1", "value1");
    lruCache.put("key2", "value2");
    lruCache.put("key3", "value3");

    expect(lruCache.count()).toBe(3);
    lruCache.put("key4", "value4"); //add a new kv

    expect(lruCache.get("key1")).toBe(null);
    expect(lruCache.get("key4")).toBe("value4");
});

/*********************************************/
/*  DELETE TESTS  */
/*********************************************/
test('Delete by Valid Key', () => {
    let lruCache = new LRUCache({ size: 3 });
    lruCache.put("key1", "value1");
    lruCache.put("key2", "value2");
    lruCache.put("key3", "value3");

    expect(lruCache.count()).toBe(3);
    lruCache.del("key2");
    expect(lruCache.count()).toBe(2);
    expect(lruCache.get("key2")).toBe(null);
})

/*********************************************/
/*  RESET TEST  */
/*********************************************/
test('Reset Cache', () => {
    let lruCache = new LRUCache({ size: 3 });
    lruCache.put("key1", "value1");
    lruCache.put("key2", "value2");
    lruCache.put("key3", "value3");

    expect(lruCache.count()).toBe(3);
    lruCache.reset();
    expect(lruCache.count()).toBe(0);
})

/*********************************************/
/*  TTL TEST  */
/*********************************************/
test("TTL Tests - Delete Keys if ttl expired", async () => {
    let lruCache = new LRUCache({ size: 3 });
    lruCache.put("key1", "value1");
    lruCache.put("key2", "value2");
    lruCache.put("key3", "value3", 1);

    await new Promise((r) => setTimeout(r, 3000));
    expect(lruCache.count()).toBe(2);
});