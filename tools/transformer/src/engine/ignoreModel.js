module.exports = function isModelIgnored(name, ignoredModelNames) {
  if (!ignoredModelNames) {
    return false;
  }
  return ignoredModelNames.includes(name);
};
