/**
 * @param {number[]} arr
 * @param {Function} fn
 * @return {number[]}
 */
var map = function(arr, fn) {
    let returnedArr = []
    for (let i = 0; i < arr.length; i++) {
        returnedArr.push(fn(arr[i], i))
    }

    return returnedArr
};