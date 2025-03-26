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

## Testing if a string contains a substring

Sometimes you'll want to find if a smaller string is present inside a larger one (we generally say if a substring is present inside a string). This can be done using the [includes()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/includes) method, which takes a single [parameter](https://developer.mozilla.org/en-US/docs/Glossary/Parameter) — the substring you want to search for.

It returns `true` if the string contains the substring, and `false` otherwise.

```javascript
const browserType = "mozilla";

if (browserType.includes("zilla")) {
  console.log("Found zilla!");
} else {
  console.log("No zilla here!");
}
```

Often you'll want to know if a string starts or ends with a particular substring. This is a common enough need that there are two special methods for this: [startsWith()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/startsWith) and [endsWith()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/endsWith):

```javascript
const browserType = "mozilla";

if (browserType.startsWith("zilla")) {
  console.log("Found zilla!");
} else {
  console.log("No zilla here!");
}
```

```javascript
const browserType = "mozilla";

if (browserType.endsWith("zilla")) {
  console.log("Found zilla!");
} else {
  console.log("No zilla here!");
}
```

## Finding the position of a substring in a string

You can find the position of a substring inside a larger string using the [indexOf()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/indexOf) method. This method takes two parameters – the substring that you want to search for, and an optional parameter that specifies the starting point of the search.

If the string contains the substring, `indexOf()` returns the index of the first occurrence of the substring. If the string does not contain the substring, `indexOf()` returns -1.

```javascript
const tagline = "MDN - Resources for developers, by developers";
console.log(tagline.indexOf("developers")); // 20
```

Starting at 0, if you count the number of characters (including the whitespace) from the beginning of the string, the first occurrence of the substring "developers" is at index 20.

```javascript
console.log(tagline.indexOf("x")); // -1
```

So now that you know how to find the first occurrence of a substring, how do you go about finding subsequent occurrences? You can do that by passing in a value that's greater than the index of the previous occurrence as the second parameter to the method.

```javascript
const firstOccurrence = tagline.indexOf("developers");
const secondOccurrence = tagline.indexOf("developers", firstOccurrence + 1);

console.log(firstOccurrence); // 20
console.log(secondOccurrence); // 35
```

Here we're telling the method to search for the substring `"developers"` starting at index `21` (`firstOccurrence + 1`), and it returns the index `35`.
