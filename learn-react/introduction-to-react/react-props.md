# Props

Props (short for properties) are arguments passed into React components. They allow data to be passed from a parent component to a child component, and props are provided to components through HTML attributes.

Props Example:

```jsx
const App = () => {
  return (
    <>
      <User name="RJ Leyva" birthday="2-9-2025" age="1" />
    </>
  );
};

const User = ({ name, birthday, age }) => {
  return (
    <>
      <h1>{name}</h1>
      <p>{birthday}</p>
      <p>{age}</p>
    </>
  );
};

export default App;
```

Let's break down:

`App` Component

```jsx
const App = () => {
  return (
    <>
      <User name="RJ Leyva" birthday="2-9-2025" age="1" />
    </>
  );
};
```

- Arrow Function: `const App = () => { ... }` is an arrow function that defines a functional component in React. This component returns JSX to be rendered on the screen.
- JSX: The `return` statement inside `App` contains JSX (which looks like HTML but is actually JavaScript).
- Inside this JSX, you are rendering a `User` component with three props: `name`, `birthday`, and `age`. These props are passed as attributes, similar to HTML attributes, but they hold values that can be used inside the User component.
  - `name="RJ Leyva"`
  - `birthday="2-9-2025"`
  - `age="1"`
- When `App` renders, it passes these values to the `User` component.

`User` Component

```jsx
const User = ({ name, birthday, age }) => {
  return (
    <>
      <h1>{name}</h1>
      <p>{birthday}</p>
      <p>{age}</p>
    </>
  );
};
```

- Arrow Function with Destructuring: `const User = ({ name, birthday, age }) => { ... }` is another arrow function, but this one defines the `User` component.
  - The `({ name, birthday, age })` syntax is destructuring the props object. This means instead of accessing props with `props.name`, `props.birthday`, and `props.age`, you're directly pulling those values from the props object and assigning them to variables with the same name.
  - In short: `const User = (props) => { const { name, birthday, age } = props }` is the same as `const User = ({ name, birthday, age })`.
- Returning JSX: The `User` component returns JSX to display the values passed as props. The `{}` syntax inside JSX is used to embed JavaScript expressions, so you're inserting the values of `name`, `birthday`, and `age` directly into the HTML-like structure.

  - `<h1>{name}</h1>` will display the name prop passed from App.
  - `<p>{birthday}</p>` will display the birthday prop.
  - `<p>{age}</p>` will display the age prop.

How it works together:

1. The `App` component passes values (`name`, `birthday`, `age`) to the `User` component as props.
2. The `User` component receives those props and displays them in an HTML structure.

This is displayed because the `name`, `birthday`, and `age` props are rendered inside an `h1` and two `p` tags.

Props allow data to be passed from a parent component (App) to a child component (User). Destructuring in the User component simplifies accessing and using these props. JSX is used to render the passed data in the browser.
