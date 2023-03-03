const { parse, stringify } = require("yaml");

/**
 * Read and parse API file as json
 * @returns  {doc, apiFileLocation}
 */
function getAPI(apiFileLocation) {
  let apiFileLocation = process.argv.slice(2);
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
  return { doc, apiFileLocation };
}

/**
 * Save API to target location
 * @param {*} doc openapi doc
 * @param {*} apiFileLocation location to save
 */
function saveAPI(doc, apiFileLocation) {
  writeFileSync(apiFileLocation, stringify(doc));
}

module.exports = { getAPI, saveAPI };
