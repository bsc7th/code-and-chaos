# Introduction to JavaScript Fundamentals

JavaScript is a scripting language that enables you to create dynamically updating content, control multimedia, animate images, and pretty much everything else.

## How to run JavaScript code?

1. Inside the **body** section.

```HTML
<body>
    <script>
        let greeting = "Hello, World!"
        console.log(greeting);
    </script>
</body>
```

**console.log()** is the command to print something to the developers console in your browser.

2. Another way to include JavaScript in a webpage is through an **external script**. This is very similar to linking external CSS docs to your website.

```HTML
<script src="scripts.js"></script>
```

## Variables

These are the building blocks of any program, you can think of varaibles as **storage containers** for data in your code.

You can declare variables using the following:

1. let, which can re-assign.

```javascript
let name = "John";
let surname = "Doe";

console.log(name);
console.log(surname);
```

You can also re-assign variables:

```javascript
let age = 11;
console.log(age); // outputs 11 to the console

age = 18;

console.log(age); // outups 18 now
```

- when re-assigning a variables we don't need to explicitly use `let` since the the variable has been declared earlier, **re-assigning is allowed only within the same scope**.

- trying to declare a variable again with `let` in the same scope will throw an error.

```javascript
let age = 18;
let age = 25; // This will throw an error: "Uncaught SyntaxError: Identifier 'age' has already been declared"
```

2. const - which we can't re-assign and will throw an error if we try.

```javascript
const pi = 3.14;
pi = 10;

console.log(pi); // an error os thrown
```

3. var - which was the original way variables were declared in JavaScript. `var` is similar to `let`

- `var` is function-scoped, while `let` and `const` are block-scoped.
- `var` allows variable hoisting, which can lead to unexpected behavior.

```javascript
console.log(x); // undefined (due to hoisting)
var x = 5;
```
