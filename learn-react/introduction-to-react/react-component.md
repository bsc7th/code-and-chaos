# React Component Model

A component is a piece of the `UI` that has it's own logic and appearance.

## What is Components?

Components are independent and reusable pieces of code. They serve the same purpose as JavaScript functions but work in isolation and return HTML.

## Things to Remember When Creating a Component

1. Component names should start with an uppercase letter:
   React treats components that start with lowercase letters as DOM tags. To ensure React recognizes your component as a custom component, it should begin with an uppercase letter (e.g., `MyComponent` instead of `myComponent`).
2. Components must return JSX (or null):
   Every React component must return something, typically JSX, which defines the structure of the UI. If no UI needs to be rendered, the component can return `null`.

## How to Create a Component?

1. Regular Function

```jsx
function App() {
  return (
    <>
      <h1>Hello, World!</h1>
    </>
  );
}

export default App;
```

**Explanation**

- `function App() {}`:
  This is a regular function declaration. In JavaScript, you can define a function using the function keyword. The function name here is App.
- return (...):
  The return statement is used to specify the output of the function. In this case, the function returns JSX (JavaScript XML), which is a syntax extension used by React. It looks similar to HTML but can contain JavaScript expressions inside curly braces {}.
- `<> and </>` (React Fragment):
  The <> and </> are React fragments. A React fragment allows you to group multiple elements without adding extra nodes to the DOM. This is helpful when you want to return multiple elements from a component but don't want to wrap them in a parent HTML element like a div. The fragment is shorthand for <React.Fragment>...</React.Fragment>.

- `<h1>Hello, World!</h1>`:
  This is just standard JSX syntax for an HTML heading (h1 tag) that contains the text "Hello, World!".
- `export default App;`:
  This line exports the App function so it can be imported into other parts of the application. The default keyword means that App is the default export from this module.

2. Arrow Function

```jsx
const App = () => {
  return (
    <>
      <h1>Hello, World!</h1>
    </>
  );
};

export default App;
```

**Explanation**

- `const App = () => {}`:
  This is an arrow function. It’s a shorter and more concise way of writing functions in JavaScript. Instead of using the function keyword, you use const (or let or var) to assign a function to a variable (App in this case), and then the => syntax defines the function.
- `return (...)`:
  As with the regular function, the arrow function returns JSX. This is the output of the component.
- JSX Inside the Arrow Function:
  The JSX syntax and fragment usage inside the arrow function are the same as in the regular function.
- `export default App;`:
  This line still exports the App component. The behavior is identical to the previous code.

3. class Component(no long recommended)

```jsx
class App extends Component {
  render() {
    return (
      <main>
        <h1>Hello World!</h1>
      </main>
    );
  }
}

export default App;
```

**Key Differences Between the Two Versions:**

1. Syntax:

- The first one uses a regular function declaration: `function App() {...}`.
- The second one uses an arrow function expression: `const App = () => {...}`.

2. Binding of `this`:

- **Regular functions** have their own this context, meaning that inside the function, this refers to the context of the function (which might change depending on how the function is invoked).
- **Arrow functions** do not have their own this context. Instead, they inherit this from the surrounding lexical scope (the scope in which they are defined). This can help avoid common pitfalls when working with this inside event handlers or callback functions.

3. Return Statement:

- The regular function declaration could potentially have an implicit return if simplified (for example, omitting the `return` keyword and using parentheses around the JSX).
- The arrow function can also have an implicit return if written as `const App = () => (<h1>Hello, World!</h1>)`. This is a common pattern in React.
