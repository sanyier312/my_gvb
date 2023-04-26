<!-- eslint-disable vue/multi-word-component-names -->
<template>
  <a-button type="danger" @click="loginOut">退出</a-button>
</template>

<script>
export default {
  // eslint-disable-next-line vue/multi-word-component-names
  name: 'Header',
  methods: {
    loginOut() {
      window.sessionStorage.clear('token')
      this.$router.push('/login')
    },
  },
}
</script>
