/**
 * @param {integer} init
 * @return { increment: Function, decrement: Function, reset: Function }
 */
var createCounter = function(init) {
    let initVal = init
    return {
        increment: () => {
            return ++init
        },
        decrement: () => {
            return --init
        },
        reset: () => {
            init = initVal
            return init
        }
    }
};