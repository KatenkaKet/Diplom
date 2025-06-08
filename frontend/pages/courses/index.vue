<template>
  <v-container>
    <!-- Поиск и фильтр -->
    <div class="d-flex align-center mb-6">
      <v-btn
        class="filter-btn mr-4"
        height="55"
        width="73"
        elevation="0"
        color="#9B0003"
        rounded="pill"
      >
        <img src="@/public/icons/filter.svg" alt="filter" width="24" height="24" />
      </v-btn>

      <v-text-field
        v-model="search"
        placeholder="Поиск..."
        height="55"
        rounded="pill"
        hide-details
        clearable
        class="search-field"
        :style="{ backgroundColor: '#9B0003', color: 'white' }"
        @keyup.enter="handleSearch"
      />

      <v-btn 
        height="55" 
        color="#9B0003" 
        class="ml-2" 
        rounded="pill"
        @click="handleSearch"
        :loading="loading"
      >
        <img src="@/public/icons/лупа.svg" alt="filter" width="24" height="24" />
      </v-btn>
    </div>

    <!-- Loading state -->
    <v-progress-circular
      v-if="loading"
      indeterminate
      color="#9B0003"
      class="d-flex mx-auto my-8"
    />

    <!-- Популярные -->
    <template v-else>
      <h2 class="mb-2 section-title">Популярные ➝</h2>
      <v-slide-group show-arrows class="d-flex ga-4">
        <v-slide-item v-for="course in popularCourses" :key="course.ID">
          <CourseCard :course="course" />
        </v-slide-item>
      </v-slide-group>

      <!-- Для начинающих -->
      <h2 class="mt-8 mb-2 section-title">Для начинающих ➝</h2>
      <v-slide-group show-arrows class="d-flex ga-4">
        <v-slide-item v-for="course in beginnerCourses" :key="course.ID">
          <CourseCard :course="course" />
        </v-slide-item>
      </v-slide-group>
    </template>
  </v-container>
</template>

<script setup lang="ts">
import CourseCard from '@/components/CourseCard.vue';
import { ref, onMounted } from "vue";

interface Course {
  ID: number;
  Title: string;
  School: string;
  ImageURL: string;
  type: 'popular' | 'beginner';
}

const config = useRuntimeConfig();
const search = ref("");
const popularCourses = ref<Course[]>([]);
const beginnerCourses = ref<Course[]>([]);
const loading = ref(false);

const fetchCourses = async () => {
  try {
    loading.value = true;
    const response = await fetch(`${config.public.courseServiceUrl}/courses`);
    const courses: Course[] = await response.json();
    
    popularCourses.value = courses.filter(course => course.type === 'popular');
    beginnerCourses.value = courses.filter(course => course.type === 'beginner');
  } catch (error) {
    console.error('Failed to fetch courses:', error);
  } finally {
    loading.value = false;
  }
};

const handleSearch = async () => {
  if (!search.value) {
    await fetchCourses();
    return;
  }
  
  try {
    loading.value = true;
    const response = await fetch(`${config.public.courseServiceUrl}/courses/search?q=${encodeURIComponent(search.value)}`);
    const courses: Course[] = await response.json();
    popularCourses.value = courses.filter(course => course.type === 'popular');
    beginnerCourses.value = courses.filter(course => course.type === 'beginner');
  } catch (error) {
    console.error('Search failed:', error);
  } finally {
    loading.value = false;
  }
};

onMounted(() => {
  fetchCourses();
});
</script>

<style scoped>
.section-title {
  font-size: 20px;
  font-weight: 600;
  color: #990000;
}

.search-field {
  flex: 1;
  border-radius: 999px;
}
</style>
