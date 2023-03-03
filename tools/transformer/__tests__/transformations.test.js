const fs = require("fs");
const { test, beforeEach, expect, jest: jestGlobal } = require("@jest/globals");
const {
  applyModelNameTransformations,
  transformAllOf,
  transformOneOf,
} = require("../src/transformations");
const cases = require("./transformations-snapshots");

let api = {};
beforeEach(() => {
  api = JSON.parse(fs.readFileSync(__dirname + "/input.json", "utf8"));
});

test("Transform oneOf enum", () => {
  transformOneOf(
    ".components.schemas.ApiAtlasRegionConfigView.properties.regionName",
    api
  );
  expect(
    api.components.schemas.ApiAtlasRegionConfigView.properties.regionName
  ).toMatchInlineSnapshot(cases.EnumOneOf);
});

test("Transform oneOf model", () => {
  transformOneOf(".components.schemas.ApiAtlasHardwareSpecView", api);
  expect(
    api.components.schemas.ApiAtlasHardwareSpecView.properties
  ).toMatchInlineSnapshot(cases.PropertiesOneOf);
});

test("Transform AllOf model", () => {
  transformAllOf(".components.schemas.ApiAtlasRegionConfigView", api);
  expect(api.components.schemas.ApiAtlasRegionConfigView).toMatchInlineSnapshot(
    cases.ParentAllOf
  );
  expect(
    api.components.schemas.ApiAtlasAWSRegionConfigView
  ).toMatchInlineSnapshot(cases.ChildAllOf);
});

test("Fail Transform AllOf with wrong object structure", () => {
  transformAllOf(".components.schemas.ApiAtlasRegionConfigView", api);
  expect(() =>
    transformAllOf(".components.schemas.ApiAtlasRegionConfigView", api)
  ).toThrow(/Invalid object for AllOf Transformation/);
});

test("applyModelNameTransformations", () => {
  api = applyModelNameTransformations(api, "ApiAtlas", "View");
  api = applyModelNameTransformations(api, "Api", "View");
  for (const modelKey in api.components.schemas) {
    expect(modelKey.startsWith("ApiAtlas")).toBeFalsy();
    expect(modelKey.startsWith("Api")).toBeFalsy();
    expect(modelKey.endsWith("View")).toBeFalsy();
  }
});
