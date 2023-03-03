const { applyAllOfTransformations } = require("./allOf");
const { applyModelNameTransformations } = require("./name");
const { applyOneOfTransformations } = require("./oneOf");

module.exports = {
  applyModelNameTransformations,
  applyAllOfTransformations,
  applyOneOfTransformations,
};
