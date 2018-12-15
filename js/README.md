# UtiJs

Light-weight Js Utilities

## Installation

```bash
npm i -S uti-js
```


## Usage

```js
const uti = require('uti-js'); 

uti.isDate(new Date()) // true
```


## API

- isString
  ```js
  uti.isString("hello world") === true 
  ```

- isEmptyString
  ```js
  uti.isEmptyString("") === true
  uti.isEmptyString("hello world") === false
  ```

- isNull
  ```js
  uti.isNull(nulll) === true
  ```

- isBoolean
  ```js
  uti.isBoolean(true) === true
  ```

- isUndefined
  ```js
  uti.isUndefined(undefined) === true
  ```

- isDate
  ```js
  uti.isDate(new Date()) === true
  ```

- isObject
  ```js
  uti.isObject({}) === true
  uti.isObject([]) === false
  ```

- isNumber
  ```js
  uti.isNumber(3) === true
  ```

- isArrayAndHasLength
  ```js
  uti.isArrayAndHasLength([1,2,3]) === true
  uti.isArrayAndHasLength([]) === false
  ```

- isFunction
  ```js
  uti.isFunction(function(){return 0}) === true
  ```

- isAllPass
  ```js
  uti.isAllPass([1, 2, 3], uti.isNumber) === true
  ```

- isInRange
  ```js
  uti.isInRange([1, 12], 5) === true
  ```

- lastDayOfMonth
  ```js
  uti.lastDayOfMonth(1, 2018) === 31
  ```

- formatDateHour
  ```js
  uti.formatDateHour(new Date()) === '2018-12-24 20:59:00'
  ```

- formatDate
  ```js
  uti.formatDate(new Date()) === '2018-12-24'
  ```

- formatHour
  ```js
  uti.formatHour(new Date()) === '20:59:00'
  ```

- makeArray
  ```js
  uti.makeArray(3, 0) // [0, 0, 0]
  ```

- ascending
  ```js
  [3, 1, 2].sort(uti.ascending) // [1, 2, 3]
  ```

- ascendingByProp
  ```js
  [{age: 10}, {age: 5}, {age: 20}].sort(uti.ascendingByProp('age'))
  // [{age: 5}, {age: 10}, {age: 20}]
  ```

- descending
- descendingByProp
- descendingByIdx
  ```js
  [['bob', 5], ['alice', 10]].sort(uti.descendingByIdx(1))
  // [['bob', 10], ['alice', 5]]
  ```

- descendingByIdxIdx
  ```js
  [['bob', [5]], ['alice', [10]]].sort(uti.descendingByIdxIdx(1, 0))
  // [['bob', [10]], ['alice', [5]]]
  ```
- descendingByIdxProp
  ```js
  [['bob', {age: 5}], ['alice', {age: 10}]].sort(uti.descendingByIdxProp(1, 'age'))
  // [['bob', {age: 10}], ['alice', {age: 5}]]
  ```

- words
  ```js
  uti.words("hello world")
  // ["hello", "world"]
  ```

- uniqueWords
  ```js
  uti.words("hello hello world")
  // ["hello", "world"]
  ```

- diffOfSecs
  ```js
  uti.diffOfSecs(new Date(2018, 10, 1), new Date(2018, 10, 2))
  // 24 * 60 * 60
  ```
  
- convTimeUnitCombToNum
  ```js
  uti.convTimeUnitCombToNum('3h') === 3 * 60 * 60
  ```

- percent
  ```js
  uti.percent(0.543) === "54.3%"
  ```

- noop
  ```js
  uti.noop()
  // nothing happens
  ```

- upperFirstChar
  ```js
  uti.upperFirstChar('hello') === 'Hello'
  ```

- isTheSameStrIgnoreCase
  ```js
  uti.isTheSameStrIgnoreCase('abc', 'AbC') === true
  ```

- parseDate
  ```js
  uti.parseDate('2018-10-10').valueOf() === new Date(2018, 10, 10)
  uti.parseDate('2018-10').valueOf() === new Date(2018, 10)
  uti.parseDate('2018').valueOf() === new Date(2018,0)

  uti.parseDate('2018-2-30') instanceof Error === true
  uti.parseDate('2018-13-30') instanceof Error === true
  uti.parseDate('2018-13-30sd') instanceof Error === true
  ```

- repeat
  ```js
  uti.repeat('-', 3) // '---'
  ```

- swap
  ```js
  uti.swap([1, 2, 3], 1, 2) // [1, 3, 2]
  ```

- randRange
  ```js
  uti.randRange(1, 5) // 3
  ```

- map
  ```js
  uti.map.call([1, 2], x => 2 * x) // [2, 4]
  ```

- slice
  ```js
  uti.slice.call([1, 2, 3], 0, 1) // 1
  ```

- range
  ```js
  uti.range(1, 5) // [1, 2, 3, 4, 5]
  ```

- execute
  ```js
  uti.execute(function(a, b) {console.log(a, b)}, 2, 1, 3)
  // 1 3
  // 1 3
  ```

- inherit
  ```js
  function List() {}
  uti.inherit(List, Array)

  const list = new List()
  list.push(3)

  console.log(list)
  // List { '0': 1, length: 1}
  ```

- shuffle
  ```js
  uti.shuffle([1, 9, 8, 6])
  // [9, 6, 8, 1]
  ```

- groupArr
  ```js
  const arr = [
    {tag: 'one', content: 'A'},
    {tag: 'one', content: 'B'},
    {tag: 'two', content: 'C'},
    {tag: 'two', content: 'D'},
    {tag: 'three', content: 'E'},
    {tag: 'three', content: 'F'}
  ];

  // group by the `tag` property
  uti.groupArr(arr, 'tag')
  /*
  {
    one: [
        {tag: 'one', content: 'A'},
        {tag: 'one', content: 'B'}
    ],
    two: [
        {tag: 'two', content: 'C'},
        {tag: 'two', content: 'D'}
    ],
    three: [
        {tag: 'three', content: 'E'},
        {tag: 'three', content: 'F'}
    ]
  }
  */
  ```