import { defineStore } from 'pinia'
import { ref } from 'vue'
import { jwtDecode } from 'jwt-decode'

interface User {
  id: number
  username: string
  email: string
}

interface DecodedToken {
  user_id: number
  username: string
  email: string
  exp: number
}

export const useAuthStore = defineStore('auth', () => {
  const user = ref<User | null>(null)
  const token = ref<string | null>(null)

  const setToken = (newToken: string) => {
    token.value = newToken
    localStorage.setItem('token', newToken)
    
    try {
      const decoded = jwtDecode<DecodedToken>(newToken)
      user.value = {
        id: decoded.user_id,
        username: decoded.username,
        email: decoded.email
      }
    } catch (error) {
      console.error('Error decoding token:', error)
      user.value = null
    }
  }

  const initAuth = () => {
    const storedToken = localStorage.getItem('token')
    if (storedToken) {
      setToken(storedToken)
    }
  }

  const logout = () => {
    token.value = null
    user.value = null
    localStorage.removeItem('token')
  }

  const isAuthenticated = () => {
    if (!token.value) return false
    
    try {
      const decoded = jwtDecode<DecodedToken>(token.value)
      return decoded.exp * 1000 > Date.now()
    } catch {
      return false
    }
  }

  return {
    user,
    token,
    setToken,
    initAuth,
    logout,
    isAuthenticated
  }
}) 