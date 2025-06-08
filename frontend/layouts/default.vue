<template>
  <v-app>
    <v-navigation-drawer
      v-if="!authPages.includes(route.path)"
      app
      location="right"
      :permanent="true"
      :width="isCollapsed ? 64 : 220"
      class="sidebar"
      @mouseenter="isCollapsed = false"
      @mouseleave="isCollapsed = true"
    >
      <div class="d-flex flex-column fill-height justify-space-between pa-2">
        <!-- Верхняя часть -->
        <div>
          <div
            class="user-block d-flex align-center justify-space-between mb-4 mt-2"
          >
            <div class="d-flex align-center">
              <v-icon size="24" class="mr-2">
                {{ isAuthenticated ? "mdi-account" : "mdi-login" }}
              </v-icon>

              <span
                v-if="!isCollapsed && isAuthenticated"
                class="font-weight-medium text-white"
              >
                {{ user?.username || "Пользователь" }}
              </span>

              <NuxtLink
                v-else-if="!isCollapsed && !isAuthenticated"
                to="/login"
                class="text-white text-decoration-underline"
              >
                Войти
              </NuxtLink>
            </div>

            <v-icon
              v-if="!isCollapsed && isAuthenticated"
              size="22"
              class="logout-icon"
              @click="logout"
              style="cursor: pointer"
            >
              mdi-logout
            </v-icon>
          </div>

          <v-divider class="mb-2" />

          <!-- Меню -->
          <div v-if="isAuthenticated">
            <v-list nav dense>
              <v-list-item to="/notifications" link>
                <div class="list-item-content">
                  <v-icon>mdi-bell</v-icon>
                  <span v-if="!isCollapsed">Уведомления</span>
                </div>
              </v-list-item>

              <v-list-item to="/my-courses" link>
                <div class="list-item-content">
                  <v-icon>mdi-school</v-icon>
                  <span v-if="!isCollapsed">Мои курсы</span>
                </div>
              </v-list-item>

              <v-list-item to="/message/chats" link>
                <div class="list-item-content">
                  <v-icon>mdi-message</v-icon>
                  <span v-if="!isCollapsed">Сообщения</span>
                </div>
              </v-list-item>

              <v-list-item to="/settings" link>
                <div class="list-item-content">
                  <v-icon>mdi-cog</v-icon>
                  <span v-if="!isCollapsed">Настройки</span>
                </div>
              </v-list-item>
            </v-list>

            <div v-if="!isCollapsed" class="calendar-placeholder mt-6">
              Календарь
            </div>
          </div>
        </div>

        <!-- Логотип -->
        <div class="text-center mb-2">
          <NuxtLink to="/courses" class="logo-wrapper">
            <img
              :class="['logo-img', { expanded: !isCollapsed }]"
              src="/logo_STA.png"
              alt="logo"
            />
          </NuxtLink>
        </div>
      </div>
    </v-navigation-drawer>

    <v-main style="background-color: #FFF5EE">
      <NuxtPage />
    </v-main>
  </v-app>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'

interface UserProfile {
  id: string
  username: string
  role: string
  email?: string
  avatarUrl?: string
}

const route = useRoute()
const router = useRouter()

const authPages = ['/login', '/register']
const protectedPages = ['/my-courses', '/notifications', '/message/chats', '/settings']

const isCollapsed = ref(true)
const isAuthenticated = ref(false)
const user = ref<UserProfile | null>(null)

const logout = () => {
  if (process.client) {
    localStorage.removeItem('token')
    isAuthenticated.value = false
    user.value = null
    router.push('/courses')
  }
}

const fetchProfile = async (): Promise<void> => {
  if (!process.client) return

  const token = localStorage.getItem('token') || ''
  if (!token) {
    isAuthenticated.value = false
    user.value = null
    return
  }

  try {
    const res = await fetch('http://localhost:8080/api/profile', {
      headers: { Authorization: `Bearer ${token}` },
    })

    if (!res.ok) throw new Error('Failed to fetch profile')

    const data: UserProfile = await res.json()
    user.value = data
    isAuthenticated.value = true
  } catch (err) {
    isAuthenticated.value = false
    user.value = null
  }
}

// Следим за изменением маршрута
watch(
  () => route.path,
  async (newPath) => {
    await fetchProfile()

    if (!isAuthenticated.value && protectedPages.includes(newPath)) {
      router.push('/login')
    }
  },
  { immediate: true }
)
</script>



<style scoped>
.sidebar {
  background-color: #740101 !important;
  color: white;
  transition: width 0.3s ease;
  overflow-x: hidden;
}

.v-icon {
  color: white !important;
  font-size: 22px;
}

.v-list-item {
  padding: 0;
}

.list-item-content {
  display: flex;
  align-items: center;
  gap: 8px;
  height: 48px;
  color: white;
}

.calendar-placeholder {
  background-color: #e0e0e0;
  color: black;
  padding: 12px;
  text-align: center;
  border-radius: 8px;
  font-weight: 500;
}

.logo-img {
  transition: all 0.4s ease;
  width: 36px;
  height: auto;
  transform: rotate(0deg) scale(1);
  transform-origin: center center;
}

.logo-img.expanded {
  width: 64px;
  transform: rotate(90deg) scale(1.2);
}

.user-block {
  padding-left: 8px;
}

.logout-icon {
  margin-right: 12px;
  color: white !important;
}

.logo-wrapper {
  display: inline-block;
  cursor: pointer;
}

</style>
