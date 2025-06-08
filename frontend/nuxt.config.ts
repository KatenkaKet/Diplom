import vuetify from 'vite-plugin-vuetify'

export default defineNuxtConfig({
  css: [
    'vuetify/styles',
    '@mdi/font/css/materialdesignicons.css',
    '@/assets/styles/main.css'
  ],
  build: {
    transpile: ['vuetify'],
  },
  vite: {
    plugins: [vuetify()], 
  },
  runtimeConfig: {
    public: {
      courseServiceUrl: process.env.COURSE_SERVICE_URL || 'http://localhost:8080/api'
    }
  }
})
