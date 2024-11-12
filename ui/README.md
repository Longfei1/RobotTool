# ui

This template should help get you started developing with Vue 3 in Vite.

## Recommended IDE Setup

[VSCode](https://code.visualstudio.com/) + [Volar](https://marketplace.visualstudio.com/items?itemName=Vue.volar) (and disable Vetur).

## Type Support for `.vue` Imports in TS

TypeScript cannot handle type information for `.vue` imports by default, so we replace the `tsc` CLI with `vue-tsc` for type checking. In editors, we need [Volar](https://marketplace.visualstudio.com/items?itemName=Vue.volar) to make the TypeScript language service aware of `.vue` types.

## Customize configuration

See [Vite Configuration Reference](https://vitejs.dev/config/).

## Project Setup

```sh
npm install
```

### Compile and Hot-Reload for Development

```sh
npm run dev
```

### Type-Check, Compile and Minify for Production

```sh
npm run build
```

## 插件
1. 支持setup中name属性的语法，自动补全Ref.value
```sh
npm i vite-plugin-vue-setup-extend -D
```
2. 解决commonjs问题，具体看json-editor-vue3组件文档
```sh
npm i @originjs/vite-plugin-commonjs -D
```

## 第三方库
三方库的版本信息，可查看package.json

| 库名 | 版本 | 说明 |
|---|---|---|
| json-editor-vue3 | 1.1.1 | json编辑器，用于编辑配置、请求json |
| element-plus | 2.8.0 | ui组件 |
| pinia | 2.2.2 | 组件共享的数据中心（项目暂未使用，目前都通过prop传递数据） |
| uuid | 11.0.2 | 唯一id |
| ajv-draft-04 | 1.0.0 | ajv扩展，支持draft-04 json schema |