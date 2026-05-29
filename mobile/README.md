# 健身训练记录 App 移动端

Expo SDK 54 + React Native + TypeScript 项目，包管理工具使用 pnpm。

## 本地运行

```bash
pnpm install
pnpm start
```

## 目录说明

- `app/`：Expo Router 页面与布局。
- `assets/images/`：App 图标、启动图和 Web favicon。
- `app.json`：Expo 应用配置。
- `tsconfig.json`：TypeScript 配置，已开启 strict。

## 开发约定

- 页面和组件使用 `.tsx`。
- API、store、schema、工具函数使用 `.ts`。
- 不新增 `.js` / `.jsx` 业务文件。
- 不混用 npm、yarn、pnpm，项目只保留 `pnpm-lock.yaml`。
