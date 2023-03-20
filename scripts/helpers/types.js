import deepequal from "deep-equal";

/**
 * Returns the Go representation of a json schema property.
 * @param {object} prop json schema property 
 */
export default function getType(prop, defs) {
    if (!!prop["x-go-type"]) {
        return prop["x-go-type"]; // should this prepend pointer "*"?

    } else  if (isStringArray(prop)) {
        return "[]string";

    } else if (isStringOrSlice(prop)) {
        return "Stringorslice";

    } else if (isBytes(prop)) {
        return "MemStringorInt";

    } else if (isStringOrNumber(prop)) {
        return "StringorInt";

    } else if (isDict(prop)) {
        return "map[string]string";

    } else if (isEnum(prop)) {
        return "string";

    } else if (isMapStruct(prop)) {
        return "map[string]*" + prop.additionalProperties.$ref.slice(14);

    } else if (isMapArray(prop)) {
        return "map[string][]string";

    } else if (isArrayMap(prop)) {
        return "[]map[string]string";

    } else if (isObject(prop)) {
        return "map[string]interface{}";

    } else if (isDuration(prop)) {
        return "time.Duration";

    } else if (isDurationOrArray(prop)) {
        return "Durationorslice";

    } else if (isString(prop)) {
        return "string";

    } else if (isNumber(prop)) {
        return "int64";

    } else if (isBool(prop)) {
        return "bool";

    } else if (isRefSlice(prop)) {
        return "[]*"+prop.items.$ref.slice(14);

    } else if (isPrimative(prop)) {
        return "string";

    } else if (isPrimativeSlice(prop)) {
        return "[]string";

    } else if (isRef(prop)) {
        const name = prop.$ref.slice(14);
        if (defs[name] && defs[name].enum) {
            return name;
        }
        return "*"+name;

    } else {
        console.error("unknown type", prop.name, prop.type);
        return "interface{}";
    }   
}

//
// helper functions below
//

/**
 * Returns true if the json schema property is an object.
 */
function isObject(node) {
    return node.type === "object";
}

/**
 * Returns true if the json schema property is an array.
 */
function isArray(node) {
    return node.type === "array";
}

/**
 * Returns true if the json schema property is an enum.
 */
function isEnum(node) {
    return node.enum &&
        node.enum.length &&
        node.enum.length > 0;
}

/**
 * Returns true if the json schema property is a
 * primative type (string, number, boolean, etc).
 */
function isPrimative (node) {
    return node.type &&
        node.type.includes("string") &&
        node.type.includes("number") &&
        node.type.includes("boolean")
}

/**
 * Returns true if the json schema property is a number.
 */
function isNumber(node) {
    return node.type === "number";
}

/**
 * Returns true if the json schema property is a boolean.
 */
function isBool(node) {
    return node.type === "boolean";
}

/**
 * Returns true if the json schema property is a string.
 */
function isString(node) {
    return node.type === "string";
}

/**
 * Returns true if the json schema property is a
 * string array.
 */
function isStringArray(node) {
    return node &&
        node.type === "array" &&
        node.items &&
        node.items.type &&
        node.items.type === "string";
}

/**
 * Returns true if the json schema property is a
 * string or number.
 */
function isStringOrNumber(node) {
    return node.type &&
        node.type.length &&
        node.type.length == 2 &&
        node.type.includes("string") &&
        node.type.includes("number");
}

/**
 * Returns true if the json schema property is a
 * dictionary of string values.
 */
function isDict(node) {
    return node.type === "object" &&
        node.additionalProperties &&
        node.additionalProperties.type === "string";  
}

/**
 * Returns true if the json schema property is a
 * a string formatted as bytes.
 */
function isBytes (node) {
    return node.format === "bytes";
}

/**
 * Returns true if the json schema property is a
 * duration value.
 */
function isDuration(node) {
    return node.format === "duration";
}

/**
 * Returns true if the json schema property is a
 * map of objects / structs.
 */
function isMapStruct(node) {
    return node.type == "object" &&
        node.additionalProperties &&
        node.additionalProperties.hasOwnProperty("$ref")
}

/**
 * Returns true if the json schema property is a
 * duration value, or an array of duration values.
 * 
 * TODO simplify this code.
 */
function isDurationOrArray(node) {
    return node.anyOf && node.anyOf[0] && node.anyOf[0].items && node.anyOf[0].items.format === "duration" &&
        node.anyOf[1] && node.anyOf[1].format === "duration";
}

/**
 * Returns true if the json schema property is a
 * reference to another type.
 */
function isRef(node) {
    return node.$ref;
}

/**
 * Returns true if the json schema property is an
 * array of references to other types.
 */
function isRefSlice(node) {
    return isArray(node) && node.items && node.items.$ref;
}

/**
 * Returns true if the json schema property is an array
 * of map structures, where the value is type string.
 */
function isArrayMap(node) {
    return node.type === "array" &&
        node.items &&
        node.items.type === "object" &&
        node.items.additionalProperties &&
        node.items.additionalProperties.type === "string"
}

/**
 * Returns true if the json schema property is an map
 * type where, the the value is type []string.
 */
function isMapArray(node) {
    return node.type === "object" &&
        node.additionalProperties &&
        node.additionalProperties.type === "array" &&
        node.additionalProperties.items &&
        node.additionalProperties.items.type === "string"
}

//
// below functions need improvement
//

// HACK
const isStringOrSlice = (node) => { // SIMPLIFY
    return deepequal(node.anyOf, [
        {
            "items": {
                "type": "string"
            },
            "type": "array"
        },
        {
            "type": "string"
        }
    ]);
}

// HACK
const isPrimativeSlice = (node) => { // SIMPLIFY
    return deepequal(node, {
        "type": "array",
        "items": {
            "type": [
                "string",
                "number",
                "boolean"
            ]
        }
    });
}
