<template>
  <ContentBase>
    <div class="register-box mx-auto p-4 bg-white rounded shadow">
      <h2 class="text-center mb-4">注册</h2>
      <form @submit.prevent="handleRegister">
        <!-- 用户名 -->
        <div class="mb-3">
          <input 
            type="text" 
            v-model="username" 
            class="form-control" 
            placeholder="输入用户名"
            required
          />
        </div>

        <!-- 密码 -->
        <div class="mb-3">
          <div class="input-group">
            <input
              :type="showPassword ? 'text' : 'password'"
              v-model="password"
              class="form-control"
              placeholder="输入密码"
              required
            />
            <button
              class="btn btn-outline-secondary"
              type="button"
              @click="togglePassword"
            >
              <i :class="showPassword ? 'bi bi-eye-slash' : 'bi bi-eye'"></i>
            </button>
          </div>
        </div>

        <!-- 确认密码 -->
        <div class="mb-3">
          <div class="input-group">
            <input
              :type="showConfirmPassword ? 'text' : 'password'"
              v-model="confirmPassword"
              class="form-control"
              placeholder="确认密码"
              required
            />
            <button
              class="btn btn-outline-secondary"
              type="button"
              @click="toggleConfirmPassword"
            >
              <i :class="showConfirmPassword ? 'bi bi-eye-slash' : 'bi bi-eye'"></i>
            </button>
          </div>
        </div>

        <!-- 注册按钮 -->
        <button type="submit" class="btn btn-primary w-100">注册</button>

        <p class="text-center mt-2">
          已有账号？ 
          <router-link to="/login" class="text-blue-500 hover:underline">去登录</router-link>
        </p>
      </form>
    </div>
  </ContentBase>
</template>

<script>
import ContentBase from "@/components/ContentBase.vue"
import axios from "axios"

export default {
  components: { ContentBase },
  data() {
    return {
      username: "",
      password: "",
      confirmPassword: "",
      showPassword: false,
      showConfirmPassword: false
    }
  },
  methods: {
    togglePassword() {
      this.showPassword = !this.showPassword
    },
    toggleConfirmPassword() {
      this.showConfirmPassword = !this.showConfirmPassword
    },
    async handleRegister() {
      if (!this.username || !this.password || !this.confirmPassword) {
        alert("请完整填写注册信息")
        return
      }
      if (this.password !== this.confirmPassword) {
        alert("两次密码不一致")
        return
      }
      try {
        await axios.post("http://localhost:8080/api/auth/register", {
          username: this.username,
          password: this.password
        })
        alert("注册成功！请登录")
        this.$router.push("/login")
      } catch (err) {
        alert(err.response?.data?.error || "注册失败")
      }
    }
  }
}
</script>

<style scoped>
.register-box {
  max-width: 360px;
  width: 100%;
  padding: 2rem;
  background-color: white;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0,0,0,0.1);
}
</style>
