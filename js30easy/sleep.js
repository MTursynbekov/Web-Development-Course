/**
 * @param {number} millis
 */
async function sleep(millis) {
    return new Promise(resolve => setTimeout(() => resolve(2), millis))
}
