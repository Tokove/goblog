<template>
  <ContentBase>
    <div class="login-box mx-auto p-4 bg-white rounded shadow">
      <h2 class="text-center mb-4">登录</h2>
      <form @submit.prevent="handleLogin">
        <div class="mb-3">
          <label for="username" class="form-label"></label>
          <input type="text" v-model="username" class="form-control" id="username" placeholder="输入用户名">
        </div>

        <div class="mb-3">
          <label for="password" class="form-label"></label>
          <div class="input-group">
            <input
              :type="showPassword ? 'text' : 'password'"
              v-model="password"
              class="form-control"
              id="password"
              placeholder="输入密码"
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



        <div class="mb-3 d-flex justify-content-center align-items-center">
          <input type="checkbox" class="form-check-input me-2" id="rememberMe">
          <label class="form-check-label" for="rememberMe">记住我</label>
        </div>

        <button type="submit" class="btn btn-primary w-100">登录</button>
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
      showPassword: false  // 控制密码可见
    }
  },
  methods: {
    togglePassword() {
      this.showPassword = !this.showPassword
    },
    async handleLogin() {
  if (!this.username || !this.password) {
    alert("请输入用户名和密码")
    return
  }
  try {
    const res = await axios.post("http://localhost:8080/api/auth/login", {
      username: this.username,
      password: this.password
    });

    // 保存 token
    localStorage.setItem("token", res.data.token);

    // 保存用户信息（username + avatar）
    localStorage.setItem("user", JSON.stringify(res.data.user));

    // 通知导航栏更新
    window.dispatchEvent(new Event("storage"));

    // 跳转首页
    this.$router.push("/");
  } catch (err) {
    alert(err.response?.data?.message || "登录失败");
  }
}

  }
}
</script>

<style scoped>
.login-box {
  max-width: 360px;
  width: 100%;
  padding: 2rem;
  background-color: white;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0,0,0,0.1);
}
</style>
