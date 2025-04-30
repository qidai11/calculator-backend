// frontend/jest.config.js
module.exports = {
    testEnvironment: 'jsdom',
    setupFilesAfterEnv: ['<rootDir>/jest.setup.js'],
    transform: {
        '^.+\\.(t|j)sx?$': ['ts-jest', {
        babelConfig: true, // 启用Babel处理JSX
        diagnostics: false // 禁用TS类型检查（由ESLint处理）
        }]
    },
    moduleNameMapper: {
        '^@/(.*)$': '<rootDir>/$1',
        '\\.(css|scss)$': 'identity-obj-proxy',
        '^react$': '<rootDir>/node_modules/react' // 强制单React实例
    }
}