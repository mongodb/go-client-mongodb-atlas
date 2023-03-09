
module.exports = function isModelIgnored(name, ignoredModelNames){
    return ignoredModelNames.includes(name);
}
