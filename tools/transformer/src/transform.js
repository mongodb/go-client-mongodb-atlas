const {
  applyAllOfTransformations,
  applyOneOfTransformations,
  applyModelNameTransformations,
  applyDiscriminatorTransformations,
} = require("./transformations");
const { getAPI, saveAPI } = require("./engine/apifile");

const log = require("simple-node-logger").createSimpleLogger();

// Override default logger
global.console = log;
log.setLevel(process.env.XGEN_LOGGING_LEVEL || "warn");

let { doc, apiFileLocation } = getAPI(process.argv.slice(2));

doc = applyDiscriminatorTransformations(doc);
doc = applyOneOfTransformations(doc);
doc = applyAllOfTransformations(doc);
doc = applyModelNameTransformations(doc, "ApiAtlas", "View");

saveAPI(doc, apiFileLocation);
