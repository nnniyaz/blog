// @ts-check

import eslint from '@eslint/js'
import tseslint from 'typescript-eslint'
import globals from 'globals'
import reactHooks from 'eslint-plugin-react-hooks'
import reactRefresh from 'eslint-plugin-react-refresh'
import js from '@eslint/js'
import importPlugin from 'eslint-plugin-import'

export default tseslint.config(
    eslint.configs.recommended,
    { ignores: ['dist'] },
    {
        extends: [js.configs.recommended, ...tseslint.configs.recommended],
        files: ['**/*.{ts,tsx}', '**/*.{js,jsx}'],
        languageOptions: {
            ecmaVersion: 2020,
            globals: globals.browser,
        },
        plugins: {
            'react-hooks': reactHooks,
            'react-refresh': reactRefresh,
            import: importPlugin,
        },
        rules: {
            ...reactHooks.configs.recommended.rules,
            'react-refresh/only-export-components': [
                'warn',
                { allowConstantExport: true },
            ],
            'no-restricted-imports': [
                'error',
                {
                    patterns: ['../../*', '../../../*'],
                },
            ],
            'import/no-relative-parent-imports': 'error',
            '@typescript-eslint/no-duplicate-enum-values': 'off',
            'no-useless-catch': 'off',
        },
    },
)
