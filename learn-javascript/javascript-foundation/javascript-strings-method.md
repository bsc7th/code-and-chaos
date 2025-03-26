# Useful string method

## Strings as objects

Most values can be used as if they are objects in JavaScript. When you create a string, for example by using

```javascript
const string = "This is my string";
```

although the variable itself isn't an object, it still has a large number of properties and methods available to it, by virtue of being usable as an object when accessing properties. You can see this if you go to the [String](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String) object page and look down the list on the side of the page!

## Finding the length of a string

This is easy — you use the `length` property. Try entering the following lines:

```javascript
const browserType = "mozilla";
browserType.length; // outputs 7
```

This should return the number 7, because "mozilla" is 7 characters long. This is useful for many reasons; for example, you might want to find the lengths of a series of names so you can display them in order of length, or let a user know that a username they have entered into a form field is too long if it is over a certain length.

## Retrieving a specific string character

On a related note, you can return any character inside a string by using square bracket notation — this means you include square brackets ([]) on the end of your variable name. Inside the square brackets, you include the number of the character you want to return, so for example to retrieve the first letter you'd do this:

```javascript
browserType[0];
```

**REMEMBER**: computers count from 0, not 1!

To retrieve the last character of any string, we could use the following line, combining this technique with the `length` property we looked at above:

```javascript
browserType[browserType.length - 1];
```

The length of the string "mozilla" is 7, but because the count starts at 0, the last character's position is 6; using length-1 gets us the last character.
