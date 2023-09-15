/**
 * @param {Object | Array} obj
 * @return {boolean}
 */
var isEmpty = function(obj) {
    if (Object.entries(obj).length) return false
    return true
};