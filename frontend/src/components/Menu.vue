<template>
  <div class="dimmed">

  </div>
</template>

<style scoped>
.sidebar {
  font-family: "Quicksand", sans-serif;
}
</style>

<script>
export default {
  data() {
    return {
      
    };
  },
  methods: {
    search_text() {
      console.log('typing...');
    }
  }
};
</script>

<style>
.route {
  text-decoration: none;
}

.dimmed {
  position: relative;
  content: " ";
  z-index: 2;
  display: block;
  position: fixed;
  height: 100%;
  width: 100%;
  flex: 1 1 auto;
  top: 0;
  left: 0;
  right: 0;
  background: rgba(0, 0, 0, 0.9);
}
</style>
