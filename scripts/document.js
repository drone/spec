// This script generates Go code from the jsonschema
// bundle using handlebars templates. The templates are
// located in the ./templates subdirectory.

// This code is a work-in-progress.
// Please excuse the mess.

import fs from "fs";
import Case from "case";
import Handlebars from "handlebars";
import getType from "./helpers/types.js";
import yaml from "js-yaml";

// parse the schema
let schema = JSON.parse(fs.readFileSync("dist/schema.json"));

// for each definition
Object.entries(schema.definitions).forEach(([k, v]) => {

    if (v["x-docs-skip"] == true) {
        return
    } else {
        console.log(k, v["x-docs-skip"])
    }

    // store the go struct details.
    let struct = {
        name: v["x-docs-title"] || k,
        desc: v.description,
        enum:  [],
        props: [],
        types: [],
        examples: [],
    }

    // for each property
    Object.entries(v.properties || {}).forEach(([propkey, propval]) => {
        let type = getType(propval, schema.definitions);
        type = type.replace("[]*", "")
        type = type.replace("*", "")

        switch (type) {
        case "map[string]interface{}":
            type = "object"
            break;
        case "interface{}":
            type = "undefined";
            break;
        case "int64":
            type = "number"
            break;
        case "[]string":
            type = "string"
            break;
        }

        let enums = [];

        // for each enum value
        propval.enum && propval.enum.forEach(text => {
            enums.push(text);
        });

        // append the field to the go struct
        struct.props.push({
            name: propkey,
            desc: propval && propval.description,
            type: type,
            enum: enums,
            array: propval.type && propval.type === "array",
        });
    });

    // for each enum value
    v.enum && v.enum.forEach(text => {
        struct.enum.push(text);
    });

    // for each example value
    v.examples && v.examples.forEach(object => {
        struct.examples.push(yaml.dump(object));
    });

    // this block of code detects if we are using the type /
    // spec pattern. If yes, we store the type enum values
    // and their associated struct types.
    if (v.properties && v.properties.type && v.properties.type.enum && v.oneOf) {
        v.oneOf.forEach(({allOf}) => {
            try {
                const name = allOf[0].properties.type.const;
                const type = allOf[1].properties.spec.$ref.slice(14);
                struct.types.push({
                    name: name,
                    type: type,
                })
            } catch (e) {
                console.log(v)
            }
        });
    }

    // choose the template type
    let template = `scripts/templates/docs.handlebars`;
    if (!struct.type && struct.enum.length > 0) {
        template = `scripts/templates/docs_enum.handlebars`;
    }

    // file name
    const file = v["x-docs-title"] || Case.kebab(k);

    // parse the handlebars templates
    const text = fs.readFileSync(template);
    const tmpl = Handlebars.compile(text.toString());
        
    // execute the template and write the contents
    // to the struct filepath.
    fs.writeFileSync(`docs/content/reference/${file}.md`, tmpl(struct).replaceAll("&quot;", `"`)); 

    console.log(`docs/content/reference/${file}.md`)
});
