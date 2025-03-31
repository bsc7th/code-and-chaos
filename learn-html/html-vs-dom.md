# HTML vs DOM

- the HTML defines the structure, and the DOM allows JavaScript to manipulate the content dynamically.

## HTML (Hyper Text Markup Language)

- HTML is the standard language used to create and structure content on the web. It defines the structure of web pages using markup (tags).
- HTML specifies the elements and their content, such as paragraphs, links, headings, images, etc. It’s static and defines how the content should be structured.

example:

```html
<html>
  <body>
    <h1 id="message">Hello, World!</h1>
  </body>
</html>
```

## DOM (Document Object Model)

- The DOM is a programming interface for web documents. It represents the structure of an HTML (or XML) document as a tree of objects. The DOM is dynamic and can be manipulated using programming languages like JavaScript.
- The DOM allows you to interact with and manipulate the structure, style, and content of web pages after they have been loaded into the browser. Changes to the DOM update the page dynamically.

example: If you have the HTML <h1 id="message">Hello, World!</h1>, in the DOM, it becomes an object that JavaScript can access and modify.

```javascript
document.getElementById("message").innerHTML = "Hello, DOM!";
```

in short, the HTML defines the structure, and the DOM allows JavaScript to manipulate the content dynamically.
