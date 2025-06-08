<template>
  <v-container class="fill-height d-flex align-center justify-center">
    <v-row justify="center">
      <v-col cols="12" lg="4" md="5" sm="8">
        <v-card class="pa-6 rounded-xl login-card" elevation="12">
          <v-card-title
            class="text-h5 font-weight-bold text-center"
            style="color: #353535"
          >
            Вход в аккаунт
          </v-card-title>
          <v-form @submit.prevent="login" class="mt-4">
            <v-text-field
              class="mb-3"
              density="comfortable"
              label="Email"
              prepend-inner-icon="mdi-email"
              required=""
              rounded=""
              type="email"
              v-model="email"
              variant="outlined"
            ></v-text-field>
            <v-text-field
              class="mb-4"
              density="comfortable"
              label="Пароль"
              prepend-inner-icon="mdi-lock"
              required=""
              rounded=""
              type="password"
              v-model="password"
              variant="outlined"
            ></v-text-field>
            <v-btn
              block=""
              class="mb-3 text-white"
              rounded=""
              size="large"
              style="background-color: #353535"
              type="submit"
            >
              Войти
            </v-btn>
            <v-alert
              :icon="false"
              border="start"
              border-color="error"
              class="text-center mt-2 px-4 py-3 d-flex justify-center"
              density="comfortable"
              rounded=""
              type="error"
              v-if="error"
            >
              <div class="w-100 text-center">{{ error }}</div>
            </v-alert>
          </v-form>
          <div class="text-center mt-2">
            <NuxtLink class="register-link" to="/register">
              Зарегистрироваться
            </NuxtLink>
          </div>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const email = ref('')
const password = ref('')
const error = ref('') // ✅ Вот эта строка

const router = useRouter()
const { $authApi } = useNuxtApp()

const login = async () => {
  try {
    const res = await $authApi.post('/login', {
      email: email.value,
      password: password.value
    })

    const token = res.data.token
    if (token) {
      localStorage.setItem('token', token)
      router.push('/my-courses')
    }
  } catch (err: any) {
    error.value = err.response?.data?.error || 'Неверный логин или пароль'
  }
}
</script>


<style scoped="">
.login-card {
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.3);
  background-color: white;
}

.v-btn,
.v-text-field,
.v-alert {
  border-radius: 12px !important;
}

.register-link {
  color: #000000;
  font-weight: 500;
  text-decoration: underline;
  cursor: pointer;
  transition: 0.2s;
}

.register-link:hover {
  opacity: 0.8;
}
</style>
