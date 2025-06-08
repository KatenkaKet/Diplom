import axios from 'axios'

export default defineNuxtPlugin(() => {
  const authApi = axios.create({
    baseURL: 'http://localhost:8080/api' // auth-service
  })

  const courseApi = axios.create({
    baseURL: 'http://localhost:8081/api' // course-service
  })

  const chatApi = axios.create({
  baseURL: 'http://localhost:8082/api' // chat-service
})


  return {
    provide: {
      axios: authApi,     // общий alias для авторизации (совместимость)
      authApi,            // явный доступ к auth
      courseApi,          // доступ к курсам
      chatApi
    }
  }
})
