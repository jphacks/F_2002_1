 module.exports= {
    preset: 'ts-jest',
    moduleNameMapper: {
      '^@/(.*)$': '<rootDir>/src/$1',
      '^~/(.*)$': '<rootDir>/src/$1',
      
    },
    moduleFileExtensions: ['js', 'json', 'ts'],
    testMatch: ['<rootDir>/src/**/*.spec.ts'],
    transform: {
      '^.+\\.js$': 'babel-jest',
      '^.+\\.ts?$': 'ts-jest'
    },
    globals: {
      'ts-jest': {
        tsconfig: 'tsconfig.json'
      }
    },
  }