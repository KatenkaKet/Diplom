<template>
  <v-app>
    <!-- Сайдбар отображается только НЕ на /login и /register -->
    <v-navigation-drawer
      v-if="!authPages.includes(route.path)"
      app
      :permanent="true"
      :width="isCollapsed ? 64 : 220"
      class="sidebar"
      @mouseenter="isCollapsed = false"
      @mouseleave="isCollapsed = true"
    >
      <div class="d-flex flex-column fill-height justify-space-between pa-2">
        <!-- Верхняя часть -->
        <div>
          <div class="user-block d-flex align-center justify-space-between mb-4 mt-2">
            <div class="d-flex align-center">
              <v-icon size="24" class="mr-2">mdi-account</v-icon>
              <span v-if="!isCollapsed" class="font-weight-medium text-white">{{ username }}</span>
            </div>
            <v-icon
              v-if="!isCollapsed"
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

            <v-list-item to="/messages" link>
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

        <!-- Логотип -->
        <div class="text-center mb-2">
          <img
            :class="['logo-img', { expanded: !isCollapsed }]"
            src="/logo_STA.png"
            alt="logo"
          />
        </div>
      </div>
    </v-navigation-drawer>

    <v-main>
      <NuxtPage />
    </v-main>
  </v-app>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()
const authPages = ['/login', '/register']
const isCollapsed = ref(true)
const username = ref('User12345')

const logout = () => {
  localStorage.removeItem('token')
  window.location.href = '/login'
}
</script>

<style scoped>
.sidebar {
  background-color: #353535 !important;
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
</style>
