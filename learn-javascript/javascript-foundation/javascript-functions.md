# Functions

Let’s discuss parameters and arguments in the context of the following example function:

```javascript
function aboutMe(name) {
  return name + " is my first name";
}

console.log(aboutMe("RJ"));
```

In JavaScript, parameters are the items listed between the parentheses () in the function declaration. Function arguments are the actual values we decide to pass to the function.

In the example above, the function definition is written on the first line: `function` `aboutMe(name)`. The parameter, `name`, is found inside the parentheses.

By putting `name` inside the parentheses of `aboutMe()` function, we are telling JavaScript that we will send some value to our `aboutMe` function. This means that `name` is just **placeholder** for some future value.

The last line `(aboutMe("RJ))`, is where we are calling our `aboutMe` function and passing the value `'RJ'` inside the function. Here, `'RJ'` is our argument. We are telling the `aboutMe` function, "Please send `'RJ'` to the `aboutMe` function and use `"RJ"` whenever the `aboutMe` placeholder is". Because of the flexibility that using a parameter provides, we can declare additional information to be our `aboutMe`.
