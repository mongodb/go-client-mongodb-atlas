const fs = require("fs");
const { test, beforeEach, expect, jest: jestGlobal } = require("@jest/globals");
const {
  applyAllOfTransformations,
  applyOneOfTransformations,
  applyModelNameTransformations,
} = require("../src/transformations");
const cases = require("./transformations-snapshots");

let api = {};
beforeEach(() => {
  api = JSON.parse(fs.readFileSync(__dirname + "/input.json", "utf8"));
});

test("Empty doc. No transformations", () => {
  applyOneOfTransformations({}, []);
  applyAllOfTransformations({}, []);
});

test("Invalid Path", () => {
  expect(() => applyOneOfTransformations({}, [{ obj: undefined }])).toThrow(
    /Invalid transformation object/
  );
  expect(() => applyAllOfTransformations({}, [{ obj: undefined }])).toThrow(
    /Invalid transformation object/
  );
});

test("applyOneOfTransformations happy path", () => {
  applyOneOfTransformations(api, [
    {
      obj: api.components.schemas.ApiAtlasRegionConfigView.properties
        .regionName,
    },
    { obj: api.components.schemas.ApiAtlasHardwareSpecView },
  ]);
  expect(
    api.components.schemas.ApiAtlasRegionConfigView.properties.regionName
  ).toMatchInlineSnapshot(cases.EnumOneOf);
  expect(
    api.components.schemas.ApiAtlasHardwareSpecView.properties
  ).toMatchInlineSnapshot(cases.PropertiesOneOf);
});

test("applyAllOfTransformations happy path", () => {
  applyAllOfTransformations(api, [
    {
      obj: api.components.schemas.ApiAtlasRegionConfigView,
      path: "components.schemas.ApiAtlasRegionConfigView",
    },
  ]);
  expect(api.components.schemas.ApiAtlasRegionConfigView).toMatchInlineSnapshot(
    cases.ParentAllOf
  );
  expect(
    api.components.schemas.ApiAtlasAWSRegionConfigView
  ).toMatchInlineSnapshot(cases.ChildAllOf);
});

test("applyAllOfTransformations duplicates", () => {
  api.components.schemas.ApiAtlasAWSRegionConfigView.properties = {
    ...api.components.schemas.ApiAtlasAWSRegionConfigView.properties,
    ...api.components.schemas.ApiAtlasRegionConfigView.propertries,
  };
  global.console.warn = jestGlobal.fn();
  applyAllOfTransformations(api, [
    {
      obj: api.components.schemas.ApiAtlasRegionConfigView,
      path: "components.schemas.ApiAtlasRegionConfigView",
    },
  ]);
  applyOneOfTransformations(api, [
    { obj: api.components.schemas.ApiAtlasRegionConfigView },
  ]);
  expect(console.warn).toBeCalled();
});

test("applyAllOfTransformations wrong structure", () => {
  applyOneOfTransformations(api, [
    { obj: api.components.schemas.ApiAtlasRegionConfigView },
  ]);
  expect(() =>
    applyAllOfTransformations(api, [
      {
        obj: api.components.schemas.ApiAtlasRegionConfigView,
        path: "components.schemas.ApiAtlasRegionConfigView",
      },
    ])
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
