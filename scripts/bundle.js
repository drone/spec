// This script combines all json schema files into a single
// bundle. The reason we did not use existing libraries is
// because they nest their definitions, which makes code
// generation more complex. This script flattens defitions;
// there is no nesting.

import fs from "fs";
import yaml from "js-yaml";

// parse the root document
let root = yaml.load(fs.readFileSync("schema/config.yaml", "utf8"));

// create the root schema object. we assume the pipeline
// is the base definition.
let schema = {
    definitions: {
        "Config": root,
    },
    oneOf: [{
        $ref: "#/definitions/Config"
    }]
}

// store the pipeline file name
root["x-file"] = "config.yaml"

// walk the object tree.
let walk = (node) => {
    // if the node is null or undefined it can be ignored.
    if (!node) {
        return;
    }

    // if the node is an array, traverse and walk its
    // children.
    if (Array.isArray(node)) {
        return node.forEach(v =>  walk(v));
    }

    // if the node is an object, walk its key value paris.
    if (typeof node === "object") {
       return Object.entries(node).forEach(([k, v]) => {
            // if the key is a reference attempt to resolve
            // the reference.
            if (k === "$ref" && v.startsWith("./") && v.endsWith(".yaml")) {
                // parse the reference file
                let child = yaml.load(fs.readFileSync(`schema/${v}`, "utf8"));

                // store the resource file.
                child["x-file"] = v.slice(2);

                // ensure title is provided.
                if (!child.title) {
                    return console.error("missing title", v);
                }

                // replace external reference with internal
                // reference.
                node[k] = `#/definitions/${child.title}`;

                // avoid traversing a resource that has
                // been traversed.
                if (schema.definitions[child.title]) {
                    return;
                }

                // append resource to definitions.
                schema.definitions[child.title] = child;

                // walk the child object.
                return walk(child);
            }
            
            // else traverse the child node.
            return walk(v);
        });
    }
};

// walk the root node.
walk(root);

// serialize the schema to a file.
fs.writeFileSync('dist/schema.json', JSON.stringify(schema, null, 4)); 
