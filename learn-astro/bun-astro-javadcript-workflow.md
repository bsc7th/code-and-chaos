# Astro React Workflow Built with JavaScript and Bun

In this guide, we will walk through setting up a modern development environment for an Astro project with React, using JavaScript and Bun as the package manager. This setup ensures fast build times and modern tooling to enhance your development workflow.

## Manual Indtallation

Start by creating a new project directory.

```zsh
mkdir dir-name
z dir-name
```

This will create a new directory and move into it.

## Initialize Git (Optional)

While this step is optional, it’s always a good idea to track your project with Git.

```zsh
git init
```

This will initialize a new Git repository in your project folder.

## Configure Bun

Next, let's configure Bun to work smoothly with Git. First, create the .gitattributes file:

```zsh
vim .gitattributes
```

Add the following content to the file:

```
*.lockb binary diff=lockb
```

Then, run these commands to configure Git for Bun's lock files:

```zsh
git conifg diff.lockb.textconv bun
git config diff.lockb.binary true
```

This ensures that `.lockb files`, which are specific to Bun, are handled correctly by Git.

## Initialize Bun Project

Now, initialize a new Bun project by running:

```zsh
bun init --yes
```

This command will automatically create a `package.json` file for you.

## Add Scripts to Use Bun

Open your `package.json` file and add the following scripts section to enable Bun to run your Astro and React commands:

```json
"scripts": {
    "dev": "bunx --bun astro dev",
    "build": "bunx --bun astro check && astro build",
    "preview": "bunx --bun astro preview",
    "lint": "bunx --bun eslint .",
    "format": "bunx --bun pretttier --write .",
    "astro": "astro"
  },
```

### Explanation:

- dev: Starts the Astro development server.
- build: Builds the project for production.
- preview: Previews the build locally.
- lint: Runs ESLint to check for any code issues.
- format: Uses Prettier to format the code.
- astro: Runs Astro commands directly.

## Integrate Astro

**IMPORTANT** Astro must be installed locally, not globally. Make sure you are not running `bun add global astro`

```zsh
bun add astro
```

This will add Astro as a dependency in your project.

## Integrate React

Now, let’s integrate React into the project. Use the following command to add React to your Astro project:

```zsh
bun astro add react
```

This command will install all necessary dependencies to enable React components in your Astro project including `astro.config.mjs` file.

## Install prettier & prettier-plugin-astro

To ensure consistent code formatting, let's add Prettier along with the Astro-specific Prettier plugin:

```zsh
bun add -d prettier prettier-plugin-astro
```

Next, create a `.prettierrc.mjs` file at the root of your project and configure it to work with Astro:

```mjs
/** @type {import("prettier").Config} */
export default {
  arrowParens: "avoid",
  singleQuote: true,
  bracketSpacing: true,
  endOfLine: "lf",
  semi: false,
  tabWidth: 2,
  trailingComma: "none",
  plugins: ["prettier-plugin-astro"],
  overrides: [
    {
      files: "*.astro",
      options: {
        parser: "astro",
      },
    },
  ],
};
```

## Install ESLint for Code Linting

Next, install ESLint and the Astro ESLint plugin to maintain high-quality, error-free code:

```mjs
bun add -d eslint eslint-plugin-astro
```

Then, create an `eslint.config.mjs` file to configure ESLint:

```zsh
vim eslint.config.mjs
```

Add the following content to the file:

```mjs
import { defineConfig } from "eslint-define-config";

export default defineConfig([
  {
    files: ["*.astro"],
    parser: "astro-eslint-parser",
    parserOptions: {
      extraFileExtensions: [".astro"],
    },
    plugins: {
      astro: import("eslint-plugin-astro"),
    },
    extends: ["plugin:astro/recommended"],
  },
  {
    files: ["*.js", "*.jsx"],
    parser: "espree",
    plugins: {
      astro: import("eslint-plugin-astro"),
    },
    extends: ["eslint:recommended", "plugin:astro/recommended"],
  },
]);
```

This configuration enables linting for .astro files, helping you follow best practices and maintain code quality.

## Creating the astro.config.mjs File

In some cases, Astro might not automatically generate the `astro.config.mjs` file after integrating React. If that happens, you can create the file manually:

```zsh
vim astro.config.mjs
```

Add the following configuration:

```mjs
import { defineConfig } from "astro/config";
import react from "@astro/react";

// https://astro.build/config
export default defineConfig({
  integrations: [react()],
  devToolbar: {
    enabled: true,
  },
  server: {
    port: 3000,
  },
});
```

This configuration enables the React integration in Astro and sets up the development server to run on port `3000`.

## Create Project Directory Structure

Let’s create the directory structure for your project:

```zsh
mkdir -p src/layouts src/pages src/components/astro src/conponents/react
```

This will create the necessary folders for organizing your layouts, pages, and components.

## Create a Layout Component

Now, navigate to the `src/layouts` directory and create a layout file, such as `Base.astro`. You can use `Main.astro` or any other name based on your preference.

```zsh
z src/layouts
vim Base.astro
```

Add the following content to your Base.astro file:

```astro
<!doctype html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width" />
    <meta name="generator" content={Astro.generator} />
    <title>Astro Basics</title>
  </head>
  <body>
    <slot />
  </body>
</html>
```

This layout component defines the basic HTML structure and a `<slot />` where the page content will be injected.

## Create a Page Component

Next, navigate to the `src/pages` directory and create an `index.astro` file:

```zsh
z src/pages
vim index.astro
```

Add the following content to index.astro

```astro
---
import Base from '../layouts/Base.astro';
---

<Base>
  <h1>Welcome to Astro with React!</h1>
</Base>
```

This sets up the main page of your project, using the `Base.astro` layout and displaying a simple welcome message.

## Congratulations 🚀

You’ve successfully set up an Astro project with React, using JavaScript and Bun as the package manager! 🎉

Now you can start building your modern web applications with Astro and React, enjoying fast builds, efficient development, and excellent tooling.
