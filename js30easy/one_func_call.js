/**
 * @param {Function} fn
 * @return {Function}
 */
var once = function(fn) {
    let once =false
    return function(...args){
        if (once) return undefined
        once = true
        return fn(...args)
    }
};