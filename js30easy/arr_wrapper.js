/**
 * @param {number[]} nums
 */
var ArrayWrapper = function(nums) {
    this.nums = nums
};

ArrayWrapper.prototype.valueOf = function() {
    return this.nums.reduce((acc, n) => acc + n, 0)
}

ArrayWrapper.prototype.toString = function() {
    return `[${this.nums.join(',')}]`
} 