# List

In React, a "list" typically refers to a collection of items that are displayed on the screen. This collection can be dynamically generated using JavaScript, and React provides an efficient way to render these lists.

In practice, a list in React is often created by:

- Mapping over an array: This is a common way to render a list of items. You use JavaScript's `map()` function to iterate over an array and return an array of React elements.
- Using the `key` prop: When rendering lists in React, it’s important to assign a unique `key` to each element in the list. This helps React efficiently update and reconcile the list when the data changes.

Example of list in React

```jsx
import React from "react";

function NameList() {
  const names = ["Alice", "Bob", "Charlie"];

  return (
    <ul>
      {names.map((name, index) => (
        <li key={index}>{name}</li>
      ))}
    </ul>
  );
}

export default NameList;
```

Key points:

- The `map()` method is used to iterate over the `names` array.
- Each `<li>` element receives a `key` prop with a unique value (`index` in this case).
- The list is wrapped in a `<ul>` element, which is a common HTML element used for unordered lists.

Why `key` is important?

The `key` helps React identify which items have changed, are added, or are removed, thus optimizing performance. It should be a stable identifier, usually a unique ID from the data or, in simple cases, the index of the item (though using an index as the key isn't always the best approach when the list can change order).

Another example

```jsx
const App = () => {
  const numbers = [1, 2, 3, 4, 5];

  return (
    <main>
      <ul>
        {numbers.map((number) => (
          <li key={number}>{number}</li>
        ))}
      </ul>
    </main>
  );
};

export default App;
```
