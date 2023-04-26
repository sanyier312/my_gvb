<!-- eslint-disable vue/multi-word-component-names -->
<template>
  <a-layout class="container">
    <Nav />
    <a-layout>
      <a-layout-header class="headerBtn">
        <Header />
      </a-layout-header>
      <a-layout-content>
        <router-view></router-view>
      </a-layout-content>
      <a-layout-footer>
        <Footer />
      </a-layout-footer>
    </a-layout>
  </a-layout>
</template>

<script>
import Nav from '../components/admin/Nav'
import Footer from '../components/admin/Footer'
import Header from '../components/admin/Header'
export default {
  // eslint-disable-next-line vue/multi-word-component-names
  name: 'Admin',
  components: {
    Nav,
    Footer,
    Header,
  },
}
</script>

<style scoped>
.container {
  height: 100%;
}
.headerBtn {
  display: flex;
  justify-content: flex-end;
  align-items: center;
}
</style>
