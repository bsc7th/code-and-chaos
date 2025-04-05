# Building an Astro + React Project with TypeScript and Bun

In this tutorial, I'll walk through the steps to set up a modern Astro project with React and TypeScript, using Bun as the package manager. Bun is a fast JavaScript runtime that can significantly speed up build times, making it an excellent choice for modern web development. We'll configure the environment step-by-step, ensuring a smooth and optimized development process

## Manual Setup of the Project Directory

Begin by creating a new directory for your project. This will be where all your files are stored.

```zsh
mkdir my-astro-blog
z my-astro-blog
```

This command creates a folder named `astro-blog` and navigates into it. Replace the folder name with whatever you'd like to name your project.

## Initialize Git (Optional)

While using Git for version control is optional, it is a good practice to track changes in your project.

```zsh
git init
```

This command initializes a new Git repository within your project directory.

## Configure Bun with Git

Bun uses `.lockb` files to manage dependencies. To prevent Git from misinterpreting these files, we need to configure Git to handle them correctly.

First, create the `.gitattributes` file:

```zsh
vim .gitattributes
```

Add the following content to it:

```
*.lockb binary diff=lockb
```

Then, configure Git to handle Bun’s lock files by running the following:

```zsh
git config diff.lockb.textconv bun
git config diff.lockb.binary true
```

## Initialize Bun Project

Now that Git is set up, you can initialize a new Bun project. This command will create the `package.json` file and prepare the project for dependency management.

```zsh
bun init --yes
```

This will generate a `package.json` file that contains default settings for your project.

## Setting Up the Package Scripts

To manage common tasks like running the development server or building the project, add scripts to the `package.json` file. Open `package.json` and modify it as follows:

```json
{
  "scripts": {
    "dev": "bunx --bun astro dev",
    "build": "bunx --bun astro check && astro build",
    "preview": "bunx --bun astro preview",
    "lint": "bunx --bun eslint .",
    "format": "bunx --bun prettier --write .",
    "astro": "astro"
  }
}
```

Explanation of Scripts:

- dev: Starts the Astro development server.
- build: Builds the project for production.
- preview: Previews the production build.
- lint: Runs ESLint to check for code quality.
- format: Formats the code using Prettier.
- astro: Allows running Astro commands directly.

## Installing Astro

Astro should be installed locally, not globally. To add Astro to your project, run:

```zsh
bun add astro
```

## React Integration

Astro supports React out-of-the-box. To integrate React into your project, use the following command:

```zsh
bun astro add react
```

This command will install the necessary dependencies for React, and it will update your `astro.config.mjs` file to include the React integration.

## Install Prettier and Astro Plugin

Next, we’ll add Prettier for consistent code formatting. Install Prettier and the Astro-specific plugin:

```mjs
bun add -d prettier prettier-plugin-astro
```

## Configure Prettier

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

This setup ensures that Prettier formats your `.astro` files correctly.

run this command `bun run format` to see if it's working properly.

## Installing ESLint for Code Linting

To maintain code quality, we’ll use ESLint. Install ESLint and the Astro plugin:

```zsh
bun add -d eslint eslint-plugin-astro
```

## Configure ESLint

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
    files: ["*.ts", "*.tsx"],
    parser: "@typescript-eslint/parser",
    plugins: {
      "@typescript-eslint/parser": import("@typescript-eslint/parser"),
    },
    extends: [
      "plugin:@typescript-eslint/recommended",
      "plugin:astro/recommended",
    ],
  },
]);
```

Install eslint-define-config dependencies

```zsh
bun add eslint-define-config
```

then run `bun run lint` to check if it's working

## Install TypeScript editor plugin

```zsh
bun add @astrojs/ts-plugin
```

then add the following to `tsconfig.json` file

```json
{
  "extends": "astro/tsconfigs/strict",
  "include": [".astro/types.d.ts", "**/*"],
  "exclude": ["dist"],
  "compilerOptions": {
    "baseUrl": ".",
    "paths": {
      "@components/*": ["src/components/*"],
      "@layouts/*": ["src/layouts/*"],
      "@pages/*": ["src/pages"]
    },
    "plugins": [
      {
        "name": "@astrojs/ts-plugin"
      }
    ],
    "verbatimModuleSyntax": true,

    // Enable latest features
    "lib": ["ESNext", "DOM"],
    "target": "ESNext",
    "module": "ESNext",
    "moduleDetection": "force",
    "jsx": "react-jsx",
    "jsxImportSource": "react",
    "allowJs": true,

    // Bundler mode
    "moduleResolution": "bundler",
    "allowImportingTsExtensions": true,
    "noEmit": true,

    // Best practices
    "strict": true,
    "skipLibCheck": true,
    "noFallthroughCasesInSwitch": true,

    // Some stricter flags (disabled by default)
    "noUnusedLocals": false,
    "noUnusedParameters": false,
    "noPropertyAccessFromIndexSignature": false
  }
}
```

## Create the astro.config.mjs File

In case Astro didn’t generate the `astro.config.mjs` file after adding React, create it manually:

```zsh
vim astro.config.mjs
```

Add the following configuration:

```mjs
import { defineConfig } from "astro/config";
import react from "@astro/react";

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

This configuration ensures that React is integrated and sets up the development server to run on port `3000`.

## Creating the Project Structure

Let’s create the basic directory structure for your project:

```zsh
mkdir -p src/layouts src/pages src/components/astro src/components/react
```

This creates the necessary folders for layouts, pages, and components.

## Create a Layout Component

Navigate to the `src/layouts` directory and create a layout file. For example, create `Base.astro`:

```zsh
vim src/layouts/Base.astro
```

Add the following content:

```astro
<!doctype html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width" />
    <meta name="generator" content={Astro.generator} />
    <title>Astro with React</title>
  </head>
  <body>
    <slot />
  </body>
</html>
```

## Create the Index Page

Now, navigate to the `src/pages` directory and create the `index.astro` file:

```zsh
vim src/pages/index.astro
```

Add the following content to display a welcome message:

```astro
---
import Base from '../layouts/Base.astro';
---

<Base>
  <h1>Welcome to Astro with React!</h1>
</Base>
```

To run the development server, simply execute:

```zsh
bun run dev
```

Now you can access your project at http://localhost:3000 and start developing!

## Congratulations! 🎉

You’ve successfully set up an Astro project with React, TypeScript, and Bun. You can now start building your modern web applications with Astro and React, enjoying fast builds, efficient development, and powerful tooling.
