// This script generates Go code from the jsonschema
// bundle using handlebars templates. The templates are
// located in the ./templates subdirectory.

// This code is a work-in-progress.
// Please excuse the mess.

import fs from "fs";
import {exec} from "child_process";
import Case from "case";
import Handlebars from "handlebars";
import getType from "./helpers/types.js";

// parse the schema
let schema = JSON.parse(fs.readFileSync("dist/schema.json"));

// for each definition
Object.entries(schema.definitions).forEach(([k, v]) => {
    if (v["x-go-skip"] === true) {
        console.log("skip definition", v.title || v.name)
        return;
    }

    // name of the go file.
    const filename = v["x-file"].replace(".yaml", ".go");

    // store the go struct details.
    let struct = {
        name: k,
        desc: v.description,
        path: filename,
        enum:  [],
        props: [],
        types: [],
    }

    // for each property
    Object.entries(v.properties || {}).forEach(([propkey, propval]) => {
        const json = propkey;
        const name = Case.pascal(propkey);

        // append the field to the go struct
        struct.props.push({
            name: name,
            json: json,
            type: getType(propval, schema.definitions),
        });
    });

    Object.entries(v["x-properties"] || {}).forEach(([propkey, propval]) => {
        const json = propkey;
        const name = Case.pascal(propkey);

        // append the field to the go struct
        struct.props.push({
            name: name,
            json: json,
            type: getType(propval, schema.definitions),
        });
    });

    // for each enum value
    v.enum && v.enum.forEach(text => {
        struct.enum.push({name: struct.name + Case.pascal(text), text: text});
    });

    // this block of code detects if we are using the type /
    // spec pattern. If yes, we store the type enum values
    // and their associated struct types.
    if (v.properties && v.properties.type && v.properties.type.enum && v.oneOf) {
        v.oneOf.forEach(({allOf}) => {
            const name = allOf[0].properties.type.const;
            const type = allOf[1].properties.spec.$ref.slice(14);
            struct.types.push({
                name: name,
                type: type,
            })
        });
    }

    // parse the handlebars templates
    const text = fs.readFileSync(`scripts/templates/${struct.enum && struct.enum.length ? "enum": "struct"}.handlebars`);
    const tmpl = Handlebars.compile(text.toString());
    
    // execute the template and write the contents
    // to the struct filepath.
    fs.writeFileSync(`dist/go/${struct.path}`, tmpl(struct)); 

    console.log(`dist/go/${struct.path}`)
});

// format generated files
exec("gofmt -s -w dist/go/*.go", (err) => {
  if (err) {
    console.log("ensure gofmt is installed and in your PATH")
    console.error(err);
    process.exit(1)
  }
});

// import all dependencies if needed
exec("goimports -w dist/go/*.go", (err) => {
  if (err) {
    console.log("ensure goimports is installed and in your PATH")
    console.error(err);
    process.exit(1)
  }
});
