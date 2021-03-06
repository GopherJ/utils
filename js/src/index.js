const dayjs = require('dayjs');
const groupArray = require('group-array');

const timeUnitMap = {
    m: 60,
    h: 60  * 60,
    d: 24  * 60 * 60,
    w: 7   * 24 * 60 * 60,
    M: 30  * 24 * 60 * 60,
    y: 365 * 24 * 60 * 60
};

const monthsMap = {
    'Jan': 1,
    'Feb': 2,
    'Mar': 3,
    'Apr': 4,
    'May': 5,
    'Jun': 6,
    'Jul': 7,
    'Aug': 8,
    'Sep': 9,
    'Oct': 10,
    'Nov': 11,
    'Dec': 12
};

const isString = s => typeof s === 'string';
const isEmptyString = s => isString(s) && s === '';
const isNull = n => n === null;
const isUndefined = n => n === undefined;
const isArray = arr => Array.isArray(arr);
const isBoolean = b => typeof b === 'boolean';
const isDate = d => Object.prototype.toString.call(d) === '[object Date]' && !isNaN(d.getTime());
const isObject = o => Object.prototype.toString.call(o) === '[object Object]';
const isNumber = n => typeof n === 'number' && !isNaN(n);
const isArrayAndHasLength = arr => Array.isArray && arr.length > 0;
const isFunction = f => typeof f === 'function';
const isAllPass = (arr, det) => isArrayAndHasLength(arr) && isFunction(det) && arr.every(x => det(x));
const isInRange = (arr, n) => isArray(arr)
    && arr.length === 2
    && arr.every(x => isNumber(x))
    && isNumber(n)
    && n >= arr[0]
    && n <= arr[1];
const lastDayOfMonth = (y, m) => new Date(y, m, 0).getDate();
const formatDateHour = d => dayjs(d).format('YYYY-MM-DD HH:mm:ss');
const formatDate = d => dayjs(d).format('YYYY-MM-DD');
const formatHour = d => dayjs(d).format('HH:mm:ss');
const makeArray = (s, i) => new Array(s).fill(i);
const ascending = (b, a) => b < a ? -1 : b > a ? 1 : b >= a ? 0 : NaN;
const descending = (a, b) => b < a ? -1 : b > a ? 1 : b >= a ? 0 : NaN;
const descendingByProp = p => ((a, b) => b[p] < a[p] ? -1 : b[p] > a[p] ? 1 : b[p] >= a[p] ? 0 : NaN);
const ascendingByProp = p => ((b, a) => b[p] < a[p] ? -1 : b[p] > a[p] ? 1 : b[p] >= a[p] ? 0 : NaN);
const descendingByIdx = i => ((a, b) => b[i] < a[i] ? -1 : b[i] > a[i] ? 1 : b[i] >= a[i] ? 0 : NaN);
const descendingByIdxIdx = (i1, i2) => ((a, b) => b[i1][i2] < a[i1][i2] ? -1 : b[i1][i2] > a[i1][i2] ? 1 : b[i1][i2] >= a[i1][i2] ? 0 : NaN);
const descendingByIdxProp = (i, p) => ((a, b) => b[i][p] < a[i][p] ? -1 : b[i][p] > a[i][p] ? 1 : b[i][p] >= a[i][p] ? 0 : NaN);
const isTheSameStrIgnoreCase = (a, b) => isString(a) && isString(b) && a.toUpperCase() === b.toUpperCase();
const words = s => isEmptyString(s)
    ? []
    : /(([a-zA-Z]{2,})|([aI]))/.test(s)
    ? s.match(/(([a-zA-Z]{2,})|([aI]))/g)
    : [];
const uniqueWords = s => isEmptyString(s)
    ? []
    : /(([a-zA-Z]{2,})|([aI]))/.test(s)
    ? Array.from(new Set(s.match(/(([a-zA-Z]{2,})|([aI]))/g)))
    : [];
const percent = (v, t) => `${(v / t).toFixed(4) * 100}%`;
const repeat = (c, n) => n >= -1 ? new Array(n + 1).join(c) : '';
const noop = () => { };
const upperFirstChar = s => s[0].toUpperCase() + s.substring(1);
const diffOfSecs = (a, b) => !(isDate(a) && isDate(b))
    ? 0
    : Math.ceil(Math.abs(a.valueOf() - b.valueOf()) / 1000);
const convTimeUnitCombToNum = s => !(/^\s*[1-9]([hmdwMy])\s*$/.test(s))
    ? NaN
    : s.match(/^\s*([1-9])([hmdwMy])\s*$/).slice(1).reduce((ite, cur) => (+ite) * timeUnitMap[cur]);
const parseDate = s => {
    const re = /^\s*(?:([0-9]{4}))(?:[-/]([0-9]{1,2}))?(?:[-/]([0-9]{1,2}))?\s*$/;
    const matches = s.match(re);

    if (isNull(matches)) return new Error('Invalid Format');

    const [
        y,
        m,
        d,
    ] = matches.slice(1).map(x => +x);

    if (!isNumber(m) && !isNumber(d)) return new Date(y, 0);
    else if (!isNumber(d)) return !isInRange([1, 12], m)
            ? new Error('Out Of Range')
            : new Date(y, m - 1, 0);
    else return !isInRange([1, lastDayOfMonth(y, m)], d) || !isInRange([1, 12], m)
            ? new Error('Out Of Range')
            : new Date(y, m - 1, d);
};
const swap = (input, i, j) => isArray(input) || isString(input)
    ? ([input[i], input[j]] = [input[j], input[i]], input)
    : noop();
const randRange = (min, max) => Math.floor(Math.random(max - min + 1)) + min;
const map = Array.prototype.map;
const slice = Array.prototype.slice;
const range = (min, max) => makeArray(max - min + 1).map((_, i) => i + 1);
const execute = (task, n, ...args) => {
    if (!isFunction(task)) return;

    for (let i = 0; i < n; i += 1) {
        if (isArrayAndHasLength(args)) task(...args);
        else task();
    }
};
const inherit = (sub, sup) => isFunction(sub) && isFunction(sup)
    ? sub.prototype.__proto__ = sup.prototype
    : noop();
const shuffle = arr => {
    if (!isArrayAndHasLength(arr)) return;

    let i = arr.length;
    while (i-- > 0) {
        swap(arr, i, randInt(0, arr.length - 1));
    }
};
const groupArr = (arr, ...args)  => isArrayAndHasLength(arr)
    ? groupArray(arr, ...args)
    : noop();
const concat = (arr, ...items) => {
    if (!isArray) return;

    for (let i = 0, l = items.length; i < l; i += 1) {
        const item = items[i];
        if (isArrayAndHasLength(item)) arr.push(...item);
        else arr.push(item);
    }
};
const findMaxItemByProp = (arr, prop) => {
    if (!isArrayAndHasLength(arr) || isUndefined(prop) || !isString(prop)) return;
    let maxIdx = 0;
    for(let i = 0, l = arr.length; i < l; i += 1) {
        if (isObject(arr[i]) && isObject(arr[maxIdx]) && arr[i][prop] > arr[maxIdx][prop]) maxIdx = i;
    }
    return arr[maxIdx];
};
const findMinItemByProp = (arr, prop) => {
    if (!isArrayAndHasLength(arr) || isUndefined(prop) || !isString(prop)) return;
    let minIdx = 0;
    for(let i = 0, l = arr.length; i < l; i += 1) {
        if (isObject(arr[i]) && isObject(arr[minIdx]) && arr[i][prop] < arr[minIdx][prop]) minIdx = i;
    }
    return arr[minIdx];
};

module.exports = {
    isString,
    isEmptyString,
    isNull,
    isBoolean,
    isUndefined,
    isDate,
    isObject,
    isNumber,
    isArrayAndHasLength,
    isFunction,
    isAllPass,
    isInRange,
    lastDayOfMonth,
    formatDateHour,
    formatDate,
    formatHour,
    makeArray,
    ascending,
    ascendingByProp,
    descending,
    descendingByProp,
    descendingByIdx,
    descendingByIdxIdx,
    descendingByIdxProp,
    words,
    uniqueWords,
    diffOfSecs,
    convTimeUnitCombToNum,
    repeat,
    percent,
    noop,
    upperFirstChar,
    isTheSameStrIgnoreCase,
    parseDate,
    swap,
    randRange,
    map,
    slice,
    range,
    execute,
    inherit,
    shuffle,
    groupArr,
    concat,
    findMaxItemByProp,
    findMinItemByProp,

    timeUnitMap,
    monthsMap
}
