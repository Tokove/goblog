<template>
  <div class="home-view">
    <!-- 顶部 Banner -->
    <section class="banner d-flex flex-column justify-content-center align-items-center text-center">
      <h1 class="banner-title gradient-text" data-aos="fade-down">
        Welcome to goblog
      </h1>
      <p class="banner-subtitle gradient-text" data-aos="fade-up" data-aos-delay="200">
        在这里，你可以发现好友的每一个动态，分享属于你的精彩瞬间
      </p>
      <router-link 
        v-if="!isLoggedIn" 
        class="btn btn-primary btn-lg mt-3" 
        :to="{name:'register'}" 
        data-aos="zoom-in" 
        data-aos-delay="400"
      >
        加入我们
      </router-link>
    </section>

    <!-- 功能入口卡片 -->
    <section class="features container my-5">
      <div class="row g-4">
        <!-- 未登录显示原始卡片 -->
        <template v-if="!isLoggedIn">
          <div class="col-md-4" data-aos="fade-up">
            <div class="card feature-card text-center p-4">
              <i class="bi bi-people-fill fs-1 mb-2"></i>
              <h5>发现好友</h5>
              <p>快速找到你的好友，了解他们的最新动态</p>
            </div>
          </div>
          <div class="col-md-4" data-aos="fade-up" data-aos-delay="100">
            <div class="card feature-card text-center p-4">
              <i class="bi bi-chat-dots-fill fs-1 mb-2"></i>
              <h5>消息</h5>
              <p>随时收发消息，不错过任何重要信息</p>
            </div>
          </div>
          <div class="col-md-4" data-aos="fade-up" data-aos-delay="200">
            <div class="card feature-card text-center p-4">
              <i class="bi bi-heart-fill fs-1 mb-2"></i>
              <h5>喜欢</h5>
              <p>对喜欢的内容点赞，留下你的足迹</p>
            </div>
          </div>
        </template>

        <!-- 登录后显示发现内容 -->
        <template v-else>
          <div class="col-md-12" data-aos="fade-up">
            <div class="card feature-card text-center p-4">
              <h5>登录后的发现内容</h5>
              <p>这里显示好友动态、推荐内容、热门帖子等</p>
            </div>
          </div>
        </template>
      </div>
    </section>

    <!-- Footer -->
    <!-- Footer -->
    <footer class="text-center py-4 bg-light mt-5">
      © 2025 goblog. All Rights Reserved.
    </footer>
  </div>
</template>

<script>
export default {
  name: "HomeView",
  data() {
    return {
      isLoggedIn: !!localStorage.getItem("token")
    }
  },
  mounted() {
    if (window.AOS) window.AOS.refresh();

    // 监听 localStorage token 变化，保证登录后内容切换
    window.addEventListener("storage", () => {
      this.isLoggedIn = !!localStorage.getItem("token");
    });
  },
};
</script>

<style scoped>
.banner {
  background: #ffffff;
  height: 320px;
  padding-top: 200px;
  padding-bottom: 60px;
  border-radius: 12px;
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
  align-items: center;
}

.banner .btn {
  font-weight: 700;
  font-family: 'Segoe UI', 'PingFang SC', 'Helvetica', 'Arial', sans-serif;
  margin-top: 20px; 
}

.banner-title {
  font-size: 3rem;
  font-weight: 800;
  font-family: 'Segoe UI', 'PingFang SC', 'Helvetica', 'Arial', sans-serif;
}

.banner-subtitle {
  font-size: 1.4rem;
  font-weight: 600;
  font-family: 'Segoe UI', 'PingFang SC', 'Helvetica', 'Arial', sans-serif;
  margin-top: 10px;
}

.feature-card {
  border: none;
  border-radius: 12px;
  transition: transform 0.2s ease, box-shadow 0.2s ease;
  cursor: pointer;
  background-color: #fff;
}

.feature-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 8px 20px rgba(0, 0, 0, 0.15);
}

.gradient-text {
  background: linear-gradient(90deg, #1D62EB, #000A50);
  background-clip: text;
  -webkit-background-clip: text;
  color: transparent;
  -webkit-text-fill-color: transparent;
}
</style>
