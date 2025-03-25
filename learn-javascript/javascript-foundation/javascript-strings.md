# Handling Text - strings in JavaScript

strings, this is what pieces of text called in programming.

JavaScript offers many features for manipulating strings. These include creating custom welcome messages and prompts, showing the right text labels when needed, sorting terms into the desired order, and much more.J

## Declaring strings

Strings are dealt with similarly to numbers at first glance, but when you dig deeper you'll start to see some notable differences. The only difference is that when writing a string, you need to surround the value with quotes.

If you don't do this, or maybe you miss one, you'll get an error.

```javascript
let name = John
let username = 'Doe
let hobby = programming'
```

**NOTE**: You must use the same character for the start and end of a string, or you will get an error:

```javascript
let name = 'John" // this is not allowed
```

These lines don't work because any text without quotes around it is interpreted as a variable name, property name, reserved word, or similar. If the browser doesn't recognize the unquoted text, then an error is raised (e.g., `"missing; before statement"`).

## Single quotes, double quotes, and backticks

In JavaScript, you can choose single quotes ('), double quotes ("), or backticks (`) to wrap your strings in. All of the following will work:

```javascript
const single = "Single quotes";
const double = "Double quotes";
const backtick = `Backtick`;

console.log(single);
console.log(double);
console.log(backtick);
```

Strings declared using backticks are a special kind of string called a [template literal](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Template_literals). In most ways, template literals are like normal strings, but they have some special properties:

- you can [embed JavaScript](https://developer.mozilla.org/en-US/docs/Learn_web_development/Core/Scripting/Strings#embedding_javascript) in them.
- you can declare templates literals over [multiple lines](https://developer.mozilla.org/en-US/docs/Learn_web_development/Core/Scripting/Strings#multiline_strings).

## Embedding JavaScript

Inside a template literal, you can wrap JavaScript variables or expressions inside `${ }`, and the result will be included in the string:

```javascript
const name = "Chris";
const greeting = `Hello, ${name}`;
console.log(greeting); // "Hello, Chris"
```

You can use the same technique to join together two variables:

```javascript
const one = "Hello, ";
const two = "how are you?";
const joined = `${one}${two}`;
console.log(joined); // "Hello, how are you?"
```

Joining strings together like this is called concatenation.

## Concatenation in context

```html
<button>Press me</button>
<div id="greeting"></div>
```

```javascript
const button = document.querySelector("button");

function greet() {
  const name = prompt("What is your name?");
  const greeting = document.querySelector("#greeting");
  greeting.textContent = `Hello ${name}, nice to see you!`;
}

button.addEventListener("click", greet);
```

We are using the `window.prompt()` function, which prompts the user to answer a question via a popup dialog box and then stores the text they enter inside a given variable — in this case `name`. We then display a string that inserts the name into a generic greeting message.

## Concatenation using "+"

You can use `${}` only with template literals, not normal strings. You can concatenate normal strings using the `+` operator:

```javascript
const greeting = "Hello";
const name = "Chris";
console.log(greeting + ", " + name); // "Hello, Chris"
```

However, template literals usually give you more readable code:

```javascript
const greeting = "Hello";
const name = "Chris";
console.log(`${greeting}, ${name}`); // "Hello, Chris"
```

## Including expressions in strings

You can include JavaScript expressions in template literals, as well as just variables, and the results will be included in the result:

```javascript
const song = "Fight the Youth";
const score = 9;
const highestScore = 10;
const output = `I like the song ${song}. I gave it a score of ${
  (score / highestScore) * 100
}%.`;
console.log(output); // "I like the song Fight the Youth. I gave it a score of 90%."
```

## Multiline strings

Template literals respect the line breaks in the source code, so you can write strings that span multiple lines like this:

```javascript
const newline = `One day you finally knew
what you had to do, and began,`;
console.log(newline);

// output
/*
One day you finally knew
what you had to do, and began,
*/
```

To have the equivalent output using a normal string you'd have to include line break characters (`\n`) in the string:

```javascript
const newline = "One day you finally knew\nwhat you had to do, and began,";
console.log(newline);

// output
/*
One day you finally knew
what you had to do, and began,
*/
```

## Including quotes in strings

Since we use quotes to indicate the start and end of strings, how can we include actual quotes in strings? We know that this won't work:

```javascript
const badQuotes = "She said "I think so!"";
```

One common option is to use one of the other characters to declare the string:

```javascript
const goodQuotes1 = 'She said "I think so!"';
const goodQuotes2 = `She said "I'm not going in there!"`;
```

Another option is to escape the problem quotation mark. Escaping characters means that we do something to them to make sure they are recognized as text, not part of the code. In JavaScript, we do this by putting a backslash just before the character. Try this:

```javascript
const bigmouth = "I've got no right to take my place…"; // "I'\ve got no right to take my place…"
console.log(bigmouth);
```

You can use the same technique to insert other special characters. See [Escape sequences](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Lexical_grammar#escape_sequences) for more details.

## Numbers vs. strings

What happens when we try to concatenate a string and a number? Let's try it in our console:

```javascript
const name = "Front ";
const number = 242;
console.log(name + number); // "Front 242"
```

You might expect this to return an error, but it works just fine. How numbers should be displayed as strings is fairly well-defined, so the browser automatically converts the number to a string and concatenates the two strings.
