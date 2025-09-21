import eslint from '@eslint/js'
import tseslint from 'typescript-eslint'
import globals from "globals"

export default [
  // Base ESLint recommended rules
  eslint.configs.recommended,

  // TypeScript recommended rules
  ...tseslint.configs.recommended,

  {
    languageOptions: {
      parser: tseslint.parser,
      parserOptions: {
        ecmaVersion: 2020,
        sourceType: 'module',
      },
      globals: {
        node: true,
        es6: true,
        
      }
    },
    plugins: {
      '@typescript-eslint': tseslint.plugin,
    },
    rules: {
      '@typescript-eslint/no-unused-vars': 'off',
      '@typescript-eslint/explicit-function-return-type': 'off',
      '@typescript-eslint/explicit-module-boundary-types': 'off',
      '@typescript-eslint/no-explicit-any': 'off',
    }
  },
  {
    files: [
      '**/*.config.{js,cjs,mjs,ts}',
      'vite.config.{js,ts,mjs,cjs}'
    ],
    languageOptions: {
      ecmaVersion: 2020,
      sourceType: 'module',
      globals: {
        ...globals.node,   // defines process, module, __dirname, etc.
      },
    },
    // If some of these are CJS, ESLint handles them fine; if needed,
    // add parserOptions.sourceType per-file.
    rules: {},
  }
]