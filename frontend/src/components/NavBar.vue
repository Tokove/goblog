<template>
  <div>
    <!-- 主导航栏 -->
    <nav class="navbar navbar-expand-lg navbar-light bg-light shadow-sm main-navbar">
      <div class="container d-flex align-items-center justify-content-between">
        <!-- 品牌Logo -->
        <router-link class="navbar-brand d-flex align-items-center" :to="{ name: 'home' }">
          <img src="/favicon.ico" alt="goblog" height="52" />
          <span class="nav-logo-lang ms-2">goblog</span>
        </router-link>

        <!-- 移动端菜单按钮 -->
        <button
          class="navbar-toggler"
          type="button"
          data-bs-toggle="collapse"
          data-bs-target="#navbarNav"
          aria-controls="navbarNav"
          aria-expanded="false"
          aria-label="Toggle navigation"
        >
          <span class="navbar-toggler-icon"></span>
        </button>

        <!-- 导航菜单 -->
        <div class="collapse navbar-collapse justify-content-center" id="navbarNav">
          <ul class="navbar-nav d-flex flex-row gap-4 align-items-center">
            <li class="nav-item">
              <router-link
                class="nav-link"
                :to="{ name: 'home' }"
                :class="{ 'active': $route.name === 'home' }"
              >
                首页
              </router-link>
            </li>
            <li class="nav-item">
              <router-link
                class="nav-link"
                :to="{ name: 'moment' }"
                :class="{ 'active': $route.name === 'moment' }"
              >
                动态
              </router-link>
            </li>
            <li class="nav-item">
              <router-link
                class="nav-link"
                :to="{ name: 'friends' }"
                :class="{ 'active': $route.name === 'friends' }"
              >
                好友
              </router-link>
            </li>
            <li class="nav-item">
              <router-link
                class="nav-link"
                :to="{ name: 'messages' }"
                :class="{ 'active': $route.name === 'messages' }"
              >
                消息
              </router-link>
            </li>
            <li class="nav-item">
              <router-link
                class="nav-link"
                :to="{ name: 'likes' }"
                :class="{ 'active': $route.name === 'likes' }"
              >
                喜欢
              </router-link>
            </li>
          </ul>

          <!-- 搜索栏 -->
          <div class="search-box d-flex align-items-center ms-4">
            <input
              type="text"
              class="form-control search-input"
              placeholder="搜索..."
            />
            <i class="bi bi-search ms-2 search-icon"></i>
          </div>
        </div>

        <!-- 登录状态切换 -->
        <div class="d-none d-lg-flex align-items-center gap-3">
          <template v-if="!isLoggedIn">
            <router-link class="btn btn-primary text-white px-3" :to="{ name: 'login' }">登录</router-link>
            <router-link class="btn btn-outline-primary text-primary bg-white px-3" :to="{ name: 'register' }">注册</router-link>
          </template>

          <template v-else>
            <div class="d-flex align-items-center gap-2">
              <img :src="user.avatar || '/default_avatar.png'" class="rounded-circle" width="36" height="36" />
              <span class="fw-bold text-dark">{{ user.username || '用户' }}</span>
              <button class="btn btn-outline-danger btn-sm" @click="logout">退出</button>
            </div>
          </template>
        </div>
      </div>
    </nav>
  </div>
</template>

<script>
export default {
  name: "NavBar",
  data() {
    return {
      user: {},
      isLoggedIn: false,
    };
  },
  methods: {
    checkLoginStatus() {
      const token = localStorage.getItem("token");
      const userData = localStorage.getItem("user");
      this.isLoggedIn = !!token;
      this.user = userData ? JSON.parse(userData) : {};
    },
    logout() {
      localStorage.removeItem("token");
      localStorage.removeItem("user");
      this.checkLoginStatus();
      this.$router.push({ name: "login" });
    }
  },
  mounted() {
    this.checkLoginStatus();
    window.addEventListener("storage", this.checkLoginStatus);
    this.$watch('$route', () => {
      this.checkLoginStatus();
    });
  },
  beforeUnmount() {
    window.removeEventListener("storage", this.checkLoginStatus);
  }
};
</script>

<style scoped>
/* 主导航栏 */
.main-navbar {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  z-index: 1040;
  height: 70px;
}

/* Logo样式 */
.navbar-brand img {
  height: 52px;
  transition: transform 0.2s ease;
}
.navbar-brand img:hover {
  transform: scale(1.05);
}
.nav-logo-lang {
  color: #00153c;
  font-weight: 700;
  font-size: 1.6rem;
  margin-left: 10px;
  letter-spacing: 0.5px;
}

/* 导航链接样式 */
.nav-link {
  color: #333 !important;
  font-weight: 600;
  font-size: 1.1rem;
  border-radius: 8px;
  padding: 8px 14px;
  transition: all 0.2s ease;
  position: relative;
}
.nav-link.active {
  color: #007acc !important;
  background-color: rgba(0, 122, 204, 0.1);
}

.nav-link:hover:not(.active) {
  background-color: rgba(0, 122, 204, 0.08);
  color: #007acc !important;
}

/* 搜索框样式 */
.search-box {
  position: relative;
}
.search-input {
  width: 180px;
  border-radius: 20px;
  padding: 6px 12px;
  font-size: 0.95rem;
  border: 1px solid #ccc;
  transition: all 0.3s ease;
}
.search-input:focus {
  width: 240px;
  outline: none;
  box-shadow: 0 0 6px rgba(0, 122, 204, 0.3);
  border-color: #007acc;
}
.search-icon {
  font-size: 1.2rem;
  color: #555;
  cursor: pointer;
  transition: color 0.2s ease;
}
.search-icon:hover {
  color: #007acc;
}

/* 导航栏背景 */
.navbar {
  backdrop-filter: blur(10px);
  background: rgba(255, 255, 255, 0.9) !important;
  padding: 0.5rem 1rem;
}

/* 响应式调整 */
@media (max-width: 992px) {
  .navbar-collapse {
    position: absolute;
    top: 100%;
    left: 0;
    right: 0;
    background-color: white;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
    padding: 1rem;
    z-index: 1040;
  }
  .navbar-nav {
    flex-direction: column !important;
    align-items: flex-start !important;
    gap: 0.5rem !important;
  }
  .search-box {
    margin-left: 0 !important;
    margin-top: 1rem;
    width: 100%;
  }
  .search-input {
    width: 100% !important;
  }
  .search-input:focus {
    width: 100% !important;
  }
}

@media (max-width: 576px) {
  .nav-logo-lang {
    font-size: 1.3rem;
  }
  .main-navbar {
    height: 60px;
  }
  .navbar-brand img {
    height: 40px;
  }
}
</style>
