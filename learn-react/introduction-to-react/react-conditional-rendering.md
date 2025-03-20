# Conditional Rendering

In React, conditional rendering refers to the practice of rendering components or elements based on certain conditions or logic. This allows you to control what gets displayed to the user based on the state, props, or any other conditions that might change over time.

You can think of it like an `if-else statement` in JavaScript, but applied to JSX code within React components.

There are several ways to implement conditional rendering in React, including:

1. Using `if/else` Statements

You can use a simple JavaScript `if` statement to decide what JSX to render.

```jsx
function Greeting({ isLoggedIn }) {
  if (isLoggedIn) {
    return <h1>Welcome back!</h1>;
  } else {
    return <h1>Please log in</h1>;
  }
}
```

2. Using Ternary Operator

A more concise approach is to use the ternary operator (`condition ? true : false`), which is often used in JSX directly.

```jsx
function Greeting({ isLoggedIn }) {
  return <h1>{isLoggedIn ? "Welcome back!" : "Please log in"}</h1>;
}
```

3. Using Logical `&&` Operator

You can use the logical AND (&&) operator to conditionally render content. It works when you want to render something only if a condition is true.

```jsx
function Greeting({ isLoggedIn }) {
  return <div>{isLoggedIn && <h1>Welcome back!</h1>}</div>;
}
```

Here, `<h1>Welcome back!</h1>` will only be rendered if `isLoggedIn` is `true`.

4. Returning `null` to Render Nothing

You can return `null` to render nothing under certain conditions.

```jsx
function Greeting({ isLoggedIn }) {
  if (!isLoggedIn) {
    return null; // renders nothing
  }
  return <h1>Welcome back!</h1>;
}
```

5. Switch Case (In Functions)

For more complex conditional logic, you might use a `switch` case inside a function.

```jsx
function Greeting({ userRole }) {
  switch (userRole) {
    case "admin":
      return <h1>Admin Dashboard</h1>;
    case "user":
      return <h1>User Dashboard</h1>;
    default:
      return <h1>Welcome, Guest!</h1>;
  }
}
```

## Why use Conditional Rendering?

- It allows you to customize what users see based on their interactions with the app, such as logged-in status or user roles.
- It's great for handling different states (loading, error, success) in your app.
- It enables a dynamic and interactive user interface.

Overall, conditional rendering is an essential concept in React for building flexible and dynamic applications.
