<template>
  <v-container class="fill-height d-flex align-center justify-center">
    <v-row justify="center">
      <v-col cols="12" lg="4" md="5" sm="8">
        <v-card class="pa-6 rounded-xl register-card" elevation="12">
          <v-card-title
            class="text-h5 font-weight-bold text-center"
            style="color: #353535"
          >
            Регистрация
          </v-card-title>
          <v-form @submit.prevent="register" class="mt-4">
            <v-text-field
              :rules="[rules.required]"
              class="mb-3"
              density="comfortable"
              label="Имя"
              required=""
              rounded=""
              v-model="first_name"
              variant="outlined"
            ></v-text-field>
            <v-text-field
              :rules="[rules.required]"
              class="mb-3"
              density="comfortable"
              label="Фамилия"
              required=""
              rounded=""
              v-model="last_name"
              variant="outlined"
            ></v-text-field>
            <v-text-field
              class="mb-3"
              density="comfortable"
              label="Отчество"
              rounded=""
              v-model="middle_name"
              variant="outlined"
            ></v-text-field>
            <v-text-field
              :rules="[rules.required, rules.email]"
              class="mb-3"
              density="comfortable"
              label="Email"
              required=""
              rounded=""
              type="email"
              v-model="email"
              variant="outlined"
            ></v-text-field>
            <v-text-field
              class="mb-3"
              density="comfortable"
              label="Телефон"
              rounded=""
              type="tel"
              v-model="phone"
              variant="outlined"
            ></v-text-field>
            <v-text-field
              :rules="[rules.required]"
              class="mb-3"
              density="comfortable"
              label="Никнейм"
              required=""
              rounded=""
              v-model="username"
              variant="outlined"
            ></v-text-field>
            <v-text-field
              :rules="[rules.required, rules.passwordLength]"
              class="mb-3"
              density="comfortable"
              label="Пароль"
              required=""
              rounded=""
              type="password"
              v-model="password"
              variant="outlined"
            ></v-text-field>
            <v-text-field
              :rules="[rules.required, rules.passwordsMatch]"
              class="mb-4"
              density="comfortable"
              label="Повторите пароль"
              required=""
              rounded=""
              type="password"
              v-model="confirmPassword"
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
              Зарегистрироваться
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
            <NuxtLink  class="register-link" to="/login"> Войти </NuxtLink>
          </div>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup lang="ts">
import { ref } from "vue"
import { useRouter } from "vue-router"

const first_name = ref("")
const last_name = ref("")
const middle_name = ref("")
const email = ref("")
const phone = ref("")
const username = ref("")
const password = ref("")
const confirmPassword = ref("")
const error = ref("")
const router = useRouter()

const { $authApi } = useNuxtApp()

const rules = {
  required: (v: string) => !!v || "Обязательное поле",
  email: (v: string) => /.+@.+\..+/.test(v) || "Невалидный email",
  passwordLength: (v: string) => v.length >= 6 || "Минимум 6 символов",
  passwordsMatch: (v: string) => v === password.value || "Пароли не совпадают",
}

const register = async () => {
  if (password.value !== confirmPassword.value) {
    error.value = "Пароли не совпадают"
    return
  }

  try {
    await $authApi.post("/register", {
      first_name: first_name.value,
      last_name: last_name.value,
      middle_name: middle_name.value,
      email: email.value,
      phone: phone.value,
      username: username.value,
      password: password.value,
    })

    router.push("/login")
  } catch (err: any) {
    error.value = err.response?.data?.error || "Ошибка регистрации"
  }
}
</script>

<style scoped="">
.register-card {
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
