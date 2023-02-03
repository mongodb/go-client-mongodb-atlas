const { readFileSync, writeFileSync } = require('fs');
const {
  applyAllOfTransformations,
  applyOneOfTransformations,
  applyModelNameTransformations,
  applyResponseTransformation
} = require('./transformations');
const { getAllObjectsWitProperty } = require('./utils');
const { parse, stringify } = require('yaml');

let apiFileLocation = process.argv.slice(2);
console.error(apiFileLocation)
if(!apiFileLocation && !apiFileLocation[0]){
  throw new Error("Missing positional argument for openapi file")
}

apiFileLocation=apiFileLocation[0]

// Name of the key (OpenAPI extension) to trigger transformation
const extensionKey = 'x-xgen-go-transform';

const extensionAllOfValue = 'merge-allOf';
const extensionOneOfValue = 'merge-oneOf';

let doc;
if (apiFileLocation) {
  doc = parse(readFileSync(apiFileLocation, 'utf8'));
} else {
  throw new Error('Missing location. Please set OPENAPI_FILE env variable');
}

const allOfTransformations = getAllObjectsWitProperty(doc, extensionKey, (_k, v) => v === extensionAllOfValue);
const oneOfTransformations = getAllObjectsWitProperty(doc, extensionKey, (_k, v) => v === extensionOneOfValue);

console.info('Processing transformation rules');
console.info(
  '# AllOf transformations: ',
  allOfTransformations.map((e) => e.path)
);
console.info(
  '# OneOf transformations: ',
  oneOfTransformations.map((e) => e.path)
);

doc = applyOneOfTransformations(doc, oneOfTransformations);
doc = applyAllOfTransformations(doc, allOfTransformations);

doc = applyModelNameTransformations(doc, 'ApiAtlas', 'View');

writeFileSync(apiFileLocation, stringify(doc));

