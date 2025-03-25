# JavaScript Operators

JavaScript Operators are used to perform different types of mathematical, logical, and comparison.

1.  Arithmetic Operators - Used to perform arithmetic on numbers

- `+` Addition operator (adds values)

```javascript
let x = 5;
let y = 2;
let z = x + y; // z = 7
```

- `-` Substraction operator (substract values)

```javascript
let x = 5;
let y = 2;
let z = x - y; // z = 3
```

- `*` Multiplication operator (multiplies values)

```javascript
let x = 5;
let y = 2;
let z = x \* y; // z = 10
```

- `/` Division operator (divides values)

```javascript
let x = 5;
let y = 2;
let z = x / y; // z = 2.5
```

- `%` Modulus operator (returns the remainder of division)

```javascript
let x = 5;
let y = 2;
let z = x % y; // z = 1
```

- `**` Exponentiation operator (raises the base to the exponent power)

```javascript
let x = 5;
let z = x ** 2; // z = 25
```

**Alternative**: x \*\* y produces the same result as `Math.pow(x, y)`

- `++` Increment operator (increase the value by 1)

```javascript
let x = 5;
x++; // x = 6
```

- `--` Decrement operator (decrease value by 1)

```javascript
let x = 5;
x--; // x = 4
```

2. Assignment Operator - Used to assign values

- `=` Assignment operator (assigns a value to a variable)

```javascript
let x = 10; // x = 10
```

3. Comparison Operators - Used to compare values

- `>` Greater than operator (check if the value on the left is greater than the value on the right)

```javascript
let x = 5;
let y = 3;
console.log(x > y); // true
```

- `<` Less than operator (check if the value on the left is less than the value on the right)

- `>=` Greater than or equal to operator (checks if the value on the left is greater than or equal to the value on the right)

```javascript
let x = 5;
let y = 5;
console.log(x >= y); // true
```

- `<=` Less than or equal to operator (check if the value on the left is less than or equal to the value on the right)

```javascript
let x = 5;
let y = 8;
console.log(x <= y); // true
```

- `===` Stick equality operator (compares values and types for equality)

```javascript
let x = "5";
let y = 5;
console.log(x === y); // false (strict comparison)
```

- `!=` Not equal to operator (compares values for inequality)

```javascript
let x = 5;
let y = 8;
console.log(x != y); // true
```

- `!==` Strict not equal to operator (compares values and types for inequality)

```javascript
let x = "5";
let y = 5;
console.log(x !== y); // true
```

4. Logical Operators - Useed to perform logic operations

- `&&` Logical AND (returns true if both operands are true)

```javascript
let x = true;
let y = false;
console.log(x && y); // false
```

- `||` Logical OR (returns true if at least one operand is true)

```javascript
let x = true;
let y = false;
console.log(x || y); // true
```

- `!` Logical NOT (reverses the truth value of the operand)

```javascript
let x = true;
console.log(!x); // false
```

5. Ternary (Conditional) Operator - A shorthand for `if...else` statement

- `condition ? exprIfTrue : exprIfFalse`

```javascript
let x = 5;
let result = x > 3 ? "greater" : "lesser";
console.log(result); // 'greater'
```

6. Type Operators

- `typeof` Returns the type of a variable

```javascript
let x = 5;
console.log(typeof x); // 'number'
```

- `instanceof` Checks if an object is an instance of a specified class

```javascript
let x = [1, 2, 3];
console.log(x instanceof Array); // true
```
