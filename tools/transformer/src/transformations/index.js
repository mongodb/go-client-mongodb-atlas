const { applyAllOfTransformations, transformAllOf } = require("./allOf");
const { applyModelNameTransformations } = require("./name");
const { applyOneOfTransformations, transformOneOf } = require("./oneOf");

module.exports = {
  applyModelNameTransformations,
  transformAllOf,
  transformOneOf,
  applyAllOfTransformations,
  applyOneOfTransformations,
};
