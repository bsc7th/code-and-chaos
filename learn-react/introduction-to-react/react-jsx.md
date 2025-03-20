# What is JSX?

JSX (JavaScript XML) is a syntax extension for JavaScript that allows you to write HTML-like code within your JavaScript files. It makes writing React components more intuitive by combining JavaScript logic and HTML structure in a single file.

## How JSX is related to React?

1. HTML-like Syntax in JavaScript:
   JSX allows you to write HTML tags inside JavaScript code. It looks like regular HTML, but it is actually syntactic sugar for `React.createElement()` calls. For example, this JSX:

```jsx
<h1>Hello, World!</h1>
```

is transformed into:

```JavaScript
React.createElement('h1', null, 'Hello, World!');
```

2. React Components:
   React relies heavily on JSX because it makes defining the structure of components easier and more readable. A React component can return JSX to describe what the UI should look like, combining HTML and JavaScript logic.

3. Dynamic Rendering:
   JSX can also include JavaScript expressions inside curly braces `{}`, which allows dynamic rendering. For example:

```jsx
const name = "Alice";
<h1>Hello, {name}!</h1>;
```

This will render: `Hello, Alice!`, where `name` is a JavaScript variable.

4. Transpilation:
   JSX is not natively understood by browsers, so it needs to be transpiled (converted) into regular JavaScript. Tools like Babel are used in React projects to compile JSX into JavaScript code that browsers can interpret.

## Why Use JSX in React?

- Improves Readability: JSX allows you to write code that closely resembles the HTML structure of your components, making it easier for developers to visualize and understand the UI.
- Integrates Logic and UI: JSX enables you to easily embed JavaScript logic (like loops, conditionals, and variables) directly inside your component's structure.

## { Expressions in JSX }

In JSX, we use curly braces {} to embed JavaScript expressions within the markup. This is how you can mix JavaScript code with your HTML-like JSX code.

Example:

```jsx
const Hero = () => {
  return (
    <main>
      <form>
        <label htmlFor="name">{name}</label>
        <input type="text" placeholder="Enter your name" />
      </form>
    </main>
  );
};

export default Hero;
```

Another example when working with numbers:

In this example curly braces are necessary to evaluate and render JavaScript expressions (like `2 + 2)` inside the JSX. Without them, the JSX code will be treated as plain text, not dynamic JavaScript.

```jsx
const App = () => {
  return (
    <div>
      <p>2 + 2 = {2 + 2}</p>
    </div>
  );
};

export default App;
```

You can also create a special class like this:

This line defines a constant called `specialClass` and assigns it the string `"app-container"`. In this case, `specialClass` is just a string that contains a class name, which can be used later to apply styling or other behavior in a CSS file or JavaScript code.

```jsx
const specialClass = "app-container";

const App = () => {
  return (
    <div className={specialClass}>
      <p>2 + 2 = {2 + 2}</p>
    </div>
  );
};

export default App;
```

## Why use a variable for className?

The main reasons to use a variable (like `specialClass`) would be:

- **Reuse**: If you want to use the same class name in multiple places, defining it in one variable makes it easier to update.
- **Dynamic Class Names**: If you plan to dynamically change the class name based on some conditions or states in your component, using a variable allows you to easily switch class names
