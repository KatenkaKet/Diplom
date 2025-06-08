<template>
  <div class="course-card">
    <div class="image-section">
      <v-icon size="48" color="#990000">mdi-image-outline</v-icon>
    </div>
    <div class="info-section">
      <div class="text-content">
        <div class="course-title">{{ course.Title }}</div>
        <div class="course-author">От {{ course.School }}</div>
      </div>
      <NuxtLink class="course-link" :to="`/courses/${course.ID}`">Подробнее</NuxtLink>
    </div>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  course: {
    ID: number
    Title: string
    School: string
    ImageURL?: string
  }
}>()
</script>

<style scoped>
.course-card {
  width: 280px;
  height: 200px;
  border: 2px solid #990000;
  border-radius: 20px;
  display: flex;
  flex-direction: column;
  overflow: hidden;
  background-color: #FFF6F3;
}

.image-section {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
}

.info-section {
  background-color: #990000;
  color: white;
  padding: 12px;
  /* border-bottom-left-radius: 20px;
  border-bottom-right-radius: 20px; */
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: flex-end;
}

.text-content {
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.course-title {
  font-weight: 700;
  font-size: 14px;
}

.course-author {
  font-size: 12px;
  opacity: 0.9;
}

.course-link {
  background-color: white;
  color: #990000;
  padding: 4px 12px;
  border-radius: 999px;
  font-size: 13px;
  font-weight: 500;
  text-decoration: none;
  transition: background-color 0.2s ease;
}

.course-link:hover {
  background-color: #f4e3e3;
}
</style>
