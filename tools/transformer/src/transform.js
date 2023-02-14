const { readFileSync, writeFileSync } = require("fs");
const {
  applyAllOfTransformations,
  applyOneOfTransformations,
  applyModelNameTransformations,
} = require("./transformations");
const { getAllObjectsWitProperty } = require("./utils");
const { parse, stringify } = require("yaml");

let apiFileLocation = process.argv.slice(2);
console.error(apiFileLocation);
if (!apiFileLocation && !apiFileLocation[0]) {
  throw new Error("Missing positional argument for openapi file");
}

apiFileLocation = apiFileLocation[0];

let doc;
if (apiFileLocation) {
  doc = parse(readFileSync(apiFileLocation, "utf8"));
} else {
  throw new Error("Missing location. Please set OPENAPI_FILE env variable");
}

doc = applyOneOfTransformations(doc);
doc = applyAllOfTransformations(doc);

doc = applyModelNameTransformations(doc, "ApiAtlas", "View");

writeFileSync(apiFileLocation, stringify(doc));
