<template>
  <v-container class="py-8">
    <v-row>
      <v-col cols="12">
        <h2 class="text-h4 font-weight-bold mb-6" style="color: #353535">
          Мои курсы
        </h2>
      </v-col>

      <v-col
        v-for="course in courses"
        :key="course.id"
        cols="12"
        sm="6"
        md="4"
      >
        <v-card class="pa-4 rounded-lg" elevation="4">
          <v-card-title class="text-h6 font-weight-medium">
            {{ course.title }}
          </v-card-title>
          <v-card-subtitle class="text-caption mb-2">
            Автор: {{ course.author }}
          </v-card-subtitle>
          <v-card-text>
            {{ course.description }}
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup>
import { ref, onMounted } from 'vue'

const courses = ref([])

onMounted(async () => {
  // В будущем заменить на реальный запрос к course-service
  courses.value = [
    {
      id: 1,
      title: 'Frontend разработка с Vue 3',
      author: 'Redford School',
      description: 'Узнайте, как создавать мощные SPA с использованием Vue.js'
    },
    {
      id: 2,
      title: 'Маркетинг для начинающих',
      author: 'Redford School',
      description: 'Освойте базовые инструменты digital-маркетинга'
    }
  ]
})
</script>

<style scoped>
.v-card {
  background-color: #ffffff;
}
</style>
