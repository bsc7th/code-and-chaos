# Bun, Vite + React, TS, Markdown Workflow

## Latest React and React DOM Installation

Installing React and React DOM dependencies, specifically the TypeScript type definitions for React. The @types/react and @types/react-dom packages provide TypeScript types for React and its DOM manipulations, enabling better TypeScript integration.

```bash
bun install --save-exact @types/react@^19.0.8 @types/react-dom@^19.0.3
```

## Configure Bun

Bun is used as the package manager and bundler here, similar to tools like npm or Yarn but faster. The .gitattributes file is configured to handle Bun’s lockfile (lockb) efficiently in git, which ensures it treats the lock file as binary and applies appropriate diff rules. This is important because Bun’s lock files are optimized for speed and are binary, unlike npm's or yarn's text-based lock files.

create `.gitattributes` file

```bash
nvim .gitattributes
```

Then add the following line:

```php
*.lockb binary diff=lockb
```

Next, configure git to handle the Bun lock files properly:

```bash
git config diff.lockb.textconv bun
git config diff.lockb.binary true
```

## Prettier Configuration

Prettier is a tool to enforce consistent code formatting. You install `Prettier` and the plugin `@ianvs/prettier-plugin-sort-imports` to automatically sort imports. Then, you configure Prettier in the `.prettierrc.json` file.

```bash
bun add -d prettier @ianvs/prettier-plugin-sort-imports
```

Create the `.prettierrc.json` configuration file:

```bash
nvim .prettierrc.json
```

Then add the following line:

```json
{
  "arrowParens": "avoid",
  "singleQuote": true,
  "bracketSpacing": true,
  "endOfLine": "lf",
  "semi": false,
  "tabWidth": 2,
  "trailingComma": "none",
  "plugins": ["@ianvs/prettier-plugin-sort-imports"],
  "importOrder": [
    "<BUILTIN_MODULES>",
    "^react",
    "<THIRD_PARTY_MODULES>",
    "^@[/]",
    "^[.].*$",
    "<TYPES>",
    "^[/].*$",
    "[.]css$"
  ]
}
```

## Markdown Configuration

Install packages to process and render Markdown content into React components.

```bash
bun add rehype-react rehype-slug remark-frontmatter remark-gfm remark-parse remark-rehype
```

## YAML installation

YAML is a human-readable data serialization format. This package installation helps in reading and processing YAML files in your project.

```bash
bun add yaml
```

## Change dev and build commands to use bun

Modify the `package.json` file to update the `dev` and `build` commands. These commands use Bun to manage dependencies and the build process.

```json
"scripts": {
  "dev": "bunx --bun vite",
  "build": "tsc -b && bunx --bun vite build",
}
```

## Personal ESLint preference

Set up ESLint for linting TypeScript and React code. Add personal preferences for the linting rules, focusing on React hooks and refreshing rules, handling unused variables and explicit types, and more.

Create the `eslint.config.js` file:

Then modify it like so:

```javascript
import js from "@eslint/js";
import reactHooks from "eslint-plugin-react-hooks";
import reactRefresh from "eslint-plugin-react-refresh";
import globals from "globals";
import tseslint from "typescript-eslint";

export default tseslint.config(
  { ignores: ["dist"] },
  {
    extends: [js.configs.recommended, ...tseslint.configs.recommended],
    files: ["**/*.{ts,tsx}"],
    languageOptions: {
      ecmaVersion: 2020,
      globals: globals.browser,
    },
    plugins: {
      "react-hooks": reactHooks,
      "react-refresh": reactRefresh,
    },
    rules: {
      ...reactHooks.configs.recommended.rules,
      "@typescript-eslint/no-explicit-any": "off",
      "@typescript-eslint/no-empty-object-type": "off",
      "@typescript-eslint/no-unused-vars": [
        "warn",
        {
          argsIgnorePattern: "^_",
          varsIgnorePattern: "^_",
        },
      ],
      "react-refresh/only-export-components": [
        "warn",
        { allowConstantExport: true },
      ],
    },
  },
);
```

Run the linter:

```bash
bun run lint
```

## Add some TypeScript preference

The `tsconfig.app.json` file is configured with TypeScript settings that suit modern JavaScript development. It enables strict mode, module resolution, paths, and import handling, ensuring better type safety and a more reliable build process.

Update the file with:

```json
{
  "compilerOptions": {
    "target": "ESNext",
    "useDefineForClassFields": true,
    "lib": ["ES2020", "DOM", "DOM.Iterable"],
    "module": "ESNext",
    "skipLibCheck": true,
    "sourceMap": true,

    /* Bundler mode */
    "moduleResolution": "bundler",
    "allowImportingTsExtensions": true,
    "esModuleInterop": true,
    "resolveJsonModule": true,
    "isolatedModules": true,
    "moduleDetection": "force",
    "noEmit": true,
    "jsx": "react-jsx",

    /* Linting */
    "strict": true,
    "noUnusedLocals": false,
    "noImplicitAny": false,
    "noUnusedParameters": true,
    "noFallthroughCasesInSwitch": true,

    "baseUrl": "./" /* Base directory to resolve non-absolute module names. */,
    "paths": {
      /* A series of entries which re-map imports to lookup locations relative to the 'baseUrl'. */
      "@/*": ["src/*"]
    }
  },
  "include": ["src"]
}
```

## Prebuild Vite and TypeScript Configuration

Enabling type checking in Vite itself. This ensures that type errors are caught in development before even attempting to build. You can add `vite-plugin-checker` to your Vite configuration:

```bash
bun add -d vite-plugin-checker
```

## Configure Vite's Resolve Alias

You’re configuring Vite for module resolution. The alias `(@)` maps the `src` directory, which helps you write cleaner and easier-to-manage import paths.

Edit the `vite.config.ts` file:

```bash
nvim vite.config.ts
```

- add the following lines to the file

```typescript
import Checker from "vite-plugin-checker";

export default defineConfig({
  plugins: [react(), Checker({ typescript: true })],
});
```

```typescript
import { fileURLToPath, URL } from "url";
import react from "@vitejs/plugin-react";
import { defineConfig } from "vite";
import Checker from "vite-plugin-checker"; // Correct import of vite-plugin-checker

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react(), Checker({ typescript: true })],
  resolve: {
    alias: [
      {
        find: "@",
        replacement: fileURLToPath(new URL("./src", import.meta.url)),
      },
    ],
  },
  server: {
    port: 3000,
  },
});
```

To ensure proper import resolution, install Bun’s TypeScript types:

```bash
bun add -d @types/bun
```

## Set Up Automatic Dependency Updates (with Bun or GitHub Actions)

To ensure your dependencies are always up to date, you can set up GitHub Actions to run automatic checks and update them. You can schedule a regular check for outdated dependencies and automate updates using tools like Renovate or Dependabot.

Example GitHub Action for Bun dependency management:

```yaml
name: Bun Dependency Update

on:
  schedule:
    - cron: "0 0 * * 0" # Every Sunday at midnight

jobs:
  update-dependencies:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout repository
        uses: actions/checkout@v2
      - name: Run Bun update
        run: bun upgrade
      - name: Commit and push updates
        run: |
          git add .
          git commit -m "Automated dependency update"
          git push
```
