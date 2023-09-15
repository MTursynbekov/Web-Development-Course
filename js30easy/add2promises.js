/**
 * @param {Promise} promise1
 * @param {Promise} promise2
 * @return {Promise}
 */
var addTwoPromises = async function(promise1, promise2) {
    let p1 = await promise1
    let p2 = await promise2
    return p1+p2
};