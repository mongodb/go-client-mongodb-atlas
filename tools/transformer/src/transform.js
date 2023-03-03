const {
  applyAllOfTransformations,
  applyOneOfTransformations,
  applyModelNameTransformations,
} = require("./transformations");
const { getAPI, saveAPI } = require("./apifile");

const log = require("simple-node-logger").createSimpleLogger();

// Override default logger
global.console = log;
log.setLevel(process.env.XGEN_LOGGING_LEVEL || "warn");

let { doc, apiFileLocation } = getAPI()();

doc = applyOneOfTransformations(doc);
doc = applyAllOfTransformations(doc);

doc = applyModelNameTransformations(doc, "ApiAtlas", "View");

saveAPI(doc, apiFileLocation);
