# Astro + React TS Built with Bun Workflow

## Manual Setup

Create an empty directory with the name of your project, and then navigate into it.

```bash
mkdir my-astro-project
cd my-astro-project
```

## Create package.json

```bash
bun init --yes
```

Then, replace any placeholder “scripts” section of your `package.json` with the following to use bun.

```json
{
  "name": "rjleyva-blog",
  "author": "RJ Leyva",
  "description": "RJ Leyva's Astro React TS Blog built with Bun",
  "license": "MIT",
  "type": "module",
  "private": true,
  "engine": {
    "bun": ">=0.6.0"
  },
  "scripts": {
    "dev": "bunx --bun astro dev",
    "build": "bunx --bun astro check && astro build",
    "preview": "bunx --bun astro preview",
    "lint": "bunx --bun eslint .",
    "format": "bunx --bun pretttier --write .",
    "astro": "astro"
  },
  "devDependencies": {
    "@types/bun": "latest"
  },
  "peerDependencies": {
    "typescript": "^5"
  },
  "keywords": [
    "astro",
    "bun",
    "react",
    "blog",
    "static site",
    "static site generator",
    "frontend",
    "web development",
    "javascript",
    "typescript",
    "react blog",
    "astro blog",
    "personal blog",
    "content creation",
    "web design",
    "react components",
    "astro components",
    "modern web development",
    "responsive design",
    "seo",
    "seo optimization",
    "javascript blog",
    "typescript blog",
    "open source blog",
    "minimalistic blog",
    "fast website",
    "fast loading blog",
    "developer blog",
    "web performance",
    "front-end development",
    "web app",
    "progressive web app",
    "PWA"
  ]
}
```

## Create `git module`

```bash
nvim .gitattributes
```

add this following:

```.gitattributes
*.lockb binary diff=lockb
```

then run this following commands:

```bash
git config diff.lockb.textconv bun
```

```bash
git config diff.lockb.binary true
```

**Important**

Astro must be installed locally, not globally. Make sure you are not running `bun add global astro`

```bash
bun add astro
```

## React Integrations

```bash
bun astro add react
```

## Git init

```bash
git init
```

## Install prettier and prettier-plugin-astro.

```bash
bun add -d prettier prettier-plugin-astro
```

```bash
bun add --save-dev prettier prettier-plugin-astro
```

Create a `.prettierrc.mjs` config file at the root of your project and add prettier-plugin-astro to it.

```bash
nvim .prettierrc.mjs
```

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

```typescript
import { type Config } from "prettier";

const config: Config = {
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

export default config;
```

## TypeScript editor plugin

```bash
bun add @astrojs/ts-plugin
```

```bash
nvim tsconfig.json
```

TypeScript is configured using `tsconfig.json`.

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

## Parser Installation

```bash
bun add -D eslint @typescript-eslint/parser eslint-plugin-astro
```

```eslint.config.mjs
import { defineConfig } from 'eslint-define-config'

export default defineConfig([
  {
    files: ['*.astro'],
    parser: 'astro-eslint-parser',
    parserOptions: {
      extraFileExtensions: ['.astro']
    },
    plugins: {
      astro: import('eslint-plugin-astro')
    },
    extends: ['plugin:astro/recommended']
  },
  {
    files: ['*.ts', '*.tsx'],
    parser: '@typescript-eslint/parser',
    plugins: {
      '@typescript-eslint/parser': import('@typescript-eslint/parser')
    },
    extends: [
      'plugin:@typescript-eslint/recommended',
      'plugin:astro/recommended'
    ]
  }
])
```

## Create your first static asset `robots.txt`

You will also want to create a public/ directory to store your static assets. Astro will always include these assets in your final build, so you can safely reference them from inside your component templates.

In your text editor, create a new file in your directory at public/robots.txt. robots.txt is a simple file that most sites will include to tell search bots like Google how to treat your site.

```robots.txt
# Full syntax: https://developers.google.com/search/docs/advanced/robots/create-robots-txt
User-agent: *
Allow: /
```

## Create `astro.config.mjs or ts`

```mjs
import { defineConfig } from "astro/config";

// https://astro.build/config
export default defineConfig({
  integrations: [react()],
  devToolbar: {
    enabled: false,
  },
  server: {
    port: 3000,
  },
});
```

### Create a another file at `src/layouts/Layout.astro`

```html
<!doctype html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width" />
    <meta name="generator" content="{Astro.generator}" />
    <title>Astro Basics</title>
  </head>
  <body>
    <slot />
  </body>
</html>
```

## Create index.astro

create a new file in your directory at `src/pages/index.astro`

```astro
---
import Main from '../layouts/main.astro';
---

<Main />
```
